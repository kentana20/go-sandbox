// +build ignore
package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (retx, rety int) {
	// named return values
	retx = sum * 4 / 9
	rety = sum - retx

	// naked return
	return
}

// var c, python, java bool
var i, j int = 1, 2 // 初期化子

func main() {
	//fmt.Printf("now your have %g problems.", math.Sqrt(7))
	//fmt.Println(math.Pi)
	// fmt.Println(add(1, 20))
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// fmt.Println(split(17)) // # 7 10
	// var i int
	// fmt.Println(i, c, python, java) // 0 false false false
	var c, python, java = true, false, "no!" // 型定義の省略
	fmt.Println(i, j, c, python, java)

	k := 3 // short variable declatrations (関数の中でのみ宣言可能)
	fmt.Println(i, j, k, c, python, java)

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(x*x + y*y) // error
	// var z uint = f // error
	// fmt.Println(x, y, "hoge")

	v := 42
	fmt.Printf("v is of type %T", v) // 型推論
}
