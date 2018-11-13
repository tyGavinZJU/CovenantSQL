/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package xenomint

import (
	"database/sql"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CovenantSQL/CovenantSQL/types"
	xi "github.com/CovenantSQL/CovenantSQL/xenomint/interfaces"
	"github.com/pkg/errors"
)

type State struct {
	sync.RWMutex
	strg xi.Storage
	pool *pool

	// TODO(leventeliu): Reload savepoint from last block on chain initialization, and rollback
	// any ongoing transaction on exit.
	//
	// unc is the uncommitted transaction.
	unc     *sql.Tx
	origin  uint64 // origin is the original savepoint of the current transaction
	cmpoint uint64 // cmpoint is the last commit point of the current transaction
	current uint64 // current is the current savepoint of the current transaction
}

func newState(strg xi.Storage) (s *State, err error) {
	var t = &State{
		strg: strg,
		pool: newPool(),
	}
	if t.unc, err = t.strg.Writer().Begin(); err != nil {
		return
	}
	t.setSavepoint()
	s = t
	return
}

func (s *State) incSeq() {
	s.current++
}

func (s *State) setNextTxID() {
	var val = (s.current & uint64(0xffffffff00000000)) + uint64(1)<<32
	s.origin = val
	s.cmpoint = val
	s.current = val
}

func (s *State) rollbackID(id uint64) {
	s.current = id
}

func (s *State) getID() uint64 {
	return atomic.LoadUint64(&s.current)
}

func (s *State) close(commit bool) (err error) {
	if s.unc != nil {
		if commit {
			if err = s.unc.Commit(); err != nil {
				return
			}
		} else {
			if err = s.unc.Rollback(); err != nil {
				return
			}
		}
	}
	return s.strg.Close()
}

func buildArgsFromSQLNamedArgs(args []types.NamedArg) (ifs []interface{}) {
	ifs = make([]interface{}, len(args))
	for i, v := range args {
		ifs[i] = sql.NamedArg{
			Name:  v.Name,
			Value: v.Value,
		}
	}
	return
}

func buildTypeNamesFromSQLColumnTypes(types []*sql.ColumnType) (names []string) {
	names = make([]string, len(types))
	for i, v := range types {
		names[i] = v.DatabaseTypeName()
	}
	return
}

type sqlQuerier interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func readSingle(
	qer sqlQuerier, q *types.Query) (names []string, types []string, data [][]interface{}, err error,
) {
	var (
		rows *sql.Rows
		cols []*sql.ColumnType
	)
	if rows, err = qer.Query(
		q.Pattern, buildArgsFromSQLNamedArgs(q.Args)...,
	); err != nil {
		return
	}
	defer rows.Close()
	// Fetch column names and types
	if names, err = rows.Columns(); err != nil {
		return
	}
	if cols, err = rows.ColumnTypes(); err != nil {
		return
	}
	types = buildTypeNamesFromSQLColumnTypes(cols)
	// Scan data row by row
	data = make([][]interface{}, 0)
	for rows.Next() {
		var (
			row  = make([]interface{}, len(cols))
			dest = make([]interface{}, len(cols))
		)
		for i := range row {
			dest[i] = &row[i]
		}
		if err = rows.Scan(dest...); err != nil {
			return
		}
		data = append(data, row)
	}
	return
}

func buildRowsFromNativeData(data [][]interface{}) (rows []types.ResponseRow) {
	rows = make([]types.ResponseRow, len(data))
	for i, v := range data {
		rows[i].Values = v
	}
	return
}

func (s *State) read(req *types.Request) (ref *QueryTracker, resp *types.Response, err error) {
	var (
		ierr           error
		cnames, ctypes []string
		data           [][]interface{}
	)
	// TODO(leventeliu): no need to run every read query here.
	for i, v := range req.Payload.Queries {
		if cnames, ctypes, data, ierr = readSingle(s.strg.DirtyReader(), &v); ierr != nil {
			err = errors.Wrapf(ierr, "query at #%d failed", i)
			return
		}
	}
	// Build query response
	ref = &QueryTracker{req: req}
	resp = &types.Response{
		Header: types.SignedResponseHeader{
			ResponseHeader: types.ResponseHeader{
				Request:   req.Header,
				NodeID:    "",
				Timestamp: time.Now(),
				RowCount:  uint64(len(data)),
				LogOffset: s.getID(),
			},
		},
		Payload: types.ResponsePayload{
			Columns:   cnames,
			DeclTypes: ctypes,
			Rows:      buildRowsFromNativeData(data),
		},
	}
	return
}

