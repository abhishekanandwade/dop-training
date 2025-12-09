package main

import (
	"fmt"
	"strconv"
)

func main() {
	// datatypes.ExampleBasic()
	// datatypes.ExampleUnsignedIntegers()

	var a int = 10
	b := 20

	var fl float32 = 3.14
	var s string = "error sim"
	var boolean bool
	var trueBool bool = true

	var intVal int = 100

	var flInt int = int(fl)
	sInt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	fmt.Println("String to int:", sInt)

	strInt := strconv.Itoa(intVal)
	ftmtInt := strconv.FormatInt(int64(intVal), 10)
	fmt.Println("Int to string:", strInt)
	fmt.Println("FormatInt:", ftmtInt)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(fl)
	fmt.Println(s)
	fmt.Println(boolean)
	fmt.Println(trueBool)
	fmt.Println(flInt)

	var int16Val int16 = 32767
	var int8Val int8 = 20
	var int32Val int32 = 2147483647
	var int64Val int64 = 2147483647
	var fl64 float64 = 3.14159265359
	var sumFlInt float64 = fl64 + float64(intVal)
	var sumIntFl int = int(int64Val) + int(fl64)

	fmt.Println("Sum float64 + int:", sumFlInt)
	fmt.Println("Sum int + float64:", sumIntFl)

	ExampleBasic(int32(int16Val))
	ExampleBasic(int32(int8Val))
	ExampleBasic(int32Val)
	ExampleBasic(int32(int64Val))

	res := operators(int64Val, int64(int32Val))
	if res {
		fmt.Println("int64Val is greater than or equal to int32Val")
	} else {
		fmt.Println("int64Val is not greater than int32Val")
	}

	switchExample(3)
	switchExampleString("6")

	// fmt.Println(fl64)
}

func ExampleBasic(a int32) {
	fmt.Println("Integer: ", a)
}

func operators(a, b int64) bool {
	// a>=b
	if a > b || a == b {
		return true
	}
	return false
	//return a>b
}

func switchExample(val int) {
	switch val {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	default:
		fmt.Println("Value is something else")
	}
}
func switchExampleString(val string) {
	switch val {
	case "1":
		fmt.Println("Value is 1")
	case "2":
		fmt.Println("Value is 2")
	default:
		fmt.Println("Value is something else")
	}
}
