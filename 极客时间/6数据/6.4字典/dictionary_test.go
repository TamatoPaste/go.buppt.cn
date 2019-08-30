package dictionary_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 作为无序键值对集合，字典要求key必须是支持相等运算符(==, !=)的数据类型
// 字典是引用类型，使用make函数或初始化表达式语句创建

func TestDic01(t *testing.T) {
	a := make(map[string]int) // 使用make函数创建
	a["a"] = 1
	a["b"] = 2

	b := map[int]struct { // 使用初始化表达语句创建，这里的value是一个匿名结构体
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	fmt.Println(b)

	c := map[int]int{
		1: 1,
		2: 2,
	}

	c[1] = 2 //修改
	c[3] = 3 //新增

	if v, ok := c[1]; ok { //ok-idiom模式，判断某个键是否存在，且返回其对应的值
		println(v)
	}

	fmt.Println(c[101]) // 也可以不先判断，直接访问某个键的值，当键不存在时，不会报错
	// 所以不判断直接访问，你不知道这个键对应的值就是零值，还是这个键不存在

	delete(c, 1) //delete(dic,key)  删除字典中的指定key
	fmt.Println(c)

	delete(c, 1) //如果key不存在，不会报错

	for k, v := range c { // 对字典进行循环
		println("k:", k, "  v:", v)
	}

	println(len(c)) // len函数返回字典中键值对的数量，一对算一个
	//println(cap(c))  cap函数不能接受字典类型参数，过不了编译器
}

// 字典被设计为 not addressable ，不能直接访问vlue中的成员(当value是结构体或数组时)
// 可以整体替换value，或使用指针类型的value值，再利用指针去修改结构体或数组
func TestDic02(t *testing.T) {
	type user struct {
		name string
		age  int
	}

	a := map[int]user{
		1: {"Tom", 18},
	}
	//a[1].age = a[1].age + 1 can not assign to field a[1].age in map
	//a[1] = user{a[1].name, a[1].age + 1} // 对value进行整体替换
	//u := a[1] 这样也行，也是对value整体替换
	//u.age += 2
	//a[1] = u
	fmt.Println(a)

	b := map[int]*user{}
	b[1] = &user{
		name: "Jerry",
		age:  3,
	}

	b[1].age += 1
	fmt.Println(b[1])
}

// nil map可以读，不能写，这个有点奇葩了
func TestDic03(t *testing.T) {
	var a map[int]int // 只声明了一个dic，没有初始化，这个时候时nil字典
	println(a[1])     //nil 字典可以读，全都返回零值
	//a[1] = 999  nil 字典不能写，panic: assignment to entry in nil map
}

// nil map 和空 map 是不一样的
func TestDic04(t *testing.T) {
	var a map[int]int  //未分配内存空间
	b := map[int]int{} //已分配内存空间且初始化完成，等同于make函数

	println(a == nil, b == nil)
}

// 安全
// 在迭代期间删除或新增键值对是安全的
func TestDic05(t *testing.T) {
	var a = make(map[int]int)

	for i := 0; i < 10; i++ {
		a[i] = i + 10
	}

	for k := range a {
		if k == 5 {
			a[6] = 1000
		}
		delete(a, k)
		fmt.Println(k, a)
		// 如果6出现在5前面，，赋值语句会变成新增语句，最后会有一个[6：1000]留下
		// 如果5出现在6面前，先赋值再删除，最后留下的是一个空dic
		// 所以说，迭代过程如果删除，不能保证新增的会被删除
	}

}

// 运行时会对字典并发操作做出检测，如果某个任务正在对字典进行写操作，那么其他任何任务都不能对该字典
// 执行并发操作(读，写，删除)，否则会导致进程崩溃
// func TestDic06(t *testing.T) {
// 	a := make(map[int]int)

// 	go func() {
// 		for {
// 			a[1] += 1
// 			time.Sleep(time.Microsecond)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			_ = a[2]
// 			time.Sleep(time.Microsecond)
// 		}
// 	}()

// 	select {} //fatal error: concurrent map read and map write
// }

// 可用 sync.RWMutex 实现同步，避免读写同时操作
func TestDic07(t *testing.T) {
	var lock sync.RWMutex
	a := make(map[int]int)

	go func() {
		for {
			lock.Lock()
			a[1] += 1
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = a[2]
			lock.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()

	select {} //fatal error: concurrent map read and map write
}

// 性能
// 字典对象本身就是指针包装，传参时无需再取地址
// 创建时准备足够大空间有利于提升性能，因为减少了内存扩张和再哈希过程。
// make(map[int]int) make(map[int]int,1000) 在数量规模在1000左右时，后者性能比前者搞