func (s *State) readTx(req *types.Request) (ref *QueryTracker, resp *types.Response, err error) {
	var (
		tx             *sql.Tx
		id             uint64
		ierr           error
		cnames, ctypes []string
		data           [][]interface{}
	)
	if tx, ierr = s.strg.DirtyReader().Begin(); ierr != nil {
		err = errors.Wrap(ierr, "open tx failed")
		return
	}
	defer tx.Rollback()
	// FIXME(leventeliu): lock free but not consistent.
	id = s.getID()
	for i, v := range req.Payload.Queries {
		if cnames, ctypes, data, ierr = readSingle(tx, &v); ierr != nil {
			err = errors.Wrapf(ierr, "query at #%d failed", i)
			return
		}
	}
	// Build query response
	ref = &QueryTracker{req: req}
	resp = &types.Response{
		Header: types.SignedResponseHeader{
			ResponseHeader: types.ResponseHeader{
				Request:   req.Header,
				NodeID:    "",
				Timestamp: time.Now(),
				RowCount:  uint64(len(data)),
				LogOffset: id,
			},
		},
		Payload: types.ResponsePayload{
			Columns:   cnames,
			DeclTypes: ctypes,
			Rows:      buildRowsFromNativeData(data),
		},
	}
	return
}

func (s *State) writeSingle(q *types.Query) (res sql.Result, err error) {
	if res, err = s.unc.Exec(q.Pattern, buildArgsFromSQLNamedArgs(q.Args)...); err == nil {
		s.incSeq()
	}
	return
}

func (s *State) setSavepoint() (savepoint uint64) {
	savepoint = s.getID()
	s.unc.Exec("SAVEPOINT ?", savepoint)
	return
}

func (s *State) rollbackTo(savepoint uint64) {
	s.rollbackID(savepoint)
	s.unc.Exec("ROLLBACK TO ?", savepoint)
}

func (s *State) write(req *types.Request) (ref *QueryTracker, resp *types.Response, err error) {
	var (
		ierr      error
		savepoint uint64
		query     = &QueryTracker{req: req}
	)
	// TODO(leventeliu): savepoint is a sqlite-specified solution for nested transaction.
	func() {
		s.Lock()
		defer s.Unlock()
		savepoint = s.getID()
		for i, v := range req.Payload.Queries {
			if _, ierr = s.writeSingle(&v); ierr != nil {
				err = errors.Wrapf(ierr, "execute at #%d failed", i)
				s.rollbackTo(savepoint)
				return
			}
		}
		s.setSavepoint()
		s.pool.enqueue(savepoint, query)
	}()
	// Build query response
	ref = query
	resp = &types.Response{
		Header: types.SignedResponseHeader{
			ResponseHeader: types.ResponseHeader{
				Request:   req.Header,
				NodeID:    "",
				Timestamp: time.Now(),
				RowCount:  0,
				LogOffset: savepoint,
			},
		},
	}
	return
}

func (s *State) replay(req *types.Request, resp *types.Response) (err error) {
	var (
		ierr      error
		savepoint uint64
		query     = &QueryTracker{req: req, resp: resp}
	)
	s.Lock()
	defer s.Unlock()
	savepoint = s.getID()
	if resp.Header.ResponseHeader.LogOffset != savepoint {
		err = ErrQueryConflict
		return
	}
	for i, v := range req.Payload.Queries {
		if _, ierr = s.writeSingle(&v); ierr != nil {
			err = errors.Wrapf(ierr, "execute at #%d failed", i)
			s.rollbackTo(savepoint)
			return
		}
	}
	s.setSavepoint()
	s.pool.enqueue(savepoint, query)
	return
}

