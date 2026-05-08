---
title: Native Apple Home Firmware for the AC Infinity AirTap ESP32 Upgrade Kit with HomeSpan and PlatformIO
subtitle: A practical developer guide for bringing AirTap T4/T6 vents directly into Apple Home without Home Assistant
author: Luis Rodriguez
type: post
date: 2026-05-06
categories:
  - Projects
tags:
  - ESP32
  - airtap-t4
  - airtap-t6
  - homekit
  - apple-home
  - homespan
  - platformio
  - xiao-seeed
  - smart-vent
  - smart-home
  - silocitylabs

---

The SiloCityLabs ESP32 Module for the AC Infinity AirTap T4/T6 already gives Home Assistant users a clean way to upgrade an AirTap vent into a connected smart-home device. The current ESPHome firmware path supports the important hardware pieces: PWM fan control, onboard temperature sensing, OLED display output, physical buttons, panel lockout, and newer IR receiver support on supported builds.

But there is another smart-home audience that should care about this hardware: **Apple Home users**.

If you use an iPhone, HomePod, Apple TV, Siri, scenes, automations, and the Apple Home app, the natural question is:

**Could the AirTap ESP32 upgrade kit run native Apple Home firmware instead of ESPHome?**

The short answer is yes, it looks very possible.

The better answer is that [HomeSpan](https://github.com/HomeSpan/HomeSpan) gives developers a realistic path to build firmware that pairs the AirTap directly with Apple Home over Wi-Fi, without Home Assistant, without a Home Assistant HomeKit Bridge, and without a separate server translating ESPHome entities into HomeKit accessories.

This post is not a finished firmware release. It is a deep technical roadmap for developers, makers, and advanced AirTap owners who want to understand what it would take to build one. Think of it as the map: enough architecture to start the project, enough code structure to see the direction, and enough caution to avoid the obvious traps.

<!--more-->

**Why Apple Home Support Matters**
----------------------------------

Most custom ESP32 smart-home firmware assumes Home Assistant is the center of the house. That is fine for power users, and ESPHome is still one of the best ways to ship an ESP32-based device quickly.

But not every smart-home user wants to run Home Assistant.

A lot of Apple users want smart vents, quieter airflow, room-level automation, and better HVAC behavior, but they want the control surface to be the Apple Home app. They want to say:

- “Hey Siri, turn on the office vent.”
- “Set the bedroom vent to 40 percent.”
- “Turn off the vent display buttons.”
- “When the room gets warm, increase airflow.”
- “When the Good Night scene runs, lower the vent speed.”

That is a different market from the existing ESPHome/Home Assistant crowd.

It is not better or worse. It is just different.

For SiloCityLabs, this is the opportunity: keep serving the Home Assistant community while opening the door to Apple-first smart-home users who may never search for ESPHome, YAML, or Home Assistant integrations.

**The Existing AirTap ESP32 Hardware Is a Good Fit**
---------------------------------------------------

The reason this idea is worth exploring is that the AirTap ESP32 upgrade kit already has the right kind of hardware for native Apple Home firmware.

From the existing ESPHome configurations and current SiloCityLabs product direction, the firmware needs to deal with:

- ESP32 module running Wi-Fi firmware
- PWM output for AirTap fan speed control
- ADC input for temperature sensing
- SSD1306 128x64 OLED display over I2C
- Local GPIO buttons
- Stored fan speed state
- Panel lockout state
- Optional IR receiver support on supported board versions
- OTA/update strategy
- Wi-Fi provisioning
- A clean end-user pairing/reset story

The current ESPHome firmware proves the board can already do the important work. HomeSpan would not replace the hardware behavior. It would replace the smart-home protocol layer.

ESPHome exposes entities to Home Assistant.

HomeSpan exposes HomeKit services directly to Apple Home.

That distinction matters for SEO, positioning, and buyer intent. “ESP32 smart vent for Home Assistant” reaches one audience. “Apple HomeKit smart vent upgrade for AC Infinity AirTap” reaches another.

**What HomeSpan Brings to the Project**
--------------------------------------

HomeSpan is an Arduino library for creating HomeKit accessories on ESP32. It implements Apple’s HomeKit Accessory Protocol for ESP32-class devices and allows the device to pair directly with Apple Home over Wi-Fi.

For this project, HomeSpan is attractive because it can model the AirTap as an actual Apple Home accessory instead of a group of translated Home Assistant entities.

A first Apple Home firmware could expose:

- A Fan service for on/off and speed control
- A Temperature Sensor service for the AirTap temperature reading
- A Switch service for panel lockout
- Accessory Information for manufacturer, model, serial, and firmware version
- Optional support for physical button events or local-only button behavior

From the Apple Home user’s perspective, the AirTap would feel like a native accessory. They would not need to understand ESPHome. They would not need to copy YAML. They would not need to run a local server.

They would flash firmware, pair the device, name the vent, assign it to a room, and automate it.

That is the user story.

**Important Reality Check**
---------------------------

This should not be marketed as an official shipping firmware until it is tested on actual hardware across board revisions.

There are three reasons:

1. **HomeKit behavior has to feel polished.** Pairing, reset, Wi-Fi provisioning, and accessory naming need to be reliable.
2. **Fan control must be safe and predictable.** PWM values should match the existing firmware behavior closely enough that users are not surprised.
3. **Board revisions matter.** The 3-button and 4-button AirTap configurations are similar, but not identical. IR support also changes the firmware shape.

So the right framing is:

> This is a developer path toward native Apple Home support for the AirTap ESP32 upgrade kit.

Not:

> This is a finished Apple Home firmware release.

That honesty is important. It builds trust with advanced users and keeps expectations under control.

**What We Can Reuse from the ESPHome Firmware**
----------------------------------------------

The current ESPHome firmware is the best technical reference because it already describes the board behavior.

The 3-button firmware uses this core pattern:

- `GPIO2` for PWM fan output
- `GPIO4` for ADC temperature input
- `GPIO6` for OLED SDA
- `GPIO7` for OLED SCL
- `GPIO8`, `GPIO9`, and `GPIO10` for local buttons
- A stored integer fan speed from `0` to `10`
- A panel lockout boolean
- An NTC temperature conversion pipeline
- An SSD1306 128x64 display at I2C address `0x3C`

The 4-button firmware follows the same general structure, but adds the fourth button and IR receiver support:

- `GPIO2` for PWM fan output
- `GPIO4` for ADC temperature input
- `GPIO6` for OLED SDA
- `GPIO7` for OLED SCL
- `GPIO20` for Up
- `GPIO8` for Down
- `GPIO9` for Toggle
- `GPIO10` for Mode/Menu
- `GPIO3` for IR receiver on supported 4-button firmware
- Fan speed stored from `0` to `10`
- Panel lockout support
- IR commands that can toggle, increment, or decrement fan speed

The important point is that none of this is ESPHome-specific. ESPHome makes it easy, but the underlying behavior is normal embedded firmware:

- Read GPIO
- Write PWM
- Read ADC
- Calculate temperature
- Update display
- Store state
- Expose controls to a smart-home system

That is exactly the kind of project HomeSpan can handle.

**Recommended Firmware Scope**
------------------------------

Do not try to recreate the full ESPHome firmware in the first pass.

The first HomeSpan firmware should be intentionally boring.

Version `0.1.0` should prove only four things:

1. The firmware builds in PlatformIO.
2. The ESP32 board boots reliably.
3. The accessory pairs with Apple Home.
4. Apple Home can set AirTap fan speed through PWM.

Once that works, build upward.

A realistic development order:

1. **Board bring-up**  
   Compile and flash a basic HomeSpan project on the AirTap ESP32 hardware.

2. **Apple Home pairing**  
   Expose one Fan accessory and confirm that Apple Home can pair, unpair, and re-pair it.

3. **PWM fan output**  
   Map Apple Home speed percentage to the existing AirTap 0-10 fan speed model.

4. **Local buttons**  
   Let the physical buttons change the fan speed and keep Apple Home state in sync.

5. **Temperature sensor**  
   Read the ADC, convert to temperature, and expose it as a HomeKit Temperature Sensor.

6. **OLED display**  
   Show fan speed, temperature, Wi-Fi state, and pairing/setup hints.

7. **Panel lockout**  
   Expose lockout as an Apple Home switch and block local button changes when enabled.

8. **IR receiver**  
   Add optional IR support for board versions that include the IR sensor.

9. **Persistence**  
   Store fan speed, lockout state, display preference, and calibration values.

10. **Release workflow**  
   Add versioning, flashing instructions, recovery instructions, and board-specific builds.

This order gives users something useful early without trapping the project in display menus or edge cases before the HomeKit core is proven.

**How the AirTap Should Appear in Apple Home**
---------------------------------------------

The simplest and strongest HomeKit model is:

- Accessory: `AirTap Vent`
- Service: `Fan`
- Service: `Temperature Sensor`
- Service: `Switch` named `Panel Lockout`
- Service: `Accessory Information`

The Fan service should expose:

- On/off state
- Speed percentage

Internally, the firmware should keep the existing `0-10` speed model. Apple Home should see `0-100%`.

A simple mapping works well:

| Apple Home Speed | Internal AirTap Speed |
| --- | --- |
| 0% | 0 |
| 1-10% | 1 |
| 11-20% | 2 |
| 21-30% | 3 |
| 31-40% | 4 |
| 41-50% | 5 |
| 51-60% | 6 |
| 61-70% | 7 |
| 71-80% | 8 |
| 81-90% | 9 |
| 91-100% | 10 |

This is easy to explain, easy to debug, and easy to display on the OLED.

Later, the firmware could support a smoother percentage-to-PWM curve. For the first release, predictable behavior matters more than theoretical smoothness.

**PlatformIO Project Layout**
-----------------------------

A clean project layout could look like this:

```text
airtap-homespan/
├── platformio.ini
├── include/
│   ├── AirtapPins.h
│   ├── AirtapState.h
│   ├── AirtapFan.h
│   ├── AirtapButtons.h
│   ├── AirtapTemperature.h
│   ├── AirtapDisplay.h
│   └── AirtapHomeKit.h
├── src/
│   ├── main.cpp
│   ├── AirtapFan.cpp
│   ├── AirtapButtons.cpp
│   ├── AirtapTemperature.cpp
│   ├── AirtapDisplay.cpp
│   └── AirtapHomeKit.cpp
└── README.md
```

For the first prototype, it is fine to keep most code in `main.cpp`.

Once the device pairs and controls the fan, split the code into modules.

The firmware should avoid becoming one giant file because this project will eventually need to support multiple board revisions.

**A Starting platformio.ini**
-----------------------------

HomeSpan targets Arduino on ESP32. Recent HomeSpan releases require a recent Arduino-ESP32 core, and PlatformIO’s default Espressif platform may lag behind the newest Arduino-ESP32 support. A practical option is using the pioarduino Espressif platform package when newer Arduino-ESP32 support is needed.

For the XIAO ESP32-C3 target, a starting `platformio.ini` could look like this:

```ini
[env:airtap_xiao_esp32c3]
platform = https://github.com/pioarduino/platform-espressif32/releases/download/stable/platform-espressif32.zip
board = seeed_xiao_esp32c3
framework = arduino
monitor_speed = 115200
upload_speed = 921600

lib_deps =
  HomeSpan/HomeSpan
  adafruit/Adafruit SSD1306
  adafruit/Adafruit GFX Library

board_build.partitions = min_spiffs.csv

build_flags =
  -D AIRTAP_BOARD_3BTN_REV2
```

A future ESP32-C6 environment could look like this once the exact PlatformIO board target is confirmed:

```ini
[env:airtap_xiao_esp32c6]
platform = https://github.com/pioarduino/platform-espressif32/releases/download/stable/platform-espressif32.zip
board = seeed_xiao_esp32c6
framework = arduino
monitor_speed = 115200
upload_speed = 921600

lib_deps =
  HomeSpan/HomeSpan
  adafruit/Adafruit SSD1306
  adafruit/Adafruit GFX Library

board_build.partitions = min_spiffs.csv

build_flags =
  -D AIRTAP_BOARD_4BTN_REV1
```

The C3 target should probably be the first development target because the current ESPHome configs clearly use `seeed_xiao_esp32c3`, and HomeSpan supports ESP32-C3-class devices.

**Pin Mapping**
---------------

Board pins should live in one place.

Do not hardcode GPIO numbers throughout the firmware.

Create an `AirtapPins.h`:

```cpp
#pragma once

#if defined(AIRTAP_BOARD_3BTN_REV2)

static constexpr int PIN_FAN_PWM = 2;
static constexpr int PIN_TEMP_ADC = 4;
static constexpr int PIN_OLED_SDA = 6;
static constexpr int PIN_OLED_SCL = 7;

static constexpr int PIN_BUTTON_UP = 8;
static constexpr int PIN_BUTTON_DOWN = 9;
static constexpr int PIN_BUTTON_MODE = 10;
static constexpr int PIN_BUTTON_TOGGLE = -1;

static constexpr int PIN_IR_RX = -1;

#elif defined(AIRTAP_BOARD_4BTN_REV1)

static constexpr int PIN_FAN_PWM = 2;
static constexpr int PIN_TEMP_ADC = 4;
static constexpr int PIN_OLED_SDA = 6;
static constexpr int PIN_OLED_SCL = 7;

static constexpr int PIN_BUTTON_UP = 20;
static constexpr int PIN_BUTTON_DOWN = 8;
static constexpr int PIN_BUTTON_TOGGLE = 9;
static constexpr int PIN_BUTTON_MODE = 10;

static constexpr int PIN_IR_RX = 3;

#else
#error "Select an AirTap board revision with a build flag"
#endif
```

That header gives the firmware one central source of truth.

It also makes the project easier to explain to contributors:

- If you have a 3-button board, build with `AIRTAP_BOARD_3BTN_REV2`.
- If you have a 4-button board, build with `AIRTAP_BOARD_4BTN_REV1`.
- If a future board revision changes pins, add one block instead of rewriting firmware.

**Global State**
----------------

Keep the device state small and explicit.

```cpp
#pragma once

struct AirtapState {
  int fanSpeedStep = 0;        // 0-10
  bool panelLocked = false;
  float temperatureC = NAN;
  bool homekitPaired = false;
  bool wifiConnected = false;
};

extern AirtapState airtap;
```

The main state variables are:

- `fanSpeedStep`
- `panelLocked`
- `temperatureC`
- Wi-Fi/HomeKit setup state if you want to display it

Do not let every subsystem own its own version of the fan speed. That is how state drift happens.

Apple Home, local buttons, IR commands, and display rendering should all point back to the same state.

**PWM Fan Control**
-------------------

The existing ESPHome firmware uses LEDC PWM on `GPIO2` at `1000 Hz`. It also uses a minimum duty cycle for speed `1` so the fan starts reliably.

That logic should be preserved.

A simple fan driver can look like this:

```cpp
#include <Arduino.h>
#include "AirtapPins.h"
#include "AirtapState.h"

static constexpr int FAN_PWM_CHANNEL = 0;
static constexpr int FAN_PWM_FREQ = 1000;
static constexpr int FAN_PWM_RES_BITS = 10;
static constexpr int FAN_PWM_MAX = (1 << FAN_PWM_RES_BITS) - 1;

float fanStepToDuty(int step) {
  if (step <= 0) {
    return 0.0f;
  }

  if (step == 1) {
    return 0.38f;   // minimum-start behavior based on current ESPHome firmware
  }

#if defined(AIRTAP_BOARD_4BTN_REV1)
  return (step + 4) / 14.0f;
#else
  return (step + 3) / 13.0f;
#endif
}

void setupFanPwm() {
  ledcSetup(FAN_PWM_CHANNEL, FAN_PWM_FREQ, FAN_PWM_RES_BITS);
  ledcAttachPin(PIN_FAN_PWM, FAN_PWM_CHANNEL);
}

void applyFanSpeed() {
  if (airtap.fanSpeedStep < 0) {
    airtap.fanSpeedStep = 0;
  }

  if (airtap.fanSpeedStep > 10) {
    airtap.fanSpeedStep = 10;
  }

  float duty = fanStepToDuty(airtap.fanSpeedStep);
  int rawDuty = (int)(duty * FAN_PWM_MAX);

  ledcWrite(FAN_PWM_CHANNEL, rawDuty);
}
```

This is the first point where the firmware becomes useful. If Apple Home changes the fan speed, `applyFanSpeed()` should immediately change the AirTap hardware.

**Mapping HomeKit Speed to AirTap Speed**
----------------------------------------

Apple Home users think in percentages. The AirTap firmware thinks in steps.

Use helpers:

```cpp
int percentToFanStep(int percent) {
  if (percent <= 0) {
    return 0;
  }

  int step = (percent + 9) / 10;

  if (step < 1) {
    step = 1;
  }

  if (step > 10) {
    step = 10;
  }

  return step;
}

int fanStepToPercent(int step) {
  if (step <= 0) {
    return 0;
  }

  if (step > 10) {
    step = 10;
  }

  return step * 10;
}
```

This keeps the UX simple:

- Apple Home says `50%`.
- Firmware stores `5`.
- OLED shows `Fan Speed: 5`.
- PWM output uses the same curve as the ESPHome firmware.

That is easy for users to understand and easy for support to troubleshoot.

**Minimal HomeSpan Fan Service**
--------------------------------

The first HomeSpan class should expose a Fan service.

This is intentionally close to firmware, not marketing copy.

```cpp
#include <HomeSpan.h>
#include "AirtapState.h"

void applyFanSpeed();
int percentToFanStep(int percent);
int fanStepToPercent(int step);

struct AirtapFanService : Service::Fan {
  SpanCharacteristic *active;
  SpanCharacteristic *rotationSpeed;

  AirtapFanService() : Service::Fan() {
    active = new Characteristic::Active(0);
    rotationSpeed = new Characteristic::RotationSpeed(0);
  }

  boolean update() override {
    if (active->updated()) {
      bool on = active->getNewVal();

      if (!on) {
        airtap.fanSpeedStep = 0;
      } else if (airtap.fanSpeedStep == 0) {
        airtap.fanSpeedStep = 5;
      }

      rotationSpeed->setVal(fanStepToPercent(airtap.fanSpeedStep));
      applyFanSpeed();
    }

    if (rotationSpeed->updated()) {
      int percent = rotationSpeed->getNewVal();

      airtap.fanSpeedStep = percentToFanStep(percent);
      active->setVal(airtap.fanSpeedStep > 0 ? 1 : 0);

      applyFanSpeed();
    }

    return true;
  }

  void syncFromDeviceState() {
    active->setVal(airtap.fanSpeedStep > 0 ? 1 : 0);
    rotationSpeed->setVal(fanStepToPercent(airtap.fanSpeedStep));
  }
};
```

The important detail is `syncFromDeviceState()`.

When a physical button changes the fan speed, Apple Home needs to know. The firmware should update the HomeKit characteristics so the Home app does not show stale state.

**Main Firmware Skeleton**
--------------------------

A first `main.cpp` could look like this:

```cpp
#include <Arduino.h>
#include <HomeSpan.h>

#include "AirtapPins.h"
#include "AirtapState.h"

AirtapState airtap;
AirtapFanService *fanService = nullptr;

void setupFanPwm();
void applyFanSpeed();
void setupButtons();
void pollButtons();
void setupTemperature();
void pollTemperature();
void setupDisplay();
void updateDisplayIfNeeded();

void setup() {
  Serial.begin(115200);
  delay(1000);

  setupFanPwm();
  setupButtons();
  setupTemperature();
  setupDisplay();

  homeSpan.setLogLevel(1);

  // Development only. Do not ship every device with the same static setup code.
  homeSpan.setPairingCode("46637726");
  homeSpan.setQRID("ATAP");

  homeSpan.begin(Category::Fans, "AirTap Vent");

  new SpanAccessory();
    new Service::AccessoryInformation();
      new Characteristic::Identify();
      new Characteristic::Manufacturer("SiloCityLabs");
      new Characteristic::Name("AirTap Vent");
      new Characteristic::Model("AirTap ESP32 HomeSpan");
      new Characteristic::FirmwareRevision("0.1.0");

    fanService = new AirtapFanService();

    new AirtapTemperatureService();
    new AirtapPanelLockoutService();
}

void loop() {
  homeSpan.poll();

  pollButtons();
  pollTemperature();
  updateDisplayIfNeeded();
}
```

That is enough structure to prove the accessory model.

Do not overbuild the first version. Get the fan pairing and PWM control working first.

**Local Button Handling**
-------------------------

The local buttons should still work.

That is a major product expectation. The AirTap should not become useless just because the Home app is closed or the Wi-Fi network is temporarily unavailable.

Button behavior should follow the existing ESPHome logic:

- Mode or Toggle turns the fan on/off
- Up increases fan speed
- Down decreases fan speed
- Buttons do nothing when panel lockout is enabled

A basic structure:

```cpp
void setFanStepFromLocalInput(int step) {
  if (airtap.panelLocked) {
    return;
  }

  if (step < 0) {
    step = 0;
  }

  if (step > 10) {
    step = 10;
  }

  airtap.fanSpeedStep = step;
  applyFanSpeed();

  if (fanService) {
    fanService->syncFromDeviceState();
  }
}

void increaseFanFromButton() {
  setFanStepFromLocalInput(airtap.fanSpeedStep + 1);
}

void decreaseFanFromButton() {
  setFanStepFromLocalInput(airtap.fanSpeedStep - 1);
}

void toggleFanFromButton() {
  if (airtap.fanSpeedStep == 0) {
    setFanStepFromLocalInput(10);
  } else {
    setFanStepFromLocalInput(0);
  }
}
```

For debouncing, keep it boring. Use a small helper class or a proven debounce library. Avoid mixing debounce logic into the HomeKit service.

Pseudocode:

```cpp
void pollButtons() {
  // Read debounced edge events.
  // If Up pressed: increaseFanFromButton()
  // If Down pressed: decreaseFanFromButton()
  // If Toggle/Mode pressed: toggleFanFromButton()
}
```

This is not glamorous firmware, but it is the difference between a nice prototype and a device people can actually live with.

**Panel Lockout**
-----------------

Panel lockout already exists in the ESPHome firmware and should stay.

The Apple Home mapping is simple: expose a Switch service named `Panel Lockout`.

```cpp
struct AirtapPanelLockoutService : Service::Switch {
  SpanCharacteristic *on;

  AirtapPanelLockoutService() : Service::Switch() {
    on = new Characteristic::On(false);
  }

  boolean update() override {
    airtap.panelLocked = on->getNewVal();
    return true;
  }
};
```

When enabled:

- Physical buttons should not change fan speed.
- IR remote commands should probably not change fan speed.
- Apple Home should still be allowed to control the fan.

That gives the lockout feature a clear meaning: block local accidental changes, not remote control.

**Temperature Sensor**
----------------------

The AirTap firmware reads temperature through the board’s ADC path. ESPHome currently models it with an NTC/resistance pipeline and calibration points around:

- `3.389 kOhm -> 0°C`
- `10.0 kOhm -> 25°C`
- `27.219 kOhm -> 50°C`

A HomeSpan version can expose that as a native HomeKit Temperature Sensor.

Start with a rough implementation and tune it later. The development goal is first to get a stable, believable temperature value into Apple Home.

```cpp
float readAirtapTemperatureC() {
  int raw = analogRead(PIN_TEMP_ADC);

  // Pseudocode:
  // 1. Convert ADC raw value to voltage ratio.
  // 2. Convert voltage ratio to thermistor resistance.
  // 3. Use the existing calibration curve or Steinhart-Hart approximation.
  // 4. Return temperature in Celsius.

  return NAN;
}
```

The HomeSpan service:

```cpp
struct AirtapTemperatureService : Service::TemperatureSensor {
  SpanCharacteristic *currentTemperature;
  unsigned long lastReadMs = 0;

  AirtapTemperatureService() : Service::TemperatureSensor() {
    currentTemperature = new Characteristic::CurrentTemperature(20.0);
  }

  void loop() override {
    if (millis() - lastReadMs < 5000) {
      return;
    }

    lastReadMs = millis();

    float t = readAirtapTemperatureC();

    if (!isnan(t)) {
      airtap.temperatureC = t;
      currentTemperature->setVal(t);
    }
  }
};
```

Once this is working, the Apple Home user gets a temperature tile they can use in scenes and automations.

That is where the product starts to feel like more than a fan controller.

**OLED Display**
----------------

The existing firmware uses an SSD1306 128x64 OLED over I2C at address `0x3C`.

HomeSpan does not manage that display. Arduino libraries can.

A basic display implementation could show:

- `AirTap HomeKit`
- Fan speed `0-10`
- Temperature
- Panel locked/unlocked
- Wi-Fi or pairing status

```cpp
#include <Wire.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>

Adafruit_SSD1306 display(128, 64, &Wire, -1);

void setupDisplay() {
  Wire.begin(PIN_OLED_SDA, PIN_OLED_SCL);

  if (!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) {
    Serial.println("SSD1306 display not found");
    return;
  }

  display.clearDisplay();
  display.setTextSize(1);
  display.setTextColor(SSD1306_WHITE);
  display.display();
}

void updateDisplay() {
  display.clearDisplay();
  display.setCursor(0, 0);

  display.println("AirTap HomeKit");
  display.printf("Fan Speed: %d\n", airtap.fanSpeedStep);

  if (!isnan(airtap.temperatureC)) {
    float f = airtap.temperatureC * 9.0f / 5.0f + 32.0f;
    display.printf("Temp: %.1f F\n", f);
  } else {
    display.println("Temp: --");
  }

  display.printf("Panel: %s\n", airtap.panelLocked ? "Locked" : "Enabled");

  display.display();
}
```

Do not let display polish block firmware progress.

The first version does not need icons, animation, menus, or fancy fonts. It needs to confirm that the device is alive and show the state users care about.

**IR Receiver Support**
-----------------------

The 4-button ESPHome firmware includes IR receiver support on `GPIO3` and uses decoded Pronto data for remote buttons such as power, fan, plus, minus, mode, and refresh.

A HomeSpan firmware can add this later.

Recommended approach:

1. Ship the first HomeSpan prototype without IR.
2. Add IR receive logging.
3. Confirm the same remote codes on actual hardware.
4. Map IR plus/minus/power to the same local input functions used by physical buttons.
5. Respect panel lockout.

The IR layer should not directly manipulate HomeKit characteristics. It should call the same internal functions as the physical buttons:

```cpp
// IR plus command
increaseFanFromButton();

// IR minus command
decreaseFanFromButton();

// IR power command
toggleFanFromButton();
```

That keeps all local input behavior consistent.

**Persistence**
---------------

The device should remember useful state across reboot:

- Last fan speed
- Panel lockout state
- Display preference if added later
- Temperature calibration offset if added later

HomeSpan already has storage behavior for HomeKit pairing. For device settings, use ESP32 Preferences/NVS.

Example direction:

```cpp
#include <Preferences.h>

Preferences prefs;

void loadSettings() {
  prefs.begin("airtap", true);
  airtap.fanSpeedStep = prefs.getInt("fan", 0);
  airtap.panelLocked = prefs.getBool("lock", false);
  prefs.end();
}

void saveSettings() {
  prefs.begin("airtap", false);
  prefs.putInt("fan", airtap.fanSpeedStep);
  prefs.putBool("lock", airtap.panelLocked);
  prefs.end();
}
```

Do not write to flash every second. Save only when state changes, and consider a short delay/debounce before writing.

**Provisioning and Reset UX**
-----------------------------

This is one of the most important product details.

Developers can tolerate serial monitors and command-line reset commands. Customers cannot.

A native Apple Home firmware needs a clear setup and recovery story:

- How does the user enter pairing mode?
- How does the user reset HomeKit pairing?
- How does the user reset Wi-Fi credentials?
- What does the display show when not paired?
- What button combination triggers a factory reset?
- How does support tell a user to recover a device?

A reasonable product flow:

1. On first boot, display `Pair AirTap HomeKit`.
2. Show the HomeKit setup code or direct users to the printed/setup label.
3. If already paired, show normal fan status.
4. Holding two buttons for 10 seconds resets pairing and Wi-Fi.
5. Display confirms `Reset Complete`.

The firmware should not depend on users remembering serial commands.

**OTA and Updates**
-------------------

The current ESPHome path has a strong update story through ESPHome and GitHub-hosted firmware. A HomeSpan firmware would need its own answer.

Possible release paths:

- Manual USB flashing through PlatformIO for developer builds
- Web-based ESP flashing for public beta builds
- Arduino OTA or custom OTA for advanced users
- Versioned binaries per board revision

The first developer article should not promise automatic updates unless that system exists.

A good phrasing for early firmware:

> Initial HomeSpan builds should be treated as developer firmware and flashed over USB. A polished release would need a documented update path before being recommended for general customers.

That is the right level of honesty.

**Testing Checklist**
---------------------

Before calling the firmware usable, test these behaviors:

- Device flashes successfully.
- Serial logs show boot and HomeSpan startup.
- Apple Home can pair the device.
- Apple Home can remove and re-pair the device.
- Fan turns off at `0%`.
- Fan starts reliably at the first non-zero speed.
- Fan reaches a strong output at `100%`.
- Physical Up increases the speed.
- Physical Down decreases the speed.
- Physical Toggle/Mode turns the fan on and off.
- Apple Home updates when physical buttons are pressed.
- OLED display matches actual state.
- Panel lockout blocks physical buttons.
- Panel lockout does not block Apple Home control.
- Temperature does not jump wildly.
- Reboot restores expected state.
- Removing power and restoring power does not corrupt pairing.
- Factory reset works without a computer.
- Multiple vents can pair with unique names.

For multiple AirTaps in one home, naming matters. A user may have:

- Office AirTap
- Bedroom AirTap
- Living Room AirTap
- Workshop AirTap

The firmware and documentation should encourage users to rename each accessory in Apple Home after pairing.

**SEO Opportunity**
-------------------

The Apple-focused market will not always search for the same terms as Home Assistant users.

Home Assistant users search for:

- ESPHome AirTap
- AC Infinity AirTap ESP32
- Home Assistant smart vent
- ESP32 vent controller
- ESPHome smart vent

Apple users are more likely to search for:

- Apple Home smart vent
- HomeKit smart vent
- AC Infinity AirTap HomeKit
- AirTap Apple Home
- Siri smart vent
- ESP32 HomeKit fan controller
- HomeKit HVAC vent
- Apple HomeKit AC Infinity

That is why this topic matters.

The article should make it clear that the AirTap ESP32 upgrade kit is already valuable for Home Assistant users, but the same hardware could also become a native Apple Home accessory.

Suggested SEO phrases to include naturally:

- Apple Home smart vent
- HomeKit smart vent controller
- AC Infinity AirTap HomeKit
- AirTap T4 Apple Home
- AirTap T6 Apple Home
- ESP32 HomeKit fan controller
- HomeSpan ESP32 firmware
- PlatformIO ESP32 HomeKit
- Siri-controlled smart vent
- HomeKit temperature sensor ESP32

Avoid keyword stuffing. The article should read like a serious developer guide, not a landing page.

**Suggested Product Positioning**
---------------------------------

A clean positioning statement:

> The SiloCityLabs AirTap ESP32 upgrade kit is already a strong option for Home Assistant users through ESPHome. HomeSpan opens the possibility of a second firmware path for Apple Home users who want direct HomeKit pairing, Siri control, and native Apple Home automations.

That positioning does three things:

1. It protects the existing ESPHome value.
2. It introduces Apple Home as a new use case.
3. It does not overpromise a finished firmware.

A good call to action:

> If you are a developer interested in Apple Home, HomeKit, ESP32, or smart HVAC projects, the AirTap ESP32 hardware is a strong starting point. The foundation is already there: fan PWM, temperature sensing, display, buttons, and a proven mechanical install. The next step is firmware.

**What a First Public Developer Release Could Include**
------------------------------------------------------

A useful GitHub release would not need to be perfect.

It should include:

- `platformio.ini`
- One tested board target
- `main.cpp`
- Pin mapping header
- Fan service
- Panel lockout service
- Temperature service if ready
- OLED display if ready
- README with flashing steps
- Known limitations
- Pairing/reset instructions
- Board revision notes

The README should be direct:

```text
This is experimental HomeSpan firmware for the SiloCityLabs AirTap ESP32 upgrade kit.

Current status:
- Apple Home pairing: working
- Fan speed control: working
- Local buttons: working
- Temperature sensor: experimental
- OLED display: experimental
- IR remote: not implemented
- OTA updates: not implemented

Use ESPHome firmware for production Home Assistant installs.
Use this firmware if you want to test native Apple Home support.
```

That kind of clarity attracts the right contributors and filters out users who expect a finished product.

**Possible Repository Names**
-----------------------------

Good repo names:

- `airtap-homespan`
- `esp32-airtap-homespan`
- `airtap-homekit-firmware`
- `ac-infinity-airtap-homekit`

Best practical choice:

```text
esp32-airtap-homespan
```

It is descriptive, searchable, and fits the existing ESP32/AirTap naming pattern.

**Article Summary**
-------------------

Native Apple Home support for the AirTap ESP32 upgrade kit is realistic.

The existing ESPHome firmware already proves the important hardware path: PWM fan output, temperature sensing, display, local buttons, panel lockout, and optional IR support. HomeSpan gives developers a way to expose that same hardware directly to Apple Home as a native HomeKit accessory.

The best first firmware should be small:

- Pair with Apple Home
- Expose a Fan service
- Control the AirTap fan through PWM
- Keep local buttons working
- Add temperature, display, lockout, persistence, and IR after the core works

This would not replace ESPHome. It would expand the market.

Home Assistant users already have a strong path. Apple Home users should have one too.

For developers interested in ESP32, HomeKit, Apple Home, and smart HVAC control, the AirTap ESP32 upgrade kit is a compelling platform to build on.

**Reference Links**
-------------------

- [SiloCityLabs ESP32 Module for AC Infinity AirTap T4/T6](https://silocitylabs.com/post/2025/esp32-airtap-esphome/)
- [SiloCityLabs AirTap T4 tag page](https://silocitylabs.com/tags/airtap-t4)
- [HomeSpan GitHub repository](https://github.com/HomeSpan/HomeSpan)
- [HomeSpan Arduino library listing](https://www.arduinolibraries.info/libraries/home-span)
- [Espressif Arduino-ESP32 library support](https://docs.espressif.com/projects/arduino-esp32/en/latest/libraries.html)

**Draft Publishing Notes**
--------------------------

This post is intentionally written as a developer guide and market-expansion article. Before publishing, consider adding:

- A product photo of the AirTap ESP32 module
- A screenshot of Apple Home with a Fan tile mockup
- A simple firmware architecture diagram
- A GitHub repo link if/when the HomeSpan prototype exists
- A short disclaimer that this is experimental firmware direction, not the default shipped firmware
- Internal links to the ESPHome-certified AirTap post and Gen3 upgrade kit article
