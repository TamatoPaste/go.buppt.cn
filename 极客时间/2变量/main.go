package main

/*
	go的变量赋值，跟C++，Java等强类型语言相比，有2个特点
	1 ：类型推导
	2 ：一条赋值语句对多个变量同时赋值

	运行时内存分配操作会确保变量自动化初始化为二进制零值(zero value)
	避免出现不可预测行为。

	标准申明：
	单个变量，或多个同类型变量，只声明，数据类型放最后
	多个不同类型的变量，只声明，不能用批量声明方式。

	单个变量声明，带初始化，不要带数据类型，让他自动推断，不然会有warning
	多个同类型变量声明带初始化, 带不带都一样，没有warning，也没有error
	多个不同类型变量声明带初始化，带了数据类型就报error，不带就OK

	组申明：总结就是，不初始化，数据类型放最后，初始化，就不带数据类型
	单个变量用组申明了 只声明，数据类型放最后
					带初始化，带数据类型就warning
	多个同类型变量申明，只申明，数据类型也是放最后
					到初始化，带类型就warning
	多个不同类型变量申明，只申明，数据类型放最后
					带初始化，带类型就warning

*/
import "fmt"

func main() {
	// 1 变量声明方式一之标准声明：var 变量名  变量格式
	var a = "2"
	fmt.Println(a)
	// var a string = "a" 会有warning，提示应当省略数据类型，go会从右值自动推断类型

	// 多个同类型的变量声明的缩写形式
	var aa, ab, ac, ad = 1, 2, 3, 4
	fmt.Println(aa, ab, ac, ad)

	// var ax int, ay string, az bool 报错，不同类型的变量声明不能缩写
	//多个不同类型的变量声明的同时初始化就行，这种情况不能写类型
	var ax, ay, az = 0, "a", true
	//var ax int, ay string, az bool = 0, "a", true 报错
	fmt.Println(ax, ay, az)

	// 变量声明方式二之组声明：
	// var (
	// 		 	变量名  变量格式
	// 			变量名  变量格式
	// 			变量名  变量格式
	// 			......
	//  )
	var (
		b string
		c int
		d bool
		e bool
	)
	fmt.Println(b, c, d, e)

	// 变量声明的同时还可以赋值,即初始化
	// 2 变量声明方式三之简短模式(short variable declaration)
	// 限制：1 必须显示初始化 2 不能提供数据类型 3 只能写在函数内部(局部变量)
	f := true
	fmt.Println(f)
	// f = 1 报错，go是强类型语言，虽然声明时没有指明f是什么类型
	// 但是初始化操作会确定f的类型，这就是go的类型推导

	// 简短模式并不是总是重新定义变量,有可能会退化为赋值
	// 上面了定义了一个f，下面这行代码不会出错，不会出现重复定义的问题
	f, g := false, "a"
	// 退化条件为：最少有一个同作用于的新变量被定义
	//下面这条语句的错误是 no new variables on left side of :=，还有一个类型不匹配错误

	// 初始化退化为赋值，使得我们可以用err一个变量去重复接受函数的返回值！！！
	// 匿名变量 _ ,不占命名空间，不用分配内存，所以也不存在重复命名问题
	//g, _ := duck()
	fmt.Println(g)

	// 3多变量赋值
	h, i := 1, 4
	fmt.Println(h, i)

	// 不同类型的变量可以同时赋值
	j, k, l, m := 1, "k", true, 32.3
	fmt.Println(j, k, l, m)

	// o, p, q := 1 多个变量赋同一个值不能这么写，报错
	// 既然是赋值，当然可以把变量的值赋给变量
	o, p, q := j, k, l
	fmt.Println(o, p, q)

	//go中交换2个变量的值得简单写法,省略了中间变量
	h, i = i, h
	fmt.Println(h, i)

	// 4 未使用错误，编译器会将从未使用过的局部变量当做错误，全局变量不会

}

func duck() (int, int) {
	return 0, 1
}
