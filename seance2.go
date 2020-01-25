package main

import (
	"fmt"
)

//
//func main(){
//	sum := 0
//	for i := 1; i < 5; i++ {
//		//only odd number
//		if i%2 != 0 {
//			fmt.Println(i)
//			continue
//		}
//		sum += i
//	}
//	fmt.Println(sum)
//	//sum 6=4+2
////}
//func main(){
//
//	switch time.Now().Weekday() {
//case time.Saturday:
//fmt.Println("Today is Saturday.")
//case time.Sunday:
//fmt.Println("Today is Sunday.")
//default:
//fmt.Println("Today is a weekday.")
//}
//}
func developerLevel(s string) string{

		switch s {
	case "j1","j2","j3":
	fmt.Println("junior")
     return "junior"
	case "s1","s2","s3":
	fmt.Println("senior")
	return "senior"
	case "d1","d2","d3":
		fmt.Println("director")
		return "director"
	default:
	fmt.Println("developer")
		return s
	}
}

//
//	}
//}
func main(){
	fmt.Print(developerLevel("j1"))
}
