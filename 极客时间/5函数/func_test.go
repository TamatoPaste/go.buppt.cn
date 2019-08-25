package func_test

/*
	Go 中函数有些不太方便的限制，也借鉴了动态语言的某些有点：
		1 无需前置申明
		2 不支持命名嵌套定义(nested)
		3 不支持同名函数重载(overload)
		4 不支持默认参数
		5 支持不定长变参
		6 支持多返回值
		7 支持命名返回值
		8 支持匿名函数和闭包

	函数特点：
		有多个返回值
		所有参数都是值传递
		函数可以作为变量的值
		函数可以作为参数和返回值

	函数的左花括号不能另起一行
	函数属于第一类对象(first-class object)
	第一类对象指其可以在运行期创建，可以用作函数参数或返回值
		可以存入变量的实体，最常用的方法就是匿名函数
	具有相同前面(参数及返回值列表)的视作同一类型
	函数只能判断其是否为nil，不支持其他比较
*/

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func TestFunc(t *testing.T) {
	//fmt.Println(multiReturns())
	multiParameters()
	multiParameters(1, 10, 100, 1000, 10000)
	multiParameters(12, 342, 5, 2, 5, 25, 2, 254252, 2525, 2, 52, 52, 252)
}

func multiReturns() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 2 参数
// go对参数的处理偏保守，不支持有默认值的可选参数，不支持命名实参
// 调用时，必须按签名顺序传递指定类型和数量的实参，_也不能忽略

// 参数列表中，相邻的同类型参数可合并
func funcTest01(x, y int, s string) {

}

// 参数可视为函数的局部变量，不能在同层次定义同名变量
func funcTest02(x, y int) int {
	// x := 100 no new variabes on the left of :=
	// var y int  y redeclared in this block
	return x + y
}

// 3 变参
// 变参本质上就是一个切片，可以接受0个到多个同类型参数，且必须放在列表尾部
// 当有其他类型的参数，必须放在前面
// 测试发现，可以参数的个数可以为0
func multiParameters(paras ...int) {
	ret := 0
	for _, i := range paras {
		ret += i
	}
	fmt.Println(ret)
}

// 将切片作为参数是，须进行展开操作
// 如果是数组，先将其转换为切片
func funcTest03(a ...int) {
	for i := range a {
		a[i] += 10
	}

}

func TestFunc02(t *testing.T) {
	a := [3]int{20, 30, 40}
	//funcTest03(a) //调用失败，can not use type[3]int as type int in argument to funcTest03
	//funcTest03(a[:]) 调用失败，也是类型不对
	funcTest03(a[:]...)
	fmt.Println(a)
}

// 参数复制的仅是切片自身，实参和形参共享底层数组，因此可以修改原数据
// 因此如果有需要，可以用内置函数copy复制底层数据

// 4返回值
// 有返回值的函数，必须有明确的return终止语句
// 可以有多个返回值
// 命名返回值，优点是能让人望文生义，知道返回值得意思，让函数申明更加清晰
// 可以该晒帮助文档和代码提示。且和参数一样，可当做函数局部变量使用，有return隐式返回
func funcTest04(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("division by zero") // 直接赋值，已经声明了
		return                               // 直接返回，不用写z，err 这里的z为默认0值
	}
	z = x / y
	return
}

// 命名返回值作为一种特殊的局部变量，可以被“更局部”的局部变量遮蔽
// 编译器能检查到 此类状况
func funcTest05(x, y int) (z int, err error) {
	{
		z := x + y
		// return 报错：z is shadowed during return
		return z, nil // 写全了就没事了
	}
}

func TestFunc03(t *testing.T) {
	fmt.Println(funcTest04(1, 2)) // 0 nil
}

// defer 函数
// defer 函数总是在调用者的最后执行，类似于try-catch语句的finally语句
func TestDeferFunc(t *testing.T) {
	defer multiParameters(1, 2, 4, 5, 6) //写在前面，后执行
	fmt.Println("函数执行完啦！！！")             // 写在后面，先运行
	//panic("GG,程序出现异常")                   // 抛了个运行时异常，defer代码还是执行了
	// fmt.Print(1) 代码不可达
}
