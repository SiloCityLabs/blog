---
title: BV Tech 4MP 2K H.265 IP67 and Zoneminder
author: Luis Rodriguez
type: post
date: 2018-11-28T23:59:23+00:00
url: /post/2018/11/28/bv-tech-4mp-2k-h-265-ip67-and-zoneminder/
featured_image: /wp-content/uploads/2018/11/61qmqoijgvl._sl1000_-1-150x150.jpg
categories:
  - Side Projects
tags:
  - ca-ipbf-4mp-wifi-fba

---
[<img class="aligncenter size-medium wp-image-672" src="https://blog.silocitylabs.com/wp-content/uploads/2018/11/61qmqoijgvl._sl1000_-300x300.jpg" alt="" width="300" height="300" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/11/61qmqoijgvl._sl1000_-300x300.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/61qmqoijgvl._sl1000_-150x150.jpg 150w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/61qmqoijgvl._sl1000_-768x768.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/61qmqoijgvl._sl1000_.jpg 1000w" sizes="(max-width: 300px) 100vw, 300px" />][1]

Model: [ca-ipbf-4mp-wifi-fba][2]

This problem gave me a few issues I will go over briefly. This was fairly simple once I realized what I was doing wrong. I originally had it plugged in via an Ethernet cable to setup to WiFi like most cameras. But this camera actually just uses the microphone to transmit settings. Once I removed the Ethernet cable everything worked fine. I actually snipped off the Ethernet knowing I wouldn&#8217;t need it, saving space with the clunky cable.

<!--more-->

To get this camera working after much fooling around and talking with support. The old firmware needs to be updated to enable automatic onvif detection. Download the software [SmartPSS][3] then the [firmware &#8220;IP Wifi Bullet Camera&#8221;.][4]

Once you have configured the device in smartpss, there should be an option to upgrade. Do that and it will show up to zoneminder with the onvif selector.

&nbsp;

You can also configure it manually using the following settings:

Source Path: rtsp://admin:password@cameraip:554/cam/realmonitor?channel=1&subtype=0&unicast=true&proto=Onvif

[<img class="aligncenter wp-image-681 size-full" src="https://blog.silocitylabs.com/wp-content/uploads/2018/11/2.png" alt="" width="527" height="464" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/11/2.png 527w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/2-300x264.png 300w" sizes="(max-width: 527px) 100vw, 527px" />][5]

[<img class="aligncenter size-full wp-image-682" src="https://blog.silocitylabs.com/wp-content/uploads/2018/11/2-1.png" alt="" width="526" height="360" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/11/2-1.png 526w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/2-1-300x205.png 300w" sizes="(max-width: 526px) 100vw, 526px" />][6]

&nbsp;

 [1]: https://blog.silocitylabs.com/wp-content/uploads/2018/11/61qmqoijgvl._sl1000_.jpg
 [2]: https://www.bvsecurity.com/index.php/ca-ipbf-4mp-wifi.html
 [3]: https://www.bvsecurity.com/index.php/software
 [4]: https://www.bvsecurity.com/index.php/firmware
 [5]: https://blog.silocitylabs.com/wp-content/uploads/2018/11/2.png
 [6]: https://blog.silocitylabs.com/wp-content/uploads/2018/11/2-1.png