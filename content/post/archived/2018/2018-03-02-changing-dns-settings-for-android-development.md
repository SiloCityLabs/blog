---
title: Changing DNS settings for Android Development
author: Luis Rodriguez
type: post
date: 2018-03-02T16:29:41+00:00
url: /post/2018/03/02/changing-dns-settings-for-android-development/
featured_image: /wp-content/uploads/2018/03/dh4rb6lf2va9ydut1osj-150x150.jpg
categories:
  - Projects
tags:
  - android
  - development
  - dns
  - testing

---
[<img class="aligncenter size-full wp-image-412" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/ezgif-1-6db60a8e39-1.jpg" alt="" width="728" height="395" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/ezgif-1-6db60a8e39-1.jpg 728w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/ezgif-1-6db60a8e39-1-300x163.jpg 300w" sizes="(max-width: 728px) 100vw, 728px" />][1]

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

You&#8217;ll notice for each setup I have created pastebin links to where my hosts file is located, the format is &#8220;ip hostname&#8221; with one per line, just like a windows /etc/hosts file. You will need this in a later step for access to your settings. If you have a web server for static links, that will be even better but optional, I use http://somewebsite.com/dns-settings/live.txt. This allows me to edit without updating my android device.

&nbsp;

**Step 1 &#8211; Basic application settings:**

  * Download ([Daedalus][3]) and open the application.
  * Click on the menu then click &#8220;Servers&#8221;

[<img class="aligncenter size-medium wp-image-395" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-15-169x300.jpg" alt="" width="169" height="300" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-15-169x300.jpg 169w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-15-576x1024.jpg 576w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-15.jpg 720w" sizes="(max-width: 169px) 100vw, 169px" />][4]

  * Click the add button and add the two google servers &#8220;8.8.8.8&#8221; and &#8220;8.8.4.4&#8221;, You can substitute these for your preferred DNS provider as well.

[<img class="aligncenter wp-image-398 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-33-Copy-1024x605.jpg" alt="" width="678" height="401" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-33-Copy-1024x605.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-33-Copy-300x177.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-33-Copy-768x454.jpg 768w" sizes="(max-width: 678px) 100vw, 678px" />][5]

  * Activate these dns servers in the menu by clicking on &#8220;Settings&#8221;
  * Then select Google and Google 2 from the settings. Also activate Advance settings and select the two following options in the image below

[<img class="aligncenter wp-image-400 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-32-1024x908.jpg" alt="" width="678" height="601" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-32-1024x908.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-32-300x266.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-32-768x681.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-32.jpg 1444w" sizes="(max-width: 678px) 100vw, 678px" />][6]

&nbsp;

&nbsp;

**Step 2 &#8211; Creating hosts files:**

  * If you do not have your own server, you can use [pastebin.com][7] to host your hosts files for use in the next steps.
  * Create a [pastebin][7] link for each mode, you can add as many hosts as you need to test like below. Then hit &#8220;Create New Paste&#8221;

&nbsp;

[<img class="aligncenter size-large wp-image-405" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin-1024x662.png" alt="" width="678" height="438" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin-1024x662.png 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin-300x194.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin-768x496.png 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin.png 1154w" sizes="(max-width: 678px) 100vw, 678px" />][8]

  * Click on the Raw button and copy this link somewhere for later use, Repeat steps for each mode.

[<img class="aligncenter size-large wp-image-406" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin2-1024x518.png" alt="" width="678" height="343" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin2-1024x518.png 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin2-300x152.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin2-768x389.png 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin2.png 1154w" sizes="(max-width: 678px) 100vw, 678px" />][9]

**Step 3 &#8211; Setting up the rules:**

  * Back to the menu now, Click on &#8220;Rules&#8221;.
  * Add a new rule for each development mode, Live, staging, dev. Filename is not crucial, just make something up.
  * Hit sync rule before leaving to make sure the file is downloaded.

[<img class="aligncenter wp-image-401 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-13-Copy-1024x605.jpg" alt="" width="678" height="401" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-13-Copy-1024x605.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-13-Copy-300x177.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-13-Copy-768x453.jpg 768w" sizes="(max-width: 678px) 100vw, 678px" />][10]

&nbsp;

**Step 4 &#8211; Activating:**

  * Tap and activate the mode you wish to be in, You can activate more than one at a time so be careful as this will override one or the other.
  * In the menu go to home, and hit the &#8220;Activate&#8221; button.

[<img class="aligncenter wp-image-402 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-08-1024x906.jpg" alt="" width="678" height="600" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-08-1024x906.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-08-300x265.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-08-768x680.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-08.jpg 1442w" sizes="(max-width: 678px) 100vw, 678px" />][11]

&nbsp;

**Step 5 &#8211; Testing:**

  * The final setup, lets see if I am hitting my dev server.
  * As you can see in the app below, the short url is no longer short, I now get a full url as the dev server doesnt have an api key to tra.li link shortener.
  * Depending on the use-case it may be difficult to detect which server your hitting, You can go back into Daedalus in the menu and select test. Type your domain name and see if the ip matches your dev server.

[<img class="aligncenter size-large wp-image-403" src="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-04-1024x905.jpg" alt="" width="678" height="599" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-04-1024x905.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-04-300x265.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-04-768x679.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-04.jpg 1448w" sizes="(max-width: 678px) 100vw, 678px" />][12]

 [1]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/ezgif-1-6db60a8e39-1.jpg
 [2]: http://codrcg.com
 [3]: http://a.tra.li/SF9e
 [4]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-15.jpg
 [5]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-33-Copy.jpg
 [6]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-32.jpg
 [7]: https://pastebin.com
 [8]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin.png
 [9]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/pastebin2.png
 [10]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-13-Copy.jpg
 [11]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-08.jpg
 [12]: https://blog.silocitylabs.com/wp-content/uploads/2018/03/photo_2018-03-02_10-35-04.jpg