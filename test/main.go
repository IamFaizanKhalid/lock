package main

import (
	"fmt"
	"github.com/IamFaizanKhalid/lock"
)

func main() {
	for range lock.Notify() {
		fmt.Println("screen locked")
	}
}
