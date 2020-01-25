package main

import (
 "fmt"
 "strings"

 )
//1
func replaceAllvowel (s string) string{

s = strings.Map(func(r rune) rune {
	switch r {
	case 'a':
		return '1'
	case 'e':
		return '1'
	// etc.
	default:
		return r
	}
}, s)
return s}


//2 remove all chracters except alphabets
func removeAllcharacters(s string) string {
	s = strings.Map(func(r rune) rune{
		if((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return r
		} else {
			return rune(0)
		}
	},s)
	return s
}
//3 return the sum of vowels and consonants

func sum (s string) (int,int){
	var vow=0
	var cons=0
	s = strings.Map(func(r rune) rune {
		if(r == 'a' || r == 'i' || r == 'y' || r == 'i' || r =='o'){
			//fmt.Println(r)
			//fmt.Println(vow)
			vow=vow+1
			return r
		}else{
	         cons=cons+1
	         return r
}
	},s)
	return vow,cons
}


func main(){
	//1
	//2
	//fmt.Print(removeAllcharacters("aaaaeeee//<<,>:hhhh"))

 //3
	fmt.Println(sum("kioouyt"))

}