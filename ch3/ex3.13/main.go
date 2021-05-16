package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
)

func main() {
	fmt.Println(KB, MB, GB, TB)
}
