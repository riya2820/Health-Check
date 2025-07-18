
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// To print the type of a value, use `%T`.
	fmt.Printf("type: %T\n", p)

	// Formatting booleans 
	fmt.Printf("bool: %t\n", true)


	// `%d` for standard, base-10 formatting
	fmt.Printf("int: %d\n", 123)

	// prints a binary representation
	fmt.Printf("bin: %b\n", 14)

	// prints the character corresponding to the given integer
	fmt.Printf("char: %c\n", 33)

	// `%x` provides hex encoding.
	fmt.Printf("hex: %x\n", 456)

	// basic decimal formatting use `%f`.
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// basic string printing use `%s`.
	fmt.Printf("str1: %s\n", "\"string\"")

	// double-quote strings as in Go source, use `%q`.
	fmt.Printf("str2: %q\n", "\"string\"")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.
	fmt.Printf("str3: %x\n", "hex this")

	// prints a representation of a pointer, use `%p`.
	fmt.Printf("pointer: %p\n", &p)

	// `os.Stdout` using `Fprintf`.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
