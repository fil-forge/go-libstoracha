package space

import (
	_ "embed"
	"fmt"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/ipld/go-ipld-prime/schema"
)

//go:embed space.ipldsch
var spaceSchema []byte

var spaceTS = mustLoadTS()

func mustLoadTS() *schema.TypeSystem {
	ts, err := types.LoadSchemaBytes(spaceSchema)
	if err != nil {
		panic(fmt.Errorf("loading space schema: %w", err))
	}
	return ts
}

func InfoCaveatsType() schema.Type {
	return spaceTS.TypeByName("InfoCaveats")
}

func InfoOkType() schema.Type {
	return spaceTS.TypeByName("InfoOk")
}
