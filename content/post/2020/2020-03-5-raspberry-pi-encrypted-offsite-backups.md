---
title: Turn an Old Raspberry pi into an offsite backup
author: Luis Rodriguez
type: post
date: 2020-02-07T00:57:15+00:00
url: /post/2020/02/07/raspberry-pi-encrypted-offsite-backups
categories:
  - Projects
tags:
  - nas
  - all-in-one
  - iot
  - home-server
  - server
  - cloud
  - backups
  - offsite
  - encryption
  - raspberry
  - pi
  - upnp
  - ssh

---

### Why?

Lets start off with a few notes about how and why I would do this. I already have a fireproof/waterproof drive in my house. Sounds like im all set right? Not really.

Although im set on a disaster level there are a few things that can be an issue with onsite backups. Especially with my drive being plugged in all the time.

 - Theft, Although it is "theft resistant" the metal they want me to bolt down isnt very strong at all. Anyone with half a brain can break it in about 1 minute with all the tools I have down there.
 - Hackers, In recent times most of these hacks have targeted user data by encrypting the entire set of data and holding it hostage for bitcoin.
 - Accidents, They happen all the time to people. Commands like `rm -rf ./*` or even formatting the wrong partition. Lets just not set ourselves up for failure.

### How?

Who in the world is gonna host my raspberry pi that I can trust? Are they gonna try and charge me for electric? Are they going to want an internet fee? Will I slow down the wifi?

#### Trust?

We dont actually need to trust them all that much, They dont need any passwords and we dont even risk our data becouse we plan on using encryption.

#### Electric?

The raspberry pi is device that runs on arm, luckly smartphones also run on arm which makes them ideal for super low power usage. By stripping down a raspberry pi we can get the power usage so low that they wont even care about the half dollar its going to cost them montly. Toss them a $20 and call yourselves even for like 3-4 years.

The Math:

Currently my setup consists of idle usage of about 800mah with two USB drives attached and a micro sd card. You could probably go lower by eliminating the second usb but i dont really trust the micro sd card I attached to long live, also its 2gb haha.

800mah * 5v = 4w
4w * 24hrs * 30 days = 2880wh/month
2880wh * $0.15/kw = $0.43/Month

#### Internet?

As long as your not daily downloading extremely large amounts of data daily then you should be fine. I dont backup my media library, I can get those back easy. What i do backup is my nextcloud setup with all my personal family photos/videos and documents on my server.

To avoid hogging the network I set the backup to run daily at a time that I know they wont be on, like 4AM.

#### Wifi?

You should avoid wifi at all costs, You will increase your power usage and theres more benefit to using ethernet for reliability.

### Prerequisites

Before you get going you will need to install Rasbian on a micro sd card. For my setup I [installed it on a USB drive](/post/2019/06/01/raspberry-pi-booting-from-usb-on-older-models/) as well which is entirely optional for reliability. You may want to [minimize your setup](/post/2019/06/13/ultra-minimal-raspbian-image-for-pi-zero-and-zero-w/), just skip the part about wifi.

#### Hardware:

Raspberry Pi Model B with Raspian on SD card
16GB Sandisk USB with Raspbian
4TB Western Digital My Passport Ultra USB3 Hard Drive
HooToo Powered USB3 Hub (optional)

USB hard drives are great however becouse they dont use external power bricks you may see some issues if your pi cant provide enough power. You can try this hub to power both the Pi and the HDD.

### Setup drive

#### Install cryptsetup to encrypt drives.

```
pi@raspberrypi ~ $ sudo apt-get install cryptsetup
```

#### Find your disk

Find the /dev identifier of the HDD partition you want to encrypt, on my system, it's /dev/sdb1. Note that this output is from running the command after I've configured the encrypted partition.

```
pi@raspberrypi ~ $ lsblk
NAME                    MAJ:MIN RM   SIZE RO TYPE  MOUNTPOINT
sdb                       8:0    0   3.6T  0 disk  
```

#### Create partition

Start by creating a partition with fdisk.

