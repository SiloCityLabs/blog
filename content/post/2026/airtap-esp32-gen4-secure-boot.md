---
title: AC Infinity Almost Got the Next AirTap Right, Then Locked Us Out With Secure Boot
subtitle: Reverse engineering the ESP32-C6 Matter-capable AirTap, dumping OEM firmware, and why open source still needs a module swap or a new PCB
author: Luis Rodriguez
type: post
date: 2026-07-14
categories:
  - Projects
tags:
  - ESP32
  - ESP32-C6
  - airtap
  - airtap-gen4
  - reverse-engineering
  - secure-boot
  - flash-encryption
  - esphome
  - home-assistant
  - matter
  - open-source

---

For years the AC Infinity AirTap has been a popular register booster with a frustrating smart-home story: closed apps, Bluetooth-only workflows, and a control board that was never meant for local automations. We already solved that for earlier generations with drop-in ESPHome boards see the Gen2 install guide and shop module if you have that hardware.

Then a newer AirTap showed up with something that almost made me celebrate out loud: an **ESP32-C6-WROOM-1**.

That is the right silicon. Wi‑Fi 6, BLE, and IEEE 802.15.4 for Thread/Matter. On paper AC Infinity finally moved the AirTap onto a platform the open source community can actually live with. Dump the pads, port ESPHome, keep the fan and HVAC logic local to Home Assistant. Done.

Except they also burned **secure boot** and **flash encryption** into the eFuses.

So the device falls short for the open source community in the one place that matters: you cannot put your own firmware on the module that ships in the vent. The only practical path left is to **desolder the locked ESP32-C6 and replace it**, or design a **new plug-and-play PCB** for everyone else.

All of the dumps, serial logs, pin traces, and notes for this work live in our firmware repo:

