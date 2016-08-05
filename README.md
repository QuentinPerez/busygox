### Busygox

#### Overview

(POC) Busybox like written in golang

#### Instruction

```console
$ docker build -t busygox .
$ docker run --rm -it -v $(pwd)/build:/build busygox
# (OS X) brew install qemu
# (Debian based) apt-get install qemu-system
$ qemu-system-x86_64 -kernel build/bzImage -initrd build/initramfs.cpio -drive format=raw,file=build/disk.img
```
