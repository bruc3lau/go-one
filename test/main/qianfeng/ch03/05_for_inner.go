package main

import "fmt"

func main() {
	//fmt.Println()

	//rectangle
	//for i := 0; i < 10; i++ {
	//	for j := 0; j < 10; j++ {
	//		fmt.Print("X ")
	//	}
	//	fmt.Print("\n")
	//}

	//////triangle
	//for i := 0; i < 10; i++ {
	//	for j := 0; j < i; j++ {
	//		fmt.Print("X ")
	//	}
	//	fmt.Println()
	//}

	////triangle
	//for i := 0; i < 10; i++ {
	//	for j := 9; j > i; j-- {
	//		fmt.Print("X ")
	//	}
	//	fmt.Println()
	//}

	//右下
	//for i := 0; i < 10; i++ {
	//	for j := 9; j > i; j-- {
	//		fmt.Print("  ")
	//	}
	//	for k := 0; k < i; k++ {
	//		fmt.Print("X ")
	//	}
	//	fmt.Println()
	//}

	////右上
	//for i := 0; i < 10; i++ {
	//	for j := 0; j < i; j++ {
	//		fmt.Print("  ")
	//	}
	//	for j := 9; j > i; j-- {
	//		fmt.Print("X ")
	//	}
	//	fmt.Println()
	//}

	//等腰
	//for i := 0; i < 10; i++ {
	//	for j := 9; j > i; j-- {
	//		fmt.Print("  ")
	//	}
	//	for k := 0; k < i; k++ {
	//		fmt.Print("X ")
	//	}
	//	for j := 0; j < i; j++ {
	//		fmt.Print("X ")
	//	}
	//	fmt.Println()
	//}

	//乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

}
