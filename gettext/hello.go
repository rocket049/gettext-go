// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/rocket049/gettext-go/gettext"
)

func main() {
	gettext.SetLocale("zh_CN")
	gettext.BindTextdomain("hello", "../examples/local", nil)
	gettext.Textdomain("hello")

	fmt.Println(gettext.Gettext("Hello, world!"))
	// Output: 你好, 世界!
}
