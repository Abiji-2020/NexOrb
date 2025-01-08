package test

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
    t.Run("Addition of two numbers", func(t *testing.T) {
        assert.Equal(t, 4, 2+2)
    })
}
