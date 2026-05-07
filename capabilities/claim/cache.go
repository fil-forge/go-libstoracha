package claim

import (
	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/multiformats/go-multiaddr"
)

const CacheAbility = "claim/cache"

type Provider struct {
	Addresses []multiaddr.Multiaddr
}

type CacheCaveats struct {
	Claim    ipld.Link
	Provider Provider
}

func (cc CacheCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&cc, CacheCaveatsType(), types.Converters...)
}

var CacheCaveatsReader = schema.Struct[CacheCaveats](CacheCaveatsType(), nil, types.Converters...)

var Cache = validator.NewCapability(CacheAbility, schema.DIDString(), CacheCaveatsReader, nil)
