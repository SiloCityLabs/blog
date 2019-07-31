---
title: Geo Location Micro API Service
type: post
date: 2019-07-22T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - golang
  - geo
  - api

---
A micro api drop in service written in golang, Designed to quickly integrate into any system for any project. Uses very little resources and can handle millions of requests per second. The entire projects [source code](https://github.com/SiloCityLabs/geo.bntech.io) is posted on github.

You can include a package written for Golang projects to quickly get started with a simple function.

<!--more-->

`
go get github.com/silocitylabs/geo.bntech.io/geo
`

```
// Service url, IP to lookup, Attempts
results := geo.GetIPData("http://127.0.0.1:3000", "4.4.4.4", 3)


fmt.Printf(results)

```