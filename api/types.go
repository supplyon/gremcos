package api

import (
	"fmt"

	"github.com/spf13/cast"
)

// Property represents the cosmos db type for a property.
// As it would be returned by a call to .properties().
// As it would be returned by a call to .properties().
type Property struct {
	ID    string     `mapstructure:"id"`
	Value TypedValue `mapstructure:"value,squash"`
	Label string     `mapstructure:"label"`
}

// Edge represents the cosmos DB type for an edge.
// As it would be returned by a call to g.E().
type Edge struct {
	ID        string `mapstructure:"id"`
	Label     string `mapstructure:"label"`
	Type      Type   `mapstructure:"type"`
	InVLabel  string `mapstructure:"inVLabel"`
	InV       string `mapstructure:"inV"`
	OutVLabel string `mapstructure:"outVLabel"`
	OutV      string `mapstructure:"outV"`
}

// Vertex represents the cosmos DB type for an vertex.
// As it would be returned by a call to g.V().
type Vertex struct {
	Type       Type              `mapstructure:"type"`
	ID         string            `mapstructure:"id"`
	Label      string            `mapstructure:"label"`
	Properties VertexPropertyMap `mapstructure:"properties"`
}

// ValueWithID represents the cosmos DB type for a value in case
// it is used/ attached to a complex type.
type ValueWithID struct {
	ID    string     `mapstructure:"id"`
	Value TypedValue `mapstructure:"value,squash"`
}

type VertexPropertyMap map[string][]ValueWithID

// Type defines the cosmos db complex types
type Type string

const (
	TypeVertex Type = "vertex"
	TypeEdge   Type = "edge"
)

// TypedValue represents the cosmos DB type for a value in case
// it is not used/ attached to a complex type.
type TypedValue struct {
	Value interface{}
}

// toValue converts the given input to a TypedValue
func toValue(input interface{}) (TypedValue, error) {
	return TypedValue{Value: input}, nil
}

// converts a list of values to TypedValue
func toValues(input []interface{}) ([]TypedValue, error) {
	if input == nil {
		return nil, fmt.Errorf("Data is nil")
	}

	result := make([]TypedValue, 0, len(input))
	for _, element := range input {
		value, err := toValue(element)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}

	return result, nil
}

func (tv TypedValue) AsFloat64E() (float64, error) {
	return cast.ToFloat64E(tv.Value)
}

func (tv TypedValue) AsFloat64() float64 {
	return cast.ToFloat64(tv.Value)
}

func (tv TypedValue) AsInt32E() (int32, error) {
	return cast.ToInt32E(tv.Value)
}

func (tv TypedValue) AsInt32() int32 {
	return cast.ToInt32(tv.Value)
}

func (tv TypedValue) AsBoolE() (bool, error) {
	return cast.ToBoolE(tv.Value)
}

func (tv TypedValue) AsBool() bool {
	return cast.ToBool(tv.Value)
}

func (tv TypedValue) AsStringE() (string, error) {
	return cast.ToStringE(tv.Value)
}

func (tv TypedValue) AsString() string {
	return cast.ToString(tv.Value)
}

func (tv TypedValue) String() string {
	return fmt.Sprintf("%v", tv.Value)
}

func (v Vertex) String() string {
	return fmt.Sprintf("%s %s (props %v - type %s", v.ID, v.Label, v.Properties, v.Type)
}

func (e Edge) String() string {
	return fmt.Sprintf("%s (%s)-%s->%s (%s) - type %s", e.InVLabel, e.InV, e.Label, e.OutVLabel, e.OutV, e.Type)
}