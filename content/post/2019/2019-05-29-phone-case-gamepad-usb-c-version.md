---
title: Phone case gamepad, USB-C version
author: maave
type: post
date: 2019-05-29T17:19:34+00:00
url: /post/2019/05/29/phone-case-gamepad-usb-c-version/
featured_image: /wp-content/uploads/2019/05/featured-image.png
categories:
  - Projects
tags:
  - gamepad
  - gamepad case
  - smartphone
  - USB-C

---

![Pads](/uploads/2019/05/2560301558990512124-1.jpg)

Remember my [Bluetooth gamepad project][1]? It&#8217;s time for a USB-C version. I started with Bluetooth because I thought it would be easier to support multiple phones. However [Hackaday and OSHPark are holding a flex PCB contest][2]. This is a good material to make a gamepad with, and the perfect material for a gamepad that is supposed to resize.

<!--more-->

![cardboard gamepad mockup](/uploads/2019/05/6769521558992231498.jpg)

I tried a few mockups and settled on a 2 piece design, with the halves connected by a ribbon cable. The left and right side share the same PCB but have different components attached. The &#8220;main&#8221; side (probably right side) will have most of the components and the USB connector. The &#8220;aux&#8221; side (probably left side) will have just an LED for status indication &#8211; all the button contacts are built into the PCB.

![SNES gamepad silicone on mockup](/uploads/2019/05/2320541558989079616.jpg)

I&#8217;m using SNES buttons for this version. The contest is for 2sq inches of PCB and only the D-Pad silicone will fit in my 2&#215;1 footprint, so I&#8217;ll use two of them. For future versions I&#8217;ll check out button sets for portables like the GBA, DS, and PSP.

![](/uploads/2019/05/5843821558996883569.png)

PCB layout is in the works. I still have rearranging and routing to do. I may have to change USB-C connectors to something that will fit the trace width of OSHpark.

Follow my [project on HackADay.io][6] for finer details and updates! Source code will be available the SiloCityLabs Github soon™. This will be completely open source &#8211; PCB, firmware, and 3d printed case.

 [1]: https://blog.silocitylabs.com/post/2017/12/18/phone-case-gamepad-wip/
 [2]: https://hackaday.io/contest/163267-flexible-pcb-concept-contest
 [6]: https://hackaday.io/project/165606-usb-c-gamepad-phone-case