package main

import (
	"fmt"
	"time"

	"github.com/IamFaizanKhalid/lock"
)

func main() {
	for range lock.Notify() {
		fmt.Println(time.Now().Format(time.TimeOnly), "screen locked")
	}
}
