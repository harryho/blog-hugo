package main

import (
	"fmt"
	"math"
	"reflect"
)

// func main() {
// 	fmt.Println("hello world")

// 	// var arr1 [5]int;

// 	// for  i := range arr1 {
// 	//     fmt.Printf(" index = %d , value = %d \n", i, arr1[i])
// 	// }

// 	// var arr2 = [...]int{0,0,0,0,0}
// 	// for  i := range arr2{
// 	//     fmt.Printf(" index = %d , value = %d \n", i, arr1[i])
// 	// }

// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[0:2]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[0:]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[:2]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[1:5]
// 	fmt.Printf(" slice = %v \n", slice)
// 	slice = arr[:5]
// 	fmt.Printf(" slice = %v \n", slice)

// 	sl_from := []int{1, 2, 3}
// 	sl_to := make([]int, 10)
// 	n := copy(sl_to, sl_from)
// 	fmt.Println(sl_to)
// 	// output: [1 2 3 0 0 0 0 0 0 0]
// 	fmt.Printf("Copied %d elements\n", n) // n == 3
// 	sl3 := []int{1, 2, 3}
// 	sl3 = append(sl3, 4, 5, 6)
// 	fmt.Println(sl3)

// 	str := "golang"
// 	chars := []byte(str)
// 	chars[1] = '0'
// 	newStr := string(chars)
// 	fmt.Println(newStr)
// }

// type Stack struct {
// }

// func (*Stack) Pop(st []int) int {
// 	v := 0
// 	for ix := len(st) - 1; ix >= 0; ix-- {
// 		if v = st[ix]; v != 0 {
// 			st[ix] = 0

// 		}
// 	}
// 	return v
// }

// func main() {
//     result := 0
//     for i:=0; i <= 10; i++ {
//         result = fibonacci(i)
//         fmt.Printf("fibonacci(%d) is: %d\n", i, result)
//     }
// }

// func fibonacci(n int) (res int) {
// 	if n <= 1 {
// 		res = 1
// 	} else {
// 		res = fibonacci(n-1) + fibonacci(n-2)
// 	}
// 	return
// }

// func main() {
// 	// callback(1, Add)
// 	plus := func(x, y int) int { return x + y }
// 	fmt.Println(plus(1, 2)) // 3

// 	// invoke func immediatley
// 	val := func(x, y int) int { return x + y }(1, 2) // 3
// 	fmt.Println(val)
// }
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
	// this becomes Add(1, 2)
}

type TagType struct { //tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b int
	c float32
	int
	// anonymous field
	innerS // anonymous field
}

func Func() {
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is: ", outer2)
}

type Shape interface {
	Area() float32
}
type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}

func modifyValByReflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)                      // Pass value
	fmt.Println("settability of v:", v.CanSet()) // false
	v = reflect.ValueOf(&x)                      // Note: take the address of x.
	fmt.Println("type of v:", v.Type())          // float64
	fmt.Println("settability of v:", v.CanSet()) // false
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)         // <float64 Value>
	fmt.Println("settability of v:", v.CanSet()) // true
	v.SetFloat(3.1415)                           // this works!
	fmt.Println(v.Interface())
	fmt.Println(v) // <float64 Value>
}

func main() {
	shapes := []Shape{Rectangle{5, 3}, &Square{5}, &Circle{5}}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	var shape Shape
	shape = &Square{4}

	switch t := shape.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case *Rectangle:
		fmt.Printf("Type Rectangle %T with value %v\n", t, t)
	case nil:
		fmt.Println("nil value: nothing to check?")
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
