package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
  assert.Equal(t, 123, 123, "they should be equal")
  assert.NotEqual(t, 123, 456, "they should not be equal")
  assert.Nil(t, object)
  if assert.NotNil(t, object) {
    assert.Equal(t, "Something", object.Value)
  }
}
