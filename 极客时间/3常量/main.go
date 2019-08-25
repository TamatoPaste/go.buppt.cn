package main

// 刚在此文件打开终端的情况下，修改const文件夹名字失败，但是没有任何提示

// 常量值必须是编译期可以确定的字符,字符串,数字或布尔值
// 		或编译器能计算出结果的表达式，如unsafe.Sizeof,len,cap等
// 可写明数据类型，也可以类型推导
// 未使用的常量不会报错

// 常量不同于变量在运行期动态分配内存，常量会被编译器在预处理阶段直接展开，作为指令数据使用
// 数字常量不会分配存储空间，不能对其进行取址操作
import (
	"fmt"
	"unsafe"
)

func main() {
	// 定义常量
	const pi float32 = 3.1415926
	const pi2 = 3.14
	fmt.Println(pi)

	const ptrSize = unsafe.Sizeof(uintptr(0))

	// pi = 3 这句话会报错，常量不能被赋值

	// 多个常量也可以一起声明,被省略的常量等于之前最近非空常量右值(文本)
	const (
		e  = 2.7182
		n1 // =2.7182
		n2 = "abc"
		n3 // ="abc"
	)
	fmt.Println(e, n1, n2, n3)

	// 常量声明时，后续的赋值可以省略，变量批量声明可以不行
	// var (
	// 	z1 int = 1
	// 	z2 int        这么写相当于声明而没有初始化
	// 	z3
	// 	z4
	// 	z5
	// )
	//fmt.Println("z", z1, z2, z3, z4, z5)

	// iota
	// 是go的常量计数器，只能在常量表达式中使用,可以用于实现枚举类型，go没有定义enum
	// iota在const关键字出现时将被重置为0。
	// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行计数器)，无论该行有没有出现iota。
	// 被省略的常量的值等于之前最近赋值的值,a2 = iota,a3 = iota,a4 = iota
	const (
		a1 = iota //0
		a2        //1
		a3        //2
		a4        //3
	)
	fmt.Println("a", a1, a2, a3, a4)

	//几个常见的iota示例:
	// 1 使用 _ 跳过某些值
	const (
		b1 = iota //0
		b2        //1
		_
		b4 //3
	)
	fmt.Println("b:", b1, b2, b4)
	// 2 中间插队
	const (
		c1 = iota //0
		c2 = 100  //100
		c3 = iota // 2 不是1！！！！
		c4        //3
	)
	fmt.Println("c:", c1, c2, c3, c4)

	const d = iota
	fmt.Println("d:", d) // 0出了小括号，iota重新计数

	//3 多个iota定义在一行
	const (
		e1, e2 = iota + 1, iota + 2 // iota = 0;  1,2
		e3, e4                      // iota = 1; e3 = iota+1=2, e4 = iota+2 = 3
		e5, e6                      //3,4
	)
	fmt.Println("e:", e1, e2, e3, e6)

	// 定义了e4,e5，没有使用没报错，说明常量定义但是不使用是没有问题的，注意与变量比较。

	// iota用处示例
	// 1： iota可以快速设置连续值
	const (
		Monday = iota + 1
		Tuesday
		Wendensday
		Thursday
		Friday
		Saturday
		Sunday
	)

	// 2：iota快速设置位偏移量
	// 最低位为1表示开，倒数第二位为1表示关
	const (
		Open = 1 << iota
		Close
		Pending
	)

	// 常量不占内存，没有地址
	// fmt.Println(&Open)

}
