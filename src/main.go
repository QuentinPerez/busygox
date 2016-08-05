package main

import (
	"fmt"
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
	for {
		time.Sleep(24 * time.Hour)
	}
}
