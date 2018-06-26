package main

import (
		"fmt"
	)

func main() {
i:="which"
	switch i {
	case"who":
		fmt.Println("who")
	case"what":
		fmt.Println("what")
	case"where":
		fmt.Println("where")
	case"which":
		fmt.Print("which")
	default:
		fmt.Println("IDK")
	}
}