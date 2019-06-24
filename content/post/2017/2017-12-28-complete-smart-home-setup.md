---
title: Complete Smart Home Setup
author: Luis Rodriguez
type: post
date: 2017-12-28T18:00:51+00:00
url: /post/2017/12/28/complete-smart-home-setup/
featured_image: /wp-content/uploads/2017/12/smart-home-2769210_960_720-e1514497002881.jpg
categories:
  - Projects
tags:
  - 433mhz
  - energizer
  - esp8266
  - fire alarm
  - google home
  - google home mini
  - hue
  - more things
  - morethings
  - smart home
  - smart light
  - smart strip
  - smart switch
  - smartthings

---
![](/uploads/2017/12/smart-home-2769210_960_720-e1514497002881.jpg)

&nbsp;

I wanted to talk about my smart home setup from the ground up. I have many articles describing my setup in pieces but not everyone gets an entire overview of everything I have setup here. I don't have a mansion or anything fancy like built in sprinkler/irrigation, so this will be more catered to your average city home.

&nbsp;

**The Assistant**:
r
Google Home and/or Mini manage my random routines that aren't part of a automation rule. Going to the basement to do laundry, or going to sleep at random times. I can just send a command and google home will handle it. Link your SmartThings account and Google Home together for best results. Give all your devices aliases to make it easier to command via Google. I want one located in every central room of the house, Kitchen, Living room, Master Bedroom.

![](/uploads/2017/12/100417-google-home-mini7347-2.jpg)

**Control Center**:

What controls my entire setup and why I chose this hub vs others. My only hub is a **Samsung SmartThings hub**. This manages just about everything in my home minus the security cameras. I did heavy research on many hubs including Wink, IOT (hub-less), and Phillips hue. Why SmartThings?

<!--more-->

  * **Device compatibility** - This hub unlike others will work with any Zigbee and Z-wave certified device out of the box. Other hubs like Phillips hue are more catered to the devices those companies sell. I was even able to add a hue bulb to my hub. The only thing wink has that I would have wanted was the 433mhz range for my security sensors, but they are locked down to a specific brand that is a partner of theirs.
  * **Open Source Community** - None of these hubs have the open source community behind it like the SmartThings. Samsung even has a [public github repository][2] for people to edit.
  * **Faster Wifi** - This hub ties directly into your router via ethernet. But this isn't the main reason, hubless wifi devices may save you the cost of a hub now, but the more you have the more it bogs your main wifi. Even if you seclude them on a separate router, it can still cause interference that slows down your wifi or even lowers your wifi range. Zigbee and Z-wave Devices both operate on separate frequencies that are entirely separate from wifi.

The only downside:

  * **No Apple HomeKit support** - HomeKit SDK [really limits the capabilities][3] of what a true smart home can do. I have no real need for HomeKit being an Android/Google user though. With all the assistants you can buy out of the box, who really needs HomeKit. Its not like Siri was really useful anyway.

&nbsp;

**Smart Bulbs**:

I personally have a set budget per bulb that I would pay. I would never pay anything above $10/each under any circumstances, color changing or not. My favorite bulb at the moment seems to be the affordable Sengled Classic A19. This bulb has a very nice light output at 800 lumens. I paid just under $7/each on sale, as they very often go on sale on [Amazon][4].

![](/uploads/2017/12/sengled-led-bulbs-e11-g13w-64_1000.jpg)

**Smart Switch Control**:

Why switches vs smart bulbs? Simplicity!

  * You can wire it into a pre-existing wall switch
  * No smart bulb needed - Great for chandeliers where you have 4-6 bulbs that could be about $40-60 in bulbs vs $26 switch.
  * 2 Relay Switch - Save money where you have two switches on a single plate. Cutting the cost down in half.
  * Easier to use - No need to tape the wall switch into the always on position.
  * It just makes sense! Mix old tech with the new.

![](/uploads/2017/12/41vrEw9HFL.jpg)

**Smart Plugs**:

I use these anywhere where I have lamps, automating lighting like in my turtle tank or Christmas lights. These even record power usage so you can place them on appliances like furnaces, 3d printers, etc to monitor power usage over time. [Amazon link][7]

![](/uploads/2017/12/856418004501.jpg)

&nbsp;

**Added missing switches**:

I have 3 locations in my house that have no wall switch. All 3 have pull strings. Basement lights (x6), back hall, and front hall. I was able to put in place a Sengled smart bulb and a battery smart switch. These switches last up to 10 years on battery and can be placed in a real gang box like any other switch, or double-side taped to the wall to mimic a switch plate. The most affordable switch I found was a Lutron Smart Switch on [Amazon.][9]

![Lutron Remote with faux plate](/uploads/2017/12/Wall-Mount-Lutron-Pico-Remote-Control.jpg)

&nbsp;

**Thermostat**:

