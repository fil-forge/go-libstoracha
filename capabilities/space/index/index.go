package index

import (
	_ "embed"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/result/failure"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/ucan"
	"github.com/fil-forge/go-ucanto/validator"
)

const IndexAbility = "space/index/*"

// Index capability definition
// This capability can only be delegated (but not invoked) allowing audience to
// derive any `space/index/` prefixed capability for the space identified by the DID
// in the `with` field.
var Index = validator.NewCapability(
	IndexAbility,
	schema.DIDString(),
	schema.Struct[struct{}](nil, nil, types.Converters...),
	func(claimed, delegated ucan.Capability[struct{}]) failure.Failure {
		return equalWith(claimed, delegated)
	},
)
