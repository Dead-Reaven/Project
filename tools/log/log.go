package log

import (
	"fmt"
	"time"
)

var counter int

func Out(msg string) {
	counter++
	now := time.Now()
	fmt.Printf("\n%d. %s --- %s ---\n", counter, now.Format(time.UnixDate), msg)
}

func Success(msg string) {
	now := time.Now()
	fmt.Printf("%s ✅ %s ✅\n", now.Format(time.UnixDate), msg)
}

func Error(msg string) {
	counter++
	now := time.Now()
	fmt.Printf("%s ⛔️ %s ⛔️\n", now.Format(time.UnixDate), msg)
}
