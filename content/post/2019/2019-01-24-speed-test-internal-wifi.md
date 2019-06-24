---
title: Speed Test internal WiFi
author: Luis Rodriguez
type: post
date: 2019-01-24T23:30:39+00:00
url: /post/2019/01/24/speed-test-internal-wifi/
featured_image: /wp-content/uploads/2018/12/speedtest.png
categories:
  - Projects
tags:
  - internal network
  - lan
  - speedtest
  - wifi

---

![](/uploads/2018/12/speedtest.png)

Many have heard the term speed test, the first thing that comes to mind is speedtest.net or google's speed test tool. Many don't know that speed on your internal network also affects you. When testing internet speed your device sends data to router then out to the internet, what if your router is the issue. Even worse, your device. I will show u a very handy utility I use to make sure everything in my house runs great. Its called [iPerf][2].

<!--more-->

&nbsp;

When using iPerf its best to test it with a known high performance device as the server. In this case I have connected my phone to 5ghz WiFi and place it next to router. A hard wired device would also be much better as the server as it will have full access to speed without slowing other devices down by hogging the WiFi. Download the app on whatever type of device you are using and set it up like so.

![](/uploads/2018/12/photo_2018-12-06_11-46-01.jpg)

&nbsp;

Next step, create a client device using iPerf to connect to the supplied IP on your other device. You may also search for iPerf GUI if you are uncomfortable using command line. Open a terminal window and type the path to iPerf or you can drag it into windows command prompt as well. Follow it with &#8220;-c {ipaddress given by server} -i 1&#8221;

![](/uploads/2018/12/1-7.png)

&nbsp;

This is reporting a internal speed of 105Mbits/sec. This could be far higher but its most likely limited by wifi speeds. When testing an old device as the client I get speeds of 12Mbits/sec.

You can test multiple combinations of devices using iperf. Its supported on many systems and devices listed [here][2].

Conclusion:

There are many factors I can summarize up. Many devices on WiFi basically makes the WiFi share the available speed across them all but sometimes the router is the culprit. Other times it may just be the end devices. Not all devices support high speed WiFi, in fact if you don't see a 5ghz network then you most likely are running at a much lower rate. The fastest router money can buy will always be limited if you are using older WiFi signals like 2.4ghz BGN to connect to it. It has the range, but not the speed.

 [2]: https://iperf.fr/iperf-download.php