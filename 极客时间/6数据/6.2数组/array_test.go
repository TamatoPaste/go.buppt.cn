package data

import (
	"fmt"
	"testing"
)

// 定义数组类型时，数组长度必须时非负整形常量表达式

// 数组的长度是类型的组成部分，也就是元素类型相同，但长度不同的数组不属于同一类型
func TestArray01(t *testing.T) {
	// var a1 [3]int
	// var a2 [4]int

	// a1 = a2 cannot use a2 (type [4]int) as type [3]int in assignment
}

// 多种初始化方式
func TestArray02(t *testing.T) {
	var b1 [4]int                           // 元素自动初始化为0
	b2 := [4]int{2, 4}                      // 未提供初始化值的元素自动初始化为0
	b3 := [4]int{5, 3: 10}                  // 按索引初始化
	b4 := [...]int{1, 2, 3}                 // 编译器按初始化值数量确定数组长度
	b5 := [...]int{10, 5: 200, 3: 100, 999} // 支持索引初始化，数组长度与此相关

	fmt.Println(b1, b2, b3, b4, b5)
}

// 对于结构等符合类型，可省略元素初始化类型标签
type user struct {
	name string
	age  byte
}

func TestArray03(t *testing.T) {
	c := [...]user{
		{"Tom", 20},  // 省略了类型标签
		{"Mary", 18}, // 最后这个逗号不能省
	}

	fmt.Println(c)
}

// 定义多维数组时，仅第一维度允许使用 "..."
// 内置函数len和cap都返回第一维度长度
// 如果元素类型支持 == != 操作符，那么数组也支持此操作

// 指针数组是一个数组，只是里面的元素是指针
// 数组指针是指针，这个指针指向了一个数组，也就是数组变量的地址
func TestArray04(t *testing.T) {
	a, b, c := 10, 20, 30
	x := [...]*int{&a, &b} //指针数组
	y := &x                // 数组指针

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)

	fmt.Println(&x, &x[0], &x[1]) // 数组可以取任意元素地址
	y[0] = &c                     //数组指针可以直接操作元素，使用起来跟数组变量一样
}

// Go的数组是值类型，赋值，传参都会复制整个数组
// 下面例子可以看出， a, b, x 的地址都不同，说明发生了复制
func test(x [3]int) {
	fmt.Printf("%p, %v\n", &x, x)
}
func TestArray05(t *testing.T) {
	a := [3]int{11, 22, 33}
	b := a

	fmt.Printf("%p, %v\n", &a, a)
	fmt.Printf("%p, %v\n", &b, b)

	test(a)

}

// 如有需要，可改用指针或切片，避免数据复制
func test01(x *[3]int) {
	fmt.Printf("%p, %v\n", x, x)
}
func TestArray06(t *testing.T) {
	a := [3]int{5, 6, 7}

	fmt.Printf("%p, %v\n", &a, a)

	test01(&a)

}
