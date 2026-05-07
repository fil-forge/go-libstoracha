package blob

import (
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

const AcceptAbility = "web3.storage/blob/accept"

type AcceptCaveats struct {
	Space did.DID
	Blob  types.Blob
	TTL   int
	Put   types.Promise
}

func (ac AcceptCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ac, AcceptCaveatsType(), types.Converters...)
}

var AcceptCaveatsReader = schema.Struct[AcceptCaveats](AcceptCaveatsType(), nil, types.Converters...)

type AcceptOk struct {
	Site ucan.Link
}

func (ao AcceptOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ao, AcceptOkType(), types.Converters...)
}

type AcceptReceipt receipt.Receipt[AcceptOk, failure.Failure]
type AcceptReceiptReader receipt.ReceiptReader[AcceptOk, failure.Failure]

func NewAcceptReceiptReader() (AcceptReceiptReader, error) {
	return receipt.NewReceiptReader[AcceptOk, failure.Failure](blobSchema, types.Converters...)
}

var AcceptOkReader = schema.Struct[AcceptOk](AcceptOkType(), nil, types.Converters...)

var Accept = validator.NewCapability(
	AcceptAbility,
	schema.DIDString(),
	AcceptCaveatsReader,
	func(claimed, delegated ucan.Capability[AcceptCaveats]) failure.Failure {
		fail := equalWith(claimed.With(), delegated.With())
		if fail != nil {
			return fail
		}

		fail = equalBlob(claimed.Nb().Blob, delegated.Nb().Blob)
		if fail != nil {
			return fail
		}

		fail = equalTTL(claimed.Nb().TTL, delegated.Nb().TTL)
		if fail != nil {
			return fail
		}

		return checkSpace(claimed.Nb().Space.String(), delegated.Nb().Space.String())
	},
)
