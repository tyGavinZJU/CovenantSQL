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

package types

import (
	"reflect"

	"github.com/CovenantSQL/CovenantSQL/utils"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

// LogType defines the log type.
type LogType uint8

const (
	// LogTypePrepare defines the prepare phase of a commit.
	LogTypePrepare LogType = iota
	// LogTypeRollback defines the rollback phase of a commit.
	LogTypeRollback
	// LogTypeCommit defines the commit phase of a commit.
	LogTypeCommit
	// LogTypeNoop defines noop log.
	LogTypeNoop
	// LogTypePeers defines the epoch of the new elected term.
	LogTypePeers
	// number of log types.
	numLogTypes

	// msgpack constants, copied from go/codec/msgpack.go
	valueTypeMap   = 9
	valueTypeArray = 10
)

var (
	logTypeMap = [...]string{
		"LogTypePrepare",
		"LogTypeRollback",
		"LogTypeCommit",
		"LogTypeNoop",
		"LogTypePeers",
		"Unknown",
	}
	rtLog        = reflect.TypeOf((*Log)(nil)).Elem()
	rtLogWrapper = reflect.TypeOf((*LogWrapper)(nil))
)

func init() {
	// detect msgpack version
	if codec.GenVersion != 8 {
		panic(ErrMsgPackVersionMismatch)
	}

	// register log wrapper to msgpack handler
	if err := utils.RegisterInterfaceToMsgPack(rtLog, rtLogWrapper); err != nil {
		panic(err)
	}
}

func (t LogType) String() (s string) {
	if t >= numLogTypes {
		return logTypeMap[numLogTypes]
	}

	return logTypeMap[t]
}

// Log defines the log type interface.
type Log interface {
	GetType() LogType
	SetType(LogType)
	GetIndex() uint64
	SetIndex(uint64)
	GetTerm() uint64
	GetPeerTerm() uint64
	GetVersion() uint8
}

// LogHeader defines the checksum header structure.
type LogHeader struct {
	Version  uint8   // log version
	Type     LogType // log type
	PeerTerm uint64
	Term     uint64
	Index    uint64 // log index
}

// GetIndex returns the log index.
func (h *LogHeader) GetIndex() uint64 {
	return h.Index
}

// SetIndex updates the index.
func (h *LogHeader) SetIndex(i uint64) {
	h.Index = i
}

// GetTerm returns the term of the log.
func (h *LogHeader) GetTerm() uint64 {
	return h.Term
}

// GetPeerTerm returns the peer term of the log.
func (h *LogHeader) GetPeerTerm() uint64 {
	return h.PeerTerm
}

// GetVersion returns the log version for future compatibility.
func (h *LogHeader) GetVersion() uint8 {
	return h.Version
}

// GetType returns the log type.
func (h *LogHeader) GetType() LogType {
	return h.Type
}

// SetType update the log type.
func (h *LogHeader) SetType(logType LogType) {
	h.Type = logType
}

// LogWrapper handles the log deserialization.
type LogWrapper struct {
	Log
}

// Unwrap returns log within wrapper.
func (w *LogWrapper) Unwrap() Log {
	return w.Log
}

// CodecEncodeSelf implements codec.Selfer interface.
func (w *LogWrapper) CodecEncodeSelf(e *codec.Encoder) {
	helperEncoder, encDriver := codec.GenHelperEncoder(e)

	if w == nil || w.Log == nil {
		encDriver.EncodeNil()
		return
	}

	// translate wrapper to two fields array wrapped by map
	encDriver.WriteArrayStart(2)
	encDriver.WriteArrayElem()
	encDriver.EncodeUint(uint64(w.GetType()))
	encDriver.WriteArrayElem()
	helperEncoder.EncFallback(w.Log)
	encDriver.WriteArrayEnd()
}

// CodecDecodeSelf implements codec.Selfer interface.
func (w *LogWrapper) CodecDecodeSelf(d *codec.Decoder) {
	_, decodeDriver := codec.GenHelperDecoder(d)

	// clear fields
	w.Log = nil
	ct := decodeDriver.ContainerType()

	switch ct {
	case valueTypeArray:
		w.decodeFromWrapper(d)
	case valueTypeMap:
		w.decodeFromRaw(d)
	default:
		panic(errors.Wrapf(ErrInvalidContainerType, "type %v applied", ct))
	}
}

func (w *LogWrapper) newLog(t LogType) (l Log, err error) {
	switch t {
	case LogTypePrepare:
		l = NewLogPrepare()
	case LogTypeCommit:
		l = NewLogCommit()
	case LogTypeRollback:
		l = NewLogRollback()
	case LogTypeNoop:
		l = NewLogNoop()
	case LogTypePeers:
		l = NewLogPeers()
	default:
		err = errors.Wrapf(ErrInvalidLogType, "invalid log type %v", t)
		return
	}

	return
}

func (w *LogWrapper) decodeFromWrapper(d *codec.Decoder) {
	helperDecoder, decodeDriver := codec.GenHelperDecoder(d)
	containerLen := decodeDriver.ReadArrayStart()

	for i := 0; i < containerLen; i++ {
		if decodeDriver.CheckBreak() {
			break
		}

		decodeDriver.ReadArrayElem()

		if i == 0 {
			if decodeDriver.TryDecodeAsNil() {
				// invalid type, can not instantiate transaction
				panic(ErrInvalidLogType)
			} else {
				var (
					logType LogType
					err     error
				)
				helperDecoder.DecFallback(&logType, true)
				if w.Log, err = w.newLog(logType); err != nil {
					panic(err)
				}
			}
		} else if i == 1 {
			if ct := decodeDriver.ContainerType(); ct != valueTypeMap {
				panic(errors.Wrapf(ErrInvalidContainerType, "type %v applied", ct))
			}

			if !decodeDriver.TryDecodeAsNil() {
				// the container type should be struct
				helperDecoder.DecFallback(&w.Log, true)
			}
		} else {
			helperDecoder.DecStructFieldNotFound(i, "")
		}
	}

	decodeDriver.ReadArrayEnd()

	if containerLen < 2 {
		panic(ErrInvalidLogType)
	}
}

func (w *LogWrapper) decodeFromRaw(d *codec.Decoder) {
	helperDecoder, _ := codec.GenHelperDecoder(d)

	// read all container as raw
	rawBytes := helperDecoder.DecRaw()

	var (
		typeDetector LogHeader
		err          error
	)
	typeDetector.Type = numLogTypes

	if err = utils.DecodeMsgPack(rawBytes, &typeDetector); err != nil {
		panic(err)
	}

	if typeDetector.Type == numLogTypes {
		panic(ErrInvalidLogType)
	}

	if w.Log, err = w.newLog(typeDetector.Type); err != nil {
		panic(err)
	}

	if err = utils.DecodeMsgPack(rawBytes, w.Log); err != nil {
		panic(err)
	}
}
