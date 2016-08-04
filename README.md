### Busygox

#### Overview

(POC) Busybox like written in golang

#### Instruction

```console
$ docker build -t busygox .
$ docker run --rm -it -v $(pwd)/kernel:/kernel busygox
$ qemu-system-x86_64 -kernel kernel/bzImage -initrd kernel/initramfs.cpio
```
