package main

import "fmt"

func main() {
	var a int = 10
	var b = 20 // 타입 추론
	c := 30    // 짧은 선언 (함수 내부에서만)

	fmt.Println(a, b, c, b)
	fmt.Println("hello world")
}
