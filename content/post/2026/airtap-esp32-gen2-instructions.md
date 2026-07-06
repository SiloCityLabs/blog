---
title: How to Install PCB on a Gen2 AC Infinity AirTap T4 or T6
subtitle: 
author: Luis Rodriguez
type: post
date: 2026-06-03
categories:
  - Projects
tags:
  - ESP32
  - airtap-t4
  - airtap-t6
  - airtap-gen2
  - xiao-seeed
  - xiao
  - esphome
  - home-assistant
  - made-for-esphome

---
  
Owners of the AC Infinity AirTap T4/T6 with a Gen2 vent can replace the factory control board with our ESP32 module to add local smart home control through ESPHome and Home Assistant. The Gen2 housing keeps the original button bracket in place. The kit includes a small display cover bracket made for the Gen2 layout. This guide walks through what is in the kit, how that cover bracket fits, and the exact steps to install the PCB.

This guide explains the parts included with the module, how the display cover bracket works, and the step-by-step installation process.

**Plug and Play Option**

For the fastest install, we offer a complete ESP32 board that arrives ready to mount. No board modification is required. Connect the cables, fit the display cover bracket, and install the module in the vent.

You can purchase the complete module here:

<https://shop.silocitylabs.com/products/esp32-module-for-ac-infinity-airtap-t4-t6>

**Installation Guide**

Follow the steps below to install the ESP32 module in a Gen2 vent. The OEM button bracket stays in place. There is no button transfer.

 1. Remove the six rear screws from the vent.

    {{< image src="/uploads/2025/airtap-gen3/airtap-gen3-kit-step-9.jpg" alt="airtap gen2 rear">}}

 1. Disconnect the power, fan and NTC cables. Some may be glued, use caution removing not to tear cable. Use pry tool to seperate glue.

    {{< image src="/uploads/2026/airtap-gen2/airtap-gen2-image1.webp" alt="Disconnect power, fan, and NTC cables">}}

 1. Unscrew and remove the OEM PCB. Leave the original bracket, buttons, and clear display shield in the vent.

    {{< image src="/uploads/2026/airtap-gen2/airtap-gen2-image6.webp" alt="Remove OEM PCB, keep bracket and display shield">}}

 1. Install the display cover bracket from the kit. Align it with the clear display shield and housing.

    {{< image src="/uploads/2026/airtap-gen2/airtap-gen2-image4.webp" alt="Install display cover bracket">}}

 1. Mount the ESP32 module. Use original screws to reattach.

    {{< image src="/uploads/2026/airtap-gen2/airtap-gen2-image5.webp" alt="Mount ESP32 module with washer and screw">}}

 1. Reconnect all cables.

    {{< image src="/uploads/2026/airtap-gen2/airtap-gen2-image1.webp" alt="Reconnect all cables">}}

 1. Reinstall back pannel.

    {{< image src="/uploads/2025/airtap-gen3/airtap-gen3-kit-step-9.jpg" alt="airtap gen2 rear">}}

**Getting Your Vent Online**

With the display cover bracket fitted and cables connected, your Gen2 AirTap T4/T6 is ready for ESPHome and Home Assistant. Flash or adopt the device using the SiloCityLabs device profile for your button count, then verify fan control and temperature readings before closing up the install.
