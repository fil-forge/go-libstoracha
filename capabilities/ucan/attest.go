package ucan

import (
	"fmt"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/fil-forge/go-ucanto/core/ipld"
	"github.com/fil-forge/go-ucanto/core/result/failure"
	"github.com/fil-forge/go-ucanto/core/schema"
	"github.com/fil-forge/go-ucanto/ucan"
	"github.com/fil-forge/go-ucanto/validator"
	"github.com/ipld/go-ipld-prime/datamodel"
)

const AttestAbility = "ucan/attest"

// AttestCaveats represents the caveats of a ucan/attest delegation.
type AttestCaveats struct {
	// UCAN delegation that is being attested.
	Proof ipld.Link
}

func (ac AttestCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ac, AttestCaveatsType(), types.Converters...)
}

var AttestCaveatsReader = schema.Struct[AttestCaveats](AttestCaveatsType(), nil, types.Converters...)

// Issued by trusted authority (usually the one handling invocation) who attests
// that a specific UCAN delegation has been considered authentic.
//
// https://github.com/storacha/specs/blob/main/w3-session.md#authorization-session
var Attest = validator.NewCapability(
	AttestAbility,
	schema.DIDString(),
	AttestCaveatsReader,
	func(claimed, delegated ucan.Capability[AttestCaveats]) failure.Failure {
		// Check if the with field matches
		if fail := equalWith(claimed, delegated); fail != nil {
			return fail
		}
		// Check if the proof matches
		if fail := equalProof(claimed, delegated); fail != nil {
			return fail
		}
		return nil
	},
)

// equalProof checks if the proof matches between two capabilities.
func equalProof(claimed, delegated ucan.Capability[AttestCaveats]) failure.Failure {
	if delegated.Nb().Proof != nil && delegated.Nb().Proof.String() != claimed.Nb().Proof.String() {
		return schema.NewSchemaError(fmt.Sprintf(
			"claimed proof '%s' doesn't match delegated '%s'",
			claimed.Nb().Proof, delegated.Nb().Proof,
		))
	}
	return nil
}
