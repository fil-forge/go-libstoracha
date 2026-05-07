package access

import (
	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/result/failure"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/ucan"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"
)

const AuthorizeAbility = "access/authorize"

// AuthorizeFactKey is the key of a fact in any delegation created by a
// successful authorization. The value is a link back to the `access/authorize`
// delegation.
const AuthorizeRequestFactKey = "access/request"

// AuthorizeCaveats represents the caveats required to perform an
// access/authorize invocation.
type AuthorizeCaveats struct {
	// DID of the Account authorization is requested from.
	Iss *string
	// Capabilities agent wishes to be granted.
	Att []CapabilityRequest
}

func (pc AuthorizeCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&pc, AuthorizeCaveatsType(), types.Converters...)
}

var AuthorizeCaveatsReader = schema.Struct[AuthorizeCaveats](AuthorizeCaveatsType(), nil, types.Converters...)

// AuthorizeOk represents the successful response for a access/authorize
// invocation.
type AuthorizeOk struct {
	Request    ipld.Link
	Expiration ucan.UTCUnixTimestamp
}

func (po AuthorizeOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&po, AuthorizeOkType(), types.Converters...)
}

var AuthorizeOkReader = schema.Struct[AuthorizeOk](AuthorizeOkType(), nil, types.Converters...)

// Authorize can be invoked by an agent to request set of capabilities from the
// account.
var Authorize = validator.NewCapability(
	AuthorizeAbility,
	schema.DIDString(),
	AuthorizeCaveatsReader,
	AuthorizeDerive,
)

func AuthorizeDerive(claimed, delegated ucan.Capability[AuthorizeCaveats]) failure.Failure {
	if fail := equalWith(claimed, delegated); fail != nil {
		return fail
	}

	if fail := equal(claimed.Nb().Iss, delegated.Nb().Iss, "iss"); fail != nil {
		return fail
	}

	if fail := subsetCapabilities(claimed.Nb().Att, delegated.Nb().Att); fail != nil {
		return fail
	}

	return nil
}
