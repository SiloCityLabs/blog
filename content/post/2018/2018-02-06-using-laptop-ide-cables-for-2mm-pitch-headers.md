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
[<img class="wp-image-359" src="/uploads/2018/02/IMG_20180203_213426-Copy.jpg" width="1160" height="371" srcset="/uploads/2018/02/IMG_20180203_213426-Copy.jpg 1160w, /uploads/2018/02/IMG_20180203_213426-Copy-300x96.jpg 300w, /uploads/2018/02/IMG_20180203_213426-Copy-768x246.jpg 768w, /uploads/2018/02/IMG_20180203_213426-Copy-1024x328.jpg 1024w" sizes="(max-width: 1160px) 100vw, 1160px" />][1]

<figure id="attachment_358" aria-describedby="caption-attachment-358" style="width: 300px" class="wp-caption alignright">[<img class="wp-image-358" src="/uploads/2018/02/Core51822-size-300x255.jpg" width="300" height="255" srcset="/uploads/2018/02/Core51822-size-300x255.jpg 300w, /uploads/2018/02/Core51822-size.jpg 750w" sizes="(max-width: 300px) 100vw, 300px" />][2]<figcaption id="caption-attachment-358" class="wp-caption-text">2&#215;10 2mm pitch headers, ugh</figcaption></figure>

My core51822 module has caused me a bit of headache. The headers are 2.00mm pitch rather than the typical 1&#8243; (2.54mm) pitch. This makes it incompatible with regular 1&#8243; pitch breadboards. Being double row makes it worse.

<!--more-->

&nbsp;

After searching a while I was able to find two compatible cables. First some regular cables that fit the specs &#8211; [2mm pitch 2&#215;10 IDC connectors][3]. Second is even cheaper &#8211; [44-pin laptop IDE cable][4]​. I accidentally got the 2 inch cable, I should&#8217;ve got the 8 inch. 8 inch would&#8217;ve given me more wire to play with, enough to save both end of the cable.

If you buy a longer cable you can simply cut the cable in half and get two 2&#215;24 cables. With the 2&#8243; cable I got I couldn&#8217;t afford shortening the wires any further. I removed the connector from one end by breaking the tabs on the side

<figure id="attachment_360" aria-describedby="caption-attachment-360" style="width: 281px" class="wp-caption aligncenter">[<img class="wp-image-360" src="/uploads/2018/02/IMG_20180203_191933-281x300.jpg" width="281" height="300" srcset="/uploads/2018/02/IMG_20180203_191933-281x300.jpg 281w, /uploads/2018/02/IMG_20180203_191933.jpg 485w" sizes="(max-width: 281px) 100vw, 281px" />][5]<figcaption id="caption-attachment-360" class="wp-caption-text">Snap off this tab</figcaption></figure>

Removing the plate revealed the pins on the connector underneath. I popped off both sides to confirm which wire went to which female pin. The pins are mirrored top and bottom row so I can look at the puncture point ( top or bottom) and determine which row it leads to. I could probably read the specs but this was faster.

[<img class="wp-image-361" src="/uploads/2018/02/IMG_20180203_205214-Copy-300x237.jpg" width="300" height="237" srcset="/uploads/2018/02/IMG_20180203_205214-Copy-300x237.jpg 300w, /uploads/2018/02/IMG_20180203_205214-Copy-768x608.jpg 768w, /uploads/2018/02/IMG_20180203_205214-Copy-1024x810.jpg 1024w, /uploads/2018/02/IMG_20180203_205214-Copy.jpg 1141w" sizes="(max-width: 300px) 100vw, 300px" />][6]

I glued the plate back on one side and fully removed the connector from the other. A few seconds with a hacksaw gave me two connectors. I managed to mangle only a couple columns and was left with 2&#215;10 and 2&#215;11, slightly more than I needed for the core51822&#8217;s 2&#215;9 headers. I superglued the cut ends to keep the plate on.

[<img class="wp-image-363" src="/uploads/2018/02/IMG_20180203_213426-300x250.jpg" width="300" height="250" srcset="/uploads/2018/02/IMG_20180203_213426-300x250.jpg 300w, /uploads/2018/02/IMG_20180203_213426-768x640.jpg 768w, /uploads/2018/02/IMG_20180203_213426-1024x853.jpg 1024w, /uploads/2018/02/IMG_20180203_213426.jpg 1272w" sizes="(max-width: 300px) 100vw, 300px" />][7]

I pulled the wire apart in pairs so that each pair is a column, making it easier to handle when breadboarding. The stranded wire isn&#8217;t quite suitable for breadboards so later I soldered a tiny segment of 24AWG wire onto the end.

<figure id="attachment_366" aria-describedby="caption-attachment-366" style="width: 300px" class="wp-caption aligncenter">[<img class="wp-image-366" src="/uploads/2018/02/IMG_20180203_213608-300x127.jpg" width="300" height="127" srcset="/uploads/2018/02/IMG_20180203_213608-300x127.jpg 300w, /uploads/2018/02/IMG_20180203_213608-768x326.jpg 768w, /uploads/2018/02/IMG_20180203_213608-1024x434.jpg 1024w" sizes="(max-width: 300px) 100vw, 300px" />][8]<figcaption id="caption-attachment-366" class="wp-caption-text">Done! Breadboard-able</figcaption></figure>

End result is a pair of connectors for 2mm pitch header breakout. Conclusion? It works but doesn&#8217;t have much advantage over the regular 2mm pitch IDC connector. I saved $1.50. Do yourself a favor and get a module that uses 1&#8243; headers in the first place.

 [1]: /uploads/2018/02/IMG_20180203_213426-Copy.jpg
 [2]: /uploads/2018/02/Core51822-size.jpg
 [3]: https://www.ebay.com/itm/2/182158682043
 [4]: https://www.newegg.com/Product/Product.aspx?Item=9SIA67055T3504
 [5]: /uploads/2018/02/IMG_20180203_191933.jpg
 [6]: /uploads/2018/02/IMG_20180203_205214-Copy.jpg
 [7]: /uploads/2018/02/IMG_20180203_213426.jpg
 [8]: /uploads/2018/02/IMG_20180203_213608.jpg