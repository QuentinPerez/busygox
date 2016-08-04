FROM ubuntu:xenial

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y \
    build-essential \
	bc \
	libguestfs-tools \
  && apt-get clean

WORKDIR /kernel-build

RUN curl -O https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.7.tar.xz \
  && tar xf linux-4.7.tar.xz

RUN cd linux-4.7 \
 && make x86_64_defconfig kvmconfig \
 && make -j16



RUN apt-get install -y golang

COPY ./src /initfs
WORKDIR /initfs

RUN mkdir fs \
  && go build -o fs/init main.go \
  && cd fs && find . -print0 | cpio --null -ov --format=newc > ../initramfs.cpio \
  && ls ..

CMD ["cp", "/kernel-build/linux-4.7/arch/x86_64/boot/bzImage", "/initfs/initramfs.cpio", "/kernel"]
