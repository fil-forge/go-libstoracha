package assert

import (
	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"
)

// EqualsAbility claims data is referred to by another CID and/or multihash. e.g
// CAR CID & CommP CID
const EqualsAbility = "assert/equals"

type EqualsCaveats struct {
	Content types.HasMultihash
	Equals  ipld.Link
}

func (ec EqualsCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ec, EqualsCaveatsType(), types.Converters...)
}

var EqualsCaveatsReader = schema.Struct[EqualsCaveats](EqualsCaveatsType(), nil, types.Converters...)

var Equals = validator.NewCapability(
	EqualsAbility,
	schema.DIDString(),
	EqualsCaveatsReader,
	nil,
)