- Project: [SiloCityLabs/esp32-airtap](https://github.com/SiloCityLabs/esp32-airtap)
- Gen-4 notes: [Airtap-Tx/Gen-4](https://github.com/SiloCityLabs/esp32-airtap/tree/main/Airtap-Tx/Gen-4)
- Full OEM flash dump: [oem-firmware-full.bin](https://github.com/SiloCityLabs/esp32-airtap/tree/main/Airtap-Tx/Gen-4/firmware-dump)

## What “almost correct” looks like

The hardware header is refreshingly honest. Pads are exposed for:

- `IO9` (BOOT / download strap)
- `GND`
- `TX` / `RX`
- `3.3V`
- `EN`

UART console runs at **115200 8N1**. Download mode is the classic C6 dance: hold `IO9` low, pulse `EN`, release `IO9`, then talk to the ROM with `esptool --before no-reset`.

Boot logs identify the OEM tree immediately:

- Project: `smarthome_device_airtap` **v2.1**
- ESP-IDF **v5.4.1**, built Apr 9 2026
- Reported product versions: hardware **3.0**, software **3.0.15**
- Power-on path: BLE / Blufi-style provisioning on a Matter-capable C6

That is AC Infinity meeting the industry where it lives now. ESP32 instead of a dead-end MCU. Modern IDF. A chip that can do Matter. From a product roadmap perspective: applause.

From an open source / right-to-repair perspective: keep reading.

## Dumping the OEM firmware

With the board in download mode we pulled the entire SPI flash.

- Chip: **ESP32-C6-WROOM-1** (QFN40, rev v0.2)
- External flash detected: **8MB** (manufacturer `c8`, device `4017`)
- Artifact: `Airtap-Tx/Gen-4/firmware-dump/oem-firmware-full.bin`
- SHA256: `ff35eccd4f6936a8b0c3f98f9bc9d93c65edfefdc661ccb78f1077a39ab8743d`

Also captured:

- [oem-security-info.txt](https://github.com/SiloCityLabs/esp32-airtap/tree/main/Airtap-Tx/Gen-4/firmware-dump)
- [oem-efuse-summary.txt](https://github.com/SiloCityLabs/esp32-airtap/tree/main/Airtap-Tx/Gen-4/firmware-dump)
- [oem-partition-table.bin](https://github.com/SiloCityLabs/esp32-airtap/tree/main/Airtap-Tx/Gen-4/firmware-dump) (ciphertext under encryption)

Raw flash contents are **encrypted ciphertext**. Useful as an archival restore image. Not useful as a pile of readable strings and app code without the chip’s XTS key.

## The lock that kills community firmware

`espefuse` / `get-security-info` on this unit is unambiguous:

| Lock | State |
| --- | --- |
| Secure Boot | Enabled (RSA-PSS) |
| Flash encryption | Permanently enabled (`SPI_BOOT_CRYPT_CNT = 0b111`) |
| Flash encrypt key (`BLOCK_KEY1` / XTS-AES-128) | Present, **not readable** (`RD_DIS`) |
| HMAC / other key blocks | Read-disabled where it matters |
| JTAG | Permanently disabled |
| Download mode | Still allowed (we can write flash) |
| Download-mode encrypt | Still allowed |

Runtime logs even call flash encryption “DEVELOPMENT (not secure)” but the eFuses tell the real story. Encryption is burned on, the key is not extractable, and secure boot digests are live.

So of course we tried the thing the community always tries: build ESPHome and flash it.

Write with on-chip encrypt succeeds. Boot does not:

```text
Valid secure boot key blocks: 0
No signature block magic byte found at signature sector (found 0xff not 0xe7). Image not V2 signed?
RSA-PSS secure boot verification failed
```

Then the ROM watchdog bootloops forever. Unsigned ESPHome or any third-party image will never pass AC Infinity’s secure boot key.

We restored the OEM dump afterward (hash verified) so the board boots stock again. Scripts and notes for that path are in the same Gen-4 folder.

**Translation:** this is not “we haven’t written the ESPHome YAML yet.” This is “the silicon refuses unsigned firmware forever.”

## Reverse engineering the PCB anyway

Even when flash is a dead end for custom code, tracing the board still matters. If we replace the module or spin a compatible PCB we need the pin map.

Hand traces plus GPIO init from the OEM UART line up cleanly for most of the board:

### Confirmed / high-confidence map

| GPIO | Function |
| --- | --- |
| IO0 | Via R29 → **Q5** → piezo / buzzer path |
| IO1 | Via R6/R34 → **Q2** PWM / power drive side of the fan path |
| IO2 | **NTC** (ADC / Curve Fitting calibration in OEM logs) |
| IO5 | TTP223 touch 1 |
| IO6 | TTP223 touch 2 |
| IO7 | TTP223 touch 3 |
| IO8 | Q4 → CS1621 `/CS` |
| IO9 | BOOT pad |
| IO10 | Q13 → CS1621 `/WR` |
| IO11 | Q14 → CS1621 `DATA` |
| IO13 | **IRM3638** IR receiver (RMT RX in firmware) |
| IO15 / IO18 / IO19 / IO23 | **ULN2003A** stepper / damper motor coils |
| IO21 | TTP223 touch 4 (firmware input; hard to probe on the pad) |
| IO22 | TTP223 touch 5 |

### Display driver

The segment / LED glass is driven by a **CS1621** over a bit-banged `/CS`, `/WR`, `DATA` trio on IO8 / IO10 / IO11 familiar territory if you have worked previous AirTap RE, just now tied to C6 GPIOs.

### What serial proved live

While poking the running OEM UI we correlated:

- Local **Up Key** events → `ipcd_set_speed()`
- IR remote packets via `ct_ir_manage` (power, level, screen on/off, mute, lock)
- Five pulled-up inputs matching the five touch channels
- Motor and CS1621 pins as outputs exactly where the traces said

PWM / piezo / backlight still want a scope on an unlocked module or a purpose-built pinprobe board the OEM image will not run our GPIO test firmware but the transistor stack on IO0/IO1 is already enough to plan a replacement design.

## Why this is frustrating

AC Infinity did the hard part of the hardware generation transition:

1. Left behind the old closed MCU era.
2. Picked an Espressif part the Home Assistant / ESPHome / Matter world already understands.
3. Even left a UART header that makes RE and factory debugging possible.

Then they flipped the one switch that makes all of that irrelevant to us: **production secure boot with an unavailable private signing key**.

Secure boot is a legitimate product security feature for cloud-tied devices. It is also an intentional wall against local firmware ownership. On a vent fan whose job is “move air in my house,” locking the bootROM so Home Assistant users cannot replace Blufi/app glue with ESPHome is a policy choice, not a technical necessity of the ESP32-C6.

## Where we go from here

Three tracks, in order of how we plan to ship community hardware:

1. **Module transplant** desolder the locked ESP32-C6-WROOM-1, solder an unlocked module, flash ESPHome against the pin map above.
2. **Drop-in PCB** design a SiloCityLabs board that matches connectors, display, touches, IR, motor driver, and NTC so owners get Gen2-style plug-and-play instead of hot-air surgery.
3. **Keep publishing dumps** full flash images, eFuse summaries, serial captures, and traces stay public so nobody repeats this RE from zero.

Firmware artifacts and working notes:

- Repo: <https://github.com/SiloCityLabs/esp32-airtap>
- Gen-4 tree: <https://github.com/SiloCityLabs/esp32-airtap/tree/main/Airtap-Tx/Gen-4>
- Firmware dump folder: <https://github.com/SiloCityLabs/esp32-airtap/tree/main/Airtap-Tx/Gen-4/firmware-dump>

Earlier AirTap ESPHome installs (unlocked / replacement boards) remain the right answer for Gen1/Gen2 vents today. This post is about the new C6 generation specifically: how close it got, how hard the eFuses slammed the door, and how we still intend to get ESPHome into those registers.

If you have one of these boards and want to help validate pin leftovers (backlight, exact PWM edge cases) on an unlocked C6, open an issue on the repo. We will keep the Gen-4 folder updated as the replacement path solidifies.
