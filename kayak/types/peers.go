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

import (
	"github.com/CovenantSQL/CovenantSQL/crypto/asymmetric"
	"github.com/CovenantSQL/CovenantSQL/crypto/hash"
	"github.com/CovenantSQL/CovenantSQL/crypto/verifier"
	"github.com/CovenantSQL/CovenantSQL/proto"
	"github.com/pkg/errors"
)

//go:generate hsp

// Peers defines the kayak peers conf.
type Peers struct {
	DatabaseID         proto.DatabaseID
	Servers            *proto.Peers
	Term               uint64
	Leader             proto.NodeID
	LastCompletedIndex uint64
}

// SignedPeers defines the signed kayak peers conf.
type SignedPeers struct {
	Peers
	Hash       hash.Hash
	Signees    []*asymmetric.PublicKey
	Signatures []*asymmetric.Signature
}

// AddSign adds new signature to peers structure.
func (sp *SignedPeers) AddSign(signer *asymmetric.PrivateKey) (err error) {
	if sp == nil {
		return
	}

	if sp.Hash.IsEqual(&hash.Hash{}) {
		// build hash
		var enc []byte
		if enc, err = sp.Peers.MarshalHash(); err != nil {
			return
		}

		sp.Hash = hash.THashH(enc)
	}

	var sig *asymmetric.Signature
	if sig, err = signer.Sign(sp.Hash.AsBytes()); err != nil {
		return
	}

	sp.Signees = append(sp.Signees, signer.PubKey())
	sp.Signatures = append(sp.Signatures, sig)

	return

}

func (sp *SignedPeers) ClearSignatures() {
	if sp == nil {
		return
	}

	sp.Hash = hash.Hash{}
	sp.Signatures = sp.Signatures[:0]
}

func (sp *SignedPeers) Verify() (err error) {
	if sp == nil {
		err = errors.WithStack(verifier.ErrSignatureNotMatch)
		return
	}
	if len(sp.Signatures) == 0 {
		err = errors.WithStack(verifier.ErrSignatureNotMatch)
		return
	}
	if len(sp.Signees) != len(sp.Signatures) {
		err = errors.WithStack(verifier.ErrSignatureNotMatch)
		return
	}

	// verify hash
	var enc []byte
	if enc, err = sp.Peers.MarshalHash(); err != nil {
		return
	}

	var h = hash.THashH(enc)
	if !sp.Hash.IsEqual(&h) {
		err = errors.WithStack(verifier.ErrHashValueNotMatch)
		return
	}
	// verify all signatures
	for i, sig := range sp.Signatures {
		if !sig.Verify(h.AsBytes(), sp.Signees[i]) {
			err = errors.WithStack(verifier.ErrSignatureNotMatch)
			return
		}
	}
	// verify embedded block producer generated peers
	if err = sp.Servers.Verify(); err != nil {
		return
	}

	return
}
