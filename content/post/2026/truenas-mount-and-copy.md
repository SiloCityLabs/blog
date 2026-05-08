---
title: "TrueNAS copy from drive"
author: maave
type: post
date: 2026-05-01T00:00:00+00:00
url: /post/2026/05/01/truenas-mount-and-copy/
draft: true
categories:
  - Android
tags:
  - TrueNAS
  - Linux
---


{{< image src="/uploads/2026/truenas-shell.webp" alt="TrueNAS Linux shell command prompt">}}

Woohoo, I finally got the NAS set up! It's TrueNAS plus a JBOD enclosure for drives. Time to back up my data. However I want to copy over many gigabytes of data from an old drive and the TrueNAS web interface doesn't allow me to mount drives directly. TrueNAS wants a ZFS array, I want to mount NTFS without wiping the drive. Let's use the Linux command line to copy my data instead.

<!--more-->

## Hardware overview

Let's quickly look at the hardware setup. The server is some old desktop mobo and 2 old SSDs shoved into a 2U rack-mount case. The hard drives for storage go into this JDBO enclosure. "JBOD", that's "just a bunch of disk", which exposes all the drives to the system as-if they were plugged in directly. This JBOD enclosure has a QSFP port on the back so the server needs a host-bus adapter card 

{{< image src="/uploads/2026/hba-card.webp" alt="oops, wrong backplate on the HBA card">}}

## Mounting and copying

Since TrueNAS SCALE is based on Debian I'm quite comfortable in the command line. The truenas_admin user has sudo access so you I run any command. I'd recommend enabling SSH under System > Services, and add "truenas_admin" to "Password Login Groups". I should've done that to avoid any web UI timeout, but I already had the web UI timeout set to 5hr so I just use System > Shell.

Use `lsblk` to find the drive device name. If you can't tell which drive is yours just unplug it, check lsblk, replug it, check lsblk again. My drive is `sdf` and I can tell by size that my main partition is `sdf2`. The ~~4 TB~~ 3.6 TB drives are my storage drives that I've configured in RAID using the TrueNAS UI.

```shell
truenas_admin@truenas[~]$ lsblk
NAME   MAJ:MIN RM   SIZE RO TYPE MOUNTPOINTS
sda      8:0    0   3.6T  0 disk 
└─sda1   8:1    0   3.6T  0 part 
sdb      8:16   0   3.6T  0 disk 
└─sdb1   8:17   0   3.6T  0 part 
sdc      8:32   0 238.5G  0 disk 
├─sdc1   8:33   0     1M  0 part 
├─sdc2   8:34   0   512M  0 part 
└─sdc3   8:35   0   238G  0 part 
sdd      8:48   0 223.6G  0 disk 
├─sdd1   8:49   0     1M  0 part 
├─sdd2   8:50   0   512M  0 part 
└─sdd3   8:51   0 223.1G  0 part 
sde      8:64   0   3.6T  0 disk 
└─sde1   8:65   0   3.6T  0 part 
sdf      8:80   0 931.5G  0 disk 
├─sdf1   8:81   0   100M  0 part 
├─sdf2   8:82   0 930.8G  0 part 
├─sdf3   8:83   0   558M  0 part 
└─sdf4   8:84   0   100M  0 part 
```

Make a new directory in /mnt/ and mount the drive. I mount it read-only to avoid deleting data. I make mistake #1 here - I create the directory with truenas_admin instead of root. The rest of the dataset is root owned and then doled out with TrueNAS's ACL.

```shell
truenas_admin@truenas[~]$ ls /mnt     
storage1
truenas_admin@truenas[~]$ ls /mnt/storage1 
movies  truenas-storage1
truenas_admin@truenas[~]$ sudo mkdir /mnt/direct-mount
[sudo] password for truenas_admin: 
truenas_admin@truenas[~]$ sudo mount -o ro /dev/sdf2 /mnt/direct-mount
mount: /mnt/direct-mount: unknown filesystem type 'ntfs'.
       dmesg(1) may have more information after failed mount system call.
```

