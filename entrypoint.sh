#!/bin/bash
cp -v /kernel-build/linux-4.7/arch/x86_64/boot/bzImage /initfs/initramfs.cpio /build
if [ ! -f /build/disk.img ];
then
	cp -v /initfs/disk.img /build
fi

echo ""
echo "Done ! now you can run the following command"
echo ""
echo "$ qemu-system-x86_64  -boot once=c -nographic -kernel build/bzImage -initrd build/initramfs.cpio -drive format=raw,file=build/disk.img  --append \"console=ttyS0 nousb quiet root=/dev/sda1 rootfstype=ext4 selinux=0\"  -m 1024"
