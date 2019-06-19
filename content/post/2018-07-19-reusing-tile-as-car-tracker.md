---
title: Reusing Tile as car tracker.
author: Luis Rodriguez
type: post
date: 2018-07-19T06:36:37+00:00
url: /post/2018/07/19/reusing-tile-as-car-tracker/
featured_image: /wp-content/uploads/2018/07/use-your-phone-to-find-your-car-150x150.png
categories:
  - Side Projects
tags:
  - smart car
  - tile
  - tracker
  - tracking

---
Ever since I switched to Samsung smart things I purchased presence sensors and had no more use for my tiles after the batteries died. I retired them to a box for a long time. Today I figured out a neat project for the tile. I am going to turn it into a tracker for my cars. This is not real time tracking but can help find your car in larger cities.

<img class="aligncenter wp-image-520 size-full" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/use-your-phone-to-find-your-car3939281002465595095.png" width="231" height="231" data-temp-aztec-id="e57c71de-adda-469c-8e36-8f100a97668f" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/use-your-phone-to-find-your-car3939281002465595095.png 231w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/use-your-phone-to-find-your-car3939281002465595095-150x150.png 150w" sizes="(max-width: 231px) 100vw, 231px" />

Supplies/Tools needed:

  * Pry Tool
  * Hot Glue gun
  * Voltage Meter
  * Dead battery Tile
  * [DC to DC Buck Converter][1] or [Linear converter][2] (see update below)
  * Soldering iron
  * Wire stripper/cutter
  * Drill (optional)

<!--more-->

Before we get started I just wanted to add a **warning**. In this circuit just like most, polarities are very important. I made the mistake of flipping wires and my whole setup overheated and melted the glue. Keep this in mind when testing and before plugging it into your car and before doing the final seal.

Also note that voltage on a car&#8217;s 12v line varies. With just the battery it&#8217;s around ~12v. With the car on and the alternator running it gets up to ~14v. The Tile may report &#8220;low battery&#8221; because of the reduced voltage when the car is off. It&#8217;s not an issue.

**Wire up the tile**

Start by prying open the tile carefully with the intention of resealing it. I already removed the cell battery in the picture below.

[<img class="aligncenter size-full wp-image-507" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-14-1.jpg" alt="" width="582" height="554" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-14-1.jpg 582w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-14-1-300x286.jpg 300w" sizes="(max-width: 582px) 100vw, 582px" />][3]

Break off the battery pins and solder on a small length of wire. Reinsert it back into the body of the tile and make sure the wires dont interfere with any plastic.

[<img class="aligncenter size-large wp-image-508" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-12-1024x558.jpg" alt="" width="678" height="369" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-12-1024x558.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-12-300x163.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-12-768x418.jpg 768w" sizes="(max-width: 678px) 100vw, 678px" />][4]

Drill or solder a hole at the top of the tile for the wires to exit. Slip the wires through then seal the entire enclosure with some hot glue.

[<img class="aligncenter size-large wp-image-510" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-09-1024x692.jpg" alt="" width="678" height="458" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-09-1024x692.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-09-300x203.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-09-768x519.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-09.jpg 1112w" sizes="(max-width: 678px) 100vw, 678px" />][5]

**Preparing the DC to DC Converter**

Using a 12v power supply, attach the leads to the converters &#8220;in&#8221; pins and measure the power on the &#8220;out&#8221; pins using a voltage meter. Turn the adjustment until you are close to 3v.

[<img class="aligncenter size-large wp-image-511" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-03-2-669x1024.jpg" alt="" width="669" height="1024" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-03-2-669x1024.jpg 669w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-03-2-196x300.jpg 196w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-03-2-768x1176.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-03-2.jpg 836w" sizes="(max-width: 669px) 100vw, 669px" />][6]

Glue the bottom of the converter to the tile, then cover the converter in glue to seal it from the elements.

[<img class="aligncenter size-large wp-image-512" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-00-2-1024x642.jpg" alt="" width="678" height="425" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-00-2-1024x642.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-00-2-300x188.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-00-2-768x482.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-00-2.jpg 1138w" sizes="(max-width: 678px) 100vw, 678px" />][7]

**Powering the tile**

Where you choose to tie in the tile is up to you. Some options include:

  * Direct battery leads
  * obd II adapter
  * 12v power socket

I already have a dash cam that uses a 12v obdII connector so im just gonna tie it to that.

[<img class="aligncenter wp-image-514 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316209079375938918319396-960x1024.jpg" alt="" width="678" height="723" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316209079375938918319396-960x1024.jpg 960w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316209079375938918319396-281x300.jpg 281w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316209079375938918319396-768x819.jpg 768w" sizes="(max-width: 678px) 100vw, 678px" />][8]

I opened it up and located the pins. I soldered on my cables to the 12v pins.

[<img class="aligncenter wp-image-515 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2300467315136001876509913-890x1024.jpg" alt="" width="678" height="780" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2300467315136001876509913-890x1024.jpg 890w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2300467315136001876509913-261x300.jpg 261w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2300467315136001876509913-768x884.jpg 768w" sizes="(max-width: 678px) 100vw, 678px" />][9]

Once thats all set I closed it back up with a hole drilled for the cables.

[<img class="aligncenter wp-image-516 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316153822933817977650180-1024x847.jpg" alt="" width="678" height="561" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316153822933817977650180-1024x847.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316153822933817977650180-300x248.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316153822933817977650180-768x635.jpg 768w" sizes="(max-width: 678px) 100vw, 678px" />][10]

Reattached it to my cars obd port and hooked up my dash cam again.

[<img class="aligncenter wp-image-517 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2322341355435456704761609-1018x1024.jpg" alt="" width="678" height="682" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2322341355435456704761609-1018x1024.jpg 1018w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2322341355435456704761609-150x150.jpg 150w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2322341355435456704761609-298x300.jpg 298w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2322341355435456704761609-768x773.jpg 768w" sizes="(max-width: 678px) 100vw, 678px" />][11]

To finish it off. Press the tile button down for 10 seconds to turn it on. Rename your old tile in the app and you should be set.

<img class="aligncenter wp-image-518 size-medium" src="https://blog.silocitylabs.com/wp-content/uploads/2018/07/screenshot_20180718-225107_tile5373044240499341291-146x300.jpg" alt="" width="146" height="300" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/07/screenshot_20180718-225107_tile5373044240499341291-146x300.jpg 146w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/screenshot_20180718-225107_tile5373044240499341291-498x1024.jpg 498w, https://blog.silocitylabs.com/wp-content/uploads/2018/07/screenshot_20180718-225107_tile5373044240499341291.jpg 720w" sizes="(max-width: 146px) 100vw, 146px" />

Update:

Someone on [hackaday][12] mentioned in the [comments][12] that I could save power by using a linear regulator. I will probably be doing this for a motorcycle. If you are interested in trying this I recommend this one.

[LD1117V33 Voltage Regulator][2]

 [1]: http://a.tra.li/Tn1h
 [2]: http://a.tra.li/Tqhy
 [3]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-14-1.jpg
 [4]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-12.jpg
 [5]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-09.jpg
 [6]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-03-2.jpg
 [7]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/photo_2018-07-18_23-27-00-2.jpg
 [8]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316209079375938918319396.jpg
 [9]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2300467315136001876509913.jpg
 [10]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2316153822933817977650180.jpg
 [11]: https://blog.silocitylabs.com/wp-content/uploads/2018/07/20180718_2322341355435456704761609.jpg
 [12]: https://hackaday.com/2018/07/29/turning-a-tile-into-a-car-tracker/#comment-4796213