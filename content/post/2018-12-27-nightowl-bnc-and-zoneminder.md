---
title: Nightowl BNC and Zoneminder
author: Luis Rodriguez
type: post
date: 2018-12-27T23:30:18+00:00
url: /post/2018/12/27/nightowl-bnc-and-zoneminder/
featured_image: /wp-content/uploads/2018/08/images-150x150.jpg
categories:
  - Side Projects
tags:
  - nightowl
  - security cameras
  - zoneminder

---
### Nightowl BNC

[<img class="aligncenter size-full wp-image-567" src="https://blog.silocitylabs.com/wp-content/uploads/2018/08/images.jpg" alt="" width="201" height="251" />][1]

#### Details:

These are just about every cheap camera sold in store. You can buy these at home depot, walmart, and other retailers. Your best bet is craigslist you can get them for like $10 each. They use a basic analog signal you can hookup to a TV with the right connector. The main connector is called a BNC connector. The downside to these cameras is that they usually cap out at 720p. You will also need a PCI or usb device to hook it up to your server. Unless you already have these cameras laying around, I WOULD NOT RECOMMEND this setup for purchase. Its cheaper to get a 1080p camera on amazon for $20 with less wiring and no pci card.

<!--more-->

#### Setup:

The PCI Card I used was &#8220;Hauppauge 649000-02&#8221;. It sells on ebay for around $20 and has 4 inputs. Something to look out for in these cards, If you have one main chip on the card with 4 inputs it will likely degrade the frames per camera. Something with one chip for 2 inputs would be ideal. 8 Input cards with 4 chips would be ideal but unrealistic in price. I compromised with 4 input cards and bought two of them.

Due to the name change of the pci card and the colons in the path, you will need to make a symlink to point to it. As you can see in the image my /camera/pci0 points to the path to the device.

[<img class="aligncenter size-full wp-image-718" src="https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-4.png" alt="" width="838" height="129" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-4.png 838w, https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-4-300x46.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-4-768x118.png 768w" sizes="(max-width: 838px) 100vw, 838px" />][2]

&nbsp;

You can see I have all the cameras pointing to it as the source.

[<img class="aligncenter size-full wp-image-719" src="https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-5.png" alt="" width="386" height="121" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-5.png 386w, https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-5-300x94.png 300w" sizes="(max-width: 386px) 100vw, 386px" />][3]

With the following settings for this pci card.

[<img class="aligncenter size-full wp-image-720" src="https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-6.png" alt="" width="527" height="482" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-6.png 527w, https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-6-300x274.png 300w" sizes="(max-width: 527px) 100vw, 527px" />][4]

[<img class="aligncenter size-full wp-image-721" src="https://blog.silocitylabs.com/wp-content/uploads/2018/12/2-2.png" alt="" width="527" height="482" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/12/2-2.png 527w, https://blog.silocitylabs.com/wp-content/uploads/2018/12/2-2-300x274.png 300w" sizes="(max-width: 527px) 100vw, 527px" />][5]

 [1]: https://blog.silocitylabs.com/wp-content/uploads/2018/08/images.jpg
 [2]: https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-4.png
 [3]: https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-5.png
 [4]: https://blog.silocitylabs.com/wp-content/uploads/2018/12/1-6.png
 [5]: https://blog.silocitylabs.com/wp-content/uploads/2018/12/2-2.png