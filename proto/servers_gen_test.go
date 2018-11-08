package proto

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func TestMarshalHashPeers(t *testing.T) {
	v := Peers{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashPeers(b *testing.B) {
	v := Peers{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgPeers(b *testing.B) {
	v := Peers{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashPeersHeader(t *testing.T) {
	v := PeersHeader{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashPeersHeader(b *testing.B) {
	v := PeersHeader{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgPeersHeader(b *testing.B) {
	v := PeersHeader{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}