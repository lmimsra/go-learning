package main

import "fmt"

func main() {
	// about Print function
	fmt.Println("print ln")
	fmt.Print("hello, ")
	fmt.Println("世界")
	fmt.Println("A", 100, true, 1.5)

	// about Printf function
	fmt.Println("-----printf-----")
	fmt.Printf("Hello, 世界\n")
	fmt.Printf("%d-%s \n", 100, "偶数")

	// about Scan function
	fmt.Println("-----Scan-----")

	var txt string
	fmt.Print("何か入れて > ")
	_, err := fmt.Scan(&txt)
	if err != nil {
		return
	}

	fmt.Printf("「%s」が入力されたぞ！\n", txt)

}
