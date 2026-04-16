---
title: "Bypassing YouTube home page with Tasker"
author: maave
type: post
date: 2026-04-14T00:00:00+00:00
url: /post/2026/04/14/tasker-open-intent/
draft: true
categories:
  - Android
tags:
  - Android
  - Tasker
  - YouTube
---


{{< image src="/uploads/2026/negative-chamfer-header.webp" alt="OpenSCAD pins with positive and negative chamfer">}}

Remember when you could open your computer or browser and not immediately get "suggested content"? Most apps nowadays open to a "home" or "feed" page. I miss bookmarks, so I'm bringing them back. I'm bypassing the slop conveyor that is YouTube's home page. I'm jumping straight to the subscriptions - the content I actually asked to see.

This was easy to automate in Tasker using _Android intents_.

<!--more-->

## Intents

The foundation for these "bookmarks" is [Intent](https://developer.android.com/reference/android/content/Intent). This is Android's way of launching other apps or passing data between apps. It's also the way that apps act as the default viewer for certain URLs. If you have YouTube installed on your phone, then a link like `https://www.youtube.com/@Silocitylabs/featured` will open directly in the app rather than the browser. The app also handles `https://www.youtube.com/feed/subscriptions` and `https://www.youtube.com/feed/library`.

Want the easy mode? Just visit the link in the browser and _Add To Home Screen_. Edit the shortcut to swap the icon to YouTube. Tada. This has a downside - it leaves a tab open in the browser, and sometimes fails when that tab is still active. Also I can't find a way to edit an existing shortcut link, short of deleting and recreating.

{{< image src="/uploads/2026/yt-subscription-link.webp" alt="remaining tab on Firefox. Don't worry about the other ∞ tabs">}}

## Tasker

For a cleaner launch I used [Tasker](https://play.google.com/store/apps/details?id=net.dinglisch.android.taskerm&hl=en-US). In Tasker go to the Task tab, create a new task `+`, add an action `+`. Pick "System" then "Send Intent".

Action: `android.intent.action.VIEW`
Data: `https://www.youtube.com/feed/subscriptions`
Package: `com.google.android.youtube`

{{< image src="/uploads/2026/tasker-action-intent.webp" alt="tasker intent settings">}}