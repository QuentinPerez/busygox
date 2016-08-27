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
[    0.000000] Linux version 4.7.0 (root@827f45722fd6) (gcc version 5.4.0 20160609 (Ubuntu 5.4.0-6ubuntu1~16.04.1) ) #1 SMP Fri Aug 5 16:49:59 UTC 2016
[    0.000000] Command line: console=ttyS0 nousb root=/dev/sda1 rootfstype=ext4 selinux=0
[    0.000000] x86/fpu: Legacy x87 FPU detected.
[    0.000000] x86/fpu: Using 'eager' FPU context switches.
[    0.000000] e820: BIOS-provided physical RAM map:
[    0.000000] BIOS-e820: [mem 0x0000000000000000-0x000000000009fbff] usable
[    0.000000] BIOS-e820: [mem 0x000000000009fc00-0x000000000009ffff] reserved
[    0.000000] BIOS-e820: [mem 0x00000000000f0000-0x00000000000fffff] reserved
[    0.000000] BIOS-e820: [mem 0x0000000000100000-0x000000003ffdffff] usable
[    0.000000] BIOS-e820: [mem 0x000000003ffe0000-0x000000003fffffff] reserved
[    0.000000] BIOS-e820: [mem 0x00000000fffc0000-0x00000000ffffffff] reserved
[    0.000000] NX (Execute Disable) protection: active
[    0.000000] SMBIOS 2.8 present.
[    0.000000] e820: last_pfn = 0x3ffe0 max_arch_pfn = 0x400000000
[    0.000000] x86/PAT: Configuration [0-7]: WB  WC  UC- UC  WB  WC  UC- WT
[    0.000000] found SMP MP-table at [mem 0x000f6be0-0x000f6bef] mapped at [ffff8800000f6be0]
[    0.000000] Scanning 1 areas for low memory corruption
[    0.000000] RAMDISK: [mem 0x3fe37000-0x3ffdffff]
[    0.000000] ACPI: Early table checksum verification disabled
[    0.000000] ACPI: RSDP 0x00000000000F69F0 000014 (v00 BOCHS )
[    0.000000] ACPI: RSDT 0x000000003FFE1737 000030 (v01 BOCHS  BXPCRSDT 00000001 BXPC 00000001)
[    0.000000] ACPI: FACP 0x000000003FFE1613 000074 (v01 BOCHS  BXPCFACP 00000001 BXPC 00000001)
[    0.000000] ACPI: DSDT 0x000000003FFE0040 0015D3 (v01 BOCHS  BXPCDSDT 00000001 BXPC 00000001)
[    0.000000] ACPI: FACS 0x000000003FFE0000 000040
[    0.000000] ACPI: APIC 0x000000003FFE1687 000078 (v01 BOCHS  BXPCAPIC 00000001 BXPC 00000001)
[    0.000000] ACPI: HPET 0x000000003FFE16FF 000038 (v01 BOCHS  BXPCHPET 00000001 BXPC 00000001)
[    0.000000] No NUMA configuration found
[    0.000000] Faking a node at [mem 0x0000000000000000-0x000000003ffdffff]
[    0.000000] NODE_DATA(0) allocated [mem 0x3fe33000-0x3fe36fff]
[    0.000000] Zone ranges:
[    0.000000]   DMA      [mem 0x0000000000001000-0x0000000000ffffff]
[    0.000000]   DMA32    [mem 0x0000000001000000-0x000000003ffdffff]
[    0.000000]   Normal   empty
[    0.000000] Movable zone start for each node
[    0.000000] Early memory node ranges
[    0.000000]   node   0: [mem 0x0000000000001000-0x000000000009efff]
[    0.000000]   node   0: [mem 0x0000000000100000-0x000000003ffdffff]
[    0.000000] Initmem setup node 0 [mem 0x0000000000001000-0x000000003ffdffff]
[    0.000000] ACPI: PM-Timer IO Port: 0x608
[    0.000000] ACPI: LAPIC_NMI (acpi_id[0xff] dfl dfl lint[0x1])
[    0.000000] IOAPIC[0]: apic_id 0, version 17, address 0xfec00000, GSI 0-23
[    0.000000] ACPI: INT_SRC_OVR (bus 0 bus_irq 0 global_irq 2 dfl dfl)
[    0.000000] ACPI: INT_SRC_OVR (bus 0 bus_irq 5 global_irq 5 high level)
[    0.000000] ACPI: INT_SRC_OVR (bus 0 bus_irq 9 global_irq 9 high level)
[    0.000000] ACPI: INT_SRC_OVR (bus 0 bus_irq 10 global_irq 10 high level)
[    0.000000] ACPI: INT_SRC_OVR (bus 0 bus_irq 11 global_irq 11 high level)
[    0.000000] Using ACPI (MADT) for SMP configuration information
[    0.000000] ACPI: HPET id: 0x8086a201 base: 0xfed00000
[    0.000000] smpboot: Allowing 1 CPUs, 0 hotplug CPUs
[    0.000000] PM: Registered nosave memory: [mem 0x00000000-0x00000fff]
[    0.000000] PM: Registered nosave memory: [mem 0x0009f000-0x0009ffff]
[    0.000000] PM: Registered nosave memory: [mem 0x000a0000-0x000effff]
[    0.000000] PM: Registered nosave memory: [mem 0x000f0000-0x000fffff]
[    0.000000] e820: [mem 0x40000000-0xfffbffff] available for PCI devices
[    0.000000] Booting paravirtualized kernel on bare hardware
[    0.000000] clocksource: refined-jiffies: mask: 0xffffffff max_cycles: 0xffffffff, max_idle_ns: 1910969940391419 ns
[    0.000000] setup_percpu: NR_CPUS:64 nr_cpumask_bits:64 nr_cpu_ids:1 nr_node_ids:1
[    0.000000] percpu: Embedded 33 pages/cpu @ffff88003fc00000 s95576 r8192 d31400 u2097152
[    0.000000] Built 1 zonelists in Node order, mobility grouping on.  Total pages: 257897
[    0.000000] Policy zone: DMA32
[    0.000000] Kernel command line: console=ttyS0 nousb root=/dev/sda1 rootfstype=ext4 selinux=0
[    0.000000] PID hash table entries: 4096 (order: 3, 32768 bytes)
[    0.000000] Memory: 1011940K/1048056K available (9187K kernel code, 1197K rwdata, 2816K rodata, 1176K init, 936K bss, 36116K reserved, 0K cma-reserved)
[    0.000000] SLUB: HWalign=64, Order=0-3, MinObjects=0, CPUs=1, Nodes=1
[    0.000000] Hierarchical RCU implementation.
[    0.000000] 	Build-time adjustment of leaf fanout to 64.
[    0.000000] 	RCU restricting CPUs from NR_CPUS=64 to nr_cpu_ids=1.
[    0.000000] RCU: Adjusting geometry for rcu_fanout_leaf=64, nr_cpu_ids=1
[    0.000000] NR_IRQS:4352 nr_irqs:256 16
[    0.000000] Console: colour VGA+ 80x25
[    0.000000] console [ttyS0] enabled
[    0.000000] clocksource: hpet: mask: 0xffffffff max_cycles: 0xffffffff, max_idle_ns: 19112604467 ns
[    0.000000] tsc: Unable to calibrate against PIT
[    0.000000] tsc: using HPET reference calibration
[    0.000000] tsc: Detected 2494.223 MHz processor
[    0.008000] Calibrating delay loop (skipped), value calculated using timer frequency.. 4988.44 BogoMIPS (lpj=2494223)
[    0.008229] pid_max: default: 32768 minimum: 301
[    0.008753] ACPI: Core revision 20160422
[    0.026815] ACPI: 1 ACPI AML tables successfully acquired and loaded
[    0.026977]
[    0.027577] Security Framework initialized
[    0.027743] SELinux:  Disabled at boot.
[    0.028862] Dentry cache hash table entries: 131072 (order: 8, 1048576 bytes)
[    0.030693] Inode-cache hash table entries: 65536 (order: 7, 524288 bytes)
[    0.031647] Mount-cache hash table entries: 2048 (order: 2, 16384 bytes)
[    0.031820] Mountpoint-cache hash table entries: 2048 (order: 2, 16384 bytes)
[    0.041386] mce: CPU supports 10 MCE banks
[    0.042445] Last level iTLB entries: 4KB 0, 2MB 0, 4MB 0
[    0.042600] Last level dTLB entries: 4KB 0, 2MB 0, 4MB 0, 1GB 0
[    0.266000] Freeing SMP alternatives memory: 36K (ffffffff82053000 - ffffffff8205c000)
[    0.279373] smpboot: Max logical packages: 1
[    0.279631] smpboot: APIC(0) Converting physical 0 to logical package 0
[    0.283000] ..TIMER: vector=0x30 apic1=0 pin1=2 apic2=-1 pin2=-1
[    0.292000] APIC calibration not consistent with PM-Timer: 128ms instead of 100ms
[    0.292000] APIC delta adjusted to PM-Timer: 6249941 (8029500)
[    0.292000] smpboot: CPU0: AMD QEMU Virtual CPU version 2.5+ (family: 0x6, model: 0x6, stepping: 0x3)
[    0.292000] Performance Events: Broken PMU hardware detected, using software events only.
[    0.292429] Failed to access perfctr msr (MSR c0010007 is 0)
[    0.298860] Huh? What family is it: 0x6?!
[    0.299362] x86: Booted up 1 node, 1 CPUs
[    0.299538] smpboot: Total of 1 processors activated (4988.44 BogoMIPS)
[    0.307059] devtmpfs: initialized
[    0.311509] clocksource: jiffies: mask: 0xffffffff max_cycles: 0xffffffff, max_idle_ns: 1911260446275000 ns
[    0.312589] RTC time:  0:01:22, date: 08/27/16
[    0.318000] kworker/u2:0 (15) used greatest stack depth: 14832 bytes left
[    0.320613] NET: Registered protocol family 16
[    0.325406] cpuidle: using governor menu
[    0.326593] kworker/u2:1 (19) used greatest stack depth: 13944 bytes left
[    0.327337] ACPI: bus type PCI registered
[    0.329000] PCI: Using configuration type 1 for base access
[    0.413135] HugeTLB registered 2 MB page size, pre-allocated 0 pages
[    0.416123] ACPI: Added _OSI(Module Device)
[    0.416238] ACPI: Added _OSI(Processor Device)
[    0.416326] ACPI: Added _OSI(3.0 _SCP Extensions)
[    0.416420] ACPI: Added _OSI(Processor Aggregator Device)
[    0.432959] ACPI: Interpreter enabled
[    0.433721] ACPI: (supports S0 S3 S4 S5)
[    0.433834] ACPI: Using IOAPIC for interrupt routing
[    0.434579] PCI: Using host bridge windows from ACPI; if necessary, use "pci=nocrs" and report a bug
[    0.481000] ACPI: PCI Root Bridge [PCI0] (domain 0000 [bus 00-ff])
[    0.481719] acpi PNP0A03:00: _OSC: OS supports [ASPM ClockPM Segments MSI]
[    0.482000] acpi PNP0A03:00: _OSC failed (AE_NOT_FOUND); disabling ASPM
[    0.482471] acpi PNP0A03:00: fail to add MMCONFIG information, can't access extended PCI configuration space under this bridge.
[    0.484794] PCI host bridge to bus 0000:00
[    0.484993] pci_bus 0000:00: root bus resource [io  0x0000-0x0cf7 window]
[    0.485095] pci_bus 0000:00: root bus resource [io  0x0d00-0xffff window]
[    0.485255] pci_bus 0000:00: root bus resource [mem 0x000a0000-0x000bffff window]
[    0.485393] pci_bus 0000:00: root bus resource [mem 0x40000000-0xfebfffff window]
[    0.485634] pci_bus 0000:00: root bus resource [bus 00-ff]
[    0.493092] pci 0000:00:01.1: legacy IDE quirk: reg 0x10: [io  0x01f0-0x01f7]
[    0.493267] pci 0000:00:01.1: legacy IDE quirk: reg 0x14: [io  0x03f6]
[    0.493403] pci 0000:00:01.1: legacy IDE quirk: reg 0x18: [io  0x0170-0x0177]
[    0.493535] pci 0000:00:01.1: legacy IDE quirk: reg 0x1c: [io  0x0376]
[    0.495420] pci 0000:00:01.3: quirk: [io  0x0600-0x063f] claimed by PIIX4 ACPI
[    0.495587] pci 0000:00:01.3: quirk: [io  0x0700-0x070f] claimed by PIIX4 SMB
[    0.518323] ACPI: PCI Interrupt Link [LNKA] (IRQs 5 *10 11)
[    0.519303] ACPI: PCI Interrupt Link [LNKB] (IRQs 5 *10 11)
[    0.519958] ACPI: PCI Interrupt Link [LNKC] (IRQs 5 10 *11)
[    0.520482] ACPI: PCI Interrupt Link [LNKD] (IRQs 5 10 *11)
[    0.520875] ACPI: PCI Interrupt Link [LNKS] (IRQs *9)
[    0.523000] ACPI: Enabled 16 GPEs in block 00 to 0F
[    0.525930] vgaarb: setting as boot device: PCI:0000:00:02.0
[    0.526000] vgaarb: device added: PCI:0000:00:02.0,decodes=io+mem,owns=io+mem,locks=none
[    0.526032] vgaarb: loaded
[    0.526156] vgaarb: bridge control possible 0000:00:02.0
[    0.528000] SCSI subsystem initialized
[    0.530387] ACPI: bus type USB registered
[    0.531173] usbcore: registered new interface driver usbfs
[    0.531638] usbcore: registered new interface driver hub
[    0.532000] usbcore: registered new device driver usb
[    0.533149] pps_core: LinuxPPS API ver. 1 registered
[    0.533260] pps_core: Software ver. 5.3.6 - Copyright 2005-2007 Rodolfo Giometti <giometti@linux.it>
[    0.534908] PTP clock support registered
[    0.537066] Advanced Linux Sound Architecture Driver Initialized.
[    0.537768] PCI: Using ACPI for IRQ routing
[    0.547667] NetLabel: Initializing
[    0.547789] NetLabel:  domain hash size = 128
[    0.547874] NetLabel:  protocols = UNLABELED CIPSOv4
[    0.548601] NetLabel:  unlabeled traffic allowed by default
[    0.550671] HPET: 3 timers in total, 0 timers will be used for per-cpu timer
[    0.551000] hpet0: at MMIO 0xfed00000, IRQs 2, 8, 0
[    0.551236] hpet0: 3 comparators, 64-bit 100.000000 MHz counter
[    0.554538] amd_nb: Cannot enumerate AMD northbridges
[    0.554799] clocksource: Switched to clocksource hpet
[    0.641498] VFS: Disk quotas dquot_6.6.0
[    0.641795] VFS: Dquot-cache hash table entries: 512 (order 0, 4096 bytes)
[    0.644626] pnp: PnP ACPI init
[    0.653091] pnp: PnP ACPI: found 6 devices
[    0.669361] kworker/u2:1 (553) used greatest stack depth: 13936 bytes left
[    0.697687] clocksource: acpi_pm: mask: 0xffffff max_cycles: 0xffffff, max_idle_ns: 2085701024 ns
[    0.699855] NET: Registered protocol family 2
[    0.704906] TCP established hash table entries: 8192 (order: 4, 65536 bytes)
[    0.705624] TCP bind hash table entries: 8192 (order: 5, 131072 bytes)
[    0.706014] TCP: Hash tables configured (established 8192 bind 8192)
[    0.706786] UDP hash table entries: 512 (order: 2, 16384 bytes)
[    0.707207] UDP-Lite hash table entries: 512 (order: 2, 16384 bytes)
[    0.708528] NET: Registered protocol family 1
[    0.710449] RPC: Registered named UNIX socket transport module.
[    0.710595] RPC: Registered udp transport module.
[    0.710693] RPC: Registered tcp transport module.
[    0.710782] RPC: Registered tcp NFSv4.1 backchannel transport module.
[    0.710984] pci 0000:00:00.0: Limiting direct PCI/PCI transfers
[    0.711148] pci 0000:00:01.0: PIIX3: Enabling Passive Release
[    0.711342] pci 0000:00:01.0: Activating ISA DMA hang workarounds
[    0.711828] pci 0000:00:02.0: Video device with shadowed ROM at [mem 0x000c0000-0x000dffff]
[    0.715736] Unpacking initramfs...
[    0.726288] Freeing initrd memory: 1700K (ffff88003fe37000 - ffff88003ffe0000)
[    0.730115] Scanning for low memory corruption every 60 seconds
[    0.736693] futex hash table entries: 256 (order: 2, 16384 bytes)
[    0.737751] audit: initializing netlink subsys (disabled)
[    0.738547] audit: type=2000 audit(1472256082.737:1): initialized
[    0.745415] workingset: timestamp_bits=54 max_order=18 bucket_order=0
[    0.790058] NFS: Registering the id_resolver key type
[    0.792380] Key type id_resolver registered
[    0.792516] Key type id_legacy registered
[    0.793742] 9p: Installing v9fs 9p2000 file system support
[    0.804950] Block layer SCSI generic (bsg) driver version 0.4 loaded (major 251)
[    0.805353] io scheduler noop registered
[    0.805480] io scheduler deadline registered
[    0.805872] io scheduler cfq registered (default)
[    0.808184] pci_hotplug: PCI Hot Plug PCI Core version: 0.5
[    0.811193] input: Power Button as /devices/LNXSYSTM:00/LNXPWRBN:00/input/input0
[    0.812187] ACPI: Power Button [PWRF]
[    0.816619] Serial: 8250/16550 driver, 4 ports, IRQ sharing enabled
[    0.839994] 00:05: ttyS0 at I/O 0x3f8 (irq = 4, base_baud = 115200) is a 16550A
[    0.850534] Non-volatile memory driver v1.3
[    0.851079] Linux agpgart interface v0.103
[    0.855355] [drm] Initialized drm 1.1.0 20060810
[    0.874001] loop: module loaded
[    0.887843] scsi host0: ata_piix
[    0.890494] scsi host1: ata_piix
[    0.891265] ata1: PATA max MWDMA2 cmd 0x1f0 ctl 0x3f6 bmdma 0xc040 irq 14
[    0.891412] ata2: PATA max MWDMA2 cmd 0x170 ctl 0x376 bmdma 0xc048 irq 15
[    0.896463] e100: Intel(R) PRO/100 Network Driver, 3.5.24-k2-NAPI
[    0.896561] e100: Copyright(c) 1999-2006 Intel Corporation
[    0.896561] e1000: Intel(R) PRO/1000 Network Driver - version 7.3.21-k8-NAPI
[    0.897279] e1000: Copyright (c) 1999-2006 Intel Corporation.
[    1.050516] ata1.00: ATA-7: QEMU HARDDISK, 2.5+, max UDMA/100
[    1.050671] ata1.00: 1024000 sectors, multi 16: LBA48
[    1.052194] ata2.00: ATAPI: QEMU DVD-ROM, 2.5+, max UDMA/100
[    1.053272] ata2.00: configured for MWDMA2
[    1.058942] ata1.00: configured for MWDMA2
[    1.070259] scsi 0:0:0:0: Direct-Access     ATA      QEMU HARDDISK    2.5+ PQ: 0 ANSI: 5
[    1.077698] sd 0:0:0:0: [sda] 1024000 512-byte logical blocks: (524 MB/500 MiB)
[    1.079425] sd 0:0:0:0: [sda] Write Protect is off
[    1.079822] sd 0:0:0:0: [sda] Write cache: enabled, read cache: enabled, doesn't support DPO or FUA
[    1.085073] sd 0:0:0:0: Attached scsi generic sg0 type 0
[    1.092870] scsi 1:0:0:0: CD-ROM            QEMU     QEMU DVD-ROM     2.5+ PQ: 0 ANSI: 5
[    1.107364] sr 1:0:0:0: [sr0] scsi3-mmc drive: 4x/4x cd/rw xa/form2 tray
[    1.107740] cdrom: Uniform CD-ROM driver Revision: 3.20
[    1.111018] sr 1:0:0:0: Attached scsi generic sg1 type 5
[    1.121609] sd 0:0:0:0: [sda] Attached SCSI disk
[    1.230209] ACPI: PCI Interrupt Link [LNKC] enabled at IRQ 11
[    1.512410] e1000 0000:00:03.0 eth0: (PCI:33MHz:32-bit) 52:54:00:12:34:56
[    1.512739] e1000 0000:00:03.0 eth0: Intel(R) PRO/1000 Network Connection
[    1.513747] e1000e: Intel(R) PRO/1000 Network Driver - 3.2.6-k
[    1.513874] e1000e: Copyright(c) 1999 - 2015 Intel Corporation.
[    1.514268] sky2: driver version 1.30
[    1.516692] ehci_hcd: USB 2.0 'Enhanced' Host Controller (EHCI) Driver
[    1.516820] ehci-pci: EHCI PCI platform driver
[    1.517409] ohci_hcd: USB 1.1 'Open' Host Controller (OHCI) Driver
[    1.517725] ohci-pci: OHCI PCI platform driver
[    1.518006] uhci_hcd: USB Universal Host Controller Interface driver
[    1.520020] usbcore: registered new interface driver usblp
[    1.520443] usbcore: registered new interface driver usb-storage
[    1.521596] i8042: PNP: PS/2 Controller [PNP0303:KBD,PNP0f13:MOU] at 0x60,0x64 irq 1,12
[    1.525059] serio: i8042 KBD port at 0x60,0x64 irq 1
[    1.525463] serio: i8042 AUX port at 0x60,0x64 irq 12
[    1.527598] mousedev: PS/2 mouse device common for all mice
[    1.530956] input: AT Translated Set 2 keyboard as /devices/platform/i8042/serio0/input/input1
[    1.536055] rtc_cmos 00:00: RTC can wake from S4
[    1.540819] rtc_cmos 00:00: rtc core: registered rtc_cmos as rtc0
[    1.542237] rtc_cmos 00:00: alarms up to one day, y3k, 114 bytes nvram, hpet irqs
[    1.545182] device-mapper: ioctl: 4.34.0-ioctl (2015-10-28) initialised: dm-devel@redhat.com
[    1.546221] hidraw: raw HID events driver (C) Jiri Kosina
[    1.554664] usbcore: registered new interface driver usbhid
[    1.554801] usbhid: USB HID core driver
[    1.564305] Netfilter messages via NETLINK v0.30.
[    1.566792] nf_conntrack version 0.5.0 (8192 buckets, 32768 max)
[    1.568017] ctnetlink v0.93: registering with nfnetlink.
[    1.572840] ip_tables: (C) 2000-2006 Netfilter Core Team
[    1.575068] Initializing XFRM netlink socket
[    1.578144] NET: Registered protocol family 10
[    1.585317] ip6_tables: (C) 2000-2006 Netfilter Core Team
[    1.588780] sit: IPv6 over IPv4 tunneling driver
[    1.592464] NET: Registered protocol family 17
[    1.594342] 9pnet: Installing 9P2000 support
[    1.594850] Key type dns_resolver registered
[    1.597109] microcode: AMD CPU family 0x6 not supported
[    1.600108] registered taskstats version 1
[    1.604400]   Magic number: 12:0:2
[    1.604874] console [netcon0] enabled
[    1.605240] netconsole: network logging started
[    1.607383] ALSA device list:
[    1.607386]   No soundcards found.
[    1.656671] Freeing unused kernel memory: 1176K (ffffffff81f2d000 - ffffffff82053000)
[    1.657091] Write protecting the kernel read-only data: 14336k
[    1.659512] Freeing unused kernel memory: 1036K (ffff8800018fd000 - ffff880001a00000)
[    1.680105] Freeing unused kernel memory: 1280K (ffff880001cc0000 - ffff880001e00000)
[    1.730091] tsc: Refined TSC clocksource calibration: 2494.226 MHz
[    1.730272] clocksource: tsc: mask: 0xffffffffffffffff max_cycles: 0x23f3eb3743e, max_idle_ns: 440795257847 ns


