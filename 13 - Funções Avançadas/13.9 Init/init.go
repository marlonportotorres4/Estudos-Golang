package main

import "fmt"

func main() {
	fmt.Println("Main")
}

func init()  {
	fmt.Println("init")
} //vai ser executada antes do main