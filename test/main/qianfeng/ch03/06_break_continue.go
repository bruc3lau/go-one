package main

import "fmt"

func main() {
	fmt.Println()

	//i := 0
	//for {
	//	if i > 10 {
	//		break
	//	}
	//	i++
	//	fmt.Println(i)
	//}

	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		//break
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	//for i := 1; i <= 50; i++ {
	//	if i%10 == 4 || i/10%10 == 4 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	//LOOP: X
	//	for i := 1; i <= 50; i++ {
	//		if i%10 == 4 || i/10%10 == 4 {
	//			goto LOOP
	//		}
	//		fmt.Println(i)
	//	}

	//	i := 0
	//OUT:
	//	for i < 50 {
	//		i++
	//		if i%10 == 4 || i/10%10 == 4 {
	//			goto OUT
	//		}
	//		fmt.Println(i)
	//	}

	i := 0
OUT:
	for i < 100 {
		i++
		for j := 2; j < i; j++ {
			if i%j == 0 {
				goto OUT
			}
		}
		fmt.Println(i)
	}

}
