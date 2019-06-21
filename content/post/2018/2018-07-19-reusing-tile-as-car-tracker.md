---
title: Reusing Tile as car tracker.
author: Luis Rodriguez
type: post
date: 2018-07-19T06:36:37+00:00
url: /post/2018/07/19/reusing-tile-as-car-tracker/
featured_image: /wp-content/uploads/2018/07/use-your-phone-to-find-your-car.png
categories:
  - Projects
tags:
  - smart car
  - tile
  - tracker
  - tracking

---
Ever since I switched to Samsung smart things I purchased presence sensors and had no more use for my tiles after the batteries died. I retired them to a box for a long time. Today I figured out a neat project for the tile. I am going to turn it into a tracker for my cars. This is not real time tracking but can help find your car in larger cities.

![](/uploads/2018/07/use-your-phone-to-find-your-car3939281002465595095.png)

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

![](/uploads/2018/07/photo_2018-07-18_23-27-14-1.jpg)

Break off the battery pins and solder on a small length of wire. Reinsert it back into the body of the tile and make sure the wires dont interfere with any plastic.

![](/uploads/2018/07/photo_2018-07-18_23-27-12.jpg)

Drill or solder a hole at the top of the tile for the wires to exit. Slip the wires through then seal the entire enclosure with some hot glue.

![](/uploads/2018/07/photo_2018-07-18_23-27-09.jpg)

**Preparing the DC to DC Converter**

Using a 12v power supply, attach the leads to the converters &#8220;in&#8221; pins and measure the power on the &#8220;out&#8221; pins using a voltage meter. Turn the adjustment until you are close to 3v.

![](/uploads/2018/07/photo_2018-07-18_23-27-03-2.jpg)

Glue the bottom of the converter to the tile, then cover the converter in glue to seal it from the elements.

![](/uploads/2018/07/photo_2018-07-18_23-27-00-2.jpg)

**Powering the tile**

Where you choose to tie in the tile is up to you. Some options include:

  * Direct battery leads
  * obd II adapter
  * 12v power socket

I already have a dash cam that uses a 12v obdII connector so im just gonna tie it to that.

![](/uploads/2018/07/20180718_2316209079375938918319396.jpg)

I opened it up and located the pins. I soldered on my cables to the 12v pins.

![](/uploads/2018/07/20180718_2300467315136001876509913.jpg)

Once thats all set I closed it back up with a hole drilled for the cables.

![](/uploads/2018/07/20180718_2316153822933817977650180.jpg)

Reattached it to my cars obd port and hooked up my dash cam again.

![](/uploads/2018/07/20180718_2322341355435456704761609.jpg)

To finish it off. Press the tile button down for 10 seconds to turn it on. Rename your old tile in the app and you should be set.

![](/uploads/2018/07/screenshot_20180718-225107_tile5373044240499341291.jpg)

Update:

Someone on [hackaday][12] mentioned in the [comments][12] that I could save power by using a linear regulator. I will probably be doing this for a motorcycle. If you are interested in trying this I recommend this one.

[LD1117V33 Voltage Regulator][2]

 [1]: http://a.tra.li/Tn1h
 [2]: http://a.tra.li/Tqhy
 [3]: /uploads/2018/07/photo_2018-07-18_23-27-14-1.jpg
 [4]: /uploads/2018/07/photo_2018-07-18_23-27-12.jpg
 [5]: /uploads/2018/07/photo_2018-07-18_23-27-09.jpg
 [6]: /uploads/2018/07/photo_2018-07-18_23-27-03-2.jpg
 [7]: /uploads/2018/07/photo_2018-07-18_23-27-00-2.jpg
 [8]: /uploads/2018/07/20180718_2316209079375938918319396.jpg
 [9]: /uploads/2018/07/20180718_2300467315136001876509913.jpg
 [10]: /uploads/2018/07/20180718_2316153822933817977650180.jpg
 [11]: /uploads/2018/07/20180718_2322341355435456704761609.jpg
 [12]: https://hackaday.com/2018/07/29/turning-a-tile-into-a-car-tracker/#comment-4796213
