---
title: "OpenSCAD trick - negative rounding"
author: maave
type: post
date: 2026-01-14T00:00:00+00:00
url: /post/2026/01/14/negative-rounding/
draft: false
categories:
  - OpenSCAD
tags:
  - OpenSCAD
  - BOSL2
  - 3d-modeling
  - 3d
---


{{< image src="/uploads/2026/negative-chamfer-header.webp" alt="OpenSCAD pins with positive and negative chamfer">}}

Somehow, OpenSCAD has become my main 3D modeling program. It's incredibly powerful for parametric models but it's also brain-intensive. I speed up development with the BOSL2 library. Many BOSL2 objects have built-in rounding features so I don't need to invoke `minkowski()`.

A trick with BOSL2 is negative rounding. Instead of cutting a corner, it makes the corner stick out. This is very useful for smoothing sharp edges, tapering pin holes, and generally filleting odd shapes. Make an object with negative rounding/chamfer, then use it to cut another object with subtract().

<!--more-->

### Hinge pin hole

Here's my pin hole example. I made a small hinge that uses 1.75mm filament for the pin. If the hinges were misaligned the pin wouldn't pass through. The updated version uses negative chamfer to create a ramp. We cut the pin hole and both chamfers using a single object.

```openscad
cyl(
    d=hing_pin_diam, 
    l=hinge_width+smidge, 
    chamfer=-hing_pin_diam/6
);
```

The 3D result:

{{< image src="/uploads/2026/negative-chamfer.webp" alt="OpenSCAD pins with positive and negative chamfer">}}

### Sweep a 2D path

I used the same technique to soften a sharp edge on my phone case. This screen lip plauged me for a while: the case was sharp on my thumb and it was difficult to swipe the very edge of the screen. The original cut was sorta box shaped with flat sides so it cut a sharp 90-degree edge. The new version's negative rounding creates a smooth curved on the inside edge of the case. To make an odd-shaped object with the negative chamfer, I create a 2d object `round_rectangle()` that is the screen shape, then [offset_sweep()](https://github.com/BelfrySCAD/BOSL2/wiki/rounding.scad#functionmodule-offset_sweep) to give the screen-cutter thickness and rounding. Negative rounding creates the sticking-out lip which cuts a smooth round fillet.

```openscad
    screen_corners = [
        screen_radius + screen_extra_bottom_right,
        screen_radius + screen_extra_bottom_left,
        screen_radius + screen_extra_top_left,
        screen_radius + screen_extra_top_right,
    ];

    // 2d shape
    rectangle = square([screen_width, screen_length],center=true);
    // 2d round the corners
    round_rectangle = round_corners(rectangle, radius=screen_corners,$fn=highFn);

    // setup vars for 3d shape
    buffer = 0.05;
    smooth_edge_radius = 0.5; //TODO: auto adjust smooth_edge_radius
    smooth_edge_height = 0.1 + screen_lip;
    // set minimum size so function doesn't break, we could also simply disable
    smooth_edge_height2 = (smooth_edge_height>smooth_edge_radius)?smooth_edge_height:smooth_edge_radius+buffer;

    // 3d shape, use negative rounding for the stick-out lip
    offset_sweep( 
        round_rectangle, // sweep along the 2D shape
        height=smooth_edge_height2,
        top=os_circle(r=-smooth_edge_radius) //negative radius
     );

     // then use this object to cut a fillet
```

{{< image src="/uploads/2026/phone-case.webp" alt="phone case zoomed-out">}}

{{< image src="/uploads/2026/negative-rounding.webp" alt="phone case zoomed-in on rounded edge">}}