Dang. I truly hoped that NTFS would mount without issue in 2026. I still need to specify the file system type:

```shell
truenas_admin@truenas[~]$ sudo mount -t ntfs3 -o ro /dev/sdf2 /mnt/direct-mount
truenas_admin@truenas[~]$ ls /mnt/direct-mount
'$GetCurrent'                      AMD                       DumpStack.log       Logs            'Program Files (x86)'   Reflect_Install.log          Users              hiberfil.sys   temp
'$Recycle.Bin'                     AMFTrace.log              DumpStack.log.tmp   MSOCache         ProgramData            Sandbox                      Windows            inetpub        xampp
'$WINRE_BACKUP_PARTITION.MARKER'   BOOTNXT                   END                 Modding          Python26               SteamLibrary                 Windows10Upgrade   log
'$WinREAgent'                      Config.Msi                ESD                 PerfLogs         Python32              'System Volume Information'   bootmgr            pagefile.sys
 AI_RecycleBin                    'Documents and Settings'   HaxLogs.txt        'Program Files'   Recovery               SystemBoardInfoResult.txt    eclipse            swapfile.sys
```

Excellent, all my files show up. I actually have a few junk files like pagefile and recycle bin but I'm in read-only mode for safety, I'm _not deleting those_ right now. I'll do it after verifying the backup.

Time to copy. I'm going to use `rsync` for reliability and resumability. Option `-a` preserves timestamps (mistake #2?) and option `-P` shows a progress meter. This `-a` option will bite me later as it preserves read/write permissions. I should add here `--no-perms --no-owner --no-group`

```shell
truenas_admin@truenas[/mnt/direct-mount]$ mkdir  /mnt/storage1/truenas-storage1/hdd-backup-1/
truenas_admin@truenas[/mnt/direct-mount]$ rsync -aP /mnt/direct-mount/ /mnt/storage1/truenas-storage1/hdd-backup-1/
sending incremental file list
./
$WINRE_BACKUP_PARTITION.MARKER
              0 100%    0.00kB/s    0:00:00 (xfr#1, ir-chk=1068/1070)
AMFTrace.log
              2 100%    0.00kB/s    0:00:00 (xfr#2, ir-chk=1067/1070)
BOOTNXT
              1 100%    0.98kB/s    0:00:00 (xfr#3, ir-chk=1066/1070)
Documents and Settings -> ./Users
DumpStack.log
          8,192 100%    7.81MB/s    0:00:00 (xfr#4, ir-chk=1064/1070)
DumpStack.log.tmp
          8,192 100%    7.81MB/s    0:00:00 (xfr#5, ir-chk=1063/1070)
END
              2 100%    1.95kB/s    0:00:00 (xfr#6, ir-chk=1062/1070)
HaxLogs.txt
             83 100%   81.05kB/s    0:00:00 (xfr#7, ir-chk=1061/1070)
Reflect_Install.log
        467,360 100%  111.43MB/s    0:00:00 (xfr#8, ir-chk=1060/1070)
SystemBoardInfoResult.txt
             31 100%    7.57kB/s    0:00:00 (xfr#9, ir-chk=1059/1070)
bootmgr
        407,924 100%   64.84MB/s    0:00:00 (xfr#10, ir-chk=1058/1070)
hiberfil.sys
 25,057,157,120 100%  502.72MB/s    0:00:47 (xfr#11, ir-chk=1057/1070)
pagefile.sys
 34,359,738,368 100%  462.47MB/s    0:01:10 (xfr#12, ir-chk=1056/1070)
swapfile.sys
     16,777,216 100%   39.80MB/s    0:00:00 (xfr#13, ir-chk=1055/1070)
```

This is much faster than network transfer, it should save me a few hours. I'm glad I used `rsync` rather than `cp` - I accidentally killed the shell partway through the copy, and I simply ran the same rsync command again to resume where I left off.

At the very end I see a warning:

```
rsync error: some files/attrs were not transferred (see previous errors) (code 23) at main.c(1338) [sender=3.2.7]
```

