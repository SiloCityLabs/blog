---
title: 3d printed wrist brace – WIP instructions
author: maave
type: post
date: 2018-08-30T00:09:59+00:00
url: /post/2018/08/30/3d-printed-wrist-brace-wip-instructions/
categories:
  - Projects
tags:
  - 3d modeling
  - 3d printing
  - 3d scanning
  - Blender
  - kinect
  - wrist brace

---
I won&#8217;t be able to finish [my wrist brace project][1] for a while. These instructions are very rough because they&#8217;re not actually finished. 

Post in comments if anything needs clarification.

# Scanning

requirements:

  * Windows
  * Kinect (with power adapter)
  * beefy CPU or an NVidia GPU (with CUDA drivers)
  * Kinect for Windows SDK 2.0
  * Skanect software
  * an assistant
  * a small open area

Here&#8217;s the Skanect guide I watched:

<https://www.youtube.com/watch?v=_cKb3oEM47E>

<!--more-->

Optional:  
Scanning was initially a huge pain in the ass. Moving to much would cause Skanect to fail. My fix was to strap a screen to the back of the Kinect so that my assistant could watch the Skanect UI.

  * run VNC server (or RDP or whatever) on the computer
  * use a smartphone and connect to VNC
  * strap the smartphone to the back of the Kinect (I used velcro straps and rubber bands)

Now the assistant can watch Skanect while moving around.

alternative raspi scanner that I have not tried:  
<https://www.slashgear.com/raspberry-pi-2-kinect-make-for-a-handy-3d-scanner-06391772/>

Scanning requires either a beefy CPU or CUDA acceleration from an NVidia GPU. Slow hardware requires slooowwww movement of the Kinect. Moving too fast will cause the scan to fail.  
Install Skanect  
If you have an NVidia card, install CUDA drivers (takes like 20-30min to install)[<img class="wp-image-581 alignright" src="/uploads/2018/08/orig-scan-watertight-300x169.jpg" width="300" height="169" srcset="/uploads/2018/08/orig-scan-watertight-300x169.jpg 300w, /uploads/2018/08/orig-scan-watertight-768x432.jpg 768w, /uploads/2018/08/orig-scan-watertight-1024x576.jpg 1024w" sizes="(max-width: 300px) 100vw, 300px" />][2]

After scanning you can do some basic editing in Skanect. Make the model watertight and do a rough crop. You can edit more in Blender later. Here&#8217;s what my original scan looks like after making it watertight. There was a small indent that I had to fix in Blender.

Skanect scales in meters and I haven&#8217;t figured out how to import/export with Blender properly (you&#8217;re supposed to do it in the scene but it didn&#8217;t work for me). After editing your model in Blender, export the STL, open in Windows 3d Builder, select the &#8220;m&#8221; for meters, fix any errors, then save (as STL). Now it&#8217;s scaled.

# Blender

Blender tutorial. If you&#8217;ve never used Blender then I recommend watching up to part 5 and following along.  
<https://www.youtube.com/playlist?list=PLjEaoINr3zgHs8uzT3yqe4iHGfkCmMJ0P>

Use your new Blender skills to edit the model. Crop it as needed, fix irregularities, and bisect it.

Bisect to split the brace down the middle  
<https://www.youtube.com/watch?v=UjLaxaLJZK4><del></del>

[<img class="wp-image-579 alignright" src="/uploads/2018/08/01-wtf-300x144.png" width="300" height="144" srcset="/uploads/2018/08/01-wtf-300x144.png 300w, /uploads/2018/08/01-wtf-768x369.png 768w, /uploads/2018/08/01-wtf-1024x492.png 1024w" sizes="(max-width: 300px) 100vw, 300px" />][3]

My model had some random garbage points that were sticking off the main model. That required cleanup.

Clean up garbo:  
-Edit Mode  
-select all  
-Mesh > Cleanup > Degenerate Dissolve

Time to skeletonize it with modifiers. Were reducing the polys, then wireframing the simplified geometry, then smoothing it to remove hard corners. Switch to object mode. Add these modifiers in order and play around with all the settings.

decimate

-adjust the ratio, play with settings

wireframe

-uncheck even thickness  
-check boundary  
-adjust thickness

subdivide (smooth)

-adjust Views for smoothing

Here what my model looked like early on. I wireframed it without reducing the polys and it created this fine hole pattern:[<img class="wp-image-582 aligncenter" src="/uploads/2018/08/07-tada-its-fixed-300x144.png" width="300" height="144" srcset="/uploads/2018/08/07-tada-its-fixed-300x144.png 300w, /uploads/2018/08/07-tada-its-fixed-768x369.png 768w, /uploads/2018/08/07-tada-its-fixed-1024x492.png 1024w" sizes="(max-width: 300px) 100vw, 300px" />][4]

Here&#8217;s my latest version. This is after bisecting it, and decimate and subdivide modifiers. It&#8217;s basically done, I just have to make it at least 4x thicker because this was too flexible once printed.

[<img class="wp-image-583" src="/uploads/2018/08/final-ish-300x169.jpg" width="300" height="169" srcset="/uploads/2018/08/final-ish-300x169.jpg 300w, /uploads/2018/08/final-ish-768x432.jpg 768w, /uploads/2018/08/final-ish-1024x576.jpg 1024w" sizes="(max-width: 300px) 100vw, 300px" />][5]

I also have a method for adding holes for velcro straps. It&#8217;s not strictly necessary but it&#8217;ll make the slots line up nicely.

Boolean Modifier  
<https://www.youtube.com/watch?v=pcUe-ab_jQs>

-create some cylinders or boxes that pass through the spots you want holes  
-to the arm, add a Boolean Modifier, select Difference, select a cylinder/box  
-hide boxes  
-if it&#8217;s good hit Apply to permanently apply it  
-delete the faces of the beams that are created

export to STL

&nbsp;

In the future swap the velcro for a hinge:

<https://www.thingiverse.com/thing:1615753>

&nbsp;

 [1]: https://blog.silocitylabs.com/post/2017/12/15/3d-printed-wrist-brace-wip/
 [2]: /uploads/2018/08/orig-scan-watertight.jpg
 [3]: /uploads/2018/08/01-wtf.png
 [4]: /uploads/2018/08/07-tada-its-fixed.png
 [5]: /uploads/2018/08/final-ish.jpg