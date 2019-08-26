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
	"log"
	"math/rand"
	"runtime/debug"
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

// 5 匿名函数：匿名函数就是没有名字的函数
// 在使用上，匿名函数和普通函数完全相同，可以直接调用，保存到变量，作为参数或返回值
// 最大的区别是可以在函数内部定义匿名函数，形成类似嵌套效果
// 未曾使用的匿名函数会被编译器当作错误
func funcTest06() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

// 闭包(closure) 是在其词法上下文中引用了自由变量的函数，或者说是函数和环境的组合体。
func funcTest07(x int) func() {
	println(&x)
	return func() { //定义并返回了了一个没有参数，也没有返回值的匿名函数
		println(&x, x)
		println(x) // 那么这个x哪里来的？这个匿名函数凭什么能访问外层的x
	} //这种引用函数外环境中的变量的能力叫做闭包
}

func TestFunc04(t *testing.T) {
	a := funcTest07(5) //此处返回匿名函数赋值给a
	a()                //调用a函数，输出5，按道理，functest07运行完了
} // 中间定义的x也没了，后续再调用a，结果还是访问到了5
// 通过输出指针，发现闭包直接引用了原环境变量。
// 分析汇编代码，返回的不仅仅有匿名函数，还有所引用的环境变量的指针，本质上是一个funcval结构，再runtime/runtime2.go中定义
// 所以说，闭包是函数和引用环境的组合体更确切。

// 特性1：闭包引用环境变量，会导致其生命周期延长，甚至分配再堆上。
// 特性2：延迟求值
func funcTest08() []func() {
	var s []func()

	for i := 997; i < 1000; i++ {
		s = append(s, func() {
			println(&i, i) //3个闭包全部引用了同一个i，但是都未运行
		})
	}

	return s
}

func TestFunc05(t *testing.T) {
	for _, f := range funcTest08() {
		f() //一旦运行，就会发现，结果都一样
	}
} // 解决方法就是用不同的环境变量或传参复制
// 多个匿名函数引用同意环境变量会让事情变得很复杂，在并发模式下需要做同步处理
// 闭包使得不用传参就可以读或改环境状态，也要为此付出额外的代价
// 高性能要求场合，慎用闭包

// 6 延迟调用
// 语句defer向当前函数注册稍后执行的函数调用
// 总是在调用者的最后执行，类似于try-catch语句的finally语句，常用于资源释放、解除锁定、错误处理等
// 延迟调用注册的是调用，必须提供执行所需要的参数，参数在注册时会被复制并存起来，如果对环境敏感，可改用指针或闭包
func TestFunc06(t *testing.T) {
	x, y := 1, 2

	defer func(a int) {
		println("defer x, y = ", a, y) // y是环境变量引用，受环境变化影响
	}(x) //x是参数，直接复制

	x += 100
	y += 200
	println("x, y = ", x, y)
}

// 多个延迟注册按FILO次序执行
// 延迟调用可修改当前函数命名返回值，但其自身返回值被抛弃
// 误用：延迟调用在函数结束时才会执行，不合理的使用会浪费更多资源
// 性能：相比直接调用，延迟调用须花费更大代价，包括注册，调用，缓存，高性能场景下应避免延迟调用

// 7错误处理
// err

//panic() 和 recover() 都是内置函数，使用上更类似try/catch结构化异常
// panic会立刻中断当前函数流程，转而执行延迟调用
// 延迟调用中的recover可以捕获并返回panic提交的错误对象，recover也只能在延迟调用函数中使用
// 无论延迟调用中有没有recover，所有的延迟调用都会运行
// 		如果有，那就相当于被catch住了，啥事没有
// 		如果没有，所有的延迟调用执行完成后，panic沿调用堆栈向外传递，最终要么被捕获，要么程序崩溃

func funcTest09() {
	defer println("test.1")
	defer println("test.2")

	panic("哦嚯，GG")
}

func TestFunc07(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()

	funcTest09()
}

// 连续调用panic，只有最后一个会被recover捕获
// 下面的代码有个panic没接住，程序异常退出，我猜剩下的继续往外抛？
// func TestFunc08(t *testing.T) {
// 	defer func() {
// 		for {
// 			if err := recover(); err != nil {
// 				log.Println(err)
// 			} else {
// 				log.Fatalln(err)
// 			}
// 		}
// 	}()

// 	defer func() {
// 		panic("you are dead")
// 	}()

// 	panic("i am dead")
// }

// 延迟调用中panic，不会影响后续延迟调用
// recover之后panic，可以再次被捕获 ？？？
// recover只能在延迟调用函数中才能正常工作，你把recover当作延迟调用函数也是不行的
func funcTest10() {

	if err := recover(); err != nil { // 在延迟调用函数中使用 recover
		log.Println("问题不大，继续运行")
	}

}

func TestFunc09(t *testing.T) {
	defer funcTest10()           // recover在延迟函数中调用，才能捕捉到
	defer log.Println(recover()) //recover作为参数，捕捉不到
	defer recover()              // recover本身作为延迟调用函数，捕捉不到

	panic("i am dead ")

}

// 根据recover特性，如果要保护代码片段，只能将其重构为函数调用
// runtime/debug.PrintStack函数可以输出完整的堆栈调用信息
func funcTest11(x, y int) {
	z := 0

	func() {
		defer func() {
			if recover() != nil { // 把recover的所在的延迟调用函数设置成匿名函数
				debug.PrintStack()
				z = 0
			}
		}()

		z = x / y
	}()

	println("x / y = ", z)
}

func TestFunc10(t *testing.T) {
	funcTest11(5, 0)
}
