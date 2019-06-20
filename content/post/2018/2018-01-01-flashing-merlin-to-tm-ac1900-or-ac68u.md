---
title: Flashing Merlin to TM-AC1900 or AC68U (Updated for 384 builds)
author: Luis Rodriguez
type: post
date: 2018-01-01T03:07:34+00:00
url: /post/2018/01/01/flashing-merlin-to-tm-ac1900-or-ac68u/
featured_image: /wp-content/uploads/2018/01/asus-rt-ac5300.png
categories:
  - Projects
tags:
  - ac68u
  - asus
  - custom firmware
  - flash
  - merlin
  - router

---
Very few people know that you can save yourself some money by buying a [T-mobile router on Amazon][1]. Downside is having to flash it to have the latest firmware free of vulnerabilities. In this guide I will show you how to flash a TM-AC1900 to a RT-AC68U then into merlin custom firmware. I have tested it myself and can confirm that it works.

&nbsp;

Update: Asus began rolling back firmware on the stock setups. They are not providing Merlin with the source for Ai-Mesh so you need to be on stock to use it. I have updated the guide to provide instructions on how to prevent the rollback from happening. You can also use this guide to recover from a rollback.

<!--more-->

I will be posting soon about my friends dual router setup using two TM-AC1900&#8217;s.

## 1. Download and install Prerequisites:

You will need an active internet connection on another device to get the cfe file.

**1.1.** **Download all the needed files:**

[Download Link][2] with everything packaged together into a folder named &#8220;router&#8221; (recommended option).

Install Asus Restore Utility. Everything else packaged will run without installation.

**OR**

Individual links to latest downloads

  * [WinSCP][3]
  * [Putty][4]
  * [7zip][5]
  * [Asus Restore Utility][6] ([mirror][7]) ([Asus website][8])
  * [Flash Images AC68U][9]
  * [Latest Stock AC68U][10]
  * [Latest Merlin Firmware][11] ([mirror][12])

  1. Install WinSCP, Putty, Asus Restore Utility.
  2. Create a folder on your desktop, name it &#8220;router&#8221;.
  3. Download Flash Images AC68U and extract it to the newly made router folder. It contains a firmware to downgrade to, a firmware to upgrade to, and mtd-write.

&nbsp;

## 2. How to flash AC1900 to AC68U:

**2.1. Set static IP for PC**:

  * _start > run > ncpa.cpl > double click Ethernet > properties > IPV4 >_
  * _IP: 192.168.29.5_
  * _Subnet: Default (255.255.255.0)_
  * _Gateway: 192.168.29.1_

[<img class="wp-image-415" src="/uploads/2018/01/Untitled.png" width="400" height="455" srcset="/uploads/2018/01/Untitled.png 400w, /uploads/2018/01/Untitled-264x300.png 264w" sizes="(max-width: 400px) 100vw, 400px" />][13]

**2.2. Place router into Recovery/Restore mode:**

  * _With the router on hold reset button 10 seconds_
  * _Power off router (keep holding reset)_
  * _Wait 10 seconds, keep holding reset_
  * _Power on router holding reset for 10 more seconds (this part is very finicky, has to be exactly 10. Anymore or less wont put it into this mode)_

**2.3. Flash the older T-Mobile firmware:**

  * Go to 192.168.29.1 in a web browser
  
    _If Mini-CFE won’t load use Asus Restore Utility_

[<img class="wp-image-416" src="/uploads/2018/01/Untitled3.png" width="640" height="304" srcset="/uploads/2018/01/Untitled3.png 640w, /uploads/2018/01/Untitled3-300x143.png 300w" sizes="(max-width: 640px) 100vw, 640px" />][14]

  * Flash TM-AC1900\_3.0.0.4\_376_1703-g0ffdbba.trx from &#8220;Flash Images AC68U&#8221;
  * Wait for reboot 2-5 mins – WiFi lights will turn on when boot is complete

[<img class="wp-image-417" src="/uploads/2018/01/Untitled2.png" width="591" height="280" srcset="/uploads/2018/01/Untitled2.png 591w, /uploads/2018/01/Untitled2-300x142.png 300w" sizes="(max-width: 591px) 100vw, 591px" />][15]

**2.4. Flashing firmware, and cfe file:**

  * Log in to router at http://192.168.29.1 (admin:password or admin:admin)
  * Go to Administration > System > Enable SSH > Yes > Apply
  * Open Putty (ssh) and connect to 192.168.29.1 using a admin:password or admin:admin_
  
_ 

[<img class="wp-image-422" src="/uploads/2018/01/7.png" width="452" height="437" srcset="/uploads/2018/01/7.png 452w, /uploads/2018/01/7-300x290.png 300w" sizes="(max-width: 452px) 100vw, 452px" />][16]

  * Open WinSCP (_select SCP as file protocol)_ and connect to 192.168.29.1 using a admin:password or admin:admin

[<img class="wp-image-423" src="/uploads/2018/01/9.png" width="626" height="423" srcset="/uploads/2018/01/9.png 626w, /uploads/2018/01/9-300x203.png 300w" sizes="(max-width: 626px) 100vw, 626px" />][17]

  * In putty type:
  
    `cat /dev/mtd0 > original_cfe.bin`
  * In WinSCP refresh the window on the right side and note that original_cfe.bin is present
  * Copy original_cfe.bin to a local drive
  * Upload original_cfe.bin to <https://cfeditor.pipeline.sh/> > Select 1.0.2.0 US (AiMesh for 384+ builds) as Source CFE > Download the new .bin > rename it to new_cfe.bin (Optional, You can check Max for TX power to avoid doing step 5 again)

