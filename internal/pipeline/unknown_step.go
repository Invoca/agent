package pipeline

import (
	"encoding/json"

	"github.com/buildkite/agent/v3/internal/ordered"
	"github.com/buildkite/interpolate"
)

// UnknownStep models any step we don't know how to represent in this version.
// When future step types are added, they should be parsed with more specific
// types. UnknownStep is present to allow older parsers to preserve newer
// pipelines.
type UnknownStep struct {
	Contents any
}

// MarshalJSON marshals the contents of the step.
func (u UnknownStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Contents)
}

// MarshalYAML returns the contents of the step.
func (u UnknownStep) MarshalYAML() (any, error) {
	return u.Contents, nil
}

// unmarshalMap unmarshals an unknown step from an ordered map.
func (u *UnknownStep) unmarshalMap(m *ordered.MapSA) error {
	u.Contents = m
	return nil
}

func (u *UnknownStep) interpolate(env interpolate.Env) error {
	c, err := interpolateAny(env, u.Contents)
	if err != nil {
		return err
	}
	u.Contents = c
	return nil
}

func (UnknownStep) stepTag() {}
