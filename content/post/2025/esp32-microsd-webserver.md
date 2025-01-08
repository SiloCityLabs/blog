---
title: Hosting a Web Server on the ESP32 Using MicroSD Card Files
subtitle: Serve Files with the ESP32MiniSD
author: Luis Rodriguez
type: post
draft: true
date: 2025-01-08
categories:
  - Projects
tags:
  - ESP32
  - Open Source
  - Web Server
  - DIY Electronics
  - Smart Devices
  - IoT Projects
  - Local Control
  - Cloud-Free Technology
  - xiao esp32-c3
  - xiao esp32-c6
  - esp32minisd
  - microsd card
  - esp32 microsd
  - web server

---

The **ESP32MiniSD** allows you to serve static files, like HTML, CSS, or JavaScript, over a web server hosted by your ESP32. This is perfect for creating standalone IoT dashboards, device control panels, or file servers. In this tutorial, we’ll show you how to use the ESP32MiniSD to host a web server and serve files directly from a microSD card.

{{< image src="/uploads/2025/" alt="ESP32-C6 with ESP32MiniSD serving files">}}

---

## Features of ESP32MiniSD for Web Hosting

- **Efficient File Serving**: Serve HTML, CSS, and JavaScript files from a microSD card.
- **Compact Form Factor**: Ideal for IoT projects with limited space.
- **Battery Support**: Deploy your web server on the go with LiPo battery support.

{{< image src="/uploads/2025/" alt="Diagram showing ESP32MiniSD components for web hosting">}}

---

## What You’ll Need

1. **ESP32MiniSD** board  
2. **XIAO ESP32-C3** or **XIAO ESP32-C6**  
3. **microSD card** (formatted as FAT32)  
4. **USB-C cable**  
5. Files to serve (HTML, CSS, JavaScript)  

{{< image src="/uploads/2025/" alt="Components for ESP32 web server project">}}

---

## Step 1: Hardware Setup

### Assembling the ESP32MiniSD

1. Solder the **XIAO ESP32-C3** or **XIAO ESP32-C6** to the **ESP32MiniSD** board.  
2. Insert your microSD card with preloaded files into the slot on the ESP32MiniSD.  
3. Connect your ESP32 to your computer using a USB-C cable.

{{< image src="/uploads/2025/" alt="Assembled ESP32MiniSD ready for web server setup">}}

---

## Step 2: Prepare the MicroSD Card

1. Format your microSD card as FAT32.  
2. Create a directory structure for your web server files:  

`/index.html /style.css /script.js`

3. Copy your files to the root directory of the microSD card. For testing, you can use this simple `index.html`:

```html
<!DOCTYPE html>
<html>
<head>
<title>ESP32 Web Server</title>
<link rel="stylesheet" href="style.css">
</head>
<body>
<h1>Welcome to the ESP32 Web Server</h1>
<p>This is served from a microSD card!</p>
<script src="script.js"></script>
</body>
</html>
```

Save the file and insert the card into your ESP32MiniSD.

Step 3: Arduino IDE Configuration
---------------------------------

Follow the same Arduino IDE setup steps as in the "Getting Started with ESP32MiniSD" tutorial to prepare your environment.

Step 4: Example Code for Web Server
-----------------------------------

Here’s the code to host a web server and serve files from the microSD card:

```cpp
#include <WiFi.h>
#include <WebServer.h>
#include <SD.h>
#include <SPI.h>

const char* ssid = "Your_SSID";
const char* password = "Your_PASSWORD";
const int chipSelect = 4;

WebServer server(80);

void handleFileRequest() {
  String path = server.uri();
  if (path == "/") path = "/index.html";

  File file = SD.open(path);
  if (!file) {
    server.send(404, "text/plain", "File Not Found");
    return;
  }

  String contentType = getContentType(path);
  server.streamFile(file, contentType);
  file.close();
}

String getContentType(String filename) {
  if (filename.endsWith(".html")) return "text/html";
  if (filename.endsWith(".css")) return "text/css";
  if (filename.endsWith(".js")) return "application/javascript";
  return "text/plain";
}

void setup() {
  Serial.begin(115200);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.println("Connecting to WiFi...");
  }
  Serial.println("Connected to WiFi!");

  if (!SD.begin(chipSelect)) {
    Serial.println("Card Mount Failed");
    return;
  }

  server.onNotFound(handleFileRequest);
  server.begin();
  Serial.println("Web server started!");
}

void loop() {
  server.handleClient();
}
```

Step 5: Testing the Web Server
------------------------------

1.  Upload the code to your **XIAO ESP32-C3** or **XIAO ESP32-C6**.
    
2.  Open the Serial Monitor to find the IP address assigned to your ESP32.
    
3.  Open a web browser and enter the IP address (e.g., http://192.168.1.100).
    
4.  You should see the index.html file served by the ESP32.
    

{{< image src="/uploads/2025/" alt="ESP32 web server test page in a browser">}}

Step 6: Applications and Enhancements
-------------------------------------

*   **IoT Dashboards**: Host device control panels directly on the ESP32.
    
*   **Standalone Servers**: Create portable servers with battery support.
    
*   **Local File Access**: Serve files for offline access in remote locations.
    

You can enhance this setup by adding POST/GET request handling or integrating APIs for dynamic web content. You can additionally use static site generators like Hugo or Jekyll to create complex web pages for your ESP32.

{{< image src="/uploads/2025/" alt="ESP32MiniSD serving an IoT dashboard">}}

With the **ESP32MiniSD**, hosting a web server directly from your ESP32 is simple and efficient. This setup is perfect for IoT devices that need local control or cloud-free operation. For more details, visit the [ESP32MiniSD GitHub repository](https://github.com/SiloCityLabs/esp32-minisd) or check out [SiloCityLabs](https://shop.silocitylabs.com/) for more products.