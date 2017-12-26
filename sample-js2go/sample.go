package main

import (
	. "github.com/jingweno/godzilla/runtime"
)

func main() {
	global := NewDefaultContext().Global
	_ = global

	// line 1: console.log("hello")
	Console_Log([]Object{JSString("hello")})
}
