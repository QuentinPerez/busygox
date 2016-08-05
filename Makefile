

PWD := $(shell pwd)

all:
	docker build -t busygox .
	docker run --rm -it -v $(PWD)/build:/build busygox
	qemu-system-x86_64  -boot once=c -nographic -kernel build/bzImage -initrd build/initramfs.cpio -drive format=raw,file=build/disk.img  --append "console=ttyS0 nousb quiet root=/dev/sda1 rootfstype=ext4 selinux=0"  -m 1024
