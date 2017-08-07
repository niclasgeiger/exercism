package raindrops

import "fmt"

const testVersion = 2

func Convert(num int) string {
	s := ""
	if num%3 == 0 {
		s += "Pling"
		num /= 3
	}
	if num%5 == 0 {
		s += "Plang"
		num /= 5
	}
	if num%7 == 0 {
		s += "Plong"
		num /= 7
	}
	if(len(s)==0){
		return fmt.Sprintf("%d",num)
	}
	return s
}
