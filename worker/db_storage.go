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

package worker

import (
	"container/list"
	"context"

	"github.com/CovenantSQL/CovenantSQL/types"
	x "github.com/CovenantSQL/CovenantSQL/xenomint"
	"github.com/pkg/errors"
)

// Following contains storage related logic extracted from main database instance definition.

// Check implements kayak.types.Handler.Check.
func (db *Database) Check(req *types.Request) (err error) {
	// verify signature, check time/sequence only
	if err = req.Verify(); err != nil {
		return
	}

	// verify timestamp
	nowTime := getLocalTime()
	minTime := nowTime.Add(-db.cfg.MaxWriteTimeGap)
	maxTime := nowTime.Add(db.cfg.MaxWriteTimeGap)

	if req.Header.Timestamp.Before(minTime) || req.Header.Timestamp.After(maxTime) {
		err = errors.Wrap(ErrInvalidRequest, "invalid request time")
		return
	}

	// verify sequence
	if err = db.verifySequence(req.Header.ConnectionID, req.Header.SeqNo); err != nil {
		return
	}

	// record sequence
	db.recordSequence(req.Header.ConnectionID, req.Header.SeqNo)

	return
}

// TrackerAndResponse defines a query tracker used by xenomint and an unsigned response.
type TrackerAndResponse struct {
	Tracker  *x.QueryTracker
	Response *types.Response
}

// Commit implements kayak.types.Handler.Commit.
func (db *Database) Commit(req *types.Request, isLeader bool) (result interface{}, err error) {
	// convert query and check syntax
	var (
		response *types.Response
		tracker  *x.QueryTracker
	)

	// reset context, commit should never be canceled
	req.SetContext(context.Background())

	// execute
	if tracker, response, err = db.chain.Query(req, isLeader); err != nil {
		return
	}
	result = &TrackerAndResponse{
		Tracker:  tracker,
		Response: response,
	}
	return
}

func (db *Database) recordSequence(connID uint64, seqNo uint64) {
	db.connSeqs.Store(connID, seqNo)
}

func (db *Database) verifySequence(connID uint64, seqNo uint64) (err error) {
	var data interface{}
	var ok bool
	var lastSeq uint64

	if data, ok = db.connSeqs.Load(connID); ok {
		lastSeq, _ = data.(uint64)

		if seqNo <= lastSeq {
			return ErrInvalidRequestSeq
		}
	}

	return
}

func (db *Database) evictSequences() {
	m := make(map[uint64]*list.Element)
	l := list.New()

	for connSeq := range db.connSeqEvictCh {
		if e, ok := m[connSeq]; ok {
			l.MoveToFront(e)
			return
		}

		e := l.PushFront(connSeq)
		m[connSeq] = e

		if l.Len() > MaxRecordedConnectionSequences {
			e = l.Back()
			if e != nil {
				l.Remove(e)
				evictSeq := e.Value.(uint64)
				delete(m, evictSeq)
				db.connSeqs.Delete(evictSeq)
			}
		}
	}
}
