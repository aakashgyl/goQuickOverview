package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// make is used for slices, maps, and channels

func SimpleVariables() {
	// INT type
	var a int //declared and has default value
	var b, c int = 1, 2
	var d, e = 1, 2 //type inference
	fmt.Println(a, b, c, d, e)

	// MIX type
	var f, python, java = true, false, "no!"
	g, python3, java8 := true, false, "no!"
	fmt.Println(f, python, java, g, python3, java8)

	// CONST
	const Truth = true
	//const b = math.Sqrt(4) -> wont work as b should be known at compile time and expr will be evaluated at runtime.

	// TYPE conversion - do it explicitly
	var price float32 = float32(a)
	fmt.Println(price)
}

func ListVariables() {
	// Arrays
	var a [3]int
	a[0] = 1
	b := [3]int{1, 2}         //3rd element gets default value of 0
	c := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a, b, c)

	// Slice
	var d []int = a[1:3]
	e := []int{6, 7, 8} //creates an array and returns a slice reference
	fmt.Println(d, e)

	// copy Arrays
	arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}
	arrVal := arr1  // copy by Value
	arrRef := &arr1 // copy by Reference
	fmt.Println(arrVal, arrRef)

	// copy Slice
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1 // copy by reference
	s3 := make([]int, len(s1))
	copy(s3, s1) // copy by value
	fmt.Println(s2, s3)

	// append
	// func append(s []T, x ...T) []T
	s2 = append([]int(nil), s1...) // syntax -> <slice> = append(<slice>, elements...)

	// remove an element at index i -> re-slicing
	s := []int{1, 2, 3, 4, 5}
	i := 2
	s = append(s[:i], s[i+1:]...)

	// check if slice is nil
	if s == nil {

	}

	// check if slice has any data
	fmt.Println(len(s), cap(s))

	// 2-D slice
	two_d_slice := make([][]int, 2)
	for i := range two_d_slice {
		two_d_slice[i] = make([]int, 2)
	}
}

func Map() {
	// initialized map
	employeeSalary1 := make(map[string]int) // ready for insert
	employeeSalary2 := map[string]string{}  // ready for insert
	employeeSalary3 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	fmt.Println(employeeSalary1, employeeSalary2, employeeSalary3)

	// uninitialized map
	var employeeSalary4 map[string]int            //this map is uninitialized, initialize with make
	fmt.Println(employeeSalary4["uninitialized"]) // prints zero value of return type... panics for set operation

	// check if key exists
	value, ok := employeeSalary3["key"]
	fmt.Println(value, ok)

	// delete from map
	delete(employeeSalary1, "steve")

	// copy map by reference
	modified := employeeSalary1 //both now points to same internal DS
	fmt.Println(modified)

	// copy map by value -> create new map, use for loop and update new map
	copybyvalue := make(map[string]int)
	for index, element := range employeeSalary1 {
		copybyvalue[index] = element
	}

	// check if map is nil
	if copybyvalue == nil {

	}

	// compare two maps
	reflect.DeepEqual(employeeSalary1, copybyvalue)
}

func Structs() {

}

func Channels() {
	var a chan int // a == nil here
	b := make(chan int)
	c := make(chan int, 2) // unbuffered has value 0 by default
	fmt.Println(a, b, c)

	// channel is open or closed?
	v, ok := <-b // if ok == true, channel is open... if ok == false, channel is closed and v will be zero value of channel's type
	// Note: For buffered channels, we can still read if channel if closed. But if ok==false, it means all values are read.
	fmt.Println(v, ok)

	// len and cap
	fmt.Println(len(c)) // elements in channel
	fmt.Println(cap(c)) // how many elements can it handle

	// unidirectional channel
	sendch := make(chan<- int) // send channel, cant read from it
	fmt.Println(sendch)
}

func Switch() {
	// Constant switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	default:
		fmt.Printf("%s.", os)
	}

	// Switch with maths operation
	today := time.Monday
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}

	// Expressionless switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good evening.")
	}
}

func Loops() {
	// For - standard loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// For Range loop
	a := []int{1, 2, 3, 4, 5}
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}

	// While loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
