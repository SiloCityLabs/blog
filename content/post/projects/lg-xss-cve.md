---
title: RDNS XSS On Industry Standard Looking Glass
author: Luis Rodriguez
type: post
date: 2015-01-21T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - looking-glass
  - xss
  - rdns
  - attack
  - vulnerability
  - php

---
My commitment to security is always a goal of mine. The 'industry standard' network looking glass now faces a pretty nasty XSS vulnerability which is [listed on github](https://github.com/telephone/LookingGlass/tree/4367733198698c4d314c51731e0b480b7d0383b3) and I discovered. An RDNS XSS was disclosed which has been patched by a temporary fix I applied to it. 

<!--more-->

{{< image src="/uploads/luis-portfolio/xss-lg-before.png" caption="An example of the XSS attack">}}
&nbsp;
{{< image src="/uploads/luis-portfolio/xss-lg-after.png" caption="An example of a patched LG">}}