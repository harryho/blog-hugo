package main

import "fmt"
func main() {
    fmt.Println("hello world")

    // var arr1 [5]int;

    // for  i := range arr1 {
    //     fmt.Printf(" index = %d , value = %d \n", i, arr1[i])
    // }

    // var arr2 = [...]int{0,0,0,0,0}
    // for  i := range arr2{
    //     fmt.Printf(" index = %d , value = %d \n", i, arr1[i])
    // }

    arr := [5]int {1,2,3,4,5}
    slice := arr[0:2]
    fmt.Printf( " slice = %v \n", slice)
    slice = arr[0:]
    fmt.Printf( " slice = %v \n", slice)
    slice = arr[:2]
    fmt.Printf( " slice = %v \n", slice)
    slice = arr[1:5]
    fmt.Printf( " slice = %v \n", slice)
    slice = arr[:5]
    fmt.Printf( " slice = %v \n", slice)


}