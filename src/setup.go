package main

import "syscall"

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

func setup() (err error) {
	if err = mountPseudoFS(); err != nil {
		return
	}
	return
}
