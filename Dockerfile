FROM ubuntu:xenial

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y \
    build-essential \
    bc \
    curl \
    cpio \
    golang \
    qemu-system \
  && apt-get clean

WORKDIR /kernel-build

RUN curl -O https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.7.tar.xz \
  && tar xf linux-4.7.tar.xz

RUN cd linux-4.7 \
 && make x86_64_defconfig kvmconfig \
 && make -j16

WORKDIR /initfs/

RUN qemu-img create -f raw disk.img 10M \
  && echo -e "n\np\n1\n\n\nw" | fdisk disk.img \
  && echo "y" | mkfs.ext4 disk.img \
  && mkdir fs

# Setup go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN cd /initfs/fs \
  && mkdir -vm 0755 dev \
  && mkdir -vm 0755 run \
  && mkdir -v proc sys

COPY . $GOPATH/src/github.com/QuentinPerez/busygox

RUN cd $GOPATH/src/github.com/QuentinPerez/busygox \
  && CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o /initfs/fs/init ./cmd/init \
  && cd /initfs/fs && find . -print0 | cpio --null -ov --format=newc > /initfs/initramfs.cpio \
  && ldd init || true

CMD ["/go/src/github.com/QuentinPerez/busygox/entrypoint.sh"]
