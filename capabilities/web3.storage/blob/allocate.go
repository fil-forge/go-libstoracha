package blob

import (
	"net/http"
	"net/url"
	"time"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/receipt"
	"github.com/fil-forge/go-ucanto/core/result/failure"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/did"
	"github.com/fil-forge/go-ucanto/ucan"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"
)

const AllocateAbility = "web3.storage/blob/allocate"

type AllocateCaveats struct {
	Space did.DID
	Blob  types.Blob
	Cause ucan.Link
}

func (ac AllocateCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ac, AllocateCaveatsType(), types.Converters...)
}

var AllocateCaveatsReader = schema.Struct[AllocateCaveats](AllocateCaveatsType(), nil, types.Converters...)

type Address struct {
	URL       url.URL
	Headers   http.Header
	ExpiresAt time.Time
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

var AllocateOkReader = schema.Struct[AllocateOk](AllocateOkType(), nil, types.Converters...)

var Allocate = validator.NewCapability(
	AllocateAbility,
	schema.DIDString(),
	AllocateCaveatsReader,
	func(claimed, delegated ucan.Capability[AllocateCaveats]) failure.Failure {
		fail := equalWith(claimed.With(), delegated.With())
		if fail != nil {
			return fail
		}

		fail = equalBlob(claimed.Nb().Blob, delegated.Nb().Blob)
		if fail != nil {
			return fail
		}

		fail = checkLink(claimed.Nb().Cause, delegated.Nb().Cause)
		if fail != nil {
			return fail
		}

		return checkSpace(claimed.Nb().Space.String(), delegated.Nb().Space.String())
	},
)
