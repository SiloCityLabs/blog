---
title: "Fixing WirePlumber High CPU Usage on Ubuntu 24.04"
author: Luis Rodriguez
type: post
date: 2026-07-08T00:00:00+00:00
draft: false
categories:
  - ubuntu
tags:
  - ubuntu
---


{{< image src="/uploads/2026/wireplumber-high-cpu-usage-fix.png" alt="CPU Terminal with wireplumber using high cpu">}}

I was experiencing a frustrating issue on Ubuntu 24.04 related to Bluetooth profile switching.

**Update:** Read the whole post before copying commands. My first fix moved WirePlumber from Ubuntu's 0.4.17 package to 0.5.2 using the PipeWire Debian PPAs, and that helped for a while. The issue later returned. The current fix in this post is building WirePlumber 0.5.15 locally and holding those packages. If you stop at the PPA step, you may only end up on 0.5.2 and waste time if your Bluetooth issue is the same one I hit.

<!--more-->

Here’s how I could reliably reproduce it:

1. Play music through a Bluetooth headset using the A2DP profile.
2. Pause the music.
3. Join a conference call, which switches the headset to the HSP/HFP profile.
4. Leave the call and resume music, which switches the headset back to A2DP.

After doing this, the system fan would ramp up noticeably. Running `htop` showed:

* `wireplumber` using 100% of a single CPU core.

The problem did not only affect music playback. It also caused:

* YouTube video playback issues
* Audio glitches
* Lag in other applications using audio

The system clearly was not happy.

The Root Cause
--------------

Ubuntu 24.04 ships with **WirePlumber 0.4.17**.

Although Bluetooth fixes were already merged upstream, that does **not** mean the fixes are included in Ubuntu’s packaged version. Ubuntu freezes core components before release, and the fixes can land after the version Ubuntu shipped.

So even though the issue may be fixed upstream, the fix is not necessarily present in the Ubuntu package.

The First Fix: Use the PipeWire and WirePlumber PPAs
----------------------------------------------------

My first fix was to upgrade the PipeWire and WirePlumber stack using the PipeWire Debian PPAs.

This moved my system from Ubuntu’s default WirePlumber 0.4.17 to WirePlumber 0.5.2.

__Time needed: 2 minutes__

### 1. Add the upstream PPAs

```bash
sudo add-apt-repository ppa:pipewire-debian/pipewire-upstream
sudo add-apt-repository ppa:pipewire-debian/wireplumber-upstream
```

### 2. Update and upgrade

```bash
sudo apt update
sudo apt full-upgrade
```

### 3. Clean up unused packages

```bash
sudo apt autoremove
```

### 4. Reboot

After rebooting, verify the version:

```bash
wireplumber --version
```

In my case, the PPA installed WirePlumber 0.5.2:

```text
Compiled with libwireplumber 0.5.2
Linked with libwireplumber 0.5.2
```

Important note: the `wireplumber-upstream` PPA does **not** always mean you get the newest upstream WirePlumber release. It only means you get the newest package that PPA has published for your Ubuntu release.

For Ubuntu 24.04 Noble, the PPA currently only offered 0.5.2 for me:

```bash
apt policy wireplumber libwireplumber-0.5-0 pipewire
```

That showed:

```text
wireplumber:
  Installed: 0.5.2-8~ubuntu24.04
  Candidate: 0.5.2-8~ubuntu24.04
  Version table:
 *** 0.5.2-8~ubuntu24.04 500
        500 https://ppa.launchpadcontent.net/pipewire-debian/wireplumber-upstream/ubuntu noble/main amd64 Packages
     0.4.17-1ubuntu4.1 500
        500 http://us.archive.ubuntu.com/ubuntu noble-updates/main amd64 Packages
     0.4.17-1ubuntu4 500
        500 http://us.archive.ubuntu.com/ubuntu noble/main amd64 Packages
```

The Bluetooth CPU issue originally seemed resolved after moving to 0.5.2, but it eventually returned.

The Current Fix: Build WirePlumber 0.5.15 Locally
-------------------------------------------------

Since the issue came back on 0.5.2, I moved to WirePlumber 0.5.15.

WirePlumber 0.5.14 and 0.5.15 both include additional Bluetooth-related work. The 0.5.14 changelog mentions Bluetooth profile autoswitch robustness improvements, including better headset profile detection and race condition fixes. Since my issue appears tied to Bluetooth profile switching, this update may be relevant.

Ubuntu 24.04 does not currently provide WirePlumber 0.5.15 through the normal Noble repositories or the PPA I was using, so I built the Debian source package locally instead of adding Debian unstable as an apt source.

Do **not** add Debian unstable or sid directly to Ubuntu just to install WirePlumber. That can pull in unrelated system packages and create a mess. Building the source package locally is safer.

### 1. Install build tools and dependencies

