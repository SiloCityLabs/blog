---
title: Changing DNS settings for Android Development
author: Luis Rodriguez
type: post
date: 2018-03-02T16:29:41+00:00
url: /post/2018/03/02/changing-dns-settings-for-android-development/
featured_image: /wp-content/uploads/2018/03/dh4rb6lf2va9ydut1osj.jpg
categories:
  - Projects
tags:
  - android
  - development
  - dns
  - testing

---
![](/uploads/2018/03/ezgif-1-6db60a8e39-1.jpg)

&nbsp;

This relates to any similar development environment. This guide will help you setup your device to work for your development setup either locally or a remote server. No root required.

First I will describe the use case scenario where this is useful:

<!--more-->

**Live:**

[codrcg.com][2]

192.3.206.221

(public server)

https://pastebin.com/raw/63dFFz2C

&nbsp;

**Staging / Test:**

[codrcg.com][2]

10.5.1.143

(internal server)

https://pastebin.com/raw/txav6Yit

&nbsp;

**Local Desktop:**

[codrcg.com][2]

192.168.1.256

(static desktop ip)

https://pastebin.com/raw/FudTpY1z

&nbsp;

You'll notice for each setup I have created pastebin links to where my hosts file is located, the format is "ip hostname" with one per line, just like a windows /etc/hosts file. You will need this in a later step for access to your settings. If you have a web server for static links, that will be even better but optional, I use http://somewebsite.com/dns-settings/live.txt. This allows me to edit without updating my android device.

&nbsp;

**Step 1 &#8211; Basic application settings:**

  * Download ([Daedalus][3]) and open the application.
  * Click on the menu then click "Servers"

![](/uploads/2018/03/photo_2018-03-02_10-35-15.jpg)

  * Click the add button and add the two google servers "8.8.8.8" and "8.8.4.4", You can substitute these for your preferred DNS provider as well.

![](/uploads/2018/03/photo_2018-03-02_10-35-33-Copy.jpg)

  * Activate these dns servers in the menu by clicking on "Settings"
  * Then select Google and Google 2 from the settings. Also activate Advance settings and select the two following options in the image below

![](/uploads/2018/03/photo_2018-03-02_10-35-32.jpg)

&nbsp;

&nbsp;

**Step 2 &#8211; Creating hosts files:**

  * If you do not have your own server, you can use [pastebin.com][7] to host your hosts files for use in the next steps.
  * Create a [pastebin][7] link for each mode, you can add as many hosts as you need to test like below. Then hit "Create New Paste"

&nbsp;

![](/uploads/2018/03/pastebin.png)

  * Click on the Raw button and copy this link somewhere for later use, Repeat steps for each mode.

![](/uploads/2018/03/pastebin2.png)

**Step 3 &#8211; Setting up the rules:**

  * Back to the menu now, Click on "Rules".
  * Add a new rule for each development mode, Live, staging, dev. Filename is not crucial, just make something up.
  * Hit sync rule before leaving to make sure the file is downloaded.

![](/uploads/2018/03/photo_2018-03-02_10-35-13-Copy.jpg)

&nbsp;

**Step 4 &#8211; Activating:**

  * Tap and activate the mode you wish to be in, You can activate more than one at a time so be careful as this will override one or the other.
  * In the menu go to home, and hit the "Activate" button.

![](/uploads/2018/03/photo_2018-03-02_10-35-08.jpg)

&nbsp;

**Step 5 &#8211; Testing:**

  * The final setup, lets see if I am hitting my dev server.
  * As you can see in the app below, the short url is no longer short, I now get a full url as the dev server doesnt have an api key to tra.li link shortener.
  * Depending on the use-case it may be difficult to detect which server your hitting, You can go back into Daedalus in the menu and select test. Type your domain name and see if the ip matches your dev server.

![](/uploads/2018/03/photo_2018-03-02_10-35-04.jpg)

 [1]: /uploads/2018/03/ezgif-1-6db60a8e39-1.jpg
 [2]: http://codrcg.com
 [3]: http://a.tra.li/SF9e
 [4]: /uploads/2018/03/photo_2018-03-02_10-35-15.jpg
 [5]: /uploads/2018/03/photo_2018-03-02_10-35-33-Copy.jpg
 [6]: /uploads/2018/03/photo_2018-03-02_10-35-32.jpg
 [7]: https://pastebin.com
 [8]: /uploads/2018/03/pastebin.png
 [9]: /uploads/2018/03/pastebin2.png
 [10]: /uploads/2018/03/photo_2018-03-02_10-35-13-Copy.jpg
 [11]: /uploads/2018/03/photo_2018-03-02_10-35-08.jpg
 [12]: /uploads/2018/03/photo_2018-03-02_10-35-04.jpg
