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


{{< image src="/uploads/2026/tasker-header.webp" alt="YouTube icons labelled Library and Subscriptions">}}

Remember when you could open your computer or browser and not immediately get "suggested content"? Most apps nowadays open to a "home" or "feed" page. I miss bookmarks, so I'm bringing them back. I'm bypassing the slop conveyor that is YouTube's home page. I'm jumping straight to the subscriptions - the content I actually asked to see.

This was easy to automate in Tasker using _Android intents_.

<!--more-->

## Intents

The foundation for these "bookmarks" is [Intent](https://developer.android.com/reference/android/content/Intent). This is Android's way of launching other apps or passing data between apps. It's also the way that apps act as the default viewer for certain URLs called [deep links](https://developer.android.com/training/app-links/create-deeplinks). If you have YouTube installed on your phone, then a link like `https://www.youtube.com/@Silocitylabs/featured` will open directly in the app rather than the browser. The app also handles `https://www.youtube.com/feed/subscriptions` and `https://www.youtube.com/feed/library`.

Want the easy mode? Just visit the link in the browser and _Add To Home Screen_. Edit the shortcut to swap the icon to YouTube. Done! This has a downside - it leaves a tab open in the browser, and sometimes fails when that tab is still active. Also I can't find a way to edit an existing shortcut link, short of deleting and recreating it.

{{< image src="/uploads/2026/yt-subscription-link.webp" alt="remaining tab on Firefox. Don't worry about the other ∞ tabs">}}

## Tasker

For a cleaner launch I used [Tasker](https://play.google.com/store/apps/details?id=net.dinglisch.android.taskerm&hl=en-US). In Tasker go to the Task tab, create a new task `+`, add an action `+`. Pick "System" then "Send Intent". Then fill in these fields:

Action: `android.intent.action.VIEW`

Data: `https://www.youtube.com/feed/subscriptions`

Package: `com.google.android.youtube`

{{< image src="/uploads/2026/tasker-action-intent.webp" alt="tasker intent settings">}}

Set an icon with the grid icon. 

{{< image src="/uploads/2026/tasker-add-icon.webp" alt="grid icon at the bottom of Task Edit page">}}

Finally, under the 3-dot menu, tap _Add To Launcher_ to add the icon. Tada! 

## Demo

Here you can see both methods in-action. First I open the regular YouTube app icon, then Library (browser shortcut), then Subscriptions (Tasker shortcut). The browser shortcut failed the first time because I still had the tab open. Tasker works every time and doesn't leave a process running.

TODO: update this URL with real vid link

{{< youtube XZ2zcGssKz0 >}}