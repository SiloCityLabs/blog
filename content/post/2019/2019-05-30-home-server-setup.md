---
title: Home Server Setup.
author: Luis Rodriguez
type: post
date: 2019-05-30T23:30:33+00:00
url: /post/2019/05/30/home-server-setup/
featured_image: /wp-content/uploads/2019/02/55-6723-mediasmart_600.jpg
categories:
  - Projects
tags:
  - axis
  - file share
  - home server
  - ibaby
  - media server
  - raid
  - security
  - security cameras
  - speedtest
  - zoneminder

---
Today I will talk about my server setup by request. I setup my server originally for Security cameras because I was tired of proprietary software that every different camera has. My setup right now includes 4 entirely different types of cameras. None of them officially support ONVIF. Every one of my cameras was free so quality was not something I cared for. Today, my server now has my Windows File share (Samba), offsite Sync, Software Raid, Crontab Network Speed monitoring and [zoneminder][1].

<!--more-->

## Security Camera Software

ZoneMinder - "A full-featured, open source, state-of-the-art video surveillance software system. Monitor your home, office, or wherever you want. Using off the shelf hardware with any camera, you can design a system as large or as small as you need."

## My Cameras

  * [Axis 2130 PTZ][2]
  * [Rosewill RSCM-13101B][3]
  * [Nightowl BNC x6][4]
  * [Genius WideCam 320 Webcam][5]
  * [BV Tech 4MP 2k][6]

## Onsite Fireproof Sync

To avoid paying for cloud services I purchased a fireproof hard drive. I set it up to sync with my raid group and its all set. You can view that [article here.][7]

## Offsite Sync

I use offsite sync for the security cameras using rsync and a remote ftp server. Unfortunately it would be too expensive to backup my entire media server to the cloud so I only use it for remote storing the camera footage in the event of a theft.

To start this you will need an ftp server credentials to use with the following script. After doing so you can just add it to `crontab -e`

Crontab file:

```
@reboot /mnt/4000/home/ldrrp/Scripts/ftpsync.sh
```

Script:

```
#!/bin/bash
HOST='hostedftp.com'
USER='username'
PASS='password'
TARGETFOLDER='/'
SOURCEFOLDER='/var/cache/zoneminder/'
while inotifywait -qqr "$SOURCEFOLDER"; do
lftp -f "
open $HOST
user $USER $PASS
lcd $SOURCEFOLDER
set ftp:ssl-allow no
mirror --no-symlinks --reverse --delete --verbose $SOURCEFOLDER $TARGETFOLDER
bye
"
done
```

&nbsp;

## Software Raid

I don't want to spend too much on this setup so software raid we go. I use Raid 5 for its redundancy but increase capacity. I will talk about what I did to get this quickly setup. Depending on your drive count and size this can take anywhere from 2 hours to 24 hours. You can read more about the [setup I used][8].

## File Share

Setting up a file share is fairly simple, First download and install samba

`apt-get install samba`

Then edit the main configuration file:

```
[global]
   workgroup = WORKGROUP
   dns proxy = no
   log file = /var/log/samba/log.%m
   max log size = 1000
   syslog = 0
   panic action = /usr/share/samba/panic-action %d
   server role = standalone server
   passdb backend = tdbsam
   obey pam restrictions = yes
   unix password sync = yes
   passwd program = /usr/bin/passwd %u
   passwd chat = *Enter\snew\s*\spassword:* %n\n *Retype\snew\s*\spassword:* %n\n *password\supdated\ssuccessfully* .
   pam password change = yes
   map to guest = bad user
   usershare allow guests = yes
   socket options = TCP_NODELAY IPTOS_LOWDELAY SO_RCVBUF=1073741824 SO_SNDBUF=1073741824
   min protocol = SMB2
[homes]
   comment = Home Directories
   browseable = no
   read only = no
   create mask = 0770
   directory mask = 0770
   force group = nas
   valid users = %S
   writeable = yes
   path = /mnt/4000/home/%u
   veto files = /._*/.DS_Store/
   delete veto files = yes
[Shared]
   comment = Shared Family Folder
   browsable = yes
   writable = yes
   guest ok = yes
   read only = no
   directory mask = 0770
   create mask = 0770
   valid users = ldrrp krystal
   force group = nas
   path = /mnt/4000/home/Shared
   veto files = /._*/.DS_Store/
   delete veto files = yes
[security]
   comment = Security Cameras
   browseable = yes
   read only = no
   create mask = 0755
   directory mask = 0755
   valid users = @security
   writeable = yes
   path = /var/cache/zoneminder
   veto files = /._*/.DS_Store/
   delete veto files = yes
```

Then create the groups "nas" and "family" in linux and add your users to it. There is a [guide][9] on that.

Finally reboot or restart samba `systemctl restart samba`

**EDIT:**

The first time you add a linux user (`adduser`) you need to add them to smbpasswd as well.

`sudo smbpasswd -a <user>`

## Speedtest alerts

This one was very simple, I made a simple script in php to send me alerts when my house was not getting the speeds I pay for. You will need to have php-cli,  speedtest-cli installed and smtp credentials for email. You can [download the scripts][10]

After downloading them, add a line to your schedule tasks in linux using crontab -e

```
0 8 * * * /usr/bin/php /mnt/4000/home/ldrrp/Scripts/speedtest.php
```

 [1]: https://www.zoneminder.com/
 [2]: https://blog.silocitylabs.com/post/2018/12/13/axis-2130-ptz-and-zoneminder/
 [3]: https://blog.silocitylabs.com/post/2019/02/07/axis-2130-ptz-and-zoneminder/
 [4]: https://blog.silocitylabs.com/post/2018/12/27/nightowl-bnc-and-zoneminder/
 [5]: https://blog.silocitylabs.com/post/2019/01/10/genius-widecam-320-and-zoneminder/
 [6]: https://blog.silocitylabs.com/post/2018/11/28/bv-tech-4mp-2k-h-265-ip67-and-zoneminder/
 [7]: https://blog.silocitylabs.com/post/2018/11/23/secure-backups-without-the-cloud/
 [8]: http://a.tra.li/TPzz
 [9]: https://www.howtogeek.com/50787/add-a-user-to-a-group-or-second-group-on-linux/
 [10]: https://downloads.techreanimate.com/kzhota
