package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
	"time"
)

func banner() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("_____")
	fmt.Println("|  _ \\")
	fmt.Println("| |_) |_   _ ___ _   _  __ _  _____  __")
	fmt.Println("|  _ <| | | / __| | | |/ _` |/ _ \\ \\/ /")
	fmt.Println("| |_) | |_| \\__ \\ |_| | (_| | (_) >  <")
	fmt.Println("|____/ \\__,_|___/\\__, |\\__, |\\___/_/\\_\\")
	fmt.Println("                  __/ | __/ |")
	fmt.Println("                 |___/ |___/")
	fmt.Println("")
	fmt.Println("")
}

func entrypoint() (err error) {
	banner()
	if err = setup(); err != nil {
		return
	}

	// test
	files, err := ioutil.ReadDir("/dev")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Error -> %s\n", r)
			fmt.Fprintf(os.Stderr, "Backtrace -> %s\n", debug.Stack())
		}
		for {
			time.Sleep(24 * time.Hour)
		}
	}()
	if err := entrypoint(); err != nil {
		panic(err)
	}
}
