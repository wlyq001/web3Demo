package main

import (
	"fmt"
)

func main() {
	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf
	// 八进制
	var c uint8 = 017
	var d uint8 = 0017
	var e uint8 = 0o17
	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111
	// 十进制
	var h uint8 = 15

	fmt.Println(a == b)
	fmt.Println(b == c)
	fmt.Println(c == d)
	fmt.Println(d == e)
	fmt.Println(e == f)
	fmt.Println(f == g)
	fmt.Println(g == h)

	// 浮点数
	var float1 float32 = 10.0
	float2 := 10.0
	float2 = float64(float1)
	fmt.Println(float2)
	//fmt.Println(float1 == float2)

	// 复数
	var c1 complex64 = 1 + 0.1i
	c2 := 1 + 0.1i
	c3 := 1 + 0.1i
	fmt.Println(c1 == complex64(c2))
	fmt.Println(complex128(c1) == c2)
	fmt.Println(c3 == c2)

	x := real(c1)
	y := imag(c1)
	fmt.Println(x, y)

	// 字符串&byte&rune
	var s string = "hello world,坤"
	var bytes = []byte(s)
	fmt.Println("Convert \"hello world\" to bytes :", bytes)
	var runes = []rune(s)
	fmt.Println("Convert \"hello world\" to runes :", runes)
	fmt.Println("string length:", len(s))
	fmt.Println("bytes length:", len(bytes))
	fmt.Println("rune length:", len(runes))
	fmt.Println(s[0:12])
	fmt.Println(string(bytes[0:12]))
	fmt.Println(string(runes[0:12]))

	// bool
	var success bool
	fmt.Println(success)
}
