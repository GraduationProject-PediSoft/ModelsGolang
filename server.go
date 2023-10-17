package main

import (
	"go-microservice/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	eureka "github.com/xuanbo/eureka-client"
)

const defaultPort = "8001"

func main() {
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://127.0.0.1:8761/eureka/",
		App:                   "go-microservice",
		Port:                  8001,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"MICROSERVICE-TYPE":    "IA-MODEL",
		},
	})
	
	client.Start()

	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
