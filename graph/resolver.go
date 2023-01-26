package graph

import "github.com/fasibio/autogql_example/graph/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Sql *db.AutoGqlDB
}
