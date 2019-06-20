---
title: What can golang do?
author: Luis Rodriguez
type: post
date: 2019-06-06T20:00:28+00:00
url: /post/2019/06/06/what-can-golang-do/
featured_image: /wp-content/uploads/2019/06/1_30aoNxlSnaYrLhBT0O1lzw-150x150.png
categories:
  - Golang
  - Projects
tags:
  - coding
  - daemons
  - development
  - go
  - golang
  - scripting
  - software

---
<img class="wp-image-982 size-full aligncenter" src="https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_30aonxlsnayrlhbt0o1lzw8398331545093394431.png" width="800" height="533" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_30aonxlsnayrlhbt0o1lzw8398331545093394431.png 800w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_30aonxlsnayrlhbt0o1lzw8398331545093394431-300x200.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_30aonxlsnayrlhbt0o1lzw8398331545093394431-768x512.png 768w" sizes="(max-width: 800px) 100vw, 800px" />

<span style="color: rgba(0,0,0,0.84); font-family: medium-content-serif-font,Georgia,Cambria,;">Nothing makes a developer crazier than a new programming language, right? So, I started learning Go about 6 months ago and </span>now I am going to share with you my reasons for doing so.

## Companies using Golang

[<img class="aligncenter size-full wp-image-997" src="https://blog.silocitylabs.com/wp-content/uploads/2019/06/Golang-Users.png" alt="" width="724" height="515" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/06/Golang-Users.png 724w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/Golang-Users-300x213.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/Golang-Users-210x150.png 210w" sizes="(max-width: 724px) 100vw, 724px" />][1]

Google &#8211; Go is designed and supported by Google. Google has one of the largest cloud infrastructures in the world and it is massively scaled.

Apple &#8211; Big data, analytics, infrastructure software.

Cloudflare &#8211; Used in infrastructure and other various locations. Has decreased time to first byte by 90%.

Gitlab &#8211; Handles all RPC calls made by git

AT&T &#8211; Network management, server applications, web based API&#8217;s

Netflix &#8211; Caching management

Tesla &#8211; Server management, enterprise software, production lines.

SpaceX &#8211; Used in rocket telemetry

The [list][2] is over thousands long

## Hardware limitations

7 years ago I had a 4ghz processor with 8 cores. Today my desktop has a 4 core processor with 3.9ghz. You&#8217;re probably looking at these numbers like, &#8220;why did he downgrade?&#8221; I did not. The benchmarks actually put my newer processor ahead by 66% because of multi-threading capabilities.

Nearly one decade and the hardware has not changed much besides one key thing. Multi-threading, and now possibly even more cores. Everything else has pretty much been steady and slow growth.

<!--more-->


  
&nbsp;

[<img class="size-full wp-image-1028 aligncenter" src="https://blog.silocitylabs.com/wp-content/uploads/2019/06/chart.png" alt="" width="600" height="416" data-wp-editing="1" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/06/chart.png 600w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/chart-300x208.png 300w" sizes="(max-width: 600px) 100vw, 600px" />][3]

Even multi-threading has its limits. With that said, we fall to software to fix these issues. We must increase efficiency and performance of our software.

You&#8217;re probably thinking &#8220;well python is the ultimate language for AI.&#8221; Sure it has many existing libraries, but it’s also old and slower. On some benchmarks even 40x slower.

## Go runs direct

No crazy configuration is needed. Compile it for Linux and download any minimal Linux and let it do its thing. You don&#8217;t need any added software on the server it will run on. It&#8217;ll run identical on nearly any hardware. Yes, it even compiles for that Windows server your boss says you need for “compliance&#8217;s.”

Starting a compiled Go application is as simple as a file transfer and &#8220;./myprogram&#8221;

Go has the benefit of having its own compiler for each architecture. It is not bogged down by a VM like java or a processor like Python or PHP. It is like C/C++ with ease of learning.

## Go routines

This is where the magic happens. Routines help create concurrency you need to speed up your application where others fail.

<img class="wp-image-978 alignnone size-full" src="https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_nfojvbkdrkxz0zdbu4ysna1724252319353850356.jpeg" width="538" height="172" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_nfojvbkdrkxz0zdbu4ysna1724252319353850356.jpeg 538w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_nfojvbkdrkxz0zdbu4ysna1724252319353850356-300x96.jpeg 300w" sizes="(max-width: 538px) 100vw, 538px" />

Routines in go startup faster and even use less memory to start than other languages like python and java. They can get complicated quickly in other languages but with Go it’s as simple as:

<pre>go functionName()</pre>

You could spin up millions of routines before you start having bottlenecks in hardware. Meanwhile Languages like java will start to choke around a couple thousand.

<p id="13ea" class="graf graf--p graf-after--figure" style="margin: 30px0px0px; --x-height-multiplier: 0.375; --baseline-multiplier: 0.17; font-family: medium-content-serif-font,Georgia,Cambria,;">
  <strong class="markup--strong markup--p-strong" style="font-weight: bold;">Other benefits are :</strong>
</p>

<ul class="postList" style="margin: 21px0px0px; padding: 0px; list-style: nonenone; counter-reset: post0; color: rgba(0,0,0,0.84); font-family: medium-content-sans-serif-font,-apple-system,BlinkMacSystemFont,;">
  <li id="8c31" class="graf graf--li graf-after--p" style="margin-left: 30px; margin-bottom: 14px; --x-height-multiplier: 0.375; --baseline-multiplier: 0.17; font-family: medium-content-serif-font,Georgia,Cambria,;">
    Goroutines have grow-able segmented stacks. That means they will use more memory only when needed.
  </li>
  <li id="d180" class="graf graf--li graf-after--li" style="margin-left: 30px; margin-bottom: 14px; --x-height-multiplier: 0.375; --baseline-multiplier: 0.17; font-family: medium-content-serif-font,Georgia,Cambria,;">
    Goroutines have a faster startup time than threads.
  </li>
  <li id="9f10" class="graf graf--li graf-after--li" style="margin-left: 30px; margin-bottom: 14px; --x-height-multiplier: 0.375; --baseline-multiplier: 0.17; font-family: medium-content-serif-font,Georgia,Cambria,;">
    Goroutines come with built-in primitives to communicate safely between themselves (channels).
  </li>
  <li id="5aae" class="graf graf--li graf-after--li" style="margin-left: 30px; margin-bottom: 0px; --x-height-multiplier: 0.375; --baseline-multiplier: 0.17; font-family: medium-content-serif-font,Georgia,Cambria,;">
    Also, goroutines and OS threads do not have 1:1 mapping. A single goroutine can run on multiple threads. Goroutines are multiplexed into small number of OS threads.
  </li>
</ul>

## Easy to maintain

Go is a very easy language to read and maintain. It even has its own linter. It intentionally leaves out some features that make other languages hard to maintain.

If you look at the chart below. Its right in the middle between ease and speed. But still manages to beat java in both.

<img class="wp-image-977 alignnone size-full" src="https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_nlpyi256br71xmbwd1nlfg4271399669979853869.png" width="1010" height="610" srcset="https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_nlpyi256br71xmbwd1nlfg4271399669979853869.png 1010w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_nlpyi256br71xmbwd1nlfg4271399669979853869-300x181.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2019/06/1_nlpyi256br71xmbwd1nlfg4271399669979853869-768x464.png 768w" sizes="(max-width: 1010px) 100vw, 1010px" />

Go is a win-win for both situations.

Have you ever looked at a PHP error file with the giant list of warnings and fatal&#8217;s coming in from your live servers? You think, “how did that happen and why could I not prevent it?”

Because Go is compiled, it processes all the code on the spot. No more testing for noob issues like missing a semicolon or invalid types. You can&#8217;t compile with these kinds of basic errors. If there is an issue on live it&#8217;s usually data related which can be handled with validation.

## Where have I used go?

  * I&#8217;ve replaced most of my bash scripts 
      * Easier to write, syntax matches across the board for each flavor of Linux and even osx/windows
  * Cloud backups using a [little program][4] I am making for backblaze 
      * Simple
      * Effective
      * Free storage
  * Server monitoring daemons 
      * Raid monitoring
      * Proxy testing
      * Up-time monitor
      * Memory monitoring and clearing
      * Permission checking
      * Internet check
      * Telegram/slack notifications
      * Speed-test quality checks
      * Service restarts on failure or even a full system reboot
  * Web API&#8217;s 
      * JSON
      * Reading XML
  * Big data 
      * Concurrent processing of data and high speeds and large scales.
  * IOT software like on raspberry pi 
      * Managing io pins and reading sensor data
  * Server management 
      * Cluster reboots
      * Update Rollouts
      * Git hooks push live
  * PHP replacements 
      * Daemons
      * Time to first byte
      * Deadlines for hooks
      * Crawling web pages

## Summary

There are many uses for Go. I personally like to replace all of my bash scripts with it because it&#8217;s so simple to write and use. It has replaced my main use of PHP and I don&#8217;t think I could ever go back. Feel free to comment below. I will also be doing a blog series of tutorials followed by [meetup sessions][5] to discuss each tutorial.

Presentation &#8220;[What can golang do for you?&#8221;][6] from our meetup talk

&nbsp;

Make a donation using [Cash app][7], [Paypal][8]

 [1]: https://blog.silocitylabs.com/wp-content/uploads/2019/06/Golang-Users.png
 [2]: https://github.com/golang/go/wiki/GoUsers
 [3]: https://blog.silocitylabs.com/wp-content/uploads/2019/06/chart.png
 [4]: https://github.com/SiloCityLabs/B2Backup
 [5]: https://www.meetup.com/Buffalo-GoLang-Meetup-Group/
 [6]: https://blog.silocitylabs.com/wp-content/uploads/2019/06/What-can-golang-do-for-you.pptx
 [7]: https://cash.me/%24ldrrp/10
 [8]: https://www.paypal.me/ldrrp/10