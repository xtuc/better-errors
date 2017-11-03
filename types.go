package bettererrors

type Context map[string]string

type Chain struct {
	Context Context
	Value   Node
	Next    *Chain
}

type Node error
