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

But there is another smart-home audience that could benefit from this hardware: **Apple Home users**.

If you're in the Apple ecosystem, using HomePod, Siri on iPhone, Apple TV, scenes, automations, and the Apple Home app, the natural question is: _Could the AirTap ESP32 upgrade kit run native Apple Home firmware instead of ESPHome?_

Yes it can! [HomeSpan](https://github.com/HomeSpan/HomeSpan) provides developers with libraries to create an ESP32 firmware that pairs the AirTap directly with Apple Home over Wi-Fi. That's native Apple support without Home Assistant, without a Home Assistant HomeKit Bridge, and without a separate server translating ESPHome entities into HomeKit accessories.

This post is not a finished firmware release. It is a technical roadmap for developers, makers, and advanced AirTap owners who want to understand what it would take to build one. It's enough architecture to start the project, enough code structure to see the direction, and enough caution to avoid the obvious traps.

<!--more-->

**Why HomeSpan?**
----------------------------------

Most custom ESP32 smart-home firmware assumes that Home Assistant is the center of the house. That is fine for power users like us - we each have a home server running Home Assistant - and ESPHome is still one of the best ways to ship an ESP32-based IoT device quickly.

But not every smart-home user wants to manage Home Assistant. Apple users want the "house of the future" too: smart vents, quieter airflow, room-level automation, and better HVAC behavior, but they want the control surface to be the Apple Home app. They want to say:

- “Hey Siri, turn on the office vent.”
- “Set the bedroom vent to 40 percent.”
- “Turn off the vent display buttons.”
- “When the room gets warm, increase airflow.”
- “When the Good Night scene runs, lower the vent speed.”

That is a different market from the existing ESPHome/Home Assistant crowd. For SiloCityLabs, this is an opportunity: keep serving the Home Assistant community while opening the door to Apple-first smart-home users who may never search for ESPHome, YAML, or Home Assistant integrations.

**The existing AirTap ESP32 hardware is a good fit**
---------------------------------------------------

Fortunately the AirTap ESP32 upgrade kit already has the right kind of hardware for native Apple Home firmware. The firmware needs to handle:

- ESP32 module running Wi-Fi firmware
- PWM output for AirTap fan speed control
- ADC input for temperature sensing
- SSD1306 128x64 OLED display over I2C
- Local GPIO buttons
- Stored fan speed state
- Panel lockout state
- Optional IR receiver support on supported board versions
- Wi-Fi provisioning (SSID and password setup)
- OTA/update strategy
- A clean end-user pairing/reset story

The current ESPHome firmware proves that the board can already do the important work. HomeSpan just replaces the smart-home protocol layer. HomeSpan will expose HomeKit services directly to Apple Home.

That distinction matters for SEO, positioning, and buyer intent. “ESP32 smart vent for Home Assistant” reaches one audience. “Apple HomeKit smart vent upgrade for AC Infinity AirTap” reaches another.

**What HomeSpan brings to the project**
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

**What we can reuse from the ESPHome firmware**
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

**Recommended firmware scope**
------------------------------

The first HomeSpan firmware should be intentionally boring. Version `0.1.0` should prove only four things:

1. The firmware builds in PlatformIO.
2. The ESP32 board boots reliably.
3. The accessory pairs with Apple Home.
4. Apple Home can set AirTap fan speed through PWM.

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

This order delivers value early without trapping the project in menus before HomeKit core is proven.

**How the AirTap Should Appear in Apple Home**
---------------------------------------------

Use this HomeKit model:

- Accessory: `AirTap Vent`
- Service: `Fan`
- Service: `Temperature Sensor`
- Service: `Switch` named `Panel Lockout`
- Service: `Accessory Information`

The Fan service exposes on/off state and speed percentage. Internally keep the `0-10` speed model; Apple Home sees `0-100%`. Use this mapping:

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

Predictable behavior matters more than smoothness for the first release.

**PlatformIO Project Layout**
-----------------------------

Start with a simple layout:

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

Keep code in `main.cpp` initially, then split into modules as it grows. Modular design matters for supporting multiple board revisions.

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

Start with the C3 variant; the existing ESPHome configs use `seeed_xiao_esp32c3`.

**Pin Mapping**
---------------

Keep pins in one header file, never hardcoded. Create `AirtapPins.h`:

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

This gives contributors a clear reference:

- If you have a 3-button board, build with `AIRTAP_BOARD_3BTN_REV2`.
- If you have a 4-button board, build with `AIRTAP_BOARD_4BTN_REV1`.
- If a future board revision changes pins, add one block instead of rewriting firmware.

**Global State**
----------------

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

Single source of truth matters: Apple Home, buttons, IR, and display all read from `airtap`.

**PWM Fan Control**
-------------------

The existing ESPHome firmware uses LEDC PWM on `GPIO2` at `1000 Hz`. It also uses a minimum duty cycle for speed `1` so the fan starts reliably. If the fans stall it can burn out the motors. I'm even considering starting the fan at max power for 1 second, similar to box fans and rooms fans which have the power settings go from 0-3-2-1.

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

**Mapping HomeKit Speed to AirTap Speed**
----------------------------------------

Apple Home uses percentages while AirTap uses steps. Use converters:

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

This keeps behavior predictable and traceable.

**Minimal HomeSpan Fan Service**
--------------------------------

Expose a [Fan service in HomeSpan](https://github.com/HomeSpan/HomeSpan/blob/master/docs/ServiceList.md#fan-b7):

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

`syncFromDeviceState()` is critical: when physical buttons change the fan, Apple Home stays in sync.

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

This proves the accessory model. Get pairing and PWM working first.

**Local Button Handling**
-------------------------

Local buttons are non-negotiable. The device should work offline. Follow existing ESPHome logic:

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

I'm skipping over debounce code. There are a million debounce tuts in addition tothe proven "setDebounceTime" in Arduino.

**Panel Lockout**
-----------------

Panel lockout is a an existing feautre in my ESPHome firmware. Expose a Switch service for lockout:

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

This prevents your toddler from changing the settings.

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

The Apple Home user gets a temperature tile they can use in scenes and automations.

**OLED Display**
----------------

The existing firmware uses an SSD1306 128x64 OLED over I2C at address `0x3C`. HomeSpan does not manage that display. Arduino libraries can display the status:

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

Since this in a "MVP", plain text is fine. Show state first, polish later.

**IR Receiver Support**
-----------------------

The 4-button ESPHome firmware includes IR receiver support on `GPIO3` and uses decoded Pronto data for remote buttons such as power, fan, plus, minus, mode, and refresh. This isn't vital for the MVP and can be added later. Recommended approach:

1. Ship first prototype without IR.
2. Log IR codes to confirm remote behavior.
3. Map IR commands to existing button functions.
4. Respect lockout.

IR should call the same input handlers as buttons, so that changes propagate to HomeKit:

```cpp
// IR plus/minus/power
increaseFanFromButton();
decreaseFanFromButton();
toggleFanFromButton();
```

**Persistence**
---------------

The device should remember useful state across reboot:

- Last fan speed
- Panel lockout state
- Display preference
- Temperature calibration / offset

HomeSpan already has storage behavior for HomeKit pairing. For device settings, use ESP32 Preferences/NVS.

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

Avoid frequent writes; save only on state change, and consider a short delay/debounce before writing.

**Provisioning and Reset UX**
-----------------------------

Customers need a clear setup procedure. Not serial monitors nor "magic phrases" in command-line. Consider common steps like:

- entering pairing mode
- setting or resetting wifi credentials
- what does the display show when not paired?
- what button combination triggers a factory reset?
- how does support tell a user to recover a device?


A reasonable flow:

1. On first boot, display `Pair AirTap HomeKit`.
2. Show the HomeKit setup code or direct users to the printed/setup label.
3. If already paired, show normal fan status.
4. Holding two buttons for 10 seconds resets pairing and Wi-Fi.
5. Display confirms `Reset Complete`.


**OTA and Updates**
-------------------

Use manual USB flashing for now. Automatic updates require infrastructure that can wait.

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
- Multiple vents pair with unique names and can be renamed in Apple Home.

**Market Positioning**
-------------------

Home Assistant users search for ESPHome + AirTap. Apple users search for HomeKit + AirTap. This post opens a second market: developers and Apple-first smart-home users.

**Next Steps for Developers**
-----------------------------

If you're interested in Apple Home, HomeKit, ESP32, or smart HVAC, the AirTap ESP32 hardware is a strong foundation. It provides fan PWM, temperature sensing, display, buttons, and proven mechanics. The next step is firmware.

**What a First Release Could Include**
-------------------------------------

Ship what works:

- `platformio.ini` for one board target
- `main.cpp` with PWM fan control
- Pin mapping, fan service, lockout service
- README with flashing instructions
- Pairing/reset instructions (for users who receive a completed device)
- Known limitations clearly stated

Be direct about status:

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

And finally for the repo name: `esp32-airtap-homespan`- descriptive and searchable.

**Summary**
-------------------

The AirTap ESP32 hardware is proven for Home Assistant via ESPHome. HomeSpan offers a parallel path for native Apple Home support.

The best first firmware can be small:

- Pair with Apple Home
- Expose a Fan service
- Control the AirTap fan through PWM
- Keep local buttons working
- Add temperature, display, lockout, persistence, and IR after the core works

This expands the market rather than replacing ESPHome. Developers building for Apple Home now have an entry point.

**Reference Links**
-------------------

- [SiloCityLabs ESP32 Module for AC Infinity AirTap T4/T6](https://silocitylabs.com/post/2025/esp32-airtap-esphome/)
- [SiloCityLabs AirTap T4 tag page](https://silocitylabs.com/tags/airtap-t4)
- [HomeSpan GitHub repository](https://github.com/HomeSpan/HomeSpan)
- [HomeSpan Arduino library listing](https://www.arduinolibraries.info/libraries/home-span)
- [Espressif Arduino-ESP32 library support](https://docs.espressif.com/projects/arduino-esp32/en/latest/libraries.html)
