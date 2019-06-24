---
title: BV Tech 4MP 2K H.265 IP67 and Zoneminder
author: Luis Rodriguez
type: post
date: 2018-11-28T23:59:23+00:00
url: /post/2018/11/28/bv-tech-4mp-2k-h-265-ip67-and-zoneminder/
featured_image: /wp-content/uploads/2018/11/61qmqoijgvl._sl1000_-1.jpg
categories:
  - Projects
tags:
  - ca-ipbf-4mp-wifi-fba

---
![](/uploads/2018/11/61qmqoijgvl._sl1000_.jpg)

Model: [ca-ipbf-4mp-wifi-fba][2]

This problem gave me a few issues I will go over briefly. This was fairly simple once I realized what I was doing wrong. I originally had it plugged in via an Ethernet cable to setup to WiFi like most cameras. But this camera actually just uses the microphone to transmit settings. Once I removed the Ethernet cable everything worked fine. I actually snipped off the Ethernet knowing I wouldn't need it, saving space with the clunky cable.

<!--more-->

To get this camera working after much fooling around and talking with support. The old firmware needs to be updated to enable automatic onvif detection. Download the software [SmartPSS][3] then the [firmware &#8220;IP Wifi Bullet Camera&#8221;.][4]

Once you have configured the device in smartpss, there should be an option to upgrade. Do that and it will show up to zoneminder with the onvif selector.

&nbsp;

You can also configure it manually using the following settings:

Source Path: rtsp://admin:password@cameraip:554/cam/realmonitor?channel=1&subtype=0&unicast=true&proto=Onvif

![](/uploads/2018/11/2.png)

![](/uploads/2018/11/2-1.png)

&nbsp;

 [2]: https://www.bvsecurity.com/index.php/ca-ipbf-4mp-wifi.html
 [3]: https://www.bvsecurity.com/index.php/software
 [4]: https://www.bvsecurity.com/index.php/firmware