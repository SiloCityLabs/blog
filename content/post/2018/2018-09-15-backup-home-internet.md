---
title: Backup home internet
author: Luis Rodriguez
type: post
date: 2018-09-15T14:12:00+00:00
url: /post/2018/09/15/backup-home-internet/
featured_image: /wp-content/uploads/2018/09/20180915_0943126960527662415391769.jpg
categories:
  - Projects
tags:
  - asus
  - backups
  - internet

---
With new most asus routers you can use a feature called dual wan. I have the luxury of having a work provided cell phone I am using for this. Before having that I had a free service called [freedompop][1].

Use cases (mainly mine):

  * Watching netflix withouth interuptions.
  * Backup internet for home security system.
  * Anything internet...

To start off we plugin a cell phone to the routers usb port.

![](/uploads/2018/09/20180915_0932318158192643190909326.jpg)

<!--more-->There's an issue where Android will keep turning off tether if there is slight changes in cable movement or whenever the router reboots. We solve this with automate. We configure automate to keep enabling USB tether every X seconds.

![](/uploads/2018/09/20180915_0945388772121150635334533.png)

The root command is optional if you can't get USB tether working. Here is the full command.

![](/uploads/2018/09/20180915_094523896414538435102704.png)

After it's enabled, hit play and let it run.

![](/uploads/2018/09/20180915_0930581248984325352199968.jpg)

Place your phone in a convenient spot accessible to great reception.

![](/uploads/2018/09/20180915_094244623090915077846698.jpg)

Head over to your router settings. Under WAN options there is a tab called "Dual WAN".

![](/uploads/2018/09/20180915_0943271850039273880023030.jpg)

Enable USB tether and Android device.

![](/uploads/2018/09/20180915_0943417170682303179420009.jpg)

Configure the failover settings how you please.

You are all set! Your dashboard should look like this.

![](/uploads/2018/09/20180915_0943126960527662415391769.jpg)

 [1]: https://www.freedompop.com/phone
