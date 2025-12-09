package datatypes

import "fmt"

// | Type       | Min                        | Max                        |
// | ---------- | -------------------------- | -------------------------- |
// | **int8**   | -128                       | 127                        |
// | **uint8**  | 0                          | 255                        |
// | **int16**  | -32,768                    | 32,767                     |
// | **uint16** | 0                          | 65,535                     |
// | **int32**  | -2,147,483,648             | 2,147,483,647              |
// | **uint32** | 0                          | 4,294,967,295              |
// | **int64**  | -9,223,372,036,854,775,808`` | 9,223,372,036,854,775,807  |
// | **uint64** | 0                          | 18,446,744,073,709,551,615 |
// | **int**    | depends on system          | depends on system          |
// | **uint**   | depends on system          | depends on system          |

func BasicInteger() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}

func IntegerTypeConversion() {
	var a int = 65
	var b float32 = float32(a)
	var c string = string(a) // 65 -> 'A' when converted to a rune and then string
	var d uint = uint(a)

	fmt.Println("int a:", a)
	fmt.Println("float32 b:", b)
	fmt.Println("string c (from rune):", c)
	fmt.Println("uint d:", d)

}

// UnsignedIntegerDemo demonstrates unsigned integer types and narrowing conversions
func UnsignedIntegerDemo() {
	// Unsigned types
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	fmt.Println("uint    :", u)
	fmt.Println("uint8   :", u8)
	fmt.Println("uint16  :", u16)
	fmt.Println("uint32  :", u32)
	fmt.Println("uint64  :", u64)

	// Narrowing conversion (higher -> lower) truncates modulo 2^n
	var narrowed uint8 = uint8(u16) // 65535 -> 255
	fmt.Println("narrowed uint16->uint8:", narrowed)

	// Converting signed to unsigned preserves bit pattern for positive values
	var s int = -1
	var su8 uint8 = uint8(s) // wraps around to 255
	fmt.Println("int(-1) -> uint8:", su8)
}
