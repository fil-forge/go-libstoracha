package pdp

import (
	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-libstoracha/piece/piece"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/receipt"
	"github.com/fil-forge/go-ucanto/core/result/failure"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/filecoin-project/go-data-segment/merkletree"
	"github.com/ipld/go-ipld-prime/datamodel"
	mh "github.com/multiformats/go-multihash"
)

const InfoAbility = "pdp/info"

type InfoCaveats struct {
	Blob mh.Multihash
}

func (ic InfoCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ic, InfoCaveatsType(), types.Converters...)
}

type InfoAcceptedAggregate struct {
	Aggregate      piece.PieceLink
	InclusionProof merkletree.ProofData
}

type InfoOk struct {
	Piece      piece.PieceLink
	Aggregates []InfoAcceptedAggregate
}

func (io InfoOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&io, InfoOkType(), types.Converters...)
}

type InfoReceipt receipt.Receipt[InfoOk, failure.Failure]
type InfoReceiptReader receipt.ReceiptReader[InfoOk, failure.Failure]

func NewInfoReceiptReader() (InfoReceiptReader, error) {
	return receipt.NewReceiptReader[InfoOk, failure.Failure](pdpSchema, types.Converters...)
}

var InfoCaveatsReader = schema.Struct[InfoCaveats](InfoCaveatsType(), nil, types.Converters...)

var InfoOkReader = schema.Struct[InfoOk](InfoOkType(), nil, types.Converters...)

var Info = validator.NewCapability(
	InfoAbility,
	schema.DIDString(),
	InfoCaveatsReader,
	validator.DefaultDerives,
)
