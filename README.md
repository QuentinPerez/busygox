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

_____
|  _ \
| |_) |_   _ ___ _   _  __ _  _____  __
|  _ <| | | / __| | | |/ _` |/ _ \ \/ /
| |_) | |_| \__ \ |_| | (_| | (_) >  <
|____/ \__,_|___/\__, |\__, |\___/_/\_\
                  __/ | __/ |
                 |___/ |___/


/rootfs/  # from /dev/sda
  lost+found

# to exit ctrl-a x
```


#### Scaleway rootfs

```console
$ make scaleway
docker build -t busygox .
docker run --rm -it -v /Users/quentinperez/go/src/github.com/QuentinPerez/busygox/build:/build busygox
./import-from-scaleway.sh
qemu-system-x86_64  -boot once=c -nographic -kernel build/bzImage -initrd build/initramfs.cpio -drive format=raw,file=build/disk.img  --append "console=ttyS0 nousb quiet root=/dev/sda1 rootfstype=ext4 selinux=0"  -m 1024


_____
|  _ \
| |_) |_   _ ___ _   _  __ _  _____  __
|  _ <| | | / __| | | |/ _` |/ _ \ \/ /
| |_) | |_| \__ \ |_| | (_| | (_) >  <
|____/ \__,_|___/\__, |\__, |\___/_/\_\
                  __/ | __/ |
                 |___/ |___/


/rootfs/etc/update-motd.d/  # from /dev/sda
  50-scw
-> reading 50-scw ...
#!/bin/bash

export PATH="${PATH:+$PATH:}/bin:/usr/bin:/usr/local/bin"

[ -r /etc/lsb-release ] && . /etc/lsb-release
[ -r /etc/scw-release ] && . /etc/scw-release
if [ -z "$DISTRIB_DESCRIPTION" ] && [ -x /usr/bin/lsb_release ]; then
	# Fall back to using the very slow lsb_release utility
	DISTRIB_DESCRIPTION=$(lsb_release -s -d)
fi

date=`date`
load=`cat /proc/loadavg | awk '{print $1}'`
root_usage=`df -h / | awk '/\// {print $(NF-1)}'`
memory_usage=`free -m | awk '/Mem:/ { total=$2 } /buffers\/cache/ { used=$3 } END { printf("%3.1f%%", used/total*100)}'`
swap_usage=`free -m | awk '/Swap/ { printf("%3.1f%%", "exit !$2;$3/$2*100") }'`
users=`users | wc -w`
time=`uptime | grep -ohe 'up .*' | sed 's/,/\ hours/g' | awk '{ printf $2" "$3 }'`
processes=`ps aux | wc -l`
ip=`ifconfig $(route | grep default | awk '{ print $8 }') | grep "inet addr" | awk -F: '{print $2}' | awk '{print $1}'`
public_ip=$(scw-metadata --cached PUBLIC_IP_ADDRESS)

metadata() {
  scw-metadata --cached "$1"
}

volume_metadata() {
  device=$1
  key=$2
  scw-metadata --cached "VOLUMES_${device}_${key}"
}


KERNEL_VERSION=$(uname -r)
if [[ $KERNEL_VERSION =~ ^3\.2\.[35][24].* ]]; then
    KERNEL_TITLE="- Marvell"
fi

[ -f /etc/motd.head ] && cat /etc/motd.head || true
printf "\n"
printf "Welcome on %s (%s %s %s %s)\n" "${IMAGE_DESCRIPTION}" "$(uname -o)" "${KERNEL_VERSION}" "$(uname -m)" "$KERNEL_TITLE"
printf "\n"
printf "System information as of: %s\n" "$date"
printf "\n"
printf "System load:\t%s\t\tInt IP Address:\t%s %s\n" $load $ip
printf "Memory usage:\t%s\t\tPub IP Address:\t%s\n" $memory_usage $public_ip
printf "Usage on /:\t%s\t\tSwap usage:\t%s\n" $root_usage $swap_usage
printf "Local Users:\t%s\t\tProcesses:\t%s\n" $users $processes
printf "Image build:\t%s\tSystem uptime:\t%s\n" "${IMAGE_RELEASE}" "$time"
for i in {0..16}; do
  metadata VOLUMES_${i} | grep " " >/dev/null || continue

  SIZE=$(( $(volume_metadata $i SIZE) / 1000 / 1000 / 1000 ))G
  printf "Disk nbd%s:\t%s\n" "$i" "$(volume_metadata $i VOLUME_TYPE) ${SIZE}"
done
printf "\n"
printf "Documentation:\t%s\n" "$IMAGE_DOC_URL"
printf "Community:\t%s\n" "$IMAGE_HELP_URL"
printf "Image source:\t%s\n" "$IMAGE_SOURCE_URL"
printf "\n"
[ -f /etc/motd.tail ] && cat /etc/motd.tail || true
```