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
	switch x.(type) {
	case nil:
		return "Nil"
	case int:
		return "Int"
	case string:
		return "String"
	case bool:
		return "Bool"
	default:
		return "Unrecognized"
	}
}
