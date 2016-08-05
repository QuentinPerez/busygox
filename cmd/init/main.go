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
	fmt.Println(`  ___   wWw  wWw  oo_   wWw  wWw   \/       .-.    wW    Ww`)
	fmt.Println(` (___)__(O)  (O) /  _)-<(O)  (O)  (OO)    c(O_O)c (O)\  /(O)`)
	fmt.Println(" (O)(O) / )  ( \\ \\__ `. ( \\  / ),'.--.)  ,'.---.`, `. \\/ .'")
	fmt.Println(" /  _\\ / /    \\ \\   `. | \\ \\/ // /|_|_\\ / /|_|_|\\ \\  \\  /")
	fmt.Println(` | |_))| \____/ |   _| |  \o / | \_.--. | \_____/ |  /  \`)
	fmt.Println(" | |_))'. `--' .`,-'   | _/ /  '.   \\) \\'. `---' .`.' /\\ `.")
	fmt.Println(" (.'-'   `-..-' (_..--' (_.'     `-.(_.'  `-...-' (_.'  `._)")
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