[<img class="wp-image-421" src="/uploads/2018/01/10-1024x562.png" width="678" height="372" srcset="/uploads/2018/01/10-1024x562.png 1024w, /uploads/2018/01/10-300x165.png 300w, /uploads/2018/01/10-768x421.png 768w" sizes="(max-width: 678px) 100vw, 678px" />][18]

  * Upload **new_cfe.bin** & **mtd-write** & **FW\_RT\_AC68U_30043763626.trx **to router through WinSCP. Files located inside &#8220;Flash Images AC68U&#8221; folder.
  * In Putty type:
  
    `chmod u+x mtd-write`
  * In Putty type:
  
    `./mtd-write new_cfe.bin boot`
  * In Putty type:
  
    `mtd-write2 FW_RT_AC68U_30043763626.trx linux`
  * You may get an error from putty terminal, you can ignore it.

**2.4. NVRAM reset:**

  * _Power off router_
  * _Wait 10 seconds_
  * _Press and hold WPS button_
  * _Power up the router and continue to hold WPS button for 15-20 seconds_<del></del>
  * wait for reboot 2-5 mins

**2.4. Flash Asus stock firmware:**

  * Reset PC IP back to default from steps 2.1

[<img class="wp-image-418" src="/uploads/2018/01/Untitled5.png" width="400" height="455" srcset="/uploads/2018/01/Untitled5.png 400w, /uploads/2018/01/Untitled5-264x300.png 264w" sizes="(max-width: 400px) 100vw, 400px" />][19]

  * Log in to router using 192.168.1.1 and the router is now an AC68U with 64MB jffs
  * Under Administration > Firmware Upgrade. upload **FW\_RT\_AC68U_3.0.0.4.382.18881.trx** located in the &#8220;Latest Stock AC68U&#8221; folder.

[<img class="wp-image-420" src="/uploads/2018/01/11-1024x712.png" width="678" height="471" srcset="/uploads/2018/01/11-1024x712.png 1024w, /uploads/2018/01/11-300x208.png 300w, /uploads/2018/01/11-768x534.png 768w, /uploads/2018/01/11.png 1449w" sizes="(max-width: 678px) 100vw, 678px" />][20]

  * Complete! If you continue, you will be going from stock asus to a custom firmware that will enable hardware acceleration and provide you with more benefits in performance and features.

## 3. How to flash AC68U to Merlin:

  * Make sure you are on the latest stock version, older versions use [different file structures][21] and may affect this installation.
  * Under Administration > Firmware Upgrade. upload **RT-AC68U\_380.69\_0.trx** located in the &#8220;Latest Merlin&#8221; folder.
  * Perform an NVRAM reset from steps 2.4
  * Complete! Its that simple.

## 4. AC68U Overclock:

The T-Mobile version of the router comes clocked at 800mhz. Overclock it to 1ghz with these instructions (only tested with Merlin build):

  * Make sure to have ssh enabled. Go to Administration > System > Enable SSH > LAN only > Apply
  * Log in to router using putty and type the following commands 
      * `nvram get clkfreq`
      * `nvram set clkfreq=1000,800`
      * `nvram commit`
  * Reboot

## 5. AC68U Max TX Power:

  * Log in to router using putty
  * In putty type:
  
    `cat /dev/mtd0 > original_cfe.bin`
  * In WinSCP refresh the window on the right side and note that original_cfe.bin is present
  * Copy original_cfe.bin to a local drive
  * Upload original_cfe.bin to <https://cfeditor.pipeline.sh/> > Select 1.0.2.0 US as Source CFE > Check Max TX Power > Download the new .bin > rename it to new_cfe.bin
  * Upload **new_cfe.bin** & **mtd-write**** **to router through WinSCP.
  * In Putty type:
  
    `chmod u+x mtd-write`
  * In Putty type:
  
    `./mtd-write new_cfe.bin boot`
  * Reboot

&nbsp;

**Edit:**

I updated my router firmware to the latest 384 build of Merlin which added mesh networking. I noticed my overclock settings and tx power were reset back to stock settings. After performing the above steps all I had to add to it was to Select 1.0.2.0 US **AiMesh** as Source CFE for cfeditor.pipeline.sh website for firmware versions 384+

 [1]: http://a.tra.li/Smz7
 [2]: https://downloads.techreanimate.com/pmafuh
 [3]: http://a.tra.li/Smz6
 [4]: http://a.tra.li/Smz5
 [5]: http://a.tra.li/Smz4
 [6]: https://downloads.techreanimate.com/qsnqpu
 [7]: http://a.tra.li/Smz2
 [8]: http://a.tra.li/Smz3
 [9]: https://downloads.techreanimate.com/jijryr
 [10]: http://a.tra.li/Smz1
 [11]: http://a.tra.li/Smz0
 [12]: https://downloads.techreanimate.com/fcgxkf
 [13]: /uploads/2018/01/Untitled.png
 [14]: /uploads/2018/01/Untitled3.png
 [15]: /uploads/2018/01/Untitled2.png
 [16]: /uploads/2018/01/7.png
 [17]: /uploads/2018/01/9.png
 [18]: /uploads/2018/01/10.png
 [19]: /uploads/2018/01/Untitled5.png
 [20]: /uploads/2018/01/11.png
 [21]: https://github.com/RMerl/asuswrt-merlin/wiki/Installation