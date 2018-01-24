package main

import "fmt"


func main(){
	resultFloat := (float64(3) / float64(10))
	resultFloat *= 10
	result := int(resultFloat)
	fmt.Println(result)
}