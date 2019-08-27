package slice

import (
	"fmt"
	"testing"
)

//切片本身不是动态数组或数组指针,是一个只读对象，类似于数组指针的一种封装
//它内部通过指针指向底层数组，设定相关属性将读写操作限定在指定区域内
/*
	type slice struct {
		array unsafe.Pointer
		len int
		cap int
	}
*/

// 切片创建方法之一：通过已有的数组或数组指针创建
// 以开始和结束索引位置确定所引用的数组片段，不支持反向索引，开始结束位置左闭右开
// 切片也支持按索引访问，索引从0开始，不是底层数组真实索引
func TestSlice01(t *testing.T) {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := x[:]
	s2 := x[2:5]
	s3 := x[2:5:7]
	s4 := x[4:]
	s5 := x[:4]
	s6 := x[:4:6]

	fmt.Printf("s1:%v, len:%d, cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v, len:%d, cap:%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3:%v, len:%d, cap:%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4:%v, len:%d, cap:%d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5:%v, len:%d, cap:%d\n", s5, len(s5), cap(s5))
	fmt.Printf("s6:%v, len:%d, cap:%d\n", s6, len(s6), cap(s6))
}

// 切片创建方法之二: make函数或显示初始化
func TestSlice02(t *testing.T) {
	s1 := make([]int, 3, 5)    // 指定 len 和 cap，底层数组会被初始化为零值
	s2 := make([]int, 3)       // 省略cap，len 和 cap 相等
	s3 := []int{10, 20, 5: 30} // 按初始化元素分配底层数组，并设置 len，cap

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}

func TestSlice03(t *testing.T) {
	var a []int  // 仅申明，未分配内存，也就没初始化一说
	b := []int{} // 申明且分配了内存，并完成了初始化,内部指针被赋值了，虽然它指向了runtime.zerobase

	println(a == nil, b == nil)
	println(len(a), cap(a), len(b), cap(b))
}

// 切片之间不支持比较操作，不管内部元素类型支持不支持，仅能判断是否为nil
// 可以对nil切片执行 slice[:] 操作，返回的也是 nil
