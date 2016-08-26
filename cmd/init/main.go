package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/juju/errors"
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

func listDir(dir string) (err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	fmt.Println(dir, " # from /dev/sda")
	for _, file := range files {
		fmt.Printf("[%s]\n", file.Name())
	}
	return
}

func delete_content(newroot string) (err error) {
	err = filepath.Walk("/rootfs", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	return
}

func switch_root() (err error) {
	if err = syscall.Chdir("/rootfs"); err != nil {
		err = errors.Annotate(err, "Unable to change directory (/rootfs)")
		return
	}
	if err = syscall.Mount(".", "/", "", syscall.MS_MOVE, ""); err != nil {
		err = errors.Annotate(err, "Unable to mount /rootfs")
		return
	}
	if err = syscall.Chroot("."); err != nil {
		err = errors.Annotate(err, "Unable to chroot")
		return
	}
	if err = syscall.Chdir("/"); err != nil {
		err = errors.Annotate(err, "Unable to change directory (/rootfs)")
		return
	}
	if err = syscall.Exec("/sbin/init", []string{"/sbin/init"}, []string{}); err != nil {
		err = errors.Annotate(err, "Execve /sbin/init")
		return
	}
	return
	// fmt.Println(delete_content("/newroot"))
}

func entrypoint() (err error) {
	banner()
	if err = setup(); err != nil {
		return
	}
	// test
	// if listDir("/rootfs/etc/update-motd.d/") != nil {
	// 	err = listDir("/rootfs/")
	// } else {
	err = switch_root()
	// }
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
