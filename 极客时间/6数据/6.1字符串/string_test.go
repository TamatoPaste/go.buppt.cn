package data

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

//string 是数据类型，是一种复合结构，不是引用或指针类型
type stringStruct struct {
	str unsafe.Pointer
	len int
}

// 头部指针指向字节数组，但没有NULL结尾，默认编码utf-8，字面常量允许十六进制、八进制和utf编码

//string 是只读的byte slice，len函数可以计算它所包含的byte数,不是字符数，cap不接受字符串参数
// string 的byte数组可以存放任何数组

var a string

func TestString(t *testing.T) {
	// 字符串的默认是值是 ""，不是 nil
	fmt.Println("a: ", a)

	//允许十六进制、八进制和utf编码
	a = "hello\x61\142\u0041"
	fmt.Println("a: ", a)

	// 8,所有说返回的是byte数，不是字符数
	fmt.Println(len(a))

	// `` 定义不做转义处理的原始字符串(raw string) tab上那个键
	b := `line\n
	line2`
	fmt.Println(b)

	// 支持 ！= == < > + +=操作符
	fmt.Println(a != b)
	fmt.Println(a == b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a + b)

	// 允许以索引号访问字节数组(非字符)，但不能修改，也不能取元素地址
	fmt.Println(a[1])
	//a[1] = 'a' //只读不可变byte slice，不能赋值 cannot assign to a[1]
	//fmt.Println(&a[1]) cannot take address of a[1]

	//以切片语法返回子串时，其内部依旧指向原字节数组
	c := "abcdefghijk"
	c1, c2, c3 := c[:3], c[2:5], c[3:]
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&c)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&c1)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&c2)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&c3)))

	// 使用for循环遍历字符串时，分byte和rune两种方式
	// byte字节访问，rune字符访问
	d := "hello, 中国"

	for i := 0; i < len(d); i++ {
		fmt.Printf(" %d: [%c] ", i, d[i])
	}
	print("\n")
	for i, d1 := range d {
		fmt.Printf(" %d: [%c] ", i, d1)
	}

	//2 转换，要修改字符串，须将其转换为可变类型([]rune或[]byte)
	//不管如何转换，都须重新分配内存，并复制数据

	// 转换操作会拖累性能，可尝试用“非安全”方法改善
	// 此方法利用了[]byte和string头结构“部分相同”，将指针类型转换来实现类型变更，避免了底层数组复制
	fmt.Println()
	be := []byte(d)
	e := *(*string)(unsafe.Pointer(&be))
	fmt.Printf("be: %x\n", &be)
	fmt.Printf("e: %x\n", &e)

	// append函数，可将string直接追加到[]byte内
	var bf []byte
	//bf = append(bf, "hello, jack!") //报错，can not use string as byte in append加三点就行了
	bf = append(bf, "hello, jack!"...)
	fmt.Println("bf: ", bf)

	// 字符串是只读的，转换时要新分配内存和复制数据是可以理解的，但性能也很重要
	// 编译器会在特定场景优化，避免额外分配和数据复制
	// 1 将 []byte 转换为 string key， 去 map[string] 查询的时候
	// 2 将 string 转换成 []byte，进行 for range 迭代的时候，直接取字节赋值给局部变量

	// 3 性能
	// 用加法拼接字符串时，每次都重新分配内存，在构建超大字符串时，性能极差
	// 常用 strings.Join()方法，它会统计参数长度，只分配一次内存。
	// 编译器对 s1+s2+s3 这类表达式的处理方式和 strings.Join一样
	// bytes.Buffer 也能完成类似操作，且性能相当

	// 4 Unicode
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
