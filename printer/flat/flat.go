package flat

import (
	"fmt"

	be "github.com/xtuc/better-errors"
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
		buff += fmt.Sprintf("{%s: %s}", k, v)
	}

	if node.Next != nil {
		buff = chain(depth, buff, printChain(node.Next, depth))
	}

	return buff
}

func chain(depth int, parent, child string) string {
	return fmt.Sprintf("%s: %s", parent, child)
}
