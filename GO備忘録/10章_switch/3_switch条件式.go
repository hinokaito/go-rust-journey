package main

import "fmt"

func main() {
	score := 88

	switch { // switchの後に変数を書かない
	case score >= 90:
		fmt.Println("優")
	case score >= 80:
		fmt.Println("良")
	case score >= 70:
		fmt.Println("可")
	default:
		fmt.Println("不可")
	}

	// if-elseよりもすっきりして見えると思います
	// if-elseで書いてみましょう
	if score >= 90 {
		fmt.Println("優")
	} else if score >= 80 {
		fmt.Println("良")
	} else if score >= 70 {
		fmt.Println("可")
	} else {
		fmt.Println("不可")
	}
}
