package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
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
	err := syscall.Mount("none", "/dev", "devtmpfs", syscall.MS_MGC_VAL, "mode=755")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sorry %s", err)
	}
	files, err := ioutil.ReadDir("/dev")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	for {
		time.Sleep(24 * time.Hour)
	}
}
