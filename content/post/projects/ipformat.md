---
title: Golang IP Formatting Package
author: Luis Rodriguez
type: post
date: 2019-06-25T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - golang
  - ip
  - cidr

---
Package written in Golang to process and format ip Addresses to and from ipv4/ipv6. Can also take ranges in cidr notation. The entire projects [source code](https://github.com/SiloCityLabs/ipFormat) is posted on github.

You can include a package written for Golang projects to quickly get started with a simple function.

<!--more-->

`
go get github.com/silocitylabs/ipFormat
`

```
package main

import (
	"fmt"

	"github.com/silocitylabs/ipFormat"
)

func main() {
	ip1, _ := ipFormat.New("192.168.0.2/24")
	ip2, _ := ipFormat.New("192.168.0.6")

	ip1, _ = ip1.ToV6()
	ip2, _ = ip2.ToV6()

	fmt.Println(ip1.Address)
	fmt.Println(ip1.CIDR)
	fmt.Println(ip2.Address)
}

```