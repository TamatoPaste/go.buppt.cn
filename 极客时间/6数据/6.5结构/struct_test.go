package struct_go

import (
	"fmt"
	"testing"
	"unsafe"
)

// 结构体(strut)将多个不同类型命名字段(field)序列打包成一个符合类型
// 字段名必须唯一，可用 “_” 补位，支持使用自身指针类型成员。
// 字段名和排列顺序都是类型的组成部分。除对齐处理外，编译器不会优化、调整内存布局。

// 可以按顺序初始化所有字段，或使用命名方式初始化指定字段
// 推荐使用命名方式初始化，这样在结构体扩充或者调整顺序时，不用修改代码
func TestStruct01(t *testing.T) {
	type node struct {
		_    int // 注意：定义的时候不能写逗号，赋值的时候必须加逗号，最后一个也要加
		id   int
		next *node
	}

	a1 := node{
		id: 1, // 注意这里赋值用的是冒号，不是 :=，刚才查了好久没查出来错哪了
	}

	a2 := node{
		id:   2,
		next: &a1,
	}

	a3 := node{
		1,
		2,
		&a2,
	}

	//a4 := node{1, 2} // too few values in node literal

	fmt.Println(a1, a2, a3)
}

// 可以直接定义匿名结构变量
func TestStruct02(t *testing.T) {
	a := struct {
		name string
		age  int
	}{
		name: "Tom",
		age:  15,
	}

	fmt.Println("匿名结构体", a)
}

// 匿名结构体作为字段类型，由于缺少类型表示，作为字段类型时无法初始化
func TestStruct03(t *testing.T) {
	type file struct {
		filename string
		path     string
		attr     struct { // 匿名结构体作为字段类型
			owner       string
			accessLevel int
		}
	}

	a := file{
		filename: "xixi.jpg",
		path:     "D:\\go",
		// {   missing type in composite type缺少类型，问题时字段类型是匿名的，没有名字
		// 	"aDong",
		// 	5,
		// },
	}

	a.attr.owner = "aDong" // 匿名结构体作为字段类型时要这么赋值
	a.attr.accessLevel = 5

	fmt.Println(a)
}

// 只有在所有字段全部支持时，才可做相等操作
func TestStruct04(t *testing.T) {
	type data struct {
		x int
	}

	type complicate_data struct {
		x int
		y map[int]int
	}

	a1 := data{1}
	a2 := data{1}

	a3 := complicate_data{
		x: 1,
		y: map[int]int{
			1: 1,
			2: 2,
		},
	}
	a4 := complicate_data{
		x: 1,
		y: map[int]int{
			1: 1,
			2: 2,
		},
	}

	fmt.Println("a1 == a2 ? ", a1 == a2)
	//fmt.Println("a3 == a4 ?", a3 == a4) struct cataining map[int]int can not be compared
	fmt.Println(a3, a4)
}

// 可以使用指针直接操作结构字段，但不能是多级指针
func TestStruct05(t *testing.T) {
	type user struct {
		name string
		age  int
	}

	p := &user{
		name: "Tom",
		age:  15,
	}

	p.name = "Jerry"
	fmt.Println(p.name)

	p1 := &p
	//*p1.name = "Curry" error: p1.name undefined (type **user has no field or method name)
	fmt.Println(p1)
}

// 空结构( {} )
// 空结构是指没有字段的结构类型，无论是其自身，还是作为数组元素类型，其长度都为0
func TestStruct06(t *testing.T) {
	var a struct{}
	var b [100]struct{}

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b)) // 0 0 struct本身长度为0???
}

// 尽管没有分配数组内存，但仍然可以操作元素，对应切片len,cap属性也正常
func TestStruct07(t *testing.T) {
	var a [100]struct{}
	b := a[:]

	a[1] = struct{}{}
	b[2] = struct{}{}

	fmt.Println(b[3], len(b), cap(b))

	// 实际上，这类长度为0的对象通常指向了runtime.zerobase
	// 讲道理，b应该是一个切片，底层数组其实指向了zerobase
	c := [0]int{}
	fmt.Printf("%p, %p, %p\n", &a[0], &b[0], &c)
}

// 空结构可以作为通道元素类型，用于事件通知，不是很能理解，后续再看看，P121

// 匿名字段
// 所谓匿名字段(anonymous field)，指没有名字，仅有类型的字段，也叫嵌入字段，或嵌入类型
// 搞不懂，以后再看！！！

// 字段标签

// 内存布局
// 不管结构体包含多少字段，其内存总是一次性分配的，各字段在相邻的地址空间按定义顺序排列
// 当然，对应引用类型，字符串和指针，结构内存中只包含其基本的(头部)数据，还有匿名字段成员也被包含在内

// 借助unsafe包中的相关函数，可以输出输出所有字段的便宜量和长度
func TestStruct08(t *testing.T) {
	type point struct {
		x, y int
	}

	type value struct {
		id    int    // 基本类型
		name  string // 字符串
		data  []byte // 引用类型
		next  *value // 指针类型
		point        // 匿名字段
	}

	v := value{
		id:    1,
		name:  "test",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 200},
	}

	// 定义rune类型，用的不是单引号，而是反单引号
	s := `v: %p - %x, size: %d  align: %d'
		
		field    address         offset      size
		------+---------------+----------+--------
		id       %p              %d          %d
		name     %p              %d          %d
		data     %p              %d          %d
		next     %p              %d          %d
		x        %p              %d          %d
		y        %p              %d          %d
		`

	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.y, unsafe.Offsetof(v.y), unsafe.Sizeof(v.y))
}

// 如果仅仅只有一个空结构字段，那么同样按1对齐，只不过长度为0，指向runtime.zerobase变量
func TestStruct09(t *testing.T) {
	v := struct {
		a struct{}
	}{}

	fmt.Printf("%p, %d, %d\n", &v, unsafe.Sizeof(v), unsafe.Alignof(v))
}

// 对齐还受硬件平台和访问效率有关，有些平台只能访问特定地址，比如偶数地址
// 另一方面，CPU访问自然对齐的数据所需要的读周期最短，还可以避免拼接数据
