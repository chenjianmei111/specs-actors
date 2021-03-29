// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package market

import (
	"fmt"
	"io"

	abi "github.com/chenjianmei111/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

var lengthBufState = []byte{139}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Proposals (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Proposals); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposals: %w", err)
	}

	// t.States (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.States); err != nil {
		return xerrors.Errorf("failed to write cid field t.States: %w", err)
	}

	// t.PendingProposals (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.PendingProposals); err != nil {
		return xerrors.Errorf("failed to write cid field t.PendingProposals: %w", err)
	}

	// t.EscrowTable (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.EscrowTable); err != nil {
		return xerrors.Errorf("failed to write cid field t.EscrowTable: %w", err)
	}

	// t.LockedTable (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.LockedTable); err != nil {
		return xerrors.Errorf("failed to write cid field t.LockedTable: %w", err)
	}

	// t.NextID (abi.DealID) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.NextID)); err != nil {
		return err
	}

	// t.DealOpsByEpoch (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.DealOpsByEpoch); err != nil {
		return xerrors.Errorf("failed to write cid field t.DealOpsByEpoch: %w", err)
	}

	// t.LastCron (abi.ChainEpoch) (int64)
	if t.LastCron >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.LastCron)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.LastCron-1)); err != nil {
			return err
		}
	}

	// t.TotalClientLockedCollateral (big.Int) (struct)
	if err := t.TotalClientLockedCollateral.MarshalCBOR(w); err != nil {
		return err
	}

	// t.TotalProviderLockedCollateral (big.Int) (struct)
	if err := t.TotalProviderLockedCollateral.MarshalCBOR(w); err != nil {
		return err
	}

	// t.TotalClientStorageFee (big.Int) (struct)
	if err := t.TotalClientStorageFee.MarshalCBOR(w); err != nil {
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

	if extra != 11 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Proposals (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Proposals: %w", err)
		}

		t.Proposals = c

	}
	// t.States (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.States: %w", err)
		}

		t.States = c

	}
	// t.PendingProposals (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PendingProposals: %w", err)
		}

		t.PendingProposals = c

	}
	// t.EscrowTable (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.EscrowTable: %w", err)
		}

		t.EscrowTable = c

	}
	// t.LockedTable (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.LockedTable: %w", err)
		}

		t.LockedTable = c

	}
	// t.NextID (abi.DealID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.NextID = abi.DealID(extra)

	}
	// t.DealOpsByEpoch (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.DealOpsByEpoch: %w", err)
		}

		t.DealOpsByEpoch = c

	}
	// t.LastCron (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.LastCron = abi.ChainEpoch(extraI)
	}
	// t.TotalClientLockedCollateral (big.Int) (struct)

	{

		if err := t.TotalClientLockedCollateral.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalClientLockedCollateral: %w", err)
		}

	}
	// t.TotalProviderLockedCollateral (big.Int) (struct)

	{

		if err := t.TotalProviderLockedCollateral.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalProviderLockedCollateral: %w", err)
		}

	}
	// t.TotalClientStorageFee (big.Int) (struct)

	{

		if err := t.TotalClientStorageFee.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalClientStorageFee: %w", err)
		}

	}
	return nil
}

var lengthBufVerifyDealsForActivationReturn = []byte{131}

func (t *VerifyDealsForActivationReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufVerifyDealsForActivationReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealWeight (big.Int) (struct)
	if err := t.DealWeight.MarshalCBOR(w); err != nil {
		return err
	}

	// t.VerifiedDealWeight (big.Int) (struct)
	if err := t.VerifiedDealWeight.MarshalCBOR(w); err != nil {
		return err
	}

	// t.DealSpace (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.DealSpace)); err != nil {
		return err
	}

	return nil
}

func (t *VerifyDealsForActivationReturn) UnmarshalCBOR(r io.Reader) error {
	*t = VerifyDealsForActivationReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealWeight (big.Int) (struct)

	{

		if err := t.DealWeight.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.DealWeight: %w", err)
		}

	}
	// t.VerifiedDealWeight (big.Int) (struct)

	{

		if err := t.VerifiedDealWeight.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.VerifiedDealWeight: %w", err)
		}

	}
	// t.DealSpace (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.DealSpace = uint64(extra)

	}
	return nil
}

var lengthBufDealState = []byte{131}

func (t *DealState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDealState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.SectorStartEpoch (abi.ChainEpoch) (int64)
	if t.SectorStartEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SectorStartEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.SectorStartEpoch-1)); err != nil {
			return err
		}
	}

	// t.LastUpdatedEpoch (abi.ChainEpoch) (int64)
	if t.LastUpdatedEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.LastUpdatedEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.LastUpdatedEpoch-1)); err != nil {
			return err
		}
	}

	// t.SlashEpoch (abi.ChainEpoch) (int64)
	if t.SlashEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SlashEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.SlashEpoch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *DealState) UnmarshalCBOR(r io.Reader) error {
	*t = DealState{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.SectorStartEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SectorStartEpoch = abi.ChainEpoch(extraI)
	}
	// t.LastUpdatedEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.LastUpdatedEpoch = abi.ChainEpoch(extraI)
	}
	// t.SlashEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SlashEpoch = abi.ChainEpoch(extraI)
	}
	return nil
}
