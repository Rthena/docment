package main

import "fmt"

func main() {
	test1()
	test2()
}

func test1() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Println("bbbb")
		fmt.Printf("idx:%d, val:%s\n", i, v)
	}	
}

func test2() {
	fmt.Println("good job")
}
