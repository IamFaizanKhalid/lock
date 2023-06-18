package main

import (
	"fmt"
	"time"

	"github.com/IamFaizanKhalid/lock"
)

func main() {
	lock.HandleEvents(lockEventHandler)
}

func lockEventHandler(e lock.Event) {
	fmt.Println(e.Time.Format(time.TimeOnly), e.Locked)
}
