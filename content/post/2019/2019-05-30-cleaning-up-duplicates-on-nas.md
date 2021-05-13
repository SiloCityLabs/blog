---
title: Cleaning up NAS
author: Luis Rodriguez
type: post
date: 2021-05-13T01:30:44+00:00
url: /post/2021/05/13/cleaning-up-duplicates-on-nas/
featured_image: /wp-content/uploads/2019/02/Template_cleanup_icon.svg_.png
bigimg: [{src: "/uploads/2019/02/Template_cleanup_icon.svg_.png", desc: "Golang"}]
categories:
  - Projects

---
Updated in May 2021

Over the last few years me and my wife have been using the network server more frequently from its initial install and setup. We share 3-4 computers across the house and constantly move files onto the NAS. I have installed Nextcloud which I will detail in another post in the future. We have come across an issue that I have noticed with the network share. We moved all of our images and files over and ended up with tons of duplicates and a giant mess. I will show you what I am doing to clean that up. Our Nas is running on a ubuntu server. The following commands should work on most linux NAS servers.

<!--more-->

First off, all my files are stored in /mnt/4000/home/*

Get rid of all OS specific files that shouldnt be on the NAS. This will shorten the next step of finding duplicates.

```
find ./ -iname desktop.ini -delete
```

Lets install a duplicates scanner:

### Installation on **Debian / Ubuntu**

```
apt install fdupes
```

### Installation on **Fedora**

```
dnf install fdupes
```

After installing, we can move onto the next step. Run this command to get a list of all the duplicates.

```
fdupes -r ./ > dupes.txt
```

Open this file up in a text editor and start hacking away. Delete all the lines of the files you wish to keep. Also delete any blank lines separating them. One example here:

```
./Shared/Pictures/2016/Krystal Phone Backup 6-21-16/IMG_1622.JPG
./Shared/Pictures/2016/pets/IMG_1622.JPG
```

I will be getting rid of the phone backup copy because I like to know this is part of pets.

```
./Shared/Pictures/2016/Krystal Phone Backup 6-21-16/IMG_1428.JPG
./Shared/Pictures/2016/IMG_1428.JPG
```

This one however ill keep the phone backup copy so that I don't make a messy folder inside of 2016

Now that we are done editing our list, lets run the command on the list we want to delete.

```
while IFS= read -r file ; do rm -- "$file" ; done < dupes.txt
```

You may have to do this a few times to get any you missed.

Now that we have cleaned up all duplicates. We will be moving onto the next step. Finding empty files and folders.

Find all empty files

```
find ./ -type f -size 0 > emptyfiles.txt
```

File all empty folders

```
find ./ -type d -empty > emptyfolders.txt
```

Large files

```
find ./ -xdev -type f -size +100M > largefiles.txt
```

Delete any folders or files you want here, Be careful not to delete needed files for programs.

After all of this was done we managed to save 58GB of storage on our NAS.
