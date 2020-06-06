---
title: Recover AMLogic Android boxes everytime
author: Luis Rodriguez
type: post
date: 2020-01-24T00:57:15+00:00
url: /post/2020/01/24/recover-amlogic-android-boxes-everytime
categories:
  - Projects
tags:
  - s905
  - s905x1
  - s905x2
  - s905x3
  - s922
  - s922x1
  - s915
  - amlogic
  - recovery
  - rockchip
  - rk3399
  - rk3328

---

Theres a few ways on the amlogic chips to recover boxes, but theres also many ways to recover these units. I will talk about which I have used and others. In reality it is really difficult to fully brick one of these units permanently. I will start off with the basics and work my way up to the more advanced methods. 

Disclosure:
Everything in this document was sourced from somewhere online. If you want original links please ask me if I havent already posted them in one of the Image Source links. This document does NOT contain any proprietary knowledge that wasnt already sourced from the internet.

<!--more-->

### Recover method 1: Reset button

Yes I know, very basic and im sure most of you have tried it. However some may not know that the reset button has dual purpose and is also hidden cleverly behind the "AV" or "IR" port on some of these boxes. Feel free to shove a paperclip carefully to not miss the button and hit a component. Try it first while its unplugged so you can get the hang of it and not short anything out internally live. You have to hold it down for about 3-4 seconds to trigger recovery.

{{< image src="/uploads/2020/amlogic-recovery/reset.jpg" alt="x96 max reset button">}}
[Image Source](https://thesimplicitypost.com/all-methods-custom-rom-install/)

### Recover method 2: AMLogic Burn tool

Find a downloadable image for your specific device out in the wild. You can also try banggood.com product pages ussually seem to post download links. Then download the [usb burn tool 2.1.7.3](https://downloads.techreanimate.com/kifuob)

Go to File -> import image, Select your image, Hit Start, and finally plugin your device.

It will pickup automatically. You can also hold the reset button if it does not.

{{< image src="/uploads/2020/amlogic-recovery/burn_tool.jpg" alt="usb burn tool">}}

### Recover method 3: Short the NAND chip

This is is risky and very diffucult to spot and document as each device is different. Most new devices uses pads for this while older devices have exposed leads on the old memory chips. Shorting the NAND chip's two pins will throw the device into flash mode on some devices. Take a look at the image below for some examples. This is also called MaskRom mode on many devices including Rockchip devices. You will need to have the USB Burn tool installed and follow the steps from method 2. You can also checkout the [source document](https://forum.freaktab.com/forum/main-category/main-forum/637675-nand-chip-mask-rom-mode-short-location) which contains source images I posted here.

   
1. Open device housing and find the NAND. In some boxes you also need detached heatsink or cooler in order to reach NAND
2. Place PCB in front of you so you can clearly read NAND’s name and board number
3. You should choose 6-7 or 7-8 pin from NAND bottom on the right side or 6-7 or 7-8 pin from first pin (first pin is marked on PCBA with a triangle) Short circuit it while plugging to Host PC and powering on. Check pins for other devices online first.

{{< image src="/uploads/2020/amlogic-recovery/nand.jpg" alt="NAND example 1">}}
{{< image src="/uploads/2020/amlogic-recovery/nand2.jpg" alt="NAND example 2">}}

On this device with a newer chip there are pads on the opposite side of the board right under the NAND itself.

{{< image src="/uploads/2020/amlogic-recovery/nand-pads.jpg" alt="NAND Pads">}}


### Recover method 4: AMLogic SD Burn image

I have not used this method as it shows very little success for my devices. But I am told it works well for others. Using the [Burn card maker](https://downloads.techreanimate.com/hwpjdn) you can write your image to a micro sd card and following any of the methods above will yield you a proper flash.

### Recover method 5: Serial port jtag

You will need to open your device up for this step and locate any pins or unsoldered holes labeled GND/TX/RX/3.3V or sometimes just labeled UART or even JTAG. These will work with any USB to TTL device such as the [CH340](http://a.tra.li/UC9f) or the [CP2102](http://a.tra.li/UC9h) (more reliable). This is my primary means of recovery for any device now. Its reliable and works 100% of the time on AMLogic devices.

{{< image src="/uploads/2020/amlogic-recovery/usb-ttl.png" alt="USB to TTL device" caption="USB to TTL device with custom cable that matches the layout of my device">}}

Once located attach the TTL usb device to it, you can also solder some pins permanently if you are recovering this device often for development.

{{< image src="/uploads/2020/amlogic-recovery/pcb-ttl.jpeg" alt="TTL Header">}}
[Image Source](https://forum.armbian.com/topic/9285-proof-of-concept-realtek-1295/page/5/)

Open Putty using the following settings with your COM number identified in Device manager for Windows.

{{< image src="/uploads/2020/amlogic-recovery/putty.jpg" alt="Putty">}}

With the putty window open mash the enter key while you plug the power into the device. You will essentially stop boot and enter the boot console. You can now go ahead and type `update` and use method 2 above to flash.

{{< image src="/uploads/2020/amlogic-recovery/console.jpg" alt="Console">}}


### Recover method 6: Micro SD card serial Console

Similar to method 5, However I do not own the device in question to post any good images about it. I am told its the same and works withouth opening the box. 

You can see the device here on this [facebook post](https://www.facebook.com/tanix.box/posts/developers-and-geeks-something-interesting-here-debug-now-your-amlogic-device-wi/1210753889057831/) selling it. 


{{< image src="/uploads/2020/amlogic-recovery/debug-device.jpg" alt="debug device">}}

I have not been able to locate one to buy it myself. If you would like to sell me one reach out to me in the comments below.


Know of any other recovery methods?. Let me know below in the comments.
