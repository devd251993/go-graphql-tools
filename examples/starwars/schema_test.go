package starwars

import (
	"testing"

	tools "github.com/devd251993/go-graphql-tools"
	"github.com/devd251993/graphql"
)

func TestSchema(t *testing.T) {
	gen := tools.NewGenerator(nil)
	query := gen.GenerateObject(Query{})
	mutation := gen.GenerateObject(Mutation{})

	_, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		t.Fatal(err)
	}

}
