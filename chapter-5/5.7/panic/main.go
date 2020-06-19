// Using panic/recover to write a function that contains no return statement
// yet returns a non-zero value
package main

import "fmt"

func main() {
	err := panicAndRecover()
	fmt.Println(err)
}

func panicAndRecover() (err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
		default:
			err = fmt.Errorf("return without statement")
		}

	}()
	panic("test")
}
