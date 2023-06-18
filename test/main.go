package main

import (
	"fmt"
	"time"

	"github.com/IamFaizanKhalid/lock"
)

func main() {
	for e := range lock.Notify() {
		fmt.Println(e.Time.Format(time.TimeOnly), e.Locked)
	}
}
