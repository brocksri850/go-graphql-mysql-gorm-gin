package app

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/github.com/brocksri850/go-graphql-mysql-gorm-gin/graph"
	"github.com/github.com/brocksri850/go-graphql-mysql-gorm-gin/graph/generated"
)

const defaultPort = ":8080"

func InitApplication() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/example", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "example",
			"Engaje":  "Golang",
		})
	})

	r.Run(defaultPort)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
