package blob

import (
	"net/http"
	"net/url"

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

const AllocateAbility = "blob/allocate"

type AllocateCaveats struct {
	Space did.DID
	Blob  types.Blob
	Cause ucan.Link
}

func (ac AllocateCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ac, AllocateCaveatsType(), types.Converters...)
}

type Address struct {
	URL     url.URL
	Headers http.Header
	Expires uint64
}

type AllocateOk struct {
	Size    uint64
	Address *Address
}

func (ao AllocateOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ao, AllocateOkType(), types.Converters...)
}

type AllocateReceipt receipt.Receipt[AllocateOk, failure.Failure]
type AllocateReceiptReader receipt.ReceiptReader[AllocateOk, failure.Failure]

func NewAllocateReceiptReader() (AllocateReceiptReader, error) {
	return receipt.NewReceiptReader[AllocateOk, failure.Failure](blobSchema, types.Converters...)
}

var AllocateCaveatsReader = schema.Struct[AllocateCaveats](AllocateCaveatsType(), nil, types.Converters...)
var Allocate = validator.NewCapability(
	AllocateAbility,
	schema.DIDString(),
	AllocateCaveatsReader,
	validator.DefaultDerives,
)
