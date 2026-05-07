package assert

import (
	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"
)

// InclusionAbility claims that a CID links to other CIDs.
const RelationAbility = "assert/relation"

type RelationPartInclusion struct {
	Content ipld.Link
	Parts   *[]ipld.Link
}

type RelationPart struct {
	Content  ipld.Link
	Includes *RelationPartInclusion
}

type RelationCaveats struct {
	Content  types.HasMultihash
	Children []ipld.Link
	Parts    []RelationPart
}

func (rc RelationCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&rc, RelationCaveatsType(), types.Converters...)
}

var RelationCaveatsReader = schema.Struct[RelationCaveats](RelationCaveatsType(), nil, types.Converters...)

var Relation = validator.NewCapability(
	RelationAbility,
	schema.DIDString(),
	RelationCaveatsReader,
	nil,
)
