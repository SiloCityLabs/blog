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

---

Booting from USB is easy on a Pi1, Pi2, or Pi3. This can be useful in mamy cases to speed up the performance of your pi when using a slow sd card. It can also provide a reliable location where you operating system data resides. Sd cards are prone to data corruption from writes a operating system ussually does.

&nbsp;

1. Write the latest Raspbian image to an SD card and boot it.

2. Write the latest Raspbian image to a USB drive and plug it into the Pi.

3. Run `blkid` to determine the PARTUUID of the USB drive.

4. Edit /boot/cmdline.txt and change root=PARTUUID=xxxxxxxx to match the PARTUUID of the USB drive.

5. Reboot. The Pi should be running from the USB drive.

If you ever need to boot from a different USB drive or from the SD card, it's necessary to FIRST mount the boot partition of the SD card (/dev/mmcblk0p1) and edit the boot partition's cmdline.txt so that root=PARTUUID=xxxxxxxx matches the device you wish to boot from.


Update: 02-29-2020

It seems that the partuuid is not automaticaly regenerating on the latest raspbian image. To fix this we will need to generate a new UUID manually.

After the first boot find your usb drive, for example /dev/sda

The partition we want to regenerate is partition #2

```
apt install gdisk
sgdisk --partition-guid=1:R /dev/sda2
```

Now reboot and run `blkid`. Instead of using PARTUUID wil will switch this over to UUID.

Edit /boot/cmdline.txt and change `root=PARTUUID=xxxxxxxx` to match the UUID of the usb drive partition like this `root=UUID=xxxxxxxx`

One laat reboot and you will be all set.

