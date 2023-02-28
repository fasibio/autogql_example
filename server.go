package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fasibio/autogql_example/graph"
	"github.com/fasibio/autogql_example/graph/db"
	"github.com/fasibio/autogql_example/graph/generated"
	"github.com/fasibio/autogql_example/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	dbCon, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbCon = dbCon.Debug()
	dborm := db.NewAutoGqlDB(dbCon)
	// db.AddGetHook[model.User](&dborm, "GetUser", GetUserHook{})
	db.AddQueryHook[model.Company, model.CompanyFiltersInput, model.CompanyOrder](&dborm, "QueryCompany", QueryCompanyHook{})
	db.AddAddHook[model.Company, model.CompanyInput, model.AddCompanyPayload](&dborm, "AddCompany", AddCompanyHook{})
	db.AddUpdateHook[model.Company, model.UpdateCompanyInput, model.UpdateCompanyPayload](&dborm, "UpdateCompany", UpdateCompanyHook{})
	db.AddDeleteHook[model.Company, model.CompanyFiltersInput, model.DeleteCompanyPayload](&dborm, "DeleteCompany", DeleteCompanyHook{})
	dborm.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Sql: &dborm}, Directives: generated.DirectiveRoot{
		Validated: DirectiveImpl,
		A:         DirectiveImpl,
		B:         DirectiveImpl,
		C:         DirectiveImpl,
		D:         DirectiveImpl,
		E:         DirectiveImpl,
		F:         DirectiveImpl,
		G:         DirectiveImpl,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func DirectiveImpl(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return next(ctx)
}
