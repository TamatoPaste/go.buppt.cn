package main

/*
	go的变量赋值，跟C++，Java等强类型语言相比，有2个特点
	1 ：类型推导
	2 ：一条赋值语句对多个变量同时赋值
*/
import "fmt"

func main() {
	// 变量声明方式一之标准声明：var 变量名  变量格式
	var a string
	var aa, ab, ac, ad int // 多个同类型的变量声明的缩写形式
	fmt.Println(a)
	fmt.Println(aa, ab, ac, ad)

	// 变量声明方式二之批量声明：
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
	// 变量声明方式三之短变量声明，只能在函数内部，且必须赋值
	f := true
	fmt.Println(f)
	// f = 1 报错，go是强类型语言，虽然声明时没有指明f是什么类型
	// 但是初始化操作会确定f的类型，这就是go的类型推导

	// 匿名变量 _ ,不占命名空间，不用分配内存，所以也不存在重复命名问题
	g, _ := duck()
	fmt.Println(g)

	// 变量的赋值可以同时给多个变量赋值
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

}

func duck() (int, int) {
	return 0, 1
}
