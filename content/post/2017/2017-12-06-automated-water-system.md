---
title: Automated water system?
author: Luis Rodriguez
type: post
date: 2017-12-06T20:20:16+00:00
url: /post/2017/12/06/automated-water-system/
featured_image: /wp-content/uploads/2017/12/water-system.png
categories:
  - Projects
tags:
  - esp8266

---
A friend at work asked if I could help solve his issues with his water intake from the well. At first I was like um why? Living in the city I did not know anything about how country side water systems worked. Water pumped from a well has a chance that it may contain too much sulfur making the water smell. He wanted to pull rain water into a tank and use it as primary with the well water as a failsafe.

I am thinking about using a basic ESP8266 along with a couple of water pressure sensors and relays triggering the water pumps he already has. The code should be very simple to write and I already have most of the hardware here minus the sensor.

<!--more-->

![Prototype Diagram](/uploads/2017/12/water-system.png)

1: Insert water sensor into an empty main tank, Get the zero point of water pressure.

2: Insert water sensor into an empty rain tank, Get the zero point of water pressure.

Have the user either fill up both tanks or continue the process at another time when the tanks are full?

3: Relay activated pump when it detects the main tank being low on water but rain tank has water

4: Relay activated pump to trigger well water (WP) if both tanks are low/empty

&nbsp;

Other possible features, Smart apps, Rain notification reminders to tell the users it wont be raining for a while and use minimal water. Detect fill rate using pump model numbers. Monitor water usage over time.

&nbsp;

I will be testing this in small scale with two fish pumps, and 3 milk gallons of water to simulate the water sources and main tank. The larger scale version would have 500+ gallon tanks, don't want to wait for that to fill up. Just waiting for my friend to send me some startup funds to try this out.

&nbsp;

Parts List:

1x [ESP12F][2] $2.95

1x [Valve controls][3] $6.95

2-3x gallon jugs Free

2x [Relays][4] $8.99

2x [Water Pressure sensors][5] $10.30

2x Pumps (Using fish tank pumps i have around)

?x Assorted tubes and connectors ~$10

~$40 total

 [2]: https://www.ebay.com/itm/ESP8266-ESP-12F-WIFI-Microcontroller-802-11N-Module-Arduino-NodeMCU-MicroPython/152766041599
 [3]: https://www.adafruit.com/product/997
 [4]: https://www.ebay.com/itm/2pcs-5V-Dual-Channel-2-Relay-Module-Arduino-Relays-Switch-110V-115V-120V-220V-US/292230695516
 [5]: https://www.ebay.com/itm/372158322453