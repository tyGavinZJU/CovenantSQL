/*
 * Copyright 2019 The CovenantSQL Authors.
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

package types

import "github.com/CovenantSQL/CovenantSQL/types"

// LogPrepare is the prepare operation log type.
type LogPrepare struct {
	LogHeader
	Request *types.Request
}

// LogCommit is the commit operation log type.
type LogCommit struct {
	LogHeader
	PrepareIndex  uint64
	LastCommitted uint64
}

// LogRollback is the rollback operation log type.
type LogRollback struct {
	LogHeader
	PrepareIndex uint64
}

// LogNoop is the noop/heartbeat log type.
type LogNoop struct {
	LogHeader
}

// LogPeers is the new service term epoch log type.
type LogPeers struct {
	LogHeader
	Peers *SignedPeers
}

// NewLogPrepare returns new prepare log.
func NewLogPrepare() (l *LogPrepare) {
	l = new(LogPrepare)
	l.Type = LogTypePrepare
	return
}

// NewLogCommit returns the new commit log.
func NewLogCommit() (l *LogCommit) {
	l = new(LogCommit)
	l.Type = LogTypeCommit
	return
}

// NewLogRollback returns the new rollback log.
func NewLogRollback() (l *LogRollback) {
	l = new(LogRollback)
	l.Type = LogTypeRollback
	return
}

// NewLogNoop returns the new noop log.
func NewLogNoop() (l *LogNoop) {
	l = new(LogNoop)
	l.Type = LogTypeNoop
	return
}

// NewLogPeers returns the new peers update log.
func NewLogPeers() (l *LogPeers) {
	l = new(LogPeers)
	l.Type = LogTypePeers
	return
}
