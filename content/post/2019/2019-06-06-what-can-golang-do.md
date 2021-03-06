---
title: What can golang do?
author: Luis Rodriguez
type: post
date: 2019-06-06T20:00:28+00:00
url: /post/2019/06/06/what-can-golang-do/
featured_image: /wp-content/uploads/2019/06/1_30aoNxlSnaYrLhBT0O1lzw.png
bigimg: [{src: "/uploads/2019/06/golang.png", desc: "Golang"}]
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

Nothing makes a developer crazier than a new programming language, right? So, I started learning Go about 6 months ago and now I am going to share with you my reasons for doing so.

<!--more-->

## Companies using Golang

{{< image src="/uploads/2019/06/golang-users.png" alt="Golang Users">}}

Google - Go is designed and supported by Google. Google has one of the largest cloud infrastructures in the world and it is massively scaled.

Apple - Big data, analytics, infrastructure software.

Cloudflare - Used in infrastructure and other various locations. Has decreased time to first byte by 90%.

Gitlab - Handles all RPC calls made by git

AT&T - Network management, server applications, web based API's

Netflix - Caching management

Tesla - Server management, enterprise software, production lines.

SpaceX - Used in rocket telemetry

The [list](https://github.com/golang/go/wiki/GoUsers) is over thousands long

## Hardware limitations

7 years ago I had a 4ghz processor with 8 cores. Today my desktop has a 4 core processor with 3.9ghz. You're probably looking at these numbers like, "why did he downgrade?" I did not. The benchmarks actually put my newer processor ahead by 66% because of multi-threading capabilities.

Nearly one decade and the hardware has not changed much besides one key thing. Multi-threading, and now possibly even more cores. Everything else has pretty much been steady and slow growth.

{{< image src="/uploads/2019/06/chart.png" alt="CPU Chart">}}

Even multi-threading has its limits. With that said, we fall to software to fix these issues. We must increase efficiency and performance of our software.

You're probably thinking "well python is the ultimate language for AI." Sure it has many existing libraries, but it’s also old and slower. On some benchmarks even 40x slower.

## Go runs direct

No crazy configuration is needed. Compile it for Linux and download any minimal Linux and let it do its thing. You don't need any added software on the server it will run on. It'll run identical on nearly any hardware. Yes, it even compiles for that Windows server your boss says you need for “compliance's.”

Starting a compiled Go application is as simple as a file transfer and "./myprogram"

Go has the benefit of having its own compiler for each architecture. It is not bogged down by a VM like java or a processor like Python or PHP. It is like C/C++ with ease of learning.

## Go routines

This is where the magic happens. Routines help create concurrency you need to speed up your application where others fail.

{{< image src="/uploads/2019/06/1_nfojvbkdrkxz0zdbu4ysna1724252319353850356.jpeg" alt="Go Routines">}}

Routines in go startup faster and even use less memory to start than other languages like python and java. They can get complicated quickly in other languages but with Go it’s as simple as:

```go functionName()```

You could spin up millions of routines before you start having bottlenecks in hardware. Meanwhile Languages like java will start to choke around a couple thousand.

**Other benefits are:**

 - Goroutines have grow-able segmented stacks. That means they will use more memory only when needed.
 - Goroutines have a faster startup time than threads.
 - Goroutines come with built-in primitives to communicate safely between themselves (channels).
 - Also, goroutines and OS threads do not have 1:1 mapping. A single goroutine can run on multiple threads. Goroutines are multiplexed into small number of OS threads.

## Easy to maintain

Go is a very easy language to read and maintain. It even has its own linter. It intentionally leaves out some features that make other languages hard to maintain.

If you look at the chart below. Its right in the middle between ease and speed. But still manages to beat java in both.


{{< image src="/uploads/2019/06/1_nlpyi256br71xmbwd1nlfg4271399669979853869.png" alt="Languages Chart">}}

Go is a win-win for both situations.

Have you ever looked at a PHP error file with the giant list of warnings and fatal's coming in from your live servers? You think, “how did that happen and why could I not prevent it?”

Because Go is compiled, it processes all the code on the spot. No more testing for noob issues like missing a semicolon or invalid types. You can't compile with these kinds of basic errors. If there is an issue on live it's usually data related which can be handled with validation.

## Where have I used go?

  * I've replaced most of my bash scripts 
      * Easier to write, syntax matches across the board for each flavor of Linux and even osx/windows
  * Cloud backups using a [little program](https://github.com/SiloCityLabs/B2Backup) I am making for backblaze 
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
  * Web API's 
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

There are many uses for Go. I personally like to replace all of my bash scripts with it because it's so simple to write and use. It has replaced my main use of PHP and I don't think I could ever go back. Feel free to comment below. I will also be doing a blog series of tutorials followed by [meetup sessions](https://www.meetup.com/Buffalo-GoLang-Meetup-Group/) to discuss each tutorial.

Presentation "[What can golang do for you?"](/uploads/2019/06/What-can-golang-do-for-you.pptx) from our meetup talk

&nbsp;

Make a donation using [Cash app](https://cash.me/%24ldrrp/10), [Paypal](https://www.paypal.me/ldrrp/10)
