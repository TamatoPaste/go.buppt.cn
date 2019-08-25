package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

//string 是数据类型，不是引用或指针类型
//string 是只读的byte slice，len函数可以计算它所包含的byte数
// string 的byte数组可以存放任何数组

var a string

func TestString(t *testing.T) {
	fmt.Println("a:", a)

	a = "hello"
	fmt.Println(len(a)) // 5

	//a[1] = 'a' 只读不可变byte slice，不能赋值

	a = "\xE4\xB8\xAD"
	fmt.Println("a：", a, len(a))
	// 此时a中只有一个字符“中”，但是len是3
	// 所以说字符串的len函数求出来的是字节数，而不是字符数

	b := []rune(a)
	fmt.Println("b：", b, len(b))
	fmt.Printf("中 Unicode %x\n", b[0]) // 4e2d unicode编码的十六进制
	fmt.Printf("中 UTF-8 %x\n", a)      // e4b8ad utf-8编码的十六进制
	/*
		汉字  中
		unicode   4e2d
		utf-8     e4b8ad
		string/[]byte   [\xE4,\xB8,\xAD]
	*/
}

func TestStringFunc(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",") // 字符串的分割

	for _, part := range parts {
		fmt.Println(part)
	}
	fmt.Println(strings.Join(parts, "-")) // 字符串的连接

	s1 := "1"

	if i, err := strconv.Atoi(s1); err == nil {
		fmt.Println(i)
	}
	//字符串转换函数基本都在strconv中
	// Atoi:字符串转数字，Itoa:整型转字符串

}
