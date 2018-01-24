package main

import "fmt"


func compress(message string){
	fmt.Println(len(message))

	count := 1

	for i := 0; i < len(message) - 1; i++ {
		if(message[i] == message[i + 1]){
			count++
		}else{
			count = 0
		}
	}
	fmt.Println(count)
}

func main(){
	compress("aaab")
}