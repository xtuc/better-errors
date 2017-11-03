package tree

import (
	"fmt"

	be "github.com/xtuc/better-errors"
)

const (
	PADDING = "  "
)

func PrintChain(root *be.Chain) string {
	return printChain(root, 0) + "\n"
}

func printChain(node *be.Chain, depth int) string {
	var buff string

	buff += node.Value.Error()
	depth++

	for k, v := range node.Context {
		buff += "\n"
		buff += fmt.Sprintf("%s├ %s: %s", getPadding(depth), k, v)
	}

	if node.Next != nil {
		buff = chain(depth, buff, printChain(node.Next, depth))
	}

	return buff
}

func chain(depth int, parent, child string) string {
	padding := getPadding(depth)

	return fmt.Sprintf("%s\n%s└ %s", parent, padding, child)
}

func getPadding(mx int) (s string) {
	for i := 0; i < mx; i++ {
		s += PADDING
	}

	return s
}
