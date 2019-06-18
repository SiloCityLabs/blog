---
title: Ultra minimal Raspbian Image for Pi Zero and Zero W
date: 2019-06-13
tags: ["linux", "minimal pi", "pi", "raspberry"]
---

The following are two code samples using syntax highlighting.

<!--more-->

I wanted to make a ultra lite image that I can quickly drop Go programs onto. This image had a few requirements, be under 300MB, have less than 20MB ram usage idle, startup in under 5 seconds. I will post below everything I did to strip this down so that it can be replicated in the future. Currently with the steps below I have achieved 427MB/16MB, close to my goal for disk and memory usage is below the goal. As for bootup time, I am limited by network connection time 5+ seconds.

Install a fresh Raspbian Stretch Lite image into the SD card (source).


Configure headless SSH and Wi-Fi (if necessary) before starting the SD card (source).

```
$ mount /dev/sdX /media $ :> /media/ssh $ cat > /media/wpa_supplicant.conf <<"EOF" country=GB ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev update_config=1 network={ ssid="YOUR_SSID" psk="YOUR_PSK" } EOF $ umount /media $ sync
```

## Replacing Openssh with dropbear in raspbian
### Installation
To install on Raspbian, simply run the following two commands on your Pi:

```
apt update && apt install dropbear
```

### Configuration
To enable dropbear, edit the config file with:

```
nano /etc/default/dropbear
```
and make sure NO_START matches below:
```
NO_START=0
```
and make sure the following is set:
```
DROPBEAR_PORT=22
```
22 is the default for SSH, but you can change it if you use another port for SSH(such as 443 or 80 to get around firewalls).

### Disable SSH
You could just delete the OpenSSH server with

```
apt-get purge --yes --auto-remove openssh-server
```
## Cleaning Disk Space
Get rid of some programs, I figured if I need any I will manually add them but for now kill everything for base image.

```
apt-get purge --yes --auto-remove aptitude aptitude-common apt-listchanges apt-utils blends-tasks xserver-common x11-xfs-utils x11-xserver-utils xinit libsmbclient blt gvfs gvfs-backends gvfs-daemons gvfs-fuse idle idle-python2.7 idle3 libaudio2 libice6 liblightdm-gobject-1-0 libpulse0 libqt4-svg libqtgui4 libsdl-image1.2 libsdl-mixer1.2 libsdl-ttf2.0-0 libsdl1.2debian libsm6 libsmpeg0 libwebkitgtk-1.0-0 libwebkitgtk-3.0-0 libxaw7 libxklavier16 libxmu6 libxss1 libxt6 libxtst6 lightdm lightdm-gtk-greeter lxde lxde-core midori obconf openbox python-pygame python-tk python3-tk scratch tk8.5 wpagui x11-common x11-utils x11-xkb-utils xinit xserver-common xserver-xorg xserver-xorg-core xserver-xorg-input-all xserver-xorg-input-evdev xserver-xorg-input-synaptics xserver-xorg-video-fbdev zenity xserver* x11-common x11-utils x11-xkb-utils x11-xserver-utils xarchiver xauth xkb-data console-setup xinit lightdm libx{composite,cb,cursor,damage,dmcp,ext,font,ft,i,inerama,kbfile,klavier,mu,pm,randr,render,res,t,xf86}* lxde* lx{input,menu-data,panel,polkit,randr,session,session-edit,shortcut,task,terminal} obconf openbox gtk* libgtk* alsa* nano python-pygame python-tk python3-tk scratch tsconf desktop-file-utils cifs-utils samba-common smbclient scratch debian-reference-en dillo idle3 python3-tk idle python-pygame python-tk lightdm gnome-themes-standard gnome-icon-theme raspberrypi-artwork gvfs-backends gvfs-fuse desktop-base lxpolkit netsurf-gtk zenity xdg-utils mupdf gtk2-engines alsa-utils lxde lxtask menu-xdg gksu midori xserver-xorg xinit xserver-xorg-video-fbdev libraspberrypi-dev libraspberrypi-doc dbus-x11 libx11-6 libx11-data libx11-xcb1 x11-common x11-utils lxde-icon-theme gconf-service gconf2-common udhcpd avahi-daemon gcc-arm-linux-gnueabihf sudo geoip-database libgdbm3 libperl5.24 make patch perl perl-modules-5.24 cups* rsync bluez bluez-firmware pi-bluetooth scratch debian-reference-en dillo idle3 python3-tk idle python-pygame python-tk lightdm gnome-themes-standard gnome-icon-theme raspberrypi-artwork gvfs-backends gvfs-fuse desktop-base lxpolkit netsurf-gtk zenity xdg-utils mupdf gtk2-engines alsa-utils lxde lxtask menu-xdg gksu midori xserver-xorg xinit xserver-xorg-video-fbdev libraspberrypi-dev libraspberrypi-doc dbus-x11 libx11-6 libx11-data libx11-xcb1 x11-common x11-utils lxde-icon-theme gconf-service gconf2-common firmware-atheros samba-common plymouth
```

