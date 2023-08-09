package time

import (
	"time"
	"fmt"
)

func Timer() {
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Println("It took ", elapsed)
}