#/bin/bash

set -x

if [ ! -f ./build/rootfs.tar ];
then
    curl http://163.172.156.204/x86_64-ubuntu-xenial/rootfs.tar > build/rootfs.tar
fi

echo "mkdir /media/rootfs && mount -t ext4 /scaleway/disk.img /media/rootfs && tar xpf /scaleway/rootfs.tar -C  /media/rootfs --numeric-owner; sync; umount /media/rootfs" | docker run --privileged --rm -i -v `pwd`/build:/scaleway busygox bash

# remove root password
docker run --rm --privileged -it -v `pwd`/build:/build busygox chroot passwd -d root
# remove motd to enable the login
docker run --rm --privileged -it -v `pwd`/build:/build busygox chroot rm /etc/update-motd.d/50-scw
