### Busygox

#### Overview

(POC) Busybox like written in golang

#### Instructions

```console
$ docker build -t busygox .
$ docker run --rm -it -v $(pwd)/build:/build busygox
# (OS X) brew install qemu
# (Debian based) apt-get install qemu-system
$ qemu-system-x86_64  -boot once=c -nographic -kernel build/bzImage -initrd build/initramfs.cpio -drive format=raw,file=build/disk.img --append "console=ttyS0 nousb quiet root=/dev/sda1 rootfstype=ext4 selinux=0" -m 1024
# to exit ctrl-a x
```
