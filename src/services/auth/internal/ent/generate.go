package ent

import (
	_ "github.com/luminancetech/varso/src/services/auth/internal/ent/build/runtime"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert --feature sql/versioned-migration --feature sql/execquery --feature intercept ./schema --target ./build