Cleanup any leftover crud

```
apt-get autoremove --purge
```

Delete the default pi user

```
deluser pi && rm -rf /home/pi
```

kill config comments
```
find . -type f -name "*.conf" |xargs sed -i '\#^//#d'
```

Disable dpkg docs and locale files from installing via apt
```
cat > /etc/apt/apt.conf.d/80noadditional <<"EOF"
APT::Install-Recommends "0";
APT::Install-Suggests "0";
EOF

cat > /etc/dpkg/dpkg.cfg.d/nodoc <<"EOF" path-exclude /usr/share/doc/* path-include /usr/share/doc/*/copyright path-exclude /usr/share/man/* path-exclude /usr/share/groff/* path-exclude /usr/share/info/* path-exclude /usr/share/lintian/* path-exclude /usr/share/linda/* EOF cat > /etc/dpkg/dpkg.cfg.d/nolocale <<"EOF"
path-exclude /usr/share/locale/*
EOF
```

Cleanup Logs, docs, useless files

```
find /usr/share/doc -depth -type f ! -name copyright | xargs rm
find /usr/share/doc -depth -type f ! -name copyright | xargs rm -f
find /usr/share/doc -empty | xargs rmdir
rm -rf /usr/share/man/* /usr/share/groff/* /usr/share/info/*
rm -rf /usr/share/lintian/* /usr/share/linda/* /var/cache/man/*
rm -rf /usr/share/locale/*
rm -f /var/log/{auth,boot,bootstrap,daemon,kern}.log
rm -f /var/log/{debug,dmesg,messages,syslog}
```

Disable MOTD for faster ssh connect
```
cat /dev/null > /etc/motd
```

Disable UART since we wont need it.

```
systemctl disable hciuart
```

More crud cleaning

```
apt-get clean
```

## Other cleanup and optimizations
Add the following lines to your Pi’s ``/boot/config.txt`` file and reboot:

```
# http://rpf.io/configtxt
#https://www.raspberrypi.org/documentation/configuration/config-txt/overclocking.md
arm_freq=1100

#initial_turbo=30
#dtparam=i2c_arm=on
#dtparam=i2s=on
#dtparam=spi=on

# Disable stuff
dtparam=audio=off
dtoverlay=pi3-disable-bt

# Disable the ACT LED on the Pi Zero.
dtparam=act_led_trigger=none
dtparam=act_led_activelow=on

# Disable the splash screen
disable_splash=1
 
# Overclock the SD Card from 50 to 100MHz
# This can only be done with at least a UHS Class 1 card
dtoverlay=sdtweak,overclock_50=100
 
# Set the bootloader delay to 0 seconds. The default is 1s if not specified. 
boot_delay=0
 
# Overclock the raspberry pi. This voids its warranty. Make sure you have a good power supply.
force_turbo=1
```

### Quiet boot
Make the kernel output less verbose by adding the “quiet” flag to the kernel command line in file /boot/cmdline.txt. Make sure to change the PARTUUID if you copy and paste this line directly.

```
dwc_otg.lpm_enable=0 console=tty1 root=PARTUUID=8b037108-02 rootfstype=ext4 elevator=deadline fsck.repair=no quiet rootwait
```