func (s *State) replayBlock(block *types.Block) (err error) {
	var (
		ierr      error
		savepoint uint64
	)
	s.Lock()
	defer s.Unlock()
	for i, q := range block.QueryTxs {
		var query = &QueryTracker{req: q.Request, resp: &types.Response{Header: *q.Response}}
		savepoint = s.getID()
		if q.Response.ResponseHeader.LogOffset > savepoint {
			err = ErrMissingParent
			return
		}
		// Match and skip already pooled query
		if q.Response.ResponseHeader.LogOffset < savepoint {
			if !s.pool.match(savepoint, q.Request) {
				err = ErrQueryConflict
				return
			}
			continue
		}
		// Replay query
		for j, v := range q.Request.Payload.Queries {
			if q.Request.Header.QueryType == types.ReadQuery {
				continue
			}
			if q.Request.Header.QueryType != types.WriteQuery {
				err = errors.Wrapf(ErrInvalidRequest, "replay block at %d:%d", i, j)
				s.rollbackTo(savepoint)
				return
			}
			if _, ierr = s.writeSingle(&v); ierr != nil {
				err = errors.Wrapf(ierr, "execute at %d:%d failed", i, j)
				s.rollbackTo(savepoint)
				return
			}
			s.setSavepoint()
			s.pool.enqueue(savepoint, query)
		}
	}
	// Check if the current transaction is ok to commit
	if s.pool.matchLast(savepoint) {
		if err = s.unc.Commit(); err != nil {
			// FATAL ERROR
			return
		}
		if s.unc, err = s.strg.Writer().Begin(); err != nil {
			// FATAL ERROR
			return
		}
	}
	s.setNextTxID()
	s.setSavepoint()
	// Truncate pooled queries
	s.pool.truncate(savepoint)
	return
}

func (s *State) commit(out chan<- command) (err error) {
	var queries []*QueryTracker
	s.Lock()
	defer s.Unlock()
	if err = s.unc.Commit(); err != nil {
		return
	}
	if s.unc, err = s.strg.Writer().Begin(); err != nil {
		return
	}
	s.setNextTxID()
	s.setSavepoint()
	queries = s.pool.queries
	s.pool = newPool()
	// NOTE(leventeliu): Send commit request to the outcoming command queue within locking scope to
	// prevent any extra writing before this commit.
	if out != nil {
		out <- &commitRequest{queries: queries}
	}
	return
}

func (s *State) partialCommit(qs []*types.Response) (err error) {
	var (
		i  int
		v  *types.Response
		q  *QueryTracker
		rm []*QueryTracker

		pl = len(s.pool.queries)
	)
	if len(qs) > pl {
		err = ErrLocalBehindRemote
		return
	}
	for i, v = range qs {
		var loc = s.pool.queries[i]
		if loc.req.Header.Hash() != v.Header.Request.Hash() ||
			loc.resp.Header.LogOffset != v.Header.LogOffset {
			err = ErrQueryConflict
			return
		}
	}
	if i < pl-1 {
		s.rollbackTo(s.pool.queries[i+1].resp.Header.LogOffset)
	}
	// Reset pool
	rm = s.pool.queries[:i+1]
	s.pool = newPool()
	// Rewrite
	for _, q = range rm {
		if _, _, err = s.write(q.req); err != nil {
			return
		}
	}
	return
}

func (s *State) rollback() (err error) {
	s.Lock()
	defer s.Unlock()
	s.rollbackTo(s.cmpoint)
	s.current = s.cmpoint
	return
}

func (s *State) Query(req *types.Request) (ref *QueryTracker, resp *types.Response, err error) {
	switch req.Header.QueryType {
	case types.ReadQuery:
		return s.readTx(req)
	case types.WriteQuery:
		return s.write(req)
	default:
		err = ErrInvalidRequest
	}
	return
}

// Replay replays a write log from other peer to replicate storage state.
func (s *State) Replay(req *types.Request, resp *types.Response) (err error) {
	switch req.Header.QueryType {
	case types.ReadQuery:
		return
	case types.WriteQuery:
		return s.replay(req, resp)
	default:
		err = ErrInvalidRequest
	}
	return
}
