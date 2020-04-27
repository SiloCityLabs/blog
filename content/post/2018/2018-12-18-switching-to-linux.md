---
title: Switching to Linux
author: Luis Rodriguez
type: post
date: 2018-12-19T00:30:20+00:00
url: /post/2018/12/18/switching-to-linux/
featured_image: /wp-content/uploads/2018/12/the-best-linux-distros-for-beginners-switching-from-windows-mac-670x3353674761306310635625.png
categories:
  - Projects
tags:
  - linux
  - steam
  - ubuntu

---
{{< image src="/uploads/2018/12/the-best-linux-distros-for-beginners-switching-from-windows-mac-670x3353674761306310635625.png" alt="">}}

I have just about had it with [Windows updates](https://www.softwarehow.com/fix-windows-stuck-checking-updates/) being forced on me. Graphics drivers breaking. My settings getting reset every time I get on and many [other reasons][1] I won't begin to rant about. I chose to go with a decently common operating system for better compatibility.

I have gone through the [install process and have installed Ubuntu 18.04LTS][2]. For less advanced Linux users you can go with a simpler Linux like [Solus][3] which is a little more GUI friendly system.

You can modify ubuntu to your likings using a tool called "gnome tweaks" and you can disable animations to speed things up by using a tool called "dconf-editor".

In dconf-editor, browse to **org.gnome.desktop.interface** and set **enable-animations=false**.

You can also run this command to disable animations:
  
`gsettings set org.gnome.desktop.interface enable-animations false` 

You do not need to log out and back in, it should take effect immediately.

After getting all settled and having all my software carried over. I had a few pieces of software that was not compatible and I still needed. Most of my software was available in Ubuntu application store. Everything left is listed below:

  * [Chrome][4]
  * [Steam][5]
  * Git -> installed from terminal "sudo apt install git"
  * [Java 10 (64-bit) -> installed from terminal][6]
  * [AMD Driver][7]
  * [PC Games, Steam Beta Proton][8]
  * SketchUp 2016 -> Windows VM
  * [Origin][9]
  * H&R Block Premium + Efile + State -> Windows VM
  * Microsoft office -> libreoffice

Ubuntu [Get right click new documents back][10]. Here is a link to a couple of [templates][11] we made.

To get the Windows-only software running I installed Virtualbox and installed Windows 10 and started to hack away at all its imperfections to slim it down. This virtual machine won't be used that often so its a compromise I'm willing to make. I can also set quick one click restore points just in case windows breaks. I have compiled a list for all to use:

<!--more-->

You can do much of the following using [w10privacy.exe](https://www.winprivacy.de/english-home/)

  * Enable bidirectional clipboard and drag and drop in Virtualbox
  * Disabled UAC - Annoying in VMs, slight delay in the dialog
  * Power options, enabled performance
  * [Turned off Windows defender][12]
  * Deffered windows updates to 30 days or max
  * Disabled firewall, since I will be enabling NAT in VM for Ubuntu to manage firewall
  * Disabled Microsoft smart screen
  * installed Virtualbox extensions pack
  * installed VM guest additions pack
  * Disabled sleep in power options
  * [Performance options - adjust for best performance][13]
  * Disabled cortana
  * Disabled any windows "sharing" data stuff
  * Single click shutdown - "%windir%\System32\shutdown.exe /s /t 0"
  * [Disable hibernate][14]
  * [Auto login user][13]
  * Hide taskbar items, People, cortana, task view
  * Made start menu smaller by dragging and removed most icons from right side
  * Disable system restore points
  * [Disabled system dumps and startup logs][15]
  * Disabled windows defender in task manager startup items
  * Disabled ipv6
  * Set static ip using existing ip as baseline (10.0.2.15, 255.255.255.0, 10.0.2.2, DNS Your router ip, 1.1.1.1). Since we will be disabling DHCP service.
  * Disabled windows cloud search
  * Disabled device history
  * Right click Uninstall any apps in start menu you don't need
  * Indexing options -> removed Users, I won't be using this for files. Left start menu alone
  * [Disable page file, Only do if you have enough memory][16]
  * set solid background
  * set solid taskbar colors and disable transparency
  * Turned of windows features: 
      * Internet explorer
      * Microsoft print to PDF
      * Microsoft XPS document writer
      * Powershell 2
  * Disabled services: 
      * Data usage
      * Windows Defender firewall
  * Manual services 
      * Print Spooler - I may still need it to print out tax software stuff
      * DHCP
  * Combine start menu folders by right click open file location then drag the shortcuts. Combined Ease of access, powershell and system into one
  * [Killed windows apps][17] - Powershell (admin) Run the commands below (copy and Paste) and hit Enter - 
      * get-appxpackage -allusers \*zunemusic\* | remove-appxpackage
      * get-appxpackage -allusers \*3dbuilder\* | remove-appxpackage
      * get-appxpackage -allusers \*alarms\* | remove-appxpackage
      * get-appxpackage -allusers \*communicationsapps\* | remove-appxpackage
      * get-appxpackage -allusers \*camera\* | remove-appxpackage
      * get-appxpackage -allusers \*feedback\* | remove-appxpackage
      * get-appxpackage -allusers \*officehub\* | remove-appxpackage
      * get-appxpackage -allusers \*getstarted\* | remove-appxpackage
      * get-appxpackage -allusers \*skypeapp\* | remove-appxpackage
      * get-appxpackage -allusers \*zunemusic\* | remove-appxpackage
      * get-appxpackage -allusers \*zune\* | remove-appxpackage
      * get-appxpackage -allusers \*maps\* | remove-appxpackage
      * get-appxpackage -allusers \*messaging\* | remove-appxpackage
      * get-appxpackage -allusers \*solitaire\* | remove-appxpackage
      * get-appxpackage -allusers \*wallet\* | remove-appxpackage
      * get-appxpackage -allusers \*connectivitystore\* | remove-appxpackage
      * get-appxpackage -allusers \*bingfinance\* | remove-appxpackage
      * get-appxpackage -allusers \*bing\* | remove-appxpackage
      * get-appxpackage -allusers \*zunevideo\* | remove-appxpackage
      * get-appxpackage -allusers \*bingnews\* | remove-appxpackage
      * get-appxpackage -allusers \*onenote\* | remove-appxpackage
      * get-appxpackage -allusers \*oneconnect\* | remove-appxpackage
      * get-appxpackage -allusers \*people\* | remove-appxpackage
      * get-appxpackage -allusers \*commsphone\* | remove-appxpackage
      * get-appxpackage -allusers \*windowsphone\* | remove-appxpackage
      * get-appxpackage -allusers \*phone\* | remove-appxpackage
      * get-appxpackage -allusers \*photos\* | remove-appxpackage
      * get-appxpackage -allusers \*bingsports\* | remove-appxpackage
      * get-appxpackage -allusers \*sticky\* | remove-appxpackage
      * get-appxpackage -allusers \*sway\* | remove-appxpackage
      * get-appxpackage -allusers \*3d\* | remove-appxpackage
      * get-appxpackage -allusers \*soundrecorder\* | remove-appxpackage
      * get-appxpackage -allusers \*bingweather\* | remove-appxpackage
      * get-appxpackage -allusers \*holographic\* | remove-appxpackage
      * get-appxpackage -allusers \*windowsstore\* | remove-appxpackage
      * get-appxpackage -allusers \*xbox\* | remove-appxpackage
      * get-appxpackage -allusers \*mspaint\* | remove-appxpackage
      * get-appxpackage -allusers \*holographic\* | remove-appxpackage

Other untested stuff:

  * http://a.tra.li/Uftj
  * http://a.tra.li/Ufti

 [1]: https://duckduckgo.com/?q=windows+updates+breaks&atb=v147-1__&t=cros&iar=news&ia=
 [2]: https://linuxhint.com/rufus_bootable_usb_install_ubuntu_18-04_lts/
 [3]: https://linoxide.com/distros/install-solus-usb/
 [4]: http://google.com/chrome
 [5]: https://store.steampowered.com/about/
 [6]: https://thishosting.rocks/install-java-ubuntu/
 [7]: http://amd.com
 [8]: https://steamcommunity.com/games/221410/announcements/detail/1696055855739350561
 [9]: https://www.pcsteps.com/5110-install-origin-linux-mint-ubuntu-wine/
 [10]: https://vitux.com/add-new-document-back-to-the-right-click-menu-in-ubuntu-18-04/
 [11]: https://downloads.techreanimate.com/mprryu
 [12]: https://www.ghacks.net/2015/10/25/how-to-disable-windows-defender-in-windows-10-permanently/
 [13]: https://www.itpro.co.uk/operating-systems/26138/how-to-speed-up-windows-10
 [14]: https://www.pugetsystems.com/labs/support-software/How-to-disable-Sleep-Mode-or-Hibernation-793/
 [15]: https://www.inteset.com/how-to-strip-down-windows-10-for-kiosks-digital-signage-and-other-special-purpose-systems
 [16]: https://tunecomp.net/win10-page-file-disable/
 [17]: https://www.askvg.com/guide-how-to-remove-all-built-in-apps-in-windows-10
