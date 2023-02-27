package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFoldVal(t *testing.T) {
	// GIVEN
	graphName := "mygraph"
	g := NewGraph(graphName)
	require.NotNil(t, g)
	v := g.V()
	require.NotNil(t, v)

	// WHEN
	result := v.ValuesBy("test").Fold()

	// THEN
	assert.NotNil(t, result)
	assert.Equal(t, fmt.Sprintf("%s.V().values(\"test\").fold()", graphName), result.String())
}