Nest. There is a little secret most people may not know. You can get one for under $100. My wife said I was not gonna spend more than $100. I saw an opportunity and she said a deal was a deal. Most electric or natural gas providers give incentives for using a nest smart thermostat. I purchased mine for $97, The Nest E. Its not metalic but it works just the same. If you have a complicated AC/Heating system with zones and all that crazy stuff then the E version may not have enough wiring pins. Sign up for an online account to pay your electric bill and sign up for newsletter emails. On holidays they usually send out emails giving you the chance to get these cheap. There's also an ongoing [rebate program][11] online.

&nbsp;

![](/uploads/2017/12/NestE.png)

&nbsp;

**Security Sensors**:

This part is for more experienced users. Using a wireless module I created a mini hub essentially that interfaces 433mhz to Samsung SmartThings. Using an [ESP8266][13] with our [open source code][14] and a [wireless 433mhz receiver][15], you can add super cheap security sensors from [ebay][16]. These sensors usually sell for about $1-2/each and can be purchased in bulk quantities. Other types of sensors include motion sensors, door/window sensors, leak sensors, carbon monoxide sensors and many more. Just search 433mhz on ebay and you will get a bunch of results. Unfortunately I will only use these as a one way signal, I don't trust this open frequency to be used to control devices.

![](/uploads/2017/12/s-l1600.jpg)

&nbsp;

**Automation Rules**:

These are probably the best thing about owning a smart home system. My current rules in places are below:

  * At Sunrise: Turn on Turtle Tank UV Light
  * At Sunset: Turn off Turtle Tank UV Light,
  * 4 Hours after Sunset: Turn off Turtle tank Heat lamp
  * At 1AM: turn off all designated lights, backyard, basement, etc. Save some $$
  * Arrival: When me or my wife arrive, turn on backyard/driveway lights and entryway light.

**Semi-Smart Lock**:

Nothing fancy here, just a basic pin code lock by Kwikset. Bluetooth smart locks are [inevitably hacked][18]. Wifi devices kill batteries. Nothing really interest me at the moment. I have a a Kwikset 910 pin code lock on both the front and rear entries of the home. Keyless entry along with my push-to-start car means I never see/use keys.

![](/uploads/2017/12/product_header_electronic.jpg)

**Android Auto**:

Yea my car. When I arrive I can ask Google to power on my yard lights or turn off all the lights as I am leaving right from my car. No button press needed as Google is always listening. Doesn't work for Apple users since SmartThings lacks HomeKit support.

**Batteries**:

Don't cheap out on batteries for any of this stuff. **Energizer** batteries have proven test results and outlast any other battery. Knockoff batteries are notoriously bad with fake capacity ratings. You don't need one of these smart home devices failing leaving you locked out, security-less, or even worse if you have a smart fire alarm.

**3-Way Switch**:

The simple solution for this. Using a 1 relay switch on one side. Closing the connection on the other side and placing a fake switch in the panel in place solves this problem. This is in theory. I have yet to test this on my setup yet but it will be coming soon.

![](/uploads/2017/12/two-way-light-switch-diagram-or-staircase-wiring-diagram.jpg)

**Light Switch Controlling Socket**:

In my mother's house she has rooms where the light switch controls an outlet in the wall. This does not always end up very well. They usually get taped in the _On_ position or just never used because of lamps on the opposite ends of the room.

Solution: Remote Switches. You can place smart bulbs in the lamps and close the line with a wire nut to keep the outlet always on. Then place a remote switch in the gang box to control your lamps remotely.

**Soon to be**:

Not in my arsenal yet but these are a few devices I am keeping my eyes on.

  * Z-Wave fire alarm - A little out of my budget for this kind of product. On [Amazon][21]
  * security cameras - Not really sure yet how these work, but they caught my eye. By [Netgear][22]
  * Zigbee smart strip - Seems cheaper to get a strip vs multiple plugs in some instances. On [Amazon][23]
  * Doorbell - I'm in between [Ring Pro][24] and [Aeotec][25]. Ring Pro needs an existing wire which I don't have.

**Total cost**: (market value)

  * [Smart bulbs][4] x 7 x $9.99 = $69.92 
      * Often drop by 40% off
  * [1 relay switch][26] x 10 x $26.99 = <del>$269.90</del> $251 (bulk) 
      * Bulk rates available
  * [2 relay switch][27] x 5 x $29.95 = $149.75 
      * Purchase from alternate sellers to save $7
  * [Remote switch][28] x 3 x $15.91 = $47.73
  * [Google home mini][29] x 3 x $29.00 = $87.00 
      * Can often be acquired for free ([free mini here][30])
  * [Nest E][31] x 1 x $169.90 = $169.00 (free home mini, Offer valid 12/1/17–12/31/17) 
      * Check electric/gas provider for promo offers
  * [Smart plugs][32] x 4 x $39.99 = $159.96 
      * Often seen for $25 on sale
  * [433mhz Security sensors][33] x 20 x $2.32 = $46.40
  * [Esp8266 12F][34] x 2 x $2.95 = $5.90
  * [433mhz receiver][35] x 2 x $6.95(4) = $6.95
  * [kwikset locks 909][36] x 2 x $81.41 = $162.82
  * [Energizer AA batteries][37] x 8 x $14.99(12) = $14.99 
      * Total Smart home cost at market value: $1171.42

