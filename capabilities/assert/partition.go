package assert

import (
	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"
)

// PartitionAbility claims that a CID's graph can be read from the blocks found
// in parts.
const PartitionAbility = "assert/partition"

type PartitionCaveats struct {
	Content types.HasMultihash
	Blocks  *ipld.Link
	Parts   []ipld.Link
}

func (pc PartitionCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&pc, PartitionCaveatsType(), types.Converters...)
}

var PartitionCaveatsReader = schema.Struct[PartitionCaveats](PartitionCaveatsType(), nil, types.Converters...)
var Partition = validator.NewCapability(
	PartitionAbility,
	schema.DIDString(),
	PartitionCaveatsReader, nil)
