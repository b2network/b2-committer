package errcode

import "github.com/pkg/errors"

var (
	ErrNoBlobFoundInBlock = errors.New("no blobs found in this block ")
	ErrNoBlobFound        = errors.New("no blobs found in this block ")
)
