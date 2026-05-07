package consumer

import (
	_ "embed"
	"fmt"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/ipld/go-ipld-prime/schema"
)

//go:embed consumer.ipldsch
var consumerSchema []byte

var consumerTS = mustLoadTS()

func mustLoadTS() *schema.TypeSystem {
	ts, err := types.LoadSchemaBytes(consumerSchema)
	if err != nil {
		panic(fmt.Errorf("loading consumer schema: %w", err))
	}
	return ts
}

func HasCaveatsType() schema.Type {
	return consumerTS.TypeByName("HasCaveats")
}

func HasOkType() schema.Type {
	return consumerTS.TypeByName("HasOk")
}

func GetCaveatsType() schema.Type {
	return consumerTS.TypeByName("GetCaveats")
}

func GetOkType() schema.Type {
	return consumerTS.TypeByName("GetOk")
}
