// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !compiler_bootstrap

package uint256

import (
	"errors"
)

var errOverflow error = errors.New("u256: integer overflow")
var errDivide error = errors.New("u256: integer divide by zero")
