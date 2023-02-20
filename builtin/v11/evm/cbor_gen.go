// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package evm

import (
	"fmt"
	"io"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

var lengthBufState = []byte{133}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Bytecode (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Bytecode); err != nil {
		return xerrors.Errorf("failed to write cid field t.Bytecode: %w", err)
	}

	// t.BytecodeHash ([32]uint8) (array)
	if len(t.BytecodeHash) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.BytecodeHash was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.BytecodeHash))); err != nil {
		return err
	}

	if _, err := w.Write(t.BytecodeHash[:]); err != nil {
		return err
	}

	// t.ContractState (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.ContractState); err != nil {
		return xerrors.Errorf("failed to write cid field t.ContractState: %w", err)
	}

	// t.Nonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	// t.Tombstone (evm.Tombstone) (struct)
	if err := t.Tombstone.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) error {
	*t = State{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Bytecode (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Bytecode: %w", err)
		}

		t.Bytecode = c

	}
	// t.BytecodeHash ([32]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.BytecodeHash: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 32 {
		return fmt.Errorf("expected array to have 32 elements")
	}

	t.BytecodeHash = [32]uint8{}

	if _, err := io.ReadFull(br, t.BytecodeHash[:]); err != nil {
		return err
	}
	// t.ContractState (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ContractState: %w", err)
		}

		t.ContractState = c

	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	// t.Tombstone (evm.Tombstone) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.Tombstone = new(Tombstone)
			if err := t.Tombstone.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Tombstone pointer: %w", err)
			}
		}

	}
	return nil
}

var lengthBufTombstone = []byte{130}

func (t *Tombstone) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufTombstone); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Origin (abi.ActorID) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Origin)); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	return nil
}

func (t *Tombstone) UnmarshalCBOR(r io.Reader) error {
	*t = Tombstone{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Origin (abi.ActorID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Origin = abi.ActorID(extra)

	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	return nil
}

var lengthBufConstructorParams = []byte{130}

func (t *ConstructorParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufConstructorParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Creator ([20]uint8) (array)
	if len(t.Creator) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Creator was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Creator))); err != nil {
		return err
	}

	if _, err := w.Write(t.Creator[:]); err != nil {
		return err
	}

	// t.Initcode ([]uint8) (slice)
	if len(t.Initcode) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Initcode was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Initcode))); err != nil {
		return err
	}

	if _, err := w.Write(t.Initcode[:]); err != nil {
		return err
	}
	return nil
}

func (t *ConstructorParams) UnmarshalCBOR(r io.Reader) error {
	*t = ConstructorParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Creator ([20]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Creator: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 20 {
		return fmt.Errorf("expected array to have 20 elements")
	}

	t.Creator = [20]uint8{}

	if _, err := io.ReadFull(br, t.Creator[:]); err != nil {
		return err
	}
	// t.Initcode ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Initcode: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Initcode = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Initcode[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufGetStorageAtParams = []byte{129}

func (t *GetStorageAtParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufGetStorageAtParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.StorageKey ([32]uint8) (array)
	if len(t.StorageKey) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.StorageKey was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.StorageKey))); err != nil {
		return err
	}

	if _, err := w.Write(t.StorageKey[:]); err != nil {
		return err
	}
	return nil
}

func (t *GetStorageAtParams) UnmarshalCBOR(r io.Reader) error {
	*t = GetStorageAtParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.StorageKey ([32]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.StorageKey: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 32 {
		return fmt.Errorf("expected array to have 32 elements")
	}

	t.StorageKey = [32]uint8{}

	if _, err := io.ReadFull(br, t.StorageKey[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufDelegateCallParams = []byte{132}

func (t *DelegateCallParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDelegateCallParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Code (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Code); err != nil {
		return xerrors.Errorf("failed to write cid field t.Code: %w", err)
	}

	// t.Input ([]uint8) (slice)
	if len(t.Input) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Input was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Input))); err != nil {
		return err
	}

	if _, err := w.Write(t.Input[:]); err != nil {
		return err
	}

	// t.Caller ([20]uint8) (array)
	if len(t.Caller) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Caller was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Caller))); err != nil {
		return err
	}

	if _, err := w.Write(t.Caller[:]); err != nil {
		return err
	}

	// t.Value (big.Int) (struct)
	if err := t.Value.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DelegateCallParams) UnmarshalCBOR(r io.Reader) error {
	*t = DelegateCallParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Code (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Code: %w", err)
		}

		t.Code = c

	}
	// t.Input ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Input: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Input = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Input[:]); err != nil {
		return err
	}
	// t.Caller ([20]uint8) (array)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Caller: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra != 20 {
		return fmt.Errorf("expected array to have 20 elements")
	}

	t.Caller = [20]uint8{}

	if _, err := io.ReadFull(br, t.Caller[:]); err != nil {
		return err
	}
	// t.Value (big.Int) (struct)

	{

		if err := t.Value.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Value: %w", err)
		}

	}
	return nil
}
