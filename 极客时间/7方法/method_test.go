package method

import (
	"fmt"
	"sync"
	"testing"
)

// 1 方法是与对象实例绑定的特殊函数
// 方法是面向对象编程的基本概念，用于维护和展示对象的自身状态
// 对象是收敛的，每个实例都有各自不同的独立特征，以属性和方法来暴露对外通信接口
// 普通函数专注于算法流程，通过接受参数来完成特定逻辑运算，并返回最终结构。

// 方法和函数定义语法区别在于前者有前置实例接收参数(receiver)，编译器靠此来确定方法所属类型
// 在某些语言中，虽然没有显示定义，但会在调用时隐式传递this实例参数

// 可以为当前包，出了接口和指针以外的任何类型定义方法
type A int

func (a A) toString() string {
	return fmt.Sprintf("%#x", a)
}

func TestMethod(t *testing.T) {
	var a A = 1
	fmt.Println(a.toString())

	// var b int = 2
	// b.toString()   b不能调用toString方法
}

// 方法不支持重载(overload)
// receiver 参数名没有限制，按照惯例应该选用简短有意义的名称，this self 这种不推荐使用
// 如果方法内部没有引用实例，可以省略参数名，仅仅保留类型
type B int

func (B) toString() string {
	return fmt.Sprintf("hello, int number") // 这里面没有使用对象实例，可以不写参数名，上面A的就不能省
}

func TestMethod01(t *testing.T) {
	var b B = 1
	fmt.Println(b.toString())
}

//方法可以看做是特殊的函数，那么 receiver 的类型自然可以是基础类型，也可以是指针类型
// 这会关系到调用时对象实例是否被复制
// 就是说给一个类型创建方法有2种方式，一种receiver就是这种类型的实例，这种情况下对象实例会复制
// 另一种就是传递这种类型的实例的指针，这种方法还是算成这种类型的方法，不是这种类型的指针的方法
// 第二种，传递这种类型的实例的指针，是这种类型的实例是不会被复制的
// 可以用实例的实例值或指针去调用方法，编译器会根据方法的receiver类型自动在基础类型和指针类型之间转换
// 但是不支持多级指针去调用方法
type C int

func (c C) value() {
	c++
	fmt.Printf("v: %p, %v\n", &c, c)
}

// func (c *C) value() {
// 	(*c)++
// 	fmt.Printf("v: %p, %v\n", c, *c)
// }

func (c *C) pointer() { // 这还是C类型的方法，不是C类型的指针的方法
	(*c)++
	fmt.Printf("v: %p, %v\n", c, *c)
}

func TestMethod02(t *testing.T) {
	var c C = 5
	fmt.Printf("v: %p, %v\n", &c, c)

	c.value()
	fmt.Printf("v: %p, %v\n", &c, c)

	c.pointer()
	fmt.Printf("v: %p, %v\n", &c, c)

	var c1 C = 100
	fmt.Printf("v: %p, %v\n", &c1, c1)
	var cPointer *C = &c1 //cPointer是一个指向C类型实例的指针，它是一个指针
	fmt.Printf("v: %p, %v\n", cPointer, *cPointer)

	cPointer.value() // value方法的receiver是C类型，不是指针类型，编译器会自动将cPointer转换为C类型，其实传过去的是*cPointer
	fmt.Printf("v: %p, %v\n", cPointer, *cPointer)

	cPointer.pointer()
	fmt.Printf("v: %p, %v\n", cPointer, *cPointer)
}

/* 这种转换是隐式的,那么为什么不像下面这样这样定义呢?
其实这样很蠢,所有的代码都要用按实例调用和实例指针调用写2份
并且这样是不行的,GO不支持重载,下面两个方法签名一样,编译也通不过
func (c C) value() {
	c++
	fmt.Printf("v: %p, %v\n", &c, c)
}

func (c *C) value() {
	(*c)++
	fmt.Printf("v: %p, %v\n", c, *c)
}

要么不支持实例指针调用方法,要么就只能隐式转换
可是实例指针调用可以避免实例复制的问题,有其优越性

注意， 只能用一级指针调用方法，不能用多级指针去调用方法
*/

// 指针类型的receiver必须是合法指针(包括nil)，或能获取实例地址
type X struct{}

func (x *X) test() {
	println("hi!", x)
}

func TestMethod03(t *testing.T) {
	var a *X

	a.test() // 允许，相当于 test( nil )

	// X{}.test()  不允许cannot take the address of X literalg
}

/*
如何选择方法的receiver类型？
1 要修改实例状态的，用 *T
2 不需要修改状态的小对象或固定值，建议用T
3 大对象建议用 *T, 可以减少复制成本
4 引用类型，字符串，函数等指针包装对象，直接用T
5 包含Mutex等同步字段，用 *T，避免因复制造成所操作无效
6 其他无法确定的情况，都用 *T
*/

// 匿名字段
// 可以湘访问匿名字段成员那样调用其方法，由编译器负责查找
type data struct {
	sync.Mutex // 匿名锁变量
	buf        [1024]byte
}

func TestMethod04(t *testing.T) {
	d := data{}
	d.Lock() //data类型的结构体实例，直接调用其内部匿名成员的方法，编译器会处理为sync.(*Mutex).Lock()
	defer d.Unlock()
}

// 方法也有同名遮蔽问题，利用这种特性，可以实现类似覆盖(override)操作
type user struct{}
type manager struct {
	user
}

func (u *user) toString() {
	println("user")
}
func (m *manager) toString() {
	println("manager")
}

func TestMethod05(t *testing.T) {
	m := manager{}

	m.toString() //m是manager类型实例，其内部有个user匿名变量，m.toString()既可以可是调用自己的，也可以是user的
	// 此处实际调用的是manager的，因为它自己有个toString方法
	// 如果manager类型没有toString方法，那就是调用user的，这就相当于覆盖(override)了user的方法
	// 尽管能访问匿名字段成员和方法，它们之间不属于继承关系
}

// 方法集
/*
类型有一个方法集(method set)，这决定了它是否实现了某个接口
1 类型 T 方法集包含所有 receiver T 方法
2 类型 *T 方法集包含所有 receiver T + *T 方法
3 匿名嵌入 S，T 方法集包含所有 receiver S 方法
4 匿名嵌入 *S，T 方法集包含所有所有 receiver S + *S 方法
5 匿名嵌入 S 或 *S，*T 方法集包含所有 receiver S + *S 方法

方法集仅影响接口实现和方法表达式转换，与通过实例或实例指针调用无关

很明显，匿名字段就是为方法集准备的，否则，没必要为了少写个字段名而大费周章
面向对象的三大特征：封装，继承，多态，Go只实现了部分特征，它更倾向于：组合由于继承
将模块分解成相互独立的更小的单元，分别处理不同方面的需求，最后以匿名嵌入的方式组合到一起
共同实现对外接口。
*/

// 表达式
// 方法和函数一样，除了直接调用外，还可以赋值给变量和作为参数传递。？？？能不能作为返回值？？？
// 根据具体引用方法的不同，可以分为 expression 和 value 两种状态
// 看不懂，以后再说
