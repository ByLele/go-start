package main

import (
	"fmt"
)

func main() {
	/* go语言关键字*/

	/*
	var
	import
	type
	fallthrough
	select
	interface
	defer
	func
	range
	append
	byte
	cap
	close
	complex
	copy*/
	var stockcode = 123
	var enddate = "2022"
	var url = "Code=%d&enddate=%s"
	var target_url = fmt.Sprintf(url,enddate,stockcode)
	fmt.Print(target_url)
	fmt.Print("\n")
	fmt.Print("abc")
	//go中常量
	/*const   常量名 常量类型
	const  常量名   = */
	const b string = "abc"
	const  b1 = "abc"


}