```
sudo fdisk /dev/sdb
````

 1. Press `n` option to create a new partition
 2. Press `p` for a primary partition.
 3. Press `1` to  edit the first partition
 4. Press `ENTER` twice to use the defaults for first and last sectors
 5. Press `w` to save changes to disk
 6. Press `q` to quit

It will look similar to this:

```
Command (m for help): n
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-3906916350, default 2048): 
Last sector, +sectors or +size{K,M,G,T,P} (2048-3906916350, default 3906916350): 

Created a new partition 1 of type 'Linux' and of size 4.3 TiB.

Command (m for help): w
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.
```

#### Encrypt the partition

```
sudo cryptsetup -v -y -c aes-xts-plain64 -s 512 -h sha512 -i 5000 --use-random luksFormat /dev/sdb1
```

The options of the command are as follows:

 - -v = verbose
 - -y = verify passphrase
 - -c = cipher used
 - -s = key size used
 - -h = hash used
 - -i = number of ms to spend passphrase processing
 - --use-random = which random number generator to use
 - luksFormat = initialize the partition and set a passphrase
 - /dev/sdb1 = partition to encrypt

Optionally you can [backup the luks header](https://gitlab.com/cryptsetup/cryptsetup/-/wikis/FrequentlyAskedQuestions#6-backup-and-data-recovery).

#### Add secondary passphrase for stdin use

To use luks remotely we will need to add a secondary key to the luksheader.

```
sudo cryptsetup luksAddKey /dev/sdb1
```

#### Format the volume

Mount the luks device and format the volume

```
sudo cryptsetup luksOpen /dev/sdb1 volume01
sudo mkfs.ext4 /dev/mapper/volume01
```

#### Mounting the volume

You can mount it anywhere but for this setup we will mount it to `/mnt/raid`

```
sudo mkdir -p /mnt/raid
sudo mount /dev/mapper/volume01 /mnt/raid
```

You should now be able to see it under `df -h`

#### Unmount and close

```
sudo umount /mnt/drive01
sudo cryptsetup luksClose /dev/mapper/volume01
```

### DDNS

Create a script file with the following but change `your.domain.name` and the ddns update url with your own.

`nano ddns.sh`
```
#!/bin/bash

IP_NS=$(dig +short your.domain.name)
IP_WAN=$(curl -ks http://checkip.amazonaws.com/)

if [[ "$IP_NS" == "$IP_WAN" ]]; then
    echo "No IP update required. ($IP_WAN)"
else
    echo "IP has changed"
    curl -ks https://dnsprovider.com?key=sometokentochangeip
fi
```

Now add the script to your crontab after making it executable

```
chmod +x ddns.sh
crontab -e
```

```
*/10 * * * * /path/to/ddns.sh
```

### upnp port forwarding

Becouse some routers dont allow ports below 1024 to be mounted we will just be opening up ssh to another port. View a [more detailed writeup](https://pavelfatin.com/access-your-raspberry-pi-from-anywhere/) on how to run this on ethernet connect.

```
sudo apt install miniupnpc
```

You can change port 1080 to anything you desire, it is recommended to keep it above 1024 for some routers that dont allow numbers below that.

`nano upnp.sh`
```
#!/bin/bash
ip=$($upnpc -l | grep "Local LAN ip address" | cut -d: -f2)

upnpc -e "Backups" -a $ip 22 1080 TCP
```

I will just be adding this on bootup with `crontab -e`

```
@reboot upnp.sh
```

### Script mounting

Now on your remote client create a script to start the backups

`nano remote-backup.sh`
```
#!/bin/bash

#Mount the drive
sshpass -p "password" ssh pi@yourdomainname.com << EOF
 CRYPTPASS='passphrase to descrypt drive'
 echo $CRYPTPASS | cryptsetup luksOpen /dev/sdb1 volume01 -d=-
 mount /dev/mapper/volume01 /mnt/raid
EOF

#run rsync
sshpass -p "password" rsync -a -P -e ssh /mnt/raid pi@yourdomainhere.com:/mnt/raid

#unmount the drive
sshpass -p "password" ssh pi@yourdomainname.com << EOF
 umount /mnt/raid
 cryptsetup luksClose /dev/mapper/volume01
EOF

```

Now lets make this run once a week on sunday at midnight with `crontab -e`

```
0 0 * * 0 remote-backup.sh
```
