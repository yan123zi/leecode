package main

import "fmt"

func main() {
	num:=[]string{"1","2","3","4","5","6","7","8","9"}
	index:=[]int{0,4,6,6,8,8,2,2,4,1,5}
	phone:=""
	for _, idx := range index {
		phone+=num[idx]
	}
	fmt.Println(phone)
}
