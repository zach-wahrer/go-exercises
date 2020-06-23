package main

import "fmt"

func main() {
	integer := 10
	str := "this is a test"
	boolean := true

	fmt.Printf("integer is type %v\n", TypeTest(integer))
	fmt.Printf("str is type %v\n", TypeTest(str))
	fmt.Printf("boolean is type %v\n", TypeTest(boolean))

}

func TypeTest(x interface{}) string {
	if x == nil {
		return "Nil"
	} else if _, ok := x.(int); ok {
		return "Int"
	} else if _, ok := x.(string); ok {
		return "String"
	} else if _, ok := x.(bool); ok {
		return "Bool"
	} else {
		return "Unrecognized"
	}

}