```bash
sudo apt update
sudo apt install -y \
  build-essential devscripts debhelper-compat dh-sequence-gir \
  meson ninja-build pkgconf git wget curl \
  libglib2.0-dev libpipewire-0.3-dev \
  gobject-introspection libgirepository1.0-dev \
  liblua5.4-dev libsystemd-dev \
  python3-docutils
```

### 2. Download the Debian WirePlumber 0.5.15 source package

```bash
mkdir -p ~/build/wireplumber-0.5.15
cd ~/build/wireplumber-0.5.15

wget http://deb.debian.org/debian/pool/main/w/wireplumber/wireplumber_0.5.15-1.dsc
wget http://deb.debian.org/debian/pool/main/w/wireplumber/wireplumber_0.5.15.orig.tar.gz
wget http://deb.debian.org/debian/pool/main/w/wireplumber/wireplumber_0.5.15-1.debian.tar.xz

dpkg-source -x wireplumber_0.5.15-1.dsc
cd wireplumber-0.5.15
```

### 3. Check build dependencies

```bash
dpkg-checkbuilddeps
```

If it prints missing packages, install those packages and run the check again.

### 4. Patch Debian GNOME packaging glue if needed

On Ubuntu 24.04, the build initially failed for me during the clean step with:

```text
dh_gnome_clean
No such file or directory at /usr/bin/dh_gnome_clean line 67.
```

I fixed that by disabling the Debian GNOME helper sequence for this local build.

Edit `debian/rules` so the main `dh` line uses `--without gnome`:

```makefile
%:
	dh $@ --buildsystem=meson --without gnome
```

Then remove `dh-sequence-gnome` from `debian/control` if it is listed in `Build-Depends`.

You can check for remaining references with:

```bash
grep -R "dh-sequence-gnome\|dh_gnome\|gnome" -n debian
```

### 5. Build the package

```bash
DEB_BUILD_OPTIONS=nocheck dpkg-buildpackage -us -uc -b
```

When the build completes, go back to the parent directory:

```bash
cd ..
ls -1 *.deb
```

I built these packages:

```text
gir1.2-wp-0.5_0.5.15-1_amd64.deb
libwireplumber-0.5-0_0.5.15-1_amd64.deb
wireplumber_0.5.15-1_amd64.deb
```

### 6. Install the local WirePlumber 0.5.15 packages

```bash
sudo apt install \
  ./libwireplumber-0.5-0_0.5.15-1_amd64.deb \
  ./gir1.2-wp-0.5_0.5.15-1_amd64.deb \
  ./wireplumber_0.5.15-1_amd64.deb
```

If apt prints a warning like this:

```text
N: Download is performed unsandboxed as root as file ... couldn't be accessed by user '_apt'.
```

That warning is harmless. It happens because `_apt` cannot read the `.deb` files from your home/build directory. The install can still succeed.

### 7. Restart the user audio stack

Run this as your normal desktop user:

```bash
systemctl --user daemon-reload
systemctl --user restart wireplumber pipewire pipewire-pulse
```

Then verify the version:

```bash
wireplumber --version
```

Expected output:

```text
Compiled with libwireplumber 0.5.15
Linked with libwireplumber 0.5.15
```

Also check apt:

```bash
apt policy wireplumber libwireplumber-0.5-0 gir1.2-wp-0.5
```

### 8. Put the packages on hold

Since the PPA still offers 0.5.2 for Ubuntu 24.04 Noble, I put the local 0.5.15 packages on hold so apt does not replace them later.

```bash
sudo apt-mark hold wireplumber libwireplumber-0.5-0 gir1.2-wp-0.5
```

Rollback
--------

If 0.5.15 causes problems, unhold the packages and reinstall the PPA version:

```bash
sudo apt-mark unhold wireplumber libwireplumber-0.5-0 gir1.2-wp-0.5
sudo apt install --reinstall \
  wireplumber=0.5.2-8~ubuntu24.04 \
  libwireplumber-0.5-0=0.5.2-8~ubuntu24.04 \
  gir1.2-wp-0.5=0.5.2-8~ubuntu24.04
```

Then restart the audio stack again:

```bash
systemctl --user restart wireplumber pipewire pipewire-pulse
```

Result
------

After moving to WirePlumber 0.5.15 and putting the packages on hold, the Bluetooth profile switching issue appears to be resolved again.

I am not calling this permanently fixed yet. If the high CPU issue returns, I will update this post.

Summary
-------

If you are on Ubuntu 24.04 and `wireplumber` starts using 100% of a CPU core after switching Bluetooth profiles, check your WirePlumber version first.

Ubuntu 24.04’s default package is 0.4.17. The PipeWire Debian PPA got me to 0.5.2, which helped at first, but the issue later returned.

I am now running WirePlumber 0.5.15 built locally from the Debian source package, with the packages held in apt. Since 0.5.14 and 0.5.15 include more Bluetooth profile switching fixes, this may be the better path if the PPA version is still too old for your hardware.
