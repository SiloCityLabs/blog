---
title: Blocking exit intent popups
author: maave
type: post
date: 2018-01-10T03:17:12+00:00
url: /post/2018/01/10/blocking-exit-intent-popups/
featured_image: /wp-content/uploads/2018/01/optin-monster.png
categories:
  - Projects
tags:
  - Javascript
  - Userscript

---
{{< image src="/uploads/2018/01/newsletter-popup-e1515447774854.jpg" alt="Frig off, Randy">}}

&nbsp;

I hate "exit intent" popups. The kind that automatically open when your mouse leaves the page begging "please don't go!". They're annoying. They're irrelevant (I'm often following a tutorial, not closing the tab). They interrupt the content I'm trying to read. So I made a short [Userscript to globally block them][2].

<!--more-->

{{< image src="/uploads/2018/01/optin-monster.png" alt="Soon to be blocked">}}

Specifically I want to block Ouibounce ([demo][4]) and OptIn Monster ([demo][5]). Bounce Exchange is another infamous one but I couldn't find a demo to test on. [This Metafilter topic][6] tipped me off to the main culprits: Javascript's mouseleave and mouseout events. So many sites have these popups that it's not practical to block on a site-by-site basis. Disabling all Javascript is too heavy handed.

While researching I was able to block the modals per-website using uBlock Origin's _script:inject(...)_. Specifically the _[addEventListener-defuser.js][7]_ is used in some filters to prevent JS click popups. Unfortunately the script injection intentionally doesn't accept wildcards so I can't use it globally.

I made a Userscript to do the same thing and it's working for Greasemonkey and Tampermonkey. This runs on **all** http and https pages and overrides the mouseleave and mouseout events. If some popup still appear try adding _//@run-at document-start_ so that it registers earlier. It will probably break some UIs but it's an acceptable tradeoff to me. URLs can be excluded if needed. I don't really plan on supporting such a simple script so it's "licensed" under 0BSD. You're free to use, copy, and edit. Have fun:

[https://openuserjs.org/scripts/Maave/Exit\_Intent\_defuser][2]

EDIT: I've found a few features that broke. Excludes have been added.

  * YouTube's preview on the seek bar doesn't go away.
  * Wikipedia's reference hover popup.

 [1]: /uploads/2018/01/newsletter-popup.jpg
 [2]: https://openuserjs.org/scripts/Maave/Exit_Intent_defuser
 [3]: /uploads/2018/01/optin-monster.png
 [4]: https://carlsednaoui.github.io/ouibounce/
 [5]: https://optinmonster.com/features/exit-intent/
 [6]: https://ask.metafilter.com/280379/Stop-popping-up-when-Im-about-to-switch-the-tab
 [7]: https://github.com/uBlockOrigin/uAssets/issues/692#issuecomment-329922065
