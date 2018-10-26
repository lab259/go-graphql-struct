package gqlstruct_test

import (
	"github.com/graphql-go/graphql"
	"time"
)

type UUIDTyped struct {
	Value string
}

func (t *UUIDTyped) GraphqlType() graphql.Type {
	return graphql.Float
}

type ModelComplete struct {
	ID           UUIDTyped  `graphql:"id"`
	IDPtr        *UUIDTyped `graphql:"idPtr"`
	Name         string     `graphql:"name"`
	NamePtr      *string    `graphql:"namePtr"`
	CreatedAt    time.Time  `graphql:"createdAt"`
	CreatedAtPtr *time.Time `graphql:"createdAtPtr"`
}
