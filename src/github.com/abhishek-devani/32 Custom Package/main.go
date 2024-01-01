package main

import (
	print "custom_package/custom"
)

func main() {

	print.PrintStringValue("abc")

	print.PrintWithAnyArgs(123, 456)

	print.PrintAnyValue(42, "Hello, World!", 3.14, true)

}
