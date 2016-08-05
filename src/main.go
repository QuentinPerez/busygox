package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

func mountPseudoFS() (err error) {
	// mount /dev
	if err = syscall.Mount("none", "/dev", "devtmpfs", syscall.MS_MGC_VAL, "mode=755"); err != nil {
		return
	}
	// mount /run
	if err = syscall.Mount("tmpfs", "/run", "tmpfs", syscall.MS_MGC_VAL|syscall.MS_NOSUID|syscall.MS_NODEV, "mode=0755,size=10%"); err != nil {
		return
	}
	// mount /sys
	if err = syscall.Mount("none", "/sys", "sysfs", syscall.MS_MGC_VAL|syscall.MS_NOSUID|syscall.MS_NODEV|syscall.MS_NOEXEC, ""); err != nil {
		return
	}
	// mount /proc
	if err = syscall.Mount("none", "/proc", "proc", syscall.MS_MGC_VAL|syscall.MS_NOSUID|syscall.MS_NODEV|syscall.MS_NOEXEC, ""); err != nil {
		return
	}
	return
}

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

	defer func() {
		for {
			time.Sleep(24 * time.Hour)
		}
	}()

	err := mountPseudoFS()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return
	}

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