### Disabling some unneeded services
```
systemctl disable dphys-swapfile.service
systemctl disable apt-daily.service
systemctl disable apt-daily.timer
systemctl disable nfs-client.target
systemctl disable remote-fs.target
systemctl disable apt-daily-upgrade.timer
systemctl disable nfs-config.service
```

All I am left with now for memory usage. If anyone has suggestions on how to lower any let me know.

 
```
  PID    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND         
    1    9520   5876   4788 S  0.0  1.3   0:03.50 systemd
   63   20500   4176   3792 S  0.0  0.9   0:00.83 systemd-journal
  271   10064   4004   3492 S  0.0  0.9   0:00.02 wpa_supplicant
  167   17100   3956   3524 S  0.0  0.9   0:00.22 systemd-timesyn
  219    6816   3796   3440 S  0.0  0.9   0:00.14 systemd-logind
  207    6396   3368   3000 S  0.0  0.8   0:00.28 dbus-daemon
   89   14152   3200   2604 S  0.0  0.7   0:00.71 systemd-udevd
  352    3476   2728   2316 S  0.0  0.6   0:00.13 bash
  387   22628   2708   2284 S  0.0  0.6   0:00.04 rsyslogd
  217    5200   2396   2196 S  0.0  0.5   0:00.03 cron
  345    3868   1900   1780 S  0.0  0.4   0:00.02 agetty
  351    2828   1816   1648 S  0.0  0.4   0:00.98 dropbear   
  346    4092   1804   1688 S  0.0  0.4   0:00.03 agetty
  242    2360   1524   1420 S  0.0  0.3   0:00.00 dropbear
  332    2780   1508   1192 S  0.0  0.3   0:00.00 dhcpcd
```
My disk usage unfortunately did not get as low as I wanted it.
```
Filesystem      Size  Used Avail Use% Mounted on
/dev/root       7.3G  428M  6.6G   7% /
devtmpfs        213M     0  213M   0% /dev
tmpfs           217M     0  217M   0% /dev/shm
tmpfs           217M  5.7M  212M   3% /run
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           217M     0  217M   0% /sys/fs/cgroup
/dev/mmcblk0p1   44M   22M   22M  51% /boot
```
Startup time

```
root@pi-b82xxxxxxx:~# systemd-analyze
Startup finished in 1.334s (kernel) + 16.235s (userspace) = 17.569s

root@pi-b827eb80a9a0:~# systemd-analyze blame
          9.316s dhcpcd.service
          4.058s dev-mmcblk0p2.device
          3.033s networking.service
          2.242s dropbear.service
          1.795s systemd-logind.service
           841ms systemd-journald.service
           756ms systemd-udev-trigger.service
           674ms systemd-remount-fs.service
           601ms rsyslog.service
           550ms systemd-fsck-root.service
           537ms systemd-timesyncd.service
           440ms systemd-fsck@dev-disk-by\x2dpartuuid-8b037108\x2d01.service
           409ms systemd-udevd.service
           391ms systemd-rfkill.service
           313ms systemd-tmpfiles-setup-dev.service
           309ms sys-kernel-debug.mount
           306ms systemd-update-utmp.service
           291ms dev-mqueue.mount
           286ms systemd-tmpfiles-setup.service
           285ms fake-hwclock.service
           270ms kmod-static-nodes.service
           262ms systemd-sysctl.service
           249ms systemd-modules-load.service
           244ms boot.mount
           237ms systemd-journal-flush.service
           158ms systemd-random-seed.service
           126ms sys-kernel-config.mount
           108ms systemd-user-sessions.service
            93ms rc-local.service
            90ms systemd-update-utmp-runlevel.service
            78ms wifi-country.service
```

## Download
I have uploaded the image for download. Up to date as of Jun 13th 2019
Image: ultra-minimal-pi.img.7z (113MB)
Username: root
Password: password

## Additional imaging steps
So I don’t forget how to do it later for updates

On the pi run

```
cat /dev/zero > /tmp/zeros; sync; rm /tmp/zeros
```

For the rest I took the sd card and mounted it and imaged it using ubuntu disks utility

Use the following script from the download link to shrink the .img file

The finally compress it into a zip file

Sources:
http://himeshp.blogspot.com/2018/08/fast-boot-with-raspberry-pi.html