_____
|  _ \
| |_) |_   _ ___ _   _  __ _  _____  __
|  _ <| | | / __| | | |/ _` |/ _ \ \/ /
| |_) | |_| \__ \ |_| | (_| | (_) >  <
|____/ \__,_|___/\__, |\__, |\___/_/\_\
                  __/ | __/ |
                 |___/ |___/


[    1.819013] EXT4-fs (sda): recovery complete
[    1.820683] EXT4-fs (sda): mounted filesystem with ordered data mode. Opts: (null)
[    2.065915] random: systemd urandom read with 73 bits of entropy available
[    2.091175] systemd[1]: systemd 229 running in system mode. (+PAM +AUDIT +SELINUX +IMA +APPARMOR +SMACK +SYSVINIT +UTMP +LIBCRYPTSETUP +GCRYPT +GNUTLS +ACL +XZ -LZ4 +SECCOMP +BLKID +ELFUTILS +KMOD -IDN)
[    2.093394] systemd[1]: Detected virtualization qemu.
[    2.093655] systemd[1]: Detected architecture x86-64.

Welcome to Ubuntu 16.04.1 LTS!

[    2.096991] systemd[1]: No hostname configured.
[    2.097696] systemd[1]: Set hostname to <localhost>.
[    2.342447] input: ImExPS/2 BYD TouchPad as /devices/platform/i8042/serio1/input/input3
[    2.498722] systemd-dbus1-g (973) used greatest stack depth: 13760 bytes left
[    2.569008] systemd-hiberna (977) used greatest stack depth: 13448 bytes left
[    2.572536] random: nonblocking pool is initialized
[    2.592038] systemd-rc-loca (980) used greatest stack depth: 13256 bytes left
[    2.732402] clocksource: Switched to clocksource tsc
[    2.916021] systemd[1]: Listening on udev Kernel Socket.
[  OK  ] Listening on udev Kernel Socket.
[    2.921352] systemd[1]: Listening on Journal Audit Socket.
[  OK  ] Listening on Journal Audit Socket.
[    2.923565] systemd[1]: Created slice System Slice.
[  OK  ] Created slice System Slice.
[    2.924451] systemd[1]: Created slice system-getty.slice.
[  OK  ] Created slice system-getty.slice.
[    2.925845] systemd[1]: Reached target Network is Online.
[  OK  ] Reached target Network is Online.
[    2.930588] systemd[1]: Set up automount Arbitrary Executable File Formats File System Automount Point.
[  OK  ] Set up automount Arbitrary Executab...ats File System Automount Point.
[    2.932521] systemd[1]: Listening on Journal Socket.
[  OK  ] Listening on Journal Socket.
[    2.942286] systemd[1]: Starting SCW kernel requirements checker...
         Starting SCW kernel requirements checker...
[    2.945668] systemd[1]: Started Forward Password Requests to Wall Directory Watch.
[  OK  ] Started Forward Password Requests to Wall Directory Watch.
[    2.957318] systemd[1]: Mounting POSIX Message Queue File System...
         Mounting POSIX Message Queue File System...
[    2.976165] systemd[1]: Starting Remount Root and Kernel File Systems...
         Starting Remount Root and Kernel File Systems...
[    2.984577] systemd[1]: Listening on udev Control Socket.
[  OK  ] Listening on udev Control Socket.
[    2.991040] systemd[1]: Started Dispatch Password Requests to Console Directory Watch.
[  OK  ] Started Dispatch Password Requests to Console Directory Watch.
[    2.997087] systemd[1]: Reached target Remote File Systems (Pre).
[  OK  ] Reached target Remote File Systems (Pre).
[    3.002117] systemd[1]: Reached target Remote File Systems.
[  OK  ] Reached target Remote File Systems.
[    3.012835] systemd[1]: Reached target Paths.
[  OK  ] Reached target Paths.
[    3.018239] systemd[1]: Listening on Syslog Socket.
[  OK  ] Listening on Syslog Socket.
[    3.024078] systemd[1]: Listening on /dev/initctl Compatibility Named Pipe.
[  OK  ] Listening on /dev/initctl Compatibility Named Pipe.
[    3.029685] systemd[1]: Created slice system-serial\x2dgetty.slice.
[  OK  ] Created slice system-serial\x2dgetty.slice.
[    3.036728] systemd[1]: Reached target Encrypted Volumes.
[  OK  ] Reached target Encrypted Volumes.
[    3.062324] systemd[1]: Mounting Huge Pages File System...
         Mounting Huge Pages File System...
[    3.068870] systemd[1]: Reached target Swap.
[  OK  ] Reached target Swap.
[    3.077350] systemd[1]: Created slice User and Session Slice.
[  OK  ] Created slice User and Session Slice.
[    3.115113] systemd[1]: Starting Load Kernel Modules...
         Starting Load Kernel Modules...
[    3.149310] systemd[1]: Mounting Debug File System...
         Mounting Debug File System...
[    3.190430] systemd[1]: Starting Create Static Device Nodes in /dev...
         Starting Create Static Device Nodes in /dev...
[    3.210183] systemd[1]: Reached target Slices.
[  OK  ] Reached target Slices.
[    3.221571] systemd[1]: Listening on Journal Socket (/dev/log).
[  OK  ] Listening on Journal Socket (/dev/log).
[    3.324207] systemd[1]: Mounted POSIX Message Queue File System.
[  OK  ] Mounted POSIX Message Queue File System.
[    3.385204] systemd[1]: Started Remount Root and Kernel File Systems.
[  OK  ] Started Remount Root and Kernel File Systems.
[    3.465050] systemd[1]: Mounted Huge Pages File System.
[  OK  ] Mounted Huge Pages File System.
[    3.526958] systemd[1]: Starting Load/Save Random Seed...
         Starting Load/Save Random Seed...
[    3.587037] systemd[1]: Starting udev Coldplug all Devices...
         Starting udev Coldplug all Devices...
[    3.648750] systemd[1]: Mounted Debug File System.
[  OK  ] Mounted Debug File System.
[    3.670517] systemd[1]: Started Load Kernel Modules.
[  OK  ] Started Load Kernel Modules.
[    4.015636] systemd[1]: Started Create Static Device Nodes in /dev.
[  OK  ] Started Create Static Device Nodes in /dev.
[    4.036503] systemd[1]: Started Load/Save Random Seed.
[  OK  ] Started Load/Save Random Seed.
[    4.116018] grep (1000) used greatest stack depth: 12856 bytes left
[    4.118369] scw-check-kerne (986) used greatest stack depth: 12600 bytes left
[    4.164397] systemd[1]: Started SCW kernel requirements checker.
[  OK  ] Started SCW kernel requirements checker.
[    4.587281] systemd[1]: Started Entropy daemon using the HAVEGE algorithm.
[  OK  ] Started Entropy daemon using the HAVEGE algorithm.
[    4.607854] systemd[1]: Starting Journal Service...
         Starting Journal Service...
[    4.608538] systemd[1]: Reached target Local File Systems (Pre).
[  OK  ] Reached target Local File Systems (Pre).
[    4.617290] systemd[1]: Reached target Local File Systems.
[  OK  ] Reached target Local File Systems.
[    4.640300] systemd[1]: Starting udev Kernel Device Manager...
         Starting udev Kernel Device Manager...
[    4.675841] systemd[1]: Starting Apply Kernel Variables...
         Starting Apply Kernel Variables...
[  OK  ] Started Apply Kernel Variables.
[  OK  ] Started udev Kernel Device Manager.
[  OK  ] Started Journal Service.
         Starting Flush Journal to Persistent Storage...
[  OK  ] Started Flush Journal to Persistent Storage.
         Starting Create Volatile Files and Directories...
[  OK  ] Started udev Coldplug all Devices.
[  OK  ] Started Create Volatile Files and Directories.
         Starting Update UTMP about System Boot/Shutdown...
[  OK  ] Reached target System Time Synchronized.
[  OK  ] Started Update UTMP about System Boot/Shutdown.
[  OK  ] Reached target System Initialization.
[  OK  ] Listening on D-Bus System Message Bus Socket.
[  OK  ] Started Daily apt activities.
[  OK  ] Started Daily Cleanup of Temporary Directories.
[  OK  ] Reached target Timers.
[  OK  ] Listening on UUID daemon activation socket.
[  OK  ] Reached target Sockets.
[  OK  ] Reached target Basic System.
         Starting LSB: Start NTP daemon...
         Starting Login Service...
         Starting SCW generate ssh keys on first boot...
         Starting Permit User Sessions...
         Starting LSB: Start/stop sysstat's sadc...
[  OK  ] Started Regular background program processing daemon.
         Starting SCW fetch kernel modules from Scaleway mirror...
         Starting /etc/rc.local Compatibility...
         Starting LSB: Set the CPU Frequency Scaling governor to "ondemand"...
         Starting SCW fetch ssh keys from metadata...
[  OK  ] Started D-Bus System Message Bus.
         Starting System Logging Service...
[  OK  ] Started Permit User Sessions.
[  OK  ] Started /etc/rc.local Compatibility.
[  OK  ] Started System Logging Service.
[  OK  ] Started Login Service.
[  OK  ] Started Getty on tty1.
[  OK  ] Started LSB: Start/stop sysstat's sadc.
[  OK  ] Started LSB: Set the CPU Frequency Scaling governor to "ondemand".
[FAILED] Failed to start SCW fetch kernel modules from Scaleway mirror.
See 'systemctl status scw-sync-kernel-modules.service' for details.
         Starting SCW generate machine id...
[  OK  ] Started SCW generate machine id.
[  OK  ] Started LSB: Start NTP daemon.
[  OK  ] Found device /dev/ttyS0.
[  OK  ] Started Serial Getty on ttyS0.
[  OK  ] Reached target Login Prompts.

Ubuntu 16.04 LTS localhost ttyS0

localhost login:
```

#### Good to know

This repository is for educational use only.

#### Development

Feel free to contribute 😃 🍻
