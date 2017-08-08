package luhn

import (
	"strconv"
	"strings"
	"fmt"
)

const testVersion=2

func Valid(num string) bool {
	fmt.Printf("checking: %s\n",num)
	sum := 0
	num = strings.Replace(num, " ","",-1)
	for i:=len(num)-1;i>=0;i--{
		number, _ := strconv.Atoi(string(num[i]))
		fmt.Printf("i:%d before:%d ",i, number)
		if i%2==1{
			number = (2*number)%9
		}
		fmt.Printf("now: %d\n", number)
		sum += number
	}
	if sum == 0 {
		return false
	}
	return sum%10 == 0
}
