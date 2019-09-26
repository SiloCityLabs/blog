---
title: Installing nano text editor on Android
author: Luis Rodriguez
type: post
date: 2019-09-26T23:30:40+00:00
url: /post/2019/09/26/installing-nano-android/
categories:
  - Projects
tags:
  - nano
  - android
  - adb
  - shell
  - root
  - su

---

After a while of playing around with my build.prop file in Android I needed a way to edit faster using a text editor via adb shell.

I found a way using a couple of online sources to get it going. You will need to be on a rooted device.

<!--more-->

Download and extract the [nanoforandroid.zip](https://downloads.techreanimate.com/mjksau) to your desktop the run the following commands:

```
adb push nanoforandroid /mnt/sdcard/Download
adb shell
su
mount -o rw,remount /system
cp /sdcard/Download/nanoforandroid/etc/profile /system/etc
cp -r /sdcard/Download/nanoforandroid/etc/terminfo /system/etc
cp /sdcard/Download/nanoforandroid/xbin/nano /system/xbin
chmod 755 -R /system/etc/terminfo
chmod 755 /system/xbin/nano
chmod 755 /system/etc/profile
export TERMINFO=/system/etc/terminfo
export TERM=linux
```

After this you will have a working copy of nano you can use right away. Unfortunately most setups dont have init.d to automatically export the enviormental variables needed so you will need to run the last two lines above every time you reboot. Alternative is to add the lines to a startup shell script somewhere on your system. For example I have one located inside of `/system/bin/rootsudaemon.sh`.
