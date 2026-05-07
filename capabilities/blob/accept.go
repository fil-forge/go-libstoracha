package blob

import (
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/receipt"
	"github.com/fil-forge/go-ucanto/core/result/failure"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/did"
	"github.com/fil-forge/go-ucanto/ucan"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
)

const AcceptAbility = "blob/accept"

type Await struct {
	Selector string
	Link     ipld.Link
}

type Promise struct {
	UcanAwait Await
}

type AcceptCaveats struct {
	Space did.DID
	Blob  types.Blob
	Put   Promise
}

func (ac AcceptCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ac, AcceptCaveatsType(), types.Converters...)
}

type AcceptOk struct {
	Site ucan.Link
	PDP  *ucan.Link
}

func (ao AcceptOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ao, AcceptOkType(), types.Converters...)
}

type AcceptReceipt receipt.Receipt[AcceptOk, failure.Failure]
type AcceptReceiptReader receipt.ReceiptReader[AcceptOk, failure.Failure]

func NewAcceptReceiptReader() (AcceptReceiptReader, error) {
	return receipt.NewReceiptReader[AcceptOk, failure.Failure](blobSchema, types.Converters...)
}

var AcceptCaveatsReader = schema.Struct[AcceptCaveats](AcceptCaveatsType(), nil, types.Converters...)
var Accept = validator.NewCapability(
	AcceptAbility,
	schema.DIDString(),
	AcceptCaveatsReader,
	validator.DefaultDerives,
)
