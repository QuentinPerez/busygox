#/bin/bash

set -x

if [ ! -f ./build/rootfs.tar ];
then
    curl http://163.172.156.204/x86_64-ubuntu-xenial/rootfs.tar -O ./build/rootfs.tar
fi
 
 echo "mkdir /media/rootfs && mount -t ext4 /scaleway/disk.img /media/rootfs && tar xpf /scaleway/rootfs.tar -C  /media/rootfs --numeric-owner; sync; umount /media/rootfs" | docker run --privileged --rm -i -v /Users/quentinperez/go/src/github.com/QuentinPerez/busygox/build:/scaleway busygox bash
