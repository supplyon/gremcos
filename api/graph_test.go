package api

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGraph(t *testing.T) {
	// GIVEN
	graphName := "mygraph"

	// WHEN
	g := NewGraph(graphName)

	// THEN
	assert.NotNil(t, g)
	assert.Equal(t, graphName, g.String())
}

func TestV(t *testing.T) {

	// GIVEN
	graphName := "mygraph"
	g := NewGraph(graphName)

	// WHEN
	v := g.V()

	// THEN
	assert.NotNil(t, v)
	assert.Equal(t, fmt.Sprintf("%s.V()", graphName), v.String())
}

func TestVBy(t *testing.T) {

	// GIVEN
	graphName := "mygraph"
	g := NewGraph(graphName)
	id := 1

	// WHEN
	v := g.VBy(id)

	// THEN
	assert.NotNil(t, v)
	assert.Equal(t, fmt.Sprintf("%s.V(\"%d\")", graphName, id), v.String())
}

func TestVByUUID(t *testing.T) {

	// GIVEN
	graphName := "mygraph"
	g := NewGraph(graphName)
	id, err := uuid.NewV4()
	require.NoError(t, err)

	// WHEN
	v := g.VByUUID(id)

	// THEN
	assert.NotNil(t, v)
	assert.Equal(t, fmt.Sprintf("%s.V(\"%s\")", graphName, id), v.String())
}

func TestVByStr(t *testing.T) {

	// GIVEN
	graphName := "mygraph"
	g := NewGraph(graphName)
	id := "1234ABCD"

	// WHEN
	v := g.VByStr(id)

	// THEN
	assert.NotNil(t, v)
	assert.Equal(t, fmt.Sprintf("%s.V(\"%s\")", graphName, id), v.String())
}

func TestAddV(t *testing.T) {

	// GIVEN
	graphName := "mygraph"
	g := NewGraph(graphName)
	label := "user"

	// WHEN
	v := g.AddV(label)

	// THEN
	assert.NotNil(t, v)
	assert.Equal(t, fmt.Sprintf("%s.addV(\"%s\")", graphName, label), v.String())
}

func TestE(t *testing.T) {

	// GIVEN
	graphName := "mygraph"
	g := NewGraph(graphName)

	// WHEN
	v := g.E()

	// THEN
	assert.NotNil(t, v)
	assert.Equal(t, fmt.Sprintf("%s.E()", graphName), v.String())
}

func TestMultiparamQuery(t *testing.T) {

	// GIVEN
	queryStr := ".outE"
	l1 := "label1"
	l2 := "label2"

	// WHEN
	q1 := multiParamQuery(queryStr)
	q2 := multiParamQuery(queryStr, l1)
	q3 := multiParamQuery(queryStr, l1, l2)

	// THEN
	assert.NotNil(t, q1)
	assert.Equal(t, ".outE()", q1.String())
	assert.Equal(t, ".outE(\"label1\")", q2.String())
	assert.Equal(t, ".outE(\"label1\",\"label2\")", q3.String())
}

func TestMultiparamQueryInt(t *testing.T) {

	// GIVEN
	queryStr := ".within"
	v1 := 1
	v2 := 2

	// WHEN
	q1 := multiParamQueryInt(queryStr)
	q2 := multiParamQueryInt(queryStr, v1)
	q3 := multiParamQueryInt(queryStr, v1, v2)

	// THEN
	assert.NotNil(t, q1)
	assert.Equal(t, ".within()", q1.String())
	assert.Equal(t, ".within(1)", q2.String())
	assert.Equal(t, ".within(1,2)", q3.String())
}
