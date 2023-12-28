package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.Now()
	fmt.Println(time.Format("01-02-2006 Mon"))
}
