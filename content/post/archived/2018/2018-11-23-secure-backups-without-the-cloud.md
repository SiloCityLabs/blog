---
title: Secure Backups without the cloud
author: Luis Rodriguez
type: post
date: 2018-11-23T17:30:03+00:00
url: /post/2018/11/23/secure-backups-without-the-cloud/
featured_image: /wp-content/uploads/2018/11/Untitled-150x150.png
categories:
  - Projects
tags:
  - backups
  - cloud
  - fireproof
  - iosafe
  - storage
  - toshiba
  - waterproof

---
<figure id="attachment_647" aria-describedby="caption-attachment-647" style="width: 678px" class="wp-caption aligncenter">[<img class="wp-image-647 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/11/22-501-073_R01-1024x768.jpg" alt="SOLO G3 4TB" width="678" height="509" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/11/22-501-073_R01-1024x768.jpg 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/22-501-073_R01-300x225.jpg 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/22-501-073_R01-768x576.jpg 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/22-501-073_R01.jpg 1280w" sizes="(max-width: 678px) 100vw, 678px" />][1]<figcaption id="caption-attachment-647" class="wp-caption-text">[ioSafe SOLO G3 4TB USB 3.0][2]</figcaption></figure>

I accidentally stumbled upon this solution while trying to find cloud providers to backup my 2TB&#8217;s worth of NAS backups. With every single cloud solution being out of price range for the size data I needed I had to look elsewhere. I came across the ioSafe Solo G3 enclosure. It boasts about being safe in every way possible, Fire Proof, Water Proof, and Theft Proof. I almost considered it too good to be true, seemed legit enough after watching some YouTube videos.

<!--more-->

<div style="text-align: center;">
  <iframe src="https://www.youtube.com/embed/u2yEVUMQyZY" width="560" height="315" frameborder="0" allowfullscreen="allowfullscreen"><span data-mce-type="bookmark" style="display: inline-block; width: 0px; overflow: hidden; line-height: 0;" class="mce_SELRES_start">﻿</span></iframe>
</div>

<div>
</div>

<div>
  During the video you will notice he manages to burn the enclosure under serious fire directly for a bit. There are also other videos of ioSafe products under direct flames for extended periods of time and direct water blasts from the fire department after. Here we have the drive pulled out from the crisp enclosure.
</div>

<div>
</div>

<div>
  <a href="https://blog.silocitylabs.com/wp-content/uploads/2018/11/Untitled-1.png"><img class="aligncenter wp-image-651 size-large" src="https://blog.silocitylabs.com/wp-content/uploads/2018/11/Untitled-1-1024x604.png" alt="" width="678" height="400" srcset="https://blog.silocitylabs.com/wp-content/uploads/2018/11/Untitled-1-1024x604.png 1024w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/Untitled-1-300x177.png 300w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/Untitled-1-768x453.png 768w, https://blog.silocitylabs.com/wp-content/uploads/2018/11/Untitled-1.png 1185w" sizes="(max-width: 678px) 100vw, 678px" /></a>
</div>

<div>
</div>

<div>
  Some of you may notice something you will immediately not like, TOSHIBA. I reached out to ioSafe to see if I can have it made with a brand I trust a little better. I have not had a response in 2 weeks. After some time stewing I really wanted this product but the drive inside really annoyed me. I have had more Toshiba&#8217;s die on me from experience in data centers than any other brand. I found a solution to this problem. I am going to show you the steps I am taking to protect the lifespan of this drive but be able to leave it plugged in for long term backups.
</div>

<div>
</div>

<div>
</div>

<div>
  <strong>Requirements</strong>: Linux OS, crontab, rsync, hdparm
</div>

<div>
</div>

<div>
</div>

  * I start off by plugging in my new drive and formatting it.

<pre>mkfs.ext4 /dev/sda1</pre>

  * After formatting I need to grab the device UUID for use later.

<div>
</div>

<pre>blkid /dev/sda1

/dev/sda1: UUID="fbb3b599-0a07-46f5-945d-0dbe6cac7639" TYPE="ext4" PARTLABEL="primary" PARTUUID="c15d00d1-214f-4517-8c2c-8037419ff0d4"</pre>

<div>
</div>

  * Now I have the UUID, lets move on to crontab. Create two cron tasks using `crontab -e`

<pre>0 6 * * * /mnt/4000/home/ldrrp/Scripts/fireproof-daily.sh
5 8 * * 0 /mnt/4000/home/ldrrp/Scripts/fireproof-weekly.sh
</pre>

<div>
  So according to this <code>5 8 * * 0</code> would run 8:05 AM every Sunday and <code>0 6 * * *</code> would run daily at 6 AM. Lets create the first script in the location you selected to save it.
</div>

<div>
</div>

  * This script will mount the drive, sync any new files or changed files only, then un-mount it and force it to sleep.

<pre>nano /mnt/4000/home/ldrrp/Scripts/fireproof-daily.sh</pre>

<div>
</div>

<div>
  <pre>#!/bin/bash

DRIVE_UUID=UUID-HERE

mount --uuid $DRIVE_UUID /mnt/fireproof
rsync –av /mnt/4000/ /mnt/fireproof/
umount /mnt/fireproof/
hdparm -Y /dev/disk/by-uuid/$DRIVE_UUID</pre>
</div>

<div>
</div>

  * This script will mount the drive, sync any new files or changed files, it will also delete any deleted files. Great for recovering a deleted file before sunday rolls around, then un-mount it and force it to sleep.

<div>
  <pre>nano /mnt/4000/home/ldrrp/Scripts/fireproof-weekly.sh</pre>
</div>

<div>
</div>

<div>
  <pre>#!/bin/bash

DRIVE_UUID=UUID-HERE

mount --uuid $DRIVE_UUID /mnt/fireproof
rsync –av --delete /mnt/4000/ /mnt/fireproof/
umount /mnt/fireproof/
hdparm -Y /dev/disk/by-uuid/$DRIVE_UUID</pre>
</div>

  * Last, make sure the directory exists for the mount point, run `mkdir /mnt/fireproof`

<div>
</div>

<div>
</div>

<div>
  I know some of you may point out that Linux automatically sleeps drives, but I just want that extra assurance that nothing will be waking it up creating unneeded wear on the drive. Also there will be no delay on this once the backup is complete.
</div>

 [1]: https://blog.silocitylabs.com/wp-content/uploads/2018/11/22-501-073_R01.jpg
 [2]: http://a.tra.li/TNQU