Don't get discouraged by the high price tag, You need to keep an eye out for deals. Many of these smart home items go on sale for over 50% off. A great forum I subscribed to is [SmartThings Deals][38], users post deals here all the time. This is how I got my bulbs for less than $7 each. I also received most of my google home mini's for free with select offers.

If you have any other smart home stuff you think should be on this list let me know in the comments.

 [1]: /uploads/2017/12/100417-google-home-mini7347-2.jpg
 [2]: https://github.com/SmartThingsCommunity/
 [3]: https://community.smartthings.com/t/homekit-intergration/50296/3
 [4]: https://www.amazon.com/Sengled-Compatible-SmartThings-Requires-Assistant/dp/B01N7I4X94/
 [5]: /uploads/2017/12/sengled-led-bulbs-e11-g13w-64_1000.jpg
 [6]: /uploads/2017/12/41vrEw9HFL.jpg
 [7]: https://www.amazon.com/Samsung-F-OUT-US-2-SmartThings-Outlet-White/dp/B073GV2PQY/
 [8]: /uploads/2017/12/856418004501.jpg
 [9]: https://www.amazon.com/gp/product/B00KLAXFQ0/
 [10]: /uploads/2017/12/Wall-Mount-Lutron-Pico-Remote-Control.jpg
 [11]: https://nest.com/energy-partners/national-grid/
 [12]: /uploads/2017/12/NestE.png
 [13]: https://rover.ebay.com/rover/1/711-53200-19255-0/1?icep_id=114&ipn=icep&toolid=20004&campid=5338203480&mpre=https%3A%2F%2Fwww.ebay.com%2Fitm%2F272408386985
 [14]: https://github.com/SiloCityLabs/ESP-Smartthings
 [15]: https://www.ebay.com/itm/401085388270
 [16]: https://www.ebay.com/itm/5x-433MHz-Wireless-Door-Sensor-Detector-Magnetic-Monitor-Security-Alarm-US-Stock/152127147328
 [17]: /uploads/2017/12/s-l1600.jpg
 [18]: https://www.youtube.com/watch?v=KrOReHwjCKI
 [19]: /uploads/2017/12/product_header_electronic.jpg
 [20]: /uploads/2017/12/two-way-light-switch-diagram-or-staircase-wiring-diagram.jpg
 [21]: https://www.amazon.com/First-Alert-Z-Wave-Detector-Monoxide/dp/B00KMHXFAI/
 [22]: https://www.arlo.com/en-us/default.aspx
 [23]: https://www.amazon.com/dp/B00H3RL6JW/
 [24]: https://shop.ring.com/products/video-doorbell-pro
 [25]: https://aeotec.com/z-wave-doorbell
 [26]: https://www.amazon.com/Vision-Z-Wave-Micro-Switch-1-pack/dp/B01GQX1GFC/
 [27]: https://www.amazon.com/Vision-Z-Wave-Micro-Switch-relay/dp/B00R883YKU/
 [28]: https://www.amazon.com/dp/B00KLAXFQ0/
 [29]: https://www.walmart.com/ip/Google-Home-Mini-Chalk/159013183
 [30]: https://learn.arcadiapower.com/dplat/fb/lal/ghm/?utm_source=fb&utm_medium=cpc&utm_campaign=fbq118lalcitydplatghm
 [31]: https://www.bestbuy.com/site/nest-thermostat-e-white/6051016.p?skuId=6051016
 [32]: https://www.amazon.com/Samsung-SmartThings-Outlet-Works-Amazon/dp/B00MI5V5N6/
 [33]: https://www.ebay.com/itm/5PCS-Wireless-433MHz-Door-Window-Entry-Alarm-Warning-System-Magnetic-Sensor-A0F5/172301759887
 [34]: https://www.ebay.com/itm/ESP8266-ESP-12F-WIFI-Microcontroller-802-11N-Module-Arduino-NodeMCU-MicroPython/152766041599
 [35]: https://www.ebay.com/itm/4X-433Mhz-RF-Transmitter-and-Receiver-Module-link-kit-for-Arduino-USA-seller/232298169165
 [36]: https://www.amazon.com/Kwikset-SmartCode-Electronic-Deadbolt-featuring/dp/B00NT1OX5K/
 [37]: https://www.amazon.com/Energizer-Ultimate-Lithium-Batteries-Count/dp/B071D87WPV/
 [38]: https://community.smartthings.com/c/devices-integrations/deals
