

PWD := $(shell pwd)

all: setup
	qemu-system-x86_64  -boot once=c -nographic -kernel build/bzImage -initrd build/initramfs.cpio -drive format=raw,file=build/disk.img  --append "console=ttyS0 nousb root=/dev/sda1 rootfstype=ext4 selinux=0"  -m 1024

setup:
	docker build -t busygox .
	docker run --rm -it -v $(PWD)/build:/build busygox

scaleway: setup
	./import-from-scaleway.sh
	qemu-system-x86_64  -boot once=c -nographic -kernel build/bzImage -initrd build/initramfs.cpio -drive format=raw,file=build/disk.img  --append "console=ttyS0 nousb root=/dev/sda1 rootfstype=ext4 selinux=0"  -m 1024
