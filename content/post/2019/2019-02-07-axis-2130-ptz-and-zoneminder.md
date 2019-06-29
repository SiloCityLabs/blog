---
title: Axis 2130 PTZ and Zoneminder
author: Luis Rodriguez
type: post
date: 2019-02-07T23:30:58+00:00
url: /post/2019/02/07/axis-2130-ptz-and-zoneminder/
featured_image: /wp-content/uploads/2018/08/41X1FB2KSBL._SY400_.jpg
categories:
  - Projects
tags:
  - axis
  - camera
  - security camera
  - zoneminder

---
### Axis 2130 PTZ and Zoneminder

{{< image src="/uploads/2018/08/41X1FB2KSBL._SY400_.jpg" alt="">}}

#### Details:

This camera is very dated, also using a proprietary power adapter, no WiFi and very difficult to setup. Once you are on my network you have full viewing access to my camera. My router fortunately has network isolation so I am able to block access to it except via my server that has zoneminder nvr software installed. Unfortunately you will not be able to control the ptz from Zoneminder directly due to the age of the camera. But you can however give it a fixed position in the web panel it has built in. I would recommend setting its default fixed position after a reboot or power loss.

<!--more-->

#### Setup:

I had to first hook it up to my desktop via an Ethernet cable and send the IP address it would use on my network. By default this camera has no DHCP or static IP enabled. In the [manual][2] you can find out how to send the ARP request to perform this. After you have an IP go to it on your web browser. http://youripaddress/. You will see the camera start streaming once you have done any first time setup. Head over into Zoneminder and add a new camera with the following details.

Remote Host Path: /axis-cgi/mjpg/video.cgi?camera=1&1534046539993

&nbsp;

{{< image src="/uploads/2018/12/1.png" alt="">}}

{{< image src="/uploads/2018/12/2.png" alt="">}}

 [2]: https://www.axis.com/files/usermanual/2130_um_en.pdf