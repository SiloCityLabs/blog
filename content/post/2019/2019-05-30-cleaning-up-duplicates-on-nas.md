---
title: Cleaning up duplicates on NAS
author: Luis Rodriguez
type: post
date: 2019-05-30T23:30:44+00:00
url: /post/2019/05/30/cleaning-up-duplicates-on-nas/
featured_image: /wp-content/uploads/2019/02/Template_cleanup_icon.svg_.png
bigimg: [{src: "/uploads/2019/02/Template_cleanup_icon.svg_.png", desc: "Golang"}]
categories:
  - Projects

---

Over the last year me and my wife have been using the network server more frequently from its [initial install and setup][2]. We share 3-4 computers across the house and constantly move files onto the NAS. I have installed Nextcloud which I will detail in another post in the future. We have come across an issue that I have noticed with the network share. We moved all of our images and files over and ended up with tons of duplicates and a giant mess. I will show you what I am doing to clean that up. Our Nas is running on a ubuntu server. The following commands should work on most linux NAS servers.

<!--more-->

&nbsp;

First off, all my files are stored in /mnt/4000/home/*

&nbsp;

Get rid of all OS specific files that shouldnt be on the NAS. This will shorten the next step of finding duplicates.

<pre>find /mnt/4000/home -iname desktop.ini -delete</pre>

&nbsp;

Lets install a duplicates scanner:

### Installation on **Debian / Ubuntu**

<pre>sudo apt install fdupes</pre>

### Installation on **Fedora**

<pre>dnf install fdupes</pre>

After installing, we can move onto the next step. Run this command to get a list of all the duplicates.

<pre>fdupes -r /mnt/4000/home &gt; dupes.txt</pre>

Open this file up in a text editor and start hacking away. Delete all the lines of the files you wish to keep. Also delete any blank lines separating them. One example here:

<pre>/mnt/4000/home/Shared/Pictures/2016/Krystal Phone Backup 6-21-16/IMG_1622.JPG
/mnt/4000/home/Shared/Pictures/2016/pets/IMG_1622.JPG</pre>

I will be getting rid of the phone backup copy because I like to know this is part of pets.

<pre>/mnt/4000/home/Shared/Pictures/2016/Krystal Phone Backup 6-21-16/IMG_1428.JPG
/mnt/4000/home/Shared/Pictures/2016/IMG_1428.JPG</pre>

This one however ill keep the phone backup copy so that I don't make a messy folder inside of 2016

&nbsp;

Now that we are done editing our list, lets run the command on the list we want to delete.

<pre>while IFS= read -r file ; do rm -- "$file" ; done &lt; dupes.txt</pre>

&nbsp;

You may have to do this a few times to get any you missed.

Now that we have cleaned up all duplicates. We will be moving onto the next step. Finding empty files and folders.

Find all empty files

<pre>find /mnt/4000/home -type f -size 0</pre>

File all empty folders

<pre>find /mnt/4000/home -type d -empty</pre>

&nbsp;

Delete any folders or files you want here, Be careful not to delete needed files for programs.

&nbsp;

After all of this was done we managed to save 58GB of storage on our NAS.

 [2]: https://blog.silocitylabs.com/post/2019/02/21/home-server-setup/