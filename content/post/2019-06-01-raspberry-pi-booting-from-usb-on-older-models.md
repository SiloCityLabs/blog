---
title: Raspberry Pi Booting From USB on older models
author: Luis Rodriguez
type: post
date: 2019-06-02T03:56:16+00:00
url: /post/2019/06/01/raspberry-pi-booting-from-usb-on-older-models/
featured_image: /wp-content/uploads/2019/05/Raspberry-Pi-SD-and-USB-150x150.jpg
categories:
  - Side Projects
tags:
  - boot from usb
  - linux
  - os
  - Pi

---
[<img class="aligncenter size-full wp-image-963" src="https://blog.silocitylabs.com/wp-content/uploads/2019/05/Raspberry-Pi-SD-and-USB.jpg" alt="" width="918" height="442" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/05/Raspberry-Pi-SD-and-USB.jpg 918w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/Raspberry-Pi-SD-and-USB-300x144.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/Raspberry-Pi-SD-and-USB-768x370.jpg 768w" sizes="(max-width: 918px) 100vw, 918px" />][1]

Booting from USB is easy on a Pi1, Pi2, or Pi3. This can be useful in mamy cases to speed up the performance of your pi when using a slow sd card. It can also provide a reliable location where you operating system data resides. Sd cards are prone to data corruption from writes a operating system ussually does.

&nbsp;

1. Write the latest Raspbian image to an SD card and boot it.

2. Write the latest Raspbian image to a USB drive and plug it into the Pi.

3. Run blkid to determine the PARTUUID of the USB drive.

4. Edit /boot/cmdline.txt and change root=PARTUUID=xxxxxxxx to match the PARTUUID of the USB drive.

5. Reboot. The Pi should be running from the USB drive.

If you ever need to boot from a different USB drive or from the SD card, it&#8217;s necessary to FIRST mount the boot partition of the SD card (/dev/mmcblk0p1) and edit the boot partition&#8217;s cmdline.txt so that root=PARTUUID=xxxxxxxx matches the device you wish to boot from.

&nbsp;

&nbsp;

 [1]: https://blog.silocitylabs.com/wp-content/uploads/2019/05/Raspberry-Pi-SD-and-USB.jpg