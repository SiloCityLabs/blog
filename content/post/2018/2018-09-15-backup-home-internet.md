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
  * Anything internet&#8230;

To start off we plugin a cell phone to the routers usb port.

<img class="wp-image-587" src="/uploads/2018/09/20180915_0932318158192643190909326.jpg" width="4032" height="3024" srcset="/uploads/2018/09/20180915_0932318158192643190909326.jpg 1920w, /uploads/2018/09/20180915_0932318158192643190909326-300x225.jpg 300w, /uploads/2018/09/20180915_0932318158192643190909326-768x576.jpg 768w, /uploads/2018/09/20180915_0932318158192643190909326-1024x768.jpg 1024w" sizes="(max-width: 4032px) 100vw, 4032px" />

<!--more-->There&#8217;s an issue where Android will keep turning off tether if there is slight changes in cable movement or whenever the router reboots. We solve this with automate. We configure automate to keep enabling USB tether every X seconds.

<img class="wp-image-588" src="/uploads/2018/09/20180915_0945388772121150635334533.png" width="1458" height="1561" srcset="/uploads/2018/09/20180915_0945388772121150635334533.png 1458w, /uploads/2018/09/20180915_0945388772121150635334533-280x300.png 280w, /uploads/2018/09/20180915_0945388772121150635334533-768x822.png 768w, /uploads/2018/09/20180915_0945388772121150635334533-956x1024.png 956w" sizes="(max-width: 1458px) 100vw, 1458px" />

The root command is optional if you can&#8217;t get USB tether working. Here is the full command.

<img class="wp-image-596" src="/uploads/2018/09/20180915_094523896414538435102704.png" width="573" height="573" srcset="/uploads/2018/09/20180915_094523896414538435102704.png 573w, /uploads/2018/09/20180915_094523896414538435102704.png 150w, /uploads/2018/09/20180915_094523896414538435102704-300x300.png 300w" sizes="(max-width: 573px) 100vw, 573px" />

After it&#8217;s enabled, hit play and let it run.

<img class="wp-image-589" src="/uploads/2018/09/20180915_0930581248984325352199968.jpg" width="3024" height="4032" srcset="/uploads/2018/09/20180915_0930581248984325352199968.jpg 1440w, /uploads/2018/09/20180915_0930581248984325352199968-225x300.jpg 225w, /uploads/2018/09/20180915_0930581248984325352199968-768x1024.jpg 768w" sizes="(max-width: 3024px) 100vw, 3024px" />

Place your phone in a convenient spot accessible to great reception.

<img class="wp-image-590" src="/uploads/2018/09/20180915_094244623090915077846698.jpg" width="3000" height="2680" srcset="/uploads/2018/09/20180915_094244623090915077846698.jpg 1920w, /uploads/2018/09/20180915_094244623090915077846698-300x268.jpg 300w, /uploads/2018/09/20180915_094244623090915077846698-768x686.jpg 768w, /uploads/2018/09/20180915_094244623090915077846698-1024x915.jpg 1024w" sizes="(max-width: 3000px) 100vw, 3000px" />

Head over to your router settings. Under WAN options there is a tab called &#8220;Dual WAN&#8221;.

<img class="wp-image-591" src="/uploads/2018/09/20180915_0943271850039273880023030.jpg" width="720" height="774" srcset="/uploads/2018/09/20180915_0943271850039273880023030.jpg 720w, /uploads/2018/09/20180915_0943271850039273880023030-279x300.jpg 279w" sizes="(max-width: 720px) 100vw, 720px" />

Enable USB tether and Android device.

<img class="wp-image-592" src="/uploads/2018/09/20180915_0943417170682303179420009.jpg" width="720" height="540" srcset="/uploads/2018/09/20180915_0943417170682303179420009.jpg 720w, /uploads/2018/09/20180915_0943417170682303179420009-300x225.jpg 300w" sizes="(max-width: 720px) 100vw, 720px" />

Configure the failover settings how you please.

You are all set! Your dashboard should look like this.

<img class="wp-image-593" src="/uploads/2018/09/20180915_0943126960527662415391769.jpg" width="720" height="755" srcset="/uploads/2018/09/20180915_0943126960527662415391769.jpg 720w, /uploads/2018/09/20180915_0943126960527662415391769-286x300.jpg 286w" sizes="(max-width: 720px) 100vw, 720px" />

 [1]: https://www.freedompop.com/phone