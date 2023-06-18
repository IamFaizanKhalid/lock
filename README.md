# lock
[![Go Report Card](https://goreportcard.com/badge/github.com/IamFaizanKhalid/lock)](https://goreportcard.com/report/github.com/IamFaizanKhalid/lock) [![Release](https://img.shields.io/github/v/release/IamFaizanKhalid/lock.svg?style=flat-square)](https://github.com/IamFaizanKhalid/lock/releases) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

<img align="right" width="100" src="./_lock.png" alt="Nishan Pakistan">

A simple Golang package to get notified when the screen gets locked.

⚠️ **Warnings:**
1. This was developed as an experiment and may not always work.
2. It currently supports only macOS, Windows and Linux.
3. Not tested on all linux desktops

## Installation

```console
go get -u github.com/IamFaizanKhalid/lock
```

## Usage Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/IamFaizanKhalid/lock"
)

func main() {
	if lock.IsScreenLocked() {
		fmt.Println("Screen is locked...")
	} else {
		fmt.Println("Screen is not locked...")
	}

	lock.HandleEvents(lockEventHandler)
}

func lockEventHandler(e lock.Event) {
	fmt.Print(e.Time.Format(time.TimeOnly), "\t")
	if e.Locked {
		fmt.Println("screen locked")
	} else {
		fmt.Println("screen unlocked")
	}
}
```


## License

[MIT License](./LICENSE)
