package main

import "fmt"



type Vector2f struct{
	x float32
	y float32
}

func dot(v1 Vector2f , v2 Vector2f) (float32){
	return v1.x * v2.x + v1.y * v2.y
}
func add(x int, y int) int {
	return x + y
}

func swap(x , y int) (int,int){
	return y , x
}

func changeMe(x *int){
	(x)++
}


func main(){

     fmt.Println(add(5,4))


     for i := 0; i < 10; i++{
     	fmt.Println(i)
     }

     v1 := Vector2f{1.0, 0.0}
     v2 := Vector2f{-1.0, 0.0}
     fmt.Println(dot(v1,v2))
 }