---
title: Phone case gamepad, USB-C version
author: maave
type: post
date: 2019-05-29T17:19:34+00:00
url: /post/2019/05/29/phone-case-gamepad-usb-c-version/
featured_image: /wp-content/uploads/2019/05/featured-image-150x150.png
categories:
  - Side Projects
tags:
  - gamepad
  - gamepad case
  - smartphone
  - USB-C

---
<img class="aligncenter  wp-image-949" src="https://blog.silocitylabs.com/wp-content/uploads/2019/05/2560301558990512124-1-300x53.jpg" alt="" width="537" height="95" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/05/2560301558990512124-1-300x53.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/2560301558990512124-1-768x136.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/2560301558990512124-1-1024x181.jpg 1024w" sizes="(max-width: 537px) 100vw, 537px" />

Remember my [Bluetooth gamepad project][1]? It&#8217;s time for a USB-C version. I started with Bluetooth because I thought it would be easier to support multiple phones. However [Hackaday and OSHPark are holding a flex PCB contest][2]. This is a good material to make a gamepad with, and the perfect material for a gamepad that is supposed to resize.

<!--more-->

[<img class="aligncenter size-medium wp-image-943" src="https://blog.silocitylabs.com/wp-content/uploads/2019/05/6769521558992231498-300x225.jpg" alt="cardboard gamepad mockup" width="300" height="225" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/05/6769521558992231498-300x225.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/6769521558992231498-768x576.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/6769521558992231498-1024x768.jpg 1024w" sizes="(max-width: 300px) 100vw, 300px" />][3]

I tried a few mockups and settled on a 2 piece design, with the halves connected by a ribbon cable. The left and right side share the same PCB but have different components attached. The &#8220;main&#8221; side (probably right side) will have most of the components and the USB connector. The &#8220;aux&#8221; side (probably left side) will have just an LED for status indication &#8211; all the button contacts are built into the PCB.

[<img class="aligncenter size-medium wp-image-944" src="https://blog.silocitylabs.com/wp-content/uploads/2019/05/2320541558989079616-300x225.jpg" alt="SNES gamepad silicone on mockup" width="300" height="225" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/05/2320541558989079616-300x225.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/2320541558989079616-768x576.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/2320541558989079616-1024x768.jpg 1024w" sizes="(max-width: 300px) 100vw, 300px" />][4]

I&#8217;m using SNES buttons for this version. The contest is for 2sq inches of PCB and only the D-Pad silicone will fit in my 2&#215;1 footprint, so I&#8217;ll use two of them. For future versions I&#8217;ll check out button sets for portables like the GBA, DS, and PSP.

[<img class="aligncenter size-medium wp-image-945" src="https://blog.silocitylabs.com/wp-content/uploads/2019/05/5843821558996883569-300x164.png" alt="" width="300" height="164" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/05/5843821558996883569-300x164.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/5843821558996883569-768x420.png 768w, https://blog.silocitylabs.com/wp-content/uploads/2019/05/5843821558996883569-1024x560.png 1024w" sizes="(max-width: 300px) 100vw, 300px" />][5]

PCB layout is in the works. I still have rearranging and routing to do. I may have to change USB-C connectors to something that will fit the trace width of OSHpark.

Follow my [project on HackADay.io][6] for finer details and updates! Source code will be available the SiloCityLabs Github soon™. This will be completely open source &#8211; PCB, firmware, and 3d printed case.

 [1]: https://blog.silocitylabs.com/post/2017/12/18/phone-case-gamepad-wip/
 [2]: https://hackaday.io/contest/163267-flexible-pcb-concept-contest
 [3]: https://blog.silocitylabs.com/wp-content/uploads/2019/05/6769521558992231498.jpg
 [4]: https://blog.silocitylabs.com/wp-content/uploads/2019/05/2320541558989079616.jpg
 [5]: https://blog.silocitylabs.com/wp-content/uploads/2019/05/5843821558996883569.png
 [6]: https://hackaday.io/project/165606-usb-c-gamepad-phone-case