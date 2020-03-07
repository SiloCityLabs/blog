---
title: Raspberry Pi Booting From USB on older models
author: Luis Rodriguez
type: post
date: 2019-06-02T03:56:16+00:00
url: /post/2019/06/01/raspberry-pi-booting-from-usb-on-older-models/
featured_image: /wp-content/uploads/2019/05/Raspberry-Pi-SD-and-USB.jpg
bigimg: [{src: "/uploads/2019/05/Raspberry-Pi-SD-and-USB.jpg", desc: "Pi + USB"}]
categories:
  - Projects
tags:
  - boot from usb
  - linux
  - os
  - Pi
  - raspberry

---

Booting from USB is easy on a Pi1, Pi2, or Pi3. This can be useful in mamy cases to speed up the performance of your pi when using a slow sd card. It can also provide a reliable location where you operating system data resides. Sd cards are prone to data corruption from writes a operating system ussually does.

&nbsp;

 1. Write the latest Raspbian image to an SD card and boot it.
 2. Write the latest Raspbian image to a USB drive and plug it into the Pi.
 3. Boot the raspberry pi and ssh into it.
 4. Get the UUID of your device with `sudo blkid`
    1. If you have duplicate UUID on your devices (most likely) you can run the following command on the usb drive to regenerate the UUID:
       1. `sudo apt install gdisk`
       2. `sudo sgdisk --partition-guid=1:R /dev/sda2`
 5. Create initramfs `mkinitramfs -o /boot/initrd`
 6. Add initramfs `nano /boot/config.txt`
   
```
initramfs initrd
```
 6. Add `rootdelay=5` (or `rootdelay=10` if your drive takes longer to spin up) to `/boot/cmdline.txt`
 7. Change `root=` to use the new device id, you can use any of the following methods
    1. `root=UUID=xxxeeeyy-qbbp-4444-8888-f0f0f0f03333` (best way, most reliable)
    2. `root=PTUUID=XXXXXXXXX`
    3. `root=PARTUUID=XXXX-XXXXX`
    4. `root=/dev/sda2` (not recomended)

```
dwc_otg.lpm_enable=0 console=tty1 root=UUID=xxxeeeyy-qbbp-4444-8888-f0f0f0f03333 rootfstype=ext4 elevator=deadline rootwait rootdelay=5
```

 8. change to UUID or whichever option you used above to `/etc/fstab`

```
proc            /proc           proc    defaults          0       0
/dev/mmcblk0p1  /boot           vfat    defaults          0       2
#/dev/mmcblk0p2  /               ext4    defaults,noatime  0       1
UUID=xxxeeeyy-qbbp-4444-8888-f0f0f0f03333 /     ext4    defaults,noatime       0 1
```

 9. Optinally if you have a higher power device you can increase the output power of usb by adding `max_usb_current=1` to the `/boot/config.txt` file.

 10. Reboot. The Pi should be running from the USB drive.

 11. Bonus: You can gain some usable space by mounting the sd card and usb drive on a linux machine and deleting the unused partitions on those devices.
     1.  Delete root partition on sd card
     2.  Expand boot partition on sd card
     3.  Delete boot partition on usb drive
     4.  Expand root partition on sd card
