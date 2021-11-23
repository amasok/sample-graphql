package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/amasok/sample-graphql/app/middleware"
	"github.com/amasok/sample-graphql/app/presentation/graphql"
	"github.com/amasok/sample-graphql/app/presentation/graphql/directive"
	"github.com/amasok/sample-graphql/app/presentation/graphql/generated"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter()
	r.Use(middleware.CORSForGraphql)
	r.Use(middleware.AuthForGraphql)

	config := generated.Config{
		Resolvers: &graphql.Resolver{},
		Directives: generated.DirectiveRoot{
			Auth: directive.Auth, // @authがついてるディレクティブで認証の承認を行う
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
