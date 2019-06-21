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
![Frig off, Randy](/uploads/2018/01/newsletter-popup-e1515447774854.jpg)

&nbsp;

I hate &#8220;exit intent&#8221; popups. The kind that automatically open when your mouse leaves the page begging &#8220;please don&#8217;t go!&#8221;. They&#8217;re annoying. They&#8217;re irrelevant (I&#8217;m following instructions, not closing). They interrupt the content I&#8217;m trying to read. So I made a short [Userscript to globally block them][2].

<!--more-->

![Soon to be blocked](/uploads/2018/01/optin-monster.png)

Specifically I want to block Ouibounce ([demo][4]) and OptIn Monster ([demo][5]). Bounce Exchange is another infamous one but I couldn&#8217;t find a demo to test on. [This Metafilter topic][6] tipped me off to the main culprits: Javascript&#8217;s mouseleave and mouseout events. So many sites have these popups that it&#8217;s not practical to block on a site-by-site basis. Disabling all Javascript is too heavy handed.

While researching I was able to block the modals per-website using uBlock Origin&#8217;s _script:inject(&#8230;)_. Specifically the _[addEventListener-defuser.js][7]_ is used in some filters to prevent JS click popups. Unfortunately the script injection intentionally doesn&#8217;t accept wildcards so I can&#8217;t use it globally.

I made a Userscript to do the same thing and it&#8217;s working for Greasemonkey and Tampermonkey. This runs on **all** http and https pages and overrides the mouseleave and mouseout events. If some popup still appear try adding _//@run-at document-start_ so that it registers earlier. It will probably break some UIs but it&#8217;s an acceptable tradeoff to me. URLs can be excluded if needed. I don&#8217;t really plan on supporting such a simple script so it&#8217;s &#8220;licensed&#8221; under 0BSD. You&#8217;re free to use, copy, and edit. Have fun:

[https://openuserjs.org/scripts/Maave/Exit\_Intent\_defuser][2]

EDIT: I&#8217;ve found a few features that broke. Excludes have been added.

  * YouTube&#8217;s preview on the seek bar doesn&#8217;t go away.
  * Wikipedia&#8217;s reference hover popup.

 [1]: /uploads/2018/01/newsletter-popup.jpg
 [2]: https://openuserjs.org/scripts/Maave/Exit_Intent_defuser
 [3]: /uploads/2018/01/optin-monster.png
 [4]: https://carlsednaoui.github.io/ouibounce/
 [5]: https://optinmonster.com/features/exit-intent/
 [6]: https://ask.metafilter.com/280379/Stop-popping-up-when-Im-about-to-switch-the-tab
 [7]: https://github.com/uBlockOrigin/uAssets/issues/692#issuecomment-329922065
