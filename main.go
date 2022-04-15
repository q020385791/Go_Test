package main

import (
	"fmt"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	//不定函數
	var StringSample = "this is a string" //Type is string
	var Varinferred = "This is inferred"  //type is inferred
	x :=
		3 //type is inferred

	// := 定義函數必須在functiotn裡面(區域函數)
	//var 定義則不用限制在function
	x = 18
	fmt.Println(StringSample)
	fmt.Println(Varinferred)
	fmt.Println(x)
	//若是使用 :=的話，必須給定初值

	//函數預設值
	var a string
	var b int
	var c bool
	fmt.Println(a) //a is ""
	fmt.Println(b) //b is 0
	fmt.Println(c) //c is false
}
