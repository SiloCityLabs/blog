---
title: Getting Started with the ESP32MiniSD
subtitle: Add Storage to Your ESP32 Projects
author: Luis Rodriguez
type: post
draft: true
date: 2025-01-08
categories:
  - Projects
tags:
  - ESP32
  - Open Source
  - DIY Electronics
  - Smart Devices
  - IoT Projects
  - Local Control
  - Cloud-Free Technology
  - xiao esp32-c3
  - xiao esp32-c6
  - esp32minisd
  - microsd card
  - micro sd card
  - esp32 microsd

---

The **ESP32MiniSD** is a compact add-on board designed for Seeed Studio's **XIAO ESP32-C3** or **XIAO ESP32-C6**, enabling microSD card storage, battery monitoring, and USB charging capabilities. In this tutorial, we'll guide you through setting up and using the ESP32MiniSD with an ESP32 XIAO module to read and write files to a microSD card.

<!--more-->

{{< image src="/uploads/2025/" alt="ESP32-C6 with ESP32MiniSD HAT">}}


**Features of ESP32MiniSD**
---------------------------

*   **MicroSD Card Support**: Adds storage to your ESP32 projects for logging or file storage.
    
*   **Battery Support**: Connect a LiPo battery with onboard charging via USB.
    
*   **Compact Design**: Slim and minimal, perfect for small form-factor projects.

{{< image src="/uploads/2025/" alt="">}}
    

**What You’ll Need**
--------------------

1.  **ESP32MiniSD** board
    
2.  **XIAO ESP32-C3** or **XIAO ESP32-C6**
    
3.  **microSD card** (formatted as FAT32)
    
4.  **USB-C cable**
    
5.  Optional: LiPo battery for portable operation
    

**Step 1: Hardware Setup**
--------------------------

### **Assembling the ESP32MiniSD**

1.  Solder the **XIAO ESP32-C3** or **XIAO ESP32-C6** to the **ESP32MiniSD** board.
    
2.  Insert your microSD card into the card slot on the ESP32MiniSD.
    
3.  Optionally, connect a LiPo battery to the designated pads on the board for portable use.

{{< image src="/uploads/2025/" alt="">}}
    

### **Wiring**

The ESP32MiniSD uses specific pins for the SPI interface:

*   **CS (Chip Select)**: D2
    
*   **MOSI (Master Out Slave In)**: D3
    
*   **MISO (Master In Slave Out)**: D5
    
*   **SCLK (Clock)**: D4
    

For battery voltage monitoring:

*   **METER\_EN**: D1
    
*   **METER**: D0 (Analog input)

Additional pin information available on the [github pin configuration section](https://github.com/SiloCityLabs/esp32-minisd/tree/main?tab=readme-ov-file#pin-configuration)
    

**Step 2: Software Setup**
--------------------------

### **Arduino IDE Configuration**

 1. Download and Install the latest version of Arduino IDE according to your operating system

 2. Launch the Arduino application

 3. Add ESP32 board package to your Arduino IDE:

 4. Navigate to File > Preferences, and fill "Additional Boards Manager URLs" with the url below: `https://raw.githubusercontent.com/espressif/arduino-esp32/gh-pages/package_esp32_index.json`

  {{< image src="/uploads/2025/add_board.png" alt="Adding additional board manager URL">}}

 5. Navigate to Tools > Board > Boards Manager..., type the keyword `esp32` in the search box, select the latest version of esp32, and install it.

 {{< image src="/uploads/2025/add_esp32c3.png" alt="Adding ESP32 Board in boards manager">}}
        

**Step 3: Example Code**
------------------------

Here’s a simple program to initialize the microSD card and list its files:


```c++   
 #include <SD.h>
#include <SPI.h>

// Define chip select pin for SD card
const int chipSelect = 4;

void setup() {
  Serial.begin(115200);
  delay(1000); // Allow time for the Serial Monitor to initialize

  Serial.println("Initializing SD card...");

  if (!SD.begin(chipSelect)) {
    Serial.println("Card mount failed");
    return;
  }

  Serial.println("Card initialized successfully.");
  listFiles("/");
}

void loop() {
  // Keep the loop empty for this example
}

void listFiles(const char *dirname) {
  File root = SD.open(dirname);
  if (!root) {
    Serial.println("Failed to open directory");
    return;
  }

  while (true) {
    File entry = root.openNextFile();
    if (!entry) {
      break; // No more files
    }
    Serial.print(entry.name());
    if (entry.isDirectory()) {
      Serial.println("/");
      listFiles(entry.name());
    } else {
      Serial.print("\t");
      Serial.println(entry.size());
    }
    entry.close();
  }
}
```

### **What This Code Does**

1.  Initializes the SD card.
    
2.  Lists all files and directories in the root folder.
    

Upload the code to your **XIAO ESP32-C3** or **XIAO ESP32-C6** using the **Arduino IDE**, then open the Serial Monitor to see the output.

**Step 4: Testing and Troubleshooting**
---------------------------------------

1.  **Testing**:
    
    *   After uploading the code, open the Serial Monitor at 115200 baud.
        
    *   If the SD card initializes successfully, you’ll see a list of files on the card.
        
2.  **Common Issues**:
    
    *   **Card Mount Failed**: Ensure your SD card is formatted as FAT32.
        
    *   **No Files Detected**: Double-check the configuration and card insertion.
        

**Applications of the ESP32MiniSD**
-----------------------------------

With leftover pins for additional sensors or peripherals (`D7,D8,D9`), the ESP32MiniSD is ideal for a variety of applications:

* **Data Logging**: Perfect for recording sensor data like temperature, humidity, or GPS coordinates.
    
* **File Storage**: Extend the memory of your ESP32 for larger programs or file operations.
    
* **Portable Devices**: Use the LiPo battery support for standalone IoT devices.

* **Battery Monitoring**: Keep track of your battery levels for efficient power management.

* **USB Charging**: Connect the battery pins to a battery holder and charge via USB and monitor the battery levels.
    

The **ESP32MiniSD** brings advanced storage and power capabilities to your ESP32 projects. Whether you’re logging data, creating portable devices, or managing file systems, this versatile board simplifies the process. Pair it with the **XIAO ESP32-C3/C6**, and you’re ready to tackle complex, data-heavy applications! You can get both the [ESP32MiniSD](https://shop.silocitylabs.com/products/esp32minisd) and [XIAO ESP32-C3/C6](https://shop.silocitylabs.com/products/esp32minisd?variant=50411168366892) from SiloCityLabs.

For more information, visit the [official GitHub repository](https://github.com/SiloCityLabs/esp32-minisd).