well the errors are a million lines up in the terminal and I can't scroll back that far. I'll just run rsync again, adding the flag `-v` for verbose and `--log-file=` to write to a file.

```shell
truenas_admin@truenas[/mnt/direct-mount]$ rsync -avP /mnt/direct-mount/ /mnt/storage1/truenas-storage1/hdd-backup-1/ --log-file=/mnt/storage1/truenas-storage1/hdd-backup-1/rsync.log
```

```
2026/05/01 15:13:32 [320013] rsync: [sender] read errors mapping "/mnt/direct-mount/Logs/Application.evtx": Operation not supported (95)
"/mnt/direct-mount/Windows/WinSxS/amd64_microsoft-windows-d..oryservices-ntdsapi_31bf3856ad364e35_10.0.19041.2486_none_438f7c8da6760eb7/ntdsapi.dll": Operation not supported (95)
2026/05/01 15:15:21 [320013] rsync: [sender] read errors mapping "/mnt/direct-mount/Windows/servicing/Packages/Microsoft-Windows-AppManagement-AppV-Package~31bf3856ad364e35~amd64~~10.0.19041.3208.cat": Operation not supported (95)
```

```
read errors mapping
Operation not supported (95)
```

This appears to be "NTFS reparse points". That's metadata in NTFS that can't be copied to ZFS. In my case it's basically all symlinks (or similar) for Windows OS files and some OneDrive files which get marked as from-the-cloud. I scoured the log file in Notepad++, carving it up with the Bookmark feature to mark and remove lines until I checked every error. I'm not going to copy these, I don't need them. The most basic of symlink still works:

```shell
lrwxrwxrwx  1 truenas_admin root        7 Jul 10  2015 'Documents and Settings' -> ./Users
```

## Permissions and ACL

I double check my files using SMB and the "filebrowser" app. I can't edit or delete files in the subdirectory. Oops. Let's check permissions. I have another directory "Downloads" which I created in SMB and apps can access it.

```shell
truenas_admin@truenas[~]$ cd /mnt/storage1/truenas-storage1
root@truenas[/mnt/storage1/truenas-storage1]# ls -l
total 12
drwxrwxrwx  2 apps  root  3 May  1 18:24 filebrowser-test
drwxrwxrwx 30 truenas_admin root 38 May  1 19:54 hdd-backup-1
drwxrwxrwx  3 maave root  3 Apr 30 20:15 Downloads
truenas_admin@truenas[/mnt/storage1/truenas-storage1]$ cd hdd-backup-1
truenas_admin@truenas[...torage1/truenas-storage1/hdd-backup-1]$ cd SteamLibrary
truenas_admin@truenas[...as-storage1/hdd-backup-1/SteamLibrary]$ ll
total 386
drwxr-xr-x  3 truenas_admin      5 Jul 25  2023 ./
drwxrwxrwx 30 truenas_admin     40 May  1 18:07 ../
-rwxr-xr-x  1 truenas_admin    105 Feb 13  2022 libraryfolder.vdf*
-rwxr-xr-x  1 truenas_admin 498088 Jan 16  2022 steam.dll*
drwxr-xr-x  6 truenas_admin     14 Jul 25  2023 steamapps/
```

```shell
truenas_admin@truenas[...NAL FANTASY XIV - A Realm Reborn/game]$ ls -l
total 34612
-rwxr-xr-x 1 truenas_admin root    93592 Apr  3  2022 XInputXIV3.dll
-rwxr-xr-x 1 truenas_admin root    93592 Apr  3  2022 XInputXIV9.dll
-rwxr-xr-x 1 truenas_admin root   347216 Apr  7  2019 bink2w32.dll
-rwxr-xr-x 1 truenas_admin root   436816 Apr  7  2019 bink2w64.dll
-rwxr-xr-x 1 truenas_admin root 29084056 Apr  3  2022 ffxiv.exe
-rwxr-xr-x 1 truenas_admin root 40178584 Apr  3  2022 ffxiv_dx11.exe
-rwxr-xr-x 1 truenas_admin root       20 Apr  3  2022 ffxivgame.bck
-rwxr-xr-x 1 truenas_admin root       20 Apr  3  2022 ffxivgame.ver
-rwxr-xr-x 1 truenas_admin root     1600 Apr  3  2022 fileinfo.fiin
drwxr-xr-x 6 truenas_admin root        6 May  1 20:10 movie
drwxr-xr-x 6 truenas_admin root        6 May  1 20:10 sqpack
```

