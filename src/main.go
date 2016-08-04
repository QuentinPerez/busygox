package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	for {
		fmt.Fprintln(os.Stderr, "Coucou")
		time.Sleep(1 * time.Second)
	}
}
