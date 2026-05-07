package access_test

import (
	"testing"

	"github.com/fil-forge/go-libstoracha/capabilities/access"
	"github.com/fil-forge/go-libstoracha/testutil"
	"github.com/fil-forge/go-ucanto/core/ipld/codec/cbor"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestRoundTripClaimOk(t *testing.T) {
	delegations := access.DelegationsModel{
		Keys:   make([]string, 0, 2),
		Values: make(map[string][]byte),
	}
	for range 2 {
		bytes := testutil.RandomBytes(t, 256)
		mh := testutil.MultihashFromBytes(t, bytes)
		delegations.Keys = append(delegations.Keys, cid.NewCidV1(cbor.Code, mh).String())
		delegations.Values[cid.NewCidV1(cbor.Code, mh).String()] = bytes
	}
	ok := access.ClaimOk{
		Delegations: delegations,
	}

	node, err := ok.ToIPLD()
	require.NoError(t, err)

	rok, err := access.ClaimOkReader.Read(node)
	require.NoError(t, err)
	require.Equal(t, ok, rok)
}
