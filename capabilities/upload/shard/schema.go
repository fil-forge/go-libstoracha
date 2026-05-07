package shard

import (
	_ "embed"
	"fmt"

	"github.com/fil-forge/go-libstoracha/capabilities/types"
	"github.com/ipld/go-ipld-prime/schema"
)

//go:embed shard.ipldsch
var shardSchema []byte

var shardTS = mustLoadTS()

func mustLoadTS() *schema.TypeSystem {
	ts, err := types.LoadSchemaBytes(shardSchema)
	if err != nil {
		panic(fmt.Errorf("loading shard schema: %w", err))
	}
	return ts
}

func ListCaveatsType() schema.Type {
	return shardTS.TypeByName("ListCaveats")
}

func ListOkType() schema.Type {
	return shardTS.TypeByName("ListOk")
}

func ListErrorType() schema.Type {
	return shardTS.TypeByName("ListError")
}
