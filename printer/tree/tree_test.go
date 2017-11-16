package tree

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	bettererrors "github.com/xtuc/better-errors"
)

func TestPrintChainWithContext(t *testing.T) {
	err1 := errors.New("bar")
	err := bettererrors.
		New("foo").
		SetContext("a", "b").
		With(bettererrors.NewFromErr(err1))

	actual := PrintChain(err)

	expected := `foo
  ├ a: b
  └ bar
`

	assert.Equal(t, expected, actual)
}