foolish! I've created the directory as truenas_admin _and_ copied the read-only permissions.

Let's change the owner and group with `chown` and set permissions with `chmod`. `u+rwX,g+rwX,o-rwx` is full perms for the user and group, and no permissions for others, equivalent to `770`.

```shell
root@truenas[...torage1/truenas-storage1/hdd-backup-1]# chown -R maave:root /mnt/storage1/truenas-storage1/hdd-backup-1
root@truenas[...torage1/truenas-storage1/hdd-backup-1]# chmod -R u+rwX,g+rwX,o-rwx /mnt/storage1/truenas-storage1/hdd-backup-1
```

```shell
root@truenas[/mnt/storage1/truenas-storage1]# ls -l
total 12
drwxrwxrwx  2 apps  root  3 May  1 18:24 filebrowser-test
drwxrwx--- 26 maave root 38 May  1 19:54 hdd-backup-1
drwxrwxrwx  3 maave root  3 Apr 30 20:15 Downloads
```

Then set the ACL. This will allow apps to access the directory despite the group being "root". I do this in the TrueNAS web UI, Datasets > truenas-storage1 > Permissions.

### manual mode

There's an option in the TrueNAS web UI to reset the ACL for all files but I want to try targeting just this dir. I check the existing ACL with `nfs4xdr_getfacl`

```shell
root@truenas[/mnt/storage1]# nfs4xdr_getfacl /mnt/storage1/truenas-storage1/
# File: /mnt/storage1/truenas-storage1/
# owner: 0
# group: 0
# mode: 0o40777
# trivial_acl: false
# ACL flags: none
            owner@:rwxpDdaARWcCos:fd-----:allow
            group@:rwxpDdaARWcCos:fd-----:allow
         everyone@:rwxpDdaARWc--s:fd-----:allow
root@truenas[/mnt/storage1]# nfs4xdr_getfacl /mnt/storage1/truenas-storage1/Downloads
# File: /mnt/storage1/truenas-storage1/Downloads
# owner: 3000
# group: 0
# mode: 0o40777
# trivial_acl: false
# ACL flags: none
            owner@:rwxpDdaARWcCos:fd----I:allow
            group@:rwxpDdaARWcCos:fd----I:allow
         everyone@:rwxpDdaARWc--s:fd----I:allow
root@truenas[/mnt/storage1]#
```

The ACL is actually the same for each, only the owner is different. I set these permissions with `nfs4xdr_setfacl`

```shell
root@truenas[/mnt/storage1/truenas-storage1]# nfs4xdr_setfacl -a "owner@:rwxpDdaARWcCos:fd:allow" /mnt/storage1/truenas-storage1/hdd-backup-1
root@truenas[/mnt/storage1/truenas-storage1]# nfs4xdr_setfacl -a "group@:rwxpDdaARWcCos:fd:allow" /mnt/storage1/truenas-storage1/hdd-backup-1
root@truenas[/mnt/storage1/truenas-storage1]# nfs4xdr_setfacl -a "everyone@:rwxpDdaARWc--s:fd:allow" /mnt/storage1/truenas-storage1/hdd-backup-1
root@truenas[/mnt/storage1/truenas-storage1]# ls -l
total 12
drwxrwxrwx  2 apps  root  3 May  1 18:24 filebrowser-test
drwxrwxrwx 26 maave root 34 May  1 20:55 hdd-backup-1
drwxrwxrwx  3 maave root  3 Apr 30 20:15 Downloads
```

