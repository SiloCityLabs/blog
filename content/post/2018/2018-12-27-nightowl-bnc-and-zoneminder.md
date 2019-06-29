---
title: Nightowl BNC and Zoneminder
author: Luis Rodriguez
type: post
date: 2018-12-27T23:30:18+00:00
url: /post/2018/12/27/nightowl-bnc-and-zoneminder/
featured_image: /wp-content/uploads/2018/08/images.jpg
categories:
  - Projects
tags:
  - nightowl
  - security cameras
  - zoneminder

---
### Nightowl BNC
{{< image src="/uploads/2018/08/images.jpg" alt="">}}

#### Details:

These are just about every cheap camera sold in store. You can buy these at home depot, walmart, and other retailers. Your best bet is craigslist you can get them for like $10 each. They use a basic analog signal you can hookup to a TV with the right connector. The main connector is called a BNC connector. The downside to these cameras is that they usually cap out at 720p. You will also need a PCI or usb device to hook it up to your server. Unless you already have these cameras laying around, I WOULD NOT RECOMMEND this setup for purchase. Its cheaper to get a 1080p camera on amazon for $20 with less wiring and no pci card.

<!--more-->

#### Setup:

The PCI Card I used was "Hauppauge 649000-02". It sells on ebay for around $20 and has 4 inputs. Something to look out for in these cards, If you have one main chip on the card with 4 inputs it will likely degrade the frames per camera. Something with one chip for 2 inputs would be ideal. 8 Input cards with 4 chips would be ideal but unrealistic in price. I compromised with 4 input cards and bought two of them.

Due to the name change of the pci card and the colons in the path, you will need to make a symlink to point to it. As you can see in the image my /camera/pci0 points to the path to the device.

{{< image src="uploads/2018/12/1-4.png" alt="">}}

&nbsp;

You can see I have all the cameras pointing to it as the source.

{{< image src="/uploads/2018/12/1-5.png" alt="">}}

With the following settings for this pci card.

{{< image src="/uploads/2018/12/1-6.png" alt="">}}

{{< image src="/uploads/2018/12/2-2.png" alt="">}}