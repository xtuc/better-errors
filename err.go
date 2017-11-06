package bettererrors

import "errors"

func NewFromString(str string) *Chain {
	return NewFromErr(errors.New(str))
}

func NewFromErr(err error) *Chain {
	if IsBetterError(err) {
		panic("Assert error: argument must be an instance of error, bettererrors.Chain given")
	}

	return &Chain{
		Context: make(Context),
		Value:   err,
		Next:    nil,
	}
}

func (chain *Chain) With(err error) *Chain {
	if errChain, isErrChain := err.(*Chain); isErrChain {
		newChain := chain
		newChain.Next = errChain

		return newChain
	} else {
		panic("Assert error: err must be an instance of bettererrors.Chain")
	}
}

func (chain *Chain) SetContext(key, value string) *Chain {
	chain.Context[key] = value
	return chain
}

func (chain *Chain) Error() string {
	return "NOT IMPLEMENTED"
}

func IsBetterError(err error) bool {
	_, isBetterError := err.(*Chain)

	return isBetterError
}
