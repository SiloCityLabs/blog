---
title: Using laptop IDE cables for 2mm pitch headers
author: maave
type: post
date: 2018-02-06T08:50:09+00:00
url: /post/2018/02/06/using-laptop-ide-cables-for-2mm-pitch-headers/
featured_image: /wp-content/uploads/2018/02/IMG_20180203_213426.jpg
categories:
  - Projects
tags:
  - breadboard
  - core51822
  - esp51822

---
{{< image src="/uploads/2018/02/IMG_20180203_213426-Copy.jpg" alt="">}}

{{< image src="/uploads/2018/02/Core51822-size.jpg" alt="2mm pitch headers, ugh">}}

My core51822 module has caused me a bit of headache. The headers are 2.00mm pitch rather than the typical 1" (2.54mm) pitch. This makes it incompatible with regular 1" pitch breadboards. Being double row makes it worse.

<!--more-->

&nbsp;

After searching a while I was able to find two compatible cables. First some regular cables that fit the specs - [2mm pitch 2&#215;10 IDC connectors][3]. Second is even cheaper - [44-pin laptop IDE cable][4]​. I accidentally got the 2 inch cable, I should've got the 8 inch. 8 inch would've given me more wire to play with, enough to save both end of the cable.

If you buy a longer cable you can simply cut the cable in half and get two 2&#215;24 cables. With the 2" cable I got I couldn't afford shortening the wires any further. I removed the connector from one end by breaking the tabs on the side

{{< image src="/uploads/2018/02/IMG_20180203_191933.jpg" alt="Snap off this tab">}}

Removing the plate revealed the pins on the connector underneath. I popped off both sides to confirm which wire went to which female pin. The pins are mirrored top and bottom row so I can look at the puncture point ( top or bottom) and determine which row it leads to. I could probably read the specs but this was faster.

{{< image src="/uploads/2018/02/IMG_20180203_205214-Copy.jpg" alt="">}}

I glued the plate back on one side and fully removed the connector from the other. A few seconds with a hacksaw gave me two connectors. I managed to mangle only a couple columns and was left with 2&#215;10 and 2&#215;11, slightly more than I needed for the core51822's 2&#215;9 headers. I superglued the cut ends to keep the plate on.

{{< image src="/uploads/2018/02/IMG_20180203_213426.jpg" alt="">}}

I pulled the wire apart in pairs so that each pair is a column, making it easier to handle when breadboarding. The stranded wire isn't quite suitable for breadboards so later I soldered a tiny segment of 24AWG wire onto the end.

{{< image src="/uploads/2018/02/IMG_20180203_213608.jpg" alt="Done! Breadboard-able">}}

End result is a pair of connectors for 2mm pitch header breakout. Conclusion? It works but doesn't have much advantage over the regular 2mm pitch IDC connector. I saved $1.50. Do yourself a favor and get a module that uses 1" headers in the first place.

 [3]: https://www.ebay.com/itm/2/182158682043
 [4]: https://www.newegg.com/Product/Product.aspx?Item=9SIA67055T3504
 [6]: /uploads/2018/02/IMG_20180203_205214-Copy.jpg
 [7]: /uploads/2018/02/IMG_20180203_213426.jpg
 [8]: /uploads/2018/02/IMG_20180203_213608.jpg
