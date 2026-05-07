package provider

import (
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/receipt"
	"github.com/fil-forge/go-ucanto/core/result/failure"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
)

const AddAbility = "provider/add"

type AddCaveats struct {
	Provider string
	Consumer string
}

func (ac AddCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ac, AddCaveatsType(), types.Converters...)
}

var AddCaveatsReader = schema.Struct[AddCaveats](AddCaveatsType(), nil, types.Converters...)

type AddOk struct {
	Id string
}

func (ao AddOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ao, AddOkType(), types.Converters...)
}

type AddReceipt receipt.Receipt[AddOk, failure.Failure]
type AddReceiptReader receipt.ReceiptReader[AddOk, failure.Failure]

func NewAddReceiptReader() (AddReceiptReader, error) {
	return receipt.NewReceiptReader[AddOk, failure.Failure](providerSchema, types.Converters...)
}

var AddOkReader = schema.Struct[AddOk](AddOkType(), nil, types.Converters...)

var Add = validator.NewCapability(
	AddAbility,
	schema.DIDString(),
	AddCaveatsReader,
	validator.DefaultDerives[AddCaveats],
)
