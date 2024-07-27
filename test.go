package main

import "fmt"

func main() {

	test1()
}

func test1() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Printf("idx:%d, val:%s\n", i, v)
	}	
}
