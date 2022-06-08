package main

import (
	"fmt"
	"strings"
)

func main () {
	//字符串

	var str string = "this is a simple language ."
	fmt.Println(strings.Replace(str,"this","THIS",4))

	//fmt.Printf("\t %s\n",str)

	/*
	fmt.Printf("%d\n",strings.HasPrefix("th",str))
	fmt.Printf("%d\n",strings.HasSuffix("age",str))
	//包含
	fmt.Printf("%d\n",strings.Contains("is",str))
	//子字符串出现位置
	fmt.Printf("%d\n",strings.Index("is",str))

	//子字符串最后出现位置
	fmt.Printf("%d\n",strings.LastIndex("guage",str))
	//非ascall码
	//strings.IndexRune("safas",str)//str rune

	 */



	/*/测试环境安装成功
	fmt.Printf("hello world!")

	 */
}