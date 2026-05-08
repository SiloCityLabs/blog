---
title: "Bypassing YouTube home page with Tasker"
author: maave
type: post
date: 2026-04-14T00:00:00+00:00
url: /post/2026/05/08/tasker-open-intent/
draft: false
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

The foundation for these "bookmarks" is [Intent](https://developer.android.com/reference/android/content/Intent). This is Android's way of launching other apps or passing data between apps.

It's also the way that apps act as the default viewer for certain URLs called [deep links](https://developer.android.com/training/app-links/create-deeplinks). If you have YouTube installed on your phone, then a link like `https://www.youtube.com/@Silocitylabs/featured` will open directly in the app rather than the browser. The app also handles `https://www.youtube.com/feed/subscriptions` and `https://www.youtube.com/feed/library`.

Want the easy mode? Just visit the link in the browser and _Add To Home Screen_. Edit the shortcut to swap the icon to YouTube. Done! This has a downside - it leaves a tab open in the browser, and sometimes fails when that tab is still active. Also I can't find a way to edit an existing shortcut link, short of deleting and recreating it.

{{< image src="/uploads/2026/yt-subscription-link.webp" alt="remaining tab on Firefox. Don't worry about the other ∞ tabs">}}

We can also open intent URIs from the browser. `intent://feed/subscriptions#Intent;scheme=vnd.youtube;package=com.google.android.youtube;end` can be opened in the URL bar or [from an anchor link](intent://feed/subscriptions#Intent;scheme=vnd.youtube;package=com.google.android.youtube;end). However I wasn't able to add this link to my home screen.

## Tasker

For a cleaner launch I used [Tasker](https://play.google.com/store/apps/details?id=net.dinglisch.android.taskerm&hl=en-US). In Tasker go to the Task tab, create a new task `+`, add an action `+`. Pick "System" then "Send Intent". Then fill in these fields:

Action: `android.intent.action.VIEW`

Package: `com.google.android.youtube`

Data: `https://www.youtube.com/feed/subscriptions`

OR 

Data: `vnd.youtube://feed/subscriptions`


{{< image src="/uploads/2026/tasker-action-intent.webp" alt="tasker intent settings">}}

Set an icon with the grid icon. 

{{< image src="/uploads/2026/tasker-add-icon.webp" alt="grid icon at the bottom of Task Edit page">}}

Finally, under the 3-dot menu, tap _Add To Launcher_ to add the icon. Tada! 

## Demo

Here you can see both methods in-action. I open the regular YouTube app icon, then Library (browser shortcut), then Subscriptions (Tasker shortcut). The browser shortcut failed the first time because I still had the tab open. Tasker works every time and doesn't leave a process running.

{{< youtube UZ68EdtRlnE >}}


## Finding Intents in APKs

Big whoops. YouTube already exposes the subscription Activity by long-pressing the YouTube homescreen icon. So let's find intents which aren't graciously given to the user. Like where does `feed/library` come from?

Intents are defined in the application and we can decompile the app for more info. The file `AndroidManifest.xml` contains `intent-filter` definitions and path filters, a.k.a. which URIs the app intercepts.

I started with apktool to get the manifest file but quickly switched to [JADX](https://github.com/skylot/jadx) so that I could view decompiled Java source code.

### AndroidManifest.xml

After I open the APK with JADX I can see file `res/AndroidManifest.xml` and found the URI intent-filters. However the path filter is a wildcard `android:pathPattern=".*"` so it handles every "youtube.com" link all at once and we don't see "subscriptions" yet. 

```xml
<intent-filter android:autoVerify="true">
    <action android:name="android.intent.action.VIEW"/>
    <action android:name="android.media.action.MEDIA_PLAY_FROM_SEARCH"/>
    <category android:name="android.intent.category.DEFAULT"/>
    <category android:name="android.intent.category.BROWSABLE"/>
    <data android:scheme="http"/>
    <data android:scheme="https"/>
    <data android:host="youtube.com"/>
    <data android:host="www.youtube.com"/>
    <data android:host="m.youtube.com"/>
    <data android:host="youtu.be"/>
    <data android:pathPattern=".*"/>
</intent-filter>
<intent-filter>
    <action android:name="android.intent.action.VIEW"/>
    <action android:name="android.media.action.MEDIA_PLAY_FROM_SEARCH"/>
    <category android:name="android.intent.category.DEFAULT"/>
    <category android:name="android.intent.category.BROWSABLE"/>
    <data android:scheme="vnd.youtube"/>
    <data android:scheme="vnd.youtube.launch"/>
</intent-filter>
```

But I do find the activities that YouTube exposes to the user. These are the shortcuts that can be find via the home screen icon.

```xml
<activity-alias
    android:name="com.google.android.youtube.app.honeycomb.Shell$HomeActivity"
    android:exported="true"
    android:targetActivity="com.google.android.apps.youtube.app.watchwhile.MainActivity">
    <intent-filter>
        <action android:name="android.intent.action.MAIN"/>
        <category android:name="android.intent.category.DEFAULT"/>
        <category android:name="android.intent.category.LAUNCHER"/>
    </intent-filter>
    <intent-filter>
        <action android:name="com.google.android.youtube.action.open.search"/>
        <category android:name="android.intent.category.DEFAULT"/>
    </intent-filter>
    <intent-filter>
        <action android:name="com.google.android.youtube.action.open.subscriptions"/>
        <category android:name="android.intent.category.DEFAULT"/>
    </intent-filter>
    <intent-filter>
        <action android:name="com.google.android.youtube.action.open.shorts"/>
        <category android:name="android.intent.category.DEFAULT"/>
    </intent-filter>
    <meta-data
        android:name="android.app.shortcuts"
        android:resource="@xml/main_shortcuts"/>
</activity-alias>
```

I can also see which activities are allowed to be called by other apps by searching `android:exported="true"`. 

```xml
<activity-alias
    android:name="com.google.android.apps.youtube.app.application.Shell$UrlActivity"
    android:exported="true"
    android:targetActivity="com.google.android.apps.youtube.app.application.Shell_UrlActivity"/>
```

`Shell_UrlActivity` is the main URL parsing activity.

### Half-hearted decomp

The rest of the path goes into the data field which is processed in code. Unforunately decomped code is hard to read. When the code is compiled, symbols are stripped, meaning variable and function names are not preserved. In the decompiled source files, all the var and func names are freshly-generated nonsense. It's not "obfuscated" but it is garbled. Discombobulating this is challenging.

One option for reverse engineering is to rename vars one at a time. JADX GUI is pretty good for refactoring and jumping through references. LLMs are also good at jumping through this unnamed code.

Sidenote: I tried importing the decompiled source into VS Code but refactoring wasn't fully working. I exported source from JADX, opened in VS Code, and manually added `settings.gradle` and `build.gradle`. This allowed me to jump through reference and view the source without errors but when I refactored a var name the new name didn't propagate to other parts of the code.

This is a pain in the butt and I got bored of nonsense code. I decided to search for static strings in the source. After a few combos I stumbled across `FEsubscriptions`.

```java
return "FEshared".equals(str) || "FElibrary".equals(str) || "FEoffline_what_to_watch".equals(str) || "FEsubscriptions".equals(str) || "FEwhat_to_watch".equals(str) || "FEactivity".equals(str);
```

Using these I found a few new endpoints to open:

- https://www.youtube.com/feed/offline_what_to_watch   (saved offline videos)
- https://www.youtube.com/feed/what_to_watch  (the regular home screen)
- https://www.youtube.com/feed/activity  (notifications)

So decomping code to find Activities to launch is probably more trouble than it's worth unless you're exploit-hunting. You can find starting points in AndroidManifest.xml but the rest of the data is parsed in Java.