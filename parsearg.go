package main

import (
	"os"
	"fmt"
	"reflect"
)

func main() {
	//iflen(os.Args) > 1 {
	fmt.Println(os.Args[1]) 
	fmt.Println(os.Args[2]) 
	fmt.Println(reflect.TypeOf(os.Args[1]).Kind(), "tagada")
}
