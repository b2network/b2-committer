package errcode

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestErr(t *testing.T) {
	// 测试
	err := ErrNoBlobFoundInBlock
	fmt.Println(errors.Is(err, ErrNoBlobFoundInBlock))
	fmt.Println(errors.Is(err, ErrNoBlobFound))
}
