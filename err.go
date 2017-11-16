/*
  Better error visualization for client-side applications
*/
package bettererrors

import "errors"

// Create a new chain with the given error message
func New(str string) *Chain {
	return NewFromErr(errors.New(str))
}

// Convert an error into a chain element
func NewFromErr(err error) *Chain {
	return &Chain{
		Context: make(Context),
		Value:   err,
		Next:    nil,
	}
}

// Merge both chains
func (chain *Chain) With(err error) *Chain {
	var errChain *Chain

	if IsBetterError(err) {
		errChain = err.(*Chain)
	} else {
		errChain = NewFromErr(err)
	}

	newChain := chain
	newChain.Next = errChain

	return newChain
}

// Set context on the first element of the chain
func (chain *Chain) SetContext(key, value string) *Chain {
	chain.Context[key] = value
	return chain
}

// Convert error to string. Mainly used implicitly and provides the error interface.
func (chain *Chain) Error() string {
	// TODO(sven): use the flat printer here
	return chain.Value.Error()
}

// Check if the error is a chain
func IsBetterError(err error) bool {
	_, isBetterError := err.(*Chain)

	return isBetterError
}
