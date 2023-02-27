package api

import (
	"github.com/supplyon/gremcos/interfaces"
)

type value struct {
	builders []interfaces.QueryBuilder
}

func (v *value) String() string {
	queryString := ""
	for _, queryBuilder := range v.builders {
		queryString += queryBuilder.String()
	}
	return queryString
}

func NewValueV(e interfaces.Vertex) interfaces.Value {
	queryBuilders := make([]interfaces.QueryBuilder, 0)
	queryBuilders = append(queryBuilders, e)

	return &value{
		builders: queryBuilders,
	}
}

// Add can be used to add a custom QueryBuilder
// e.g. g.V().Add(NewSimpleQB(".myCustomCall("%s")",label))
func (v *value) Add(builder interfaces.QueryBuilder) interfaces.Value {
	v.builders = append(v.builders, builder)
	return v
}

// Fold adds .fold() to the query.
func (v *value) Fold() interfaces.Value {
	return v.Add(NewSimpleQB(".fold()"))
}
