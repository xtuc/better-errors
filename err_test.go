package bettererrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromString(t *testing.T) {
	chain := NewFromString("test")

	assert.IsType(t, chain, &Chain{})
	assert.Equal(t, "test", chain.Value.Error())
}

func TestNewFromErr(t *testing.T) {
	err := errors.New("test")
	chain := NewFromErr(err)

	assert.IsType(t, chain, &Chain{})
	assert.Equal(t, "test", chain.Value.Error())
}

func TestIsBetterError(t *testing.T) {
	assert.False(t, IsBetterError(errors.New("test")))
	assert.True(t, IsBetterError(NewFromString("test")))
}
