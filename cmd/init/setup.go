package main

import (
	"syscall"

	"github.com/juju/errors"
)

func mountPseudoFS() (err error) {
	// mount /dev
	if err = syscall.Mount("none", "/dev", "devtmpfs", syscall.MS_MGC_VAL, "mode=755"); err != nil {
		err = errors.Annotate(err, "Unable to mount /dev")
		return
	}
	// create /dev/pts directory
	if err = syscall.Mkdir("/dev/pts", 0777); err != nil {
		err = errors.Annotate(err, "Unable to create /dev/pts")
		return
	}
	// mount /dev/pts
	if err = syscall.Mount("none", "/dev/pts", "devpts", syscall.MS_MGC_VAL|syscall.MS_NOSUID|syscall.MS_NOEXEC, "gid=5,mode=0620"); err != nil {
		err = errors.Annotate(err, "Unable to mount /dev/pts")
		return
	}
	// mount /run
	if err = syscall.Mount("tmpfs", "/run", "tmpfs", syscall.MS_MGC_VAL|syscall.MS_NOSUID|syscall.MS_NODEV, "mode=0755,size=10%"); err != nil {
		err = errors.Annotate(err, "Unable to mount /run")
		return
	}
	// mount /sys
	if err = syscall.Mount("none", "/sys", "sysfs", syscall.MS_MGC_VAL|syscall.MS_NOSUID|syscall.MS_NODEV|syscall.MS_NOEXEC, ""); err != nil {
		err = errors.Annotate(err, "Unable to mount /sys")
		return
	}
	// mount /proc
	if err = syscall.Mount("none", "/proc", "proc", syscall.MS_MGC_VAL|syscall.MS_NOSUID|syscall.MS_NODEV|syscall.MS_NOEXEC, ""); err != nil {
		err = errors.Annotate(err, "Unable to mount /proc")
		return
	}
	return
}

func setup() (err error) {
	if err = mountPseudoFS(); err != nil {
		return
	}
	return
}
