---
title: Setting Up an FTP Server with ESP32MiniSD
subtitle: Turn Your ESP32 into a MicroSD-Powered NAS
author: Luis Rodriguez
type: post
draft: true
date: 2025-01-08
categories:
  - Projects
tags:
  - ESP32
  - Open Source
  - FTP Server
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

---

The **ESP32MiniSD** can easily be transformed into a compact FTP server using its onboard MicroSD storage. With this setup, you can create a local network file server, upload or download files via FTP, and even monitor the device's status through a web interface. This tutorial will walk you through the steps to set up an FTP server using your ESP32MiniSD.

**Features of ESP32 FTP Server**

*   MicroSD File Access: Host files stored on the MicroSD card.
    
*   Web Interface: View battery voltage, toggle GPIO states, and monitor the SD card status.
    
*   Battery Monitoring: Check battery levels and charging status.
    
*   Standalone Mode: Run as an access point or connect to an existing Wi-Fi network.
    

**What You’ll Need**

1.  ESP32MiniSD board
    
2.  XIAO ESP32-C3 or XIAO ESP32-C6
    
3.  microSD card (formatted as FAT32)
    
4.  USB-C cable
    
5.  A PC or mobile device with FTP client software
    

**Step 1: Hardware Setup**

1.  Solder the XIAO ESP32-C3 or ESP32-C6 to the ESP32MiniSD board.
    
2.  Insert your MicroSD card into the card slot.
    
3.  Connect the ESP32MiniSD to your computer via USB-C.
    

**Step 2: Preparing the MicroSD Card**

1.  Format your microSD card as FAT32.
    
2.  Create a file named config.txt at the root of the card with the following content:

```cpp
MODE=CLIENT
HOMEWIFI=YourSSID
HOMEPASS=YourPassword
FTPUSER=ftpuser
FTPPASS=ftppass
```
    
3. Save the file and insert the card into the ESP32MiniSD.

**Step 3: Installing Dependencies in Arduino IDE**

1.  Follow the setup instructions in the "Getting Started with ESP32MiniSD" tutorial.
    
2.  Install the **SimpleFTPServer** library:
    
    *   Open Arduino IDE and go to **Sketch > Include Library > Manage Libraries**.
        
    *   Search for SimpleFTPServer and install it.
        

**Step 4: Example Code**

Below is the full code for setting up the FTP server and web interface. Latest code on [GitHub](https://github.com/SiloCityLabs/esp32-minisd/blob/main/ftp_server/ftp_server.ino).

```cpp
#include <WiFi.h>
#include <SD.h>
#include <WebServer.h>

#include <SimpleFTPServer.h>


// Uncomment or define the specific board in the IDE or here
// #define BOARD_ESP32C3
#define BOARD_ESP32C6

// Board-specific configurations
#ifdef BOARD_ESP32C3
  #define SD_CS 4
  #define SD_MOSI 5
  #define SD_SCLK 6
  #define SD_MISO 7
  #define SD_EN 21 //Input - Is SD card inserted or not
  #define CHARGE_STATUS_PIN 10 //Input - Whether device is charging or not
  #define BATTERY_METER_PIN 2 //Input - Voltage divider output with 2x 220kohm
  #define BATTERY_METER_EN_PIN 3 //Output - Enable or disable voltage divider
#elif defined(BOARD_ESP32C6)
  #define SD_CS 2
  #define SD_MOSI 21
  #define SD_SCLK 22
  #define SD_MISO 23
  #define SD_EN 16 //Input - Is SD card inserted or not
  #define CHARGE_STATUS_PIN 18 //Input - Whether device is charging or not
  #define BATTERY_METER_PIN 0 //Input - Voltage divider output with 2x 220kohm
  #define BATTERY_METER_EN_PIN 1 //Output - Enable or disable voltage divider
  #define USER_LED 15
#else
    #error "Please define a supported board."
#endif

FtpServer ftpSrv;
WebServer server(80);

char wifi_mode[8] = "CLIENT"; // CLIENT, AP
char ssid[32] = "WIFINAME";
char password[32] = "PASSWORD";
char ftp_user[32] = "user";
char ftp_pass[32] = "password";

float readBatteryVoltage() {
    uint32_t Vbatt = 0;

    // Take 16 readings for averaging
    for (int i = 0; i < 16; i++) {
        Vbatt += analogReadMilliVolts(BATTERY_METER_PIN); // Use corrected ADC readings
    }

    // Calculate the averaged voltage, accounting for attenuation ratio (1/2)
    float Vbattf = 2.0 * (Vbatt / 16.0) / 1000.0; // Convert mV to V

    return Vbattf; // Return the voltage in volts
}


void _callback(FtpOperation ftpOperation, unsigned int freeSpace, unsigned int totalSpace){
	Serial.print(">>>>>>>>>>>>>>> _callback " );
	Serial.print(ftpOperation);
	Serial.print(" ");
	Serial.print(freeSpace);
	Serial.print(" ");
	Serial.println(totalSpace);
	if (ftpOperation == FTP_CONNECT) Serial.println(F("CONNECTED"));
	if (ftpOperation == FTP_DISCONNECT) Serial.println(F("DISCONNECTED"));
};


void _transferCallback(FtpTransferOperation ftpOperation, const char* name, unsigned int transferredSize) {
    Serial.print("Transfer Operation: ");
    Serial.println(ftpOperation);
    Serial.print("File Name: ");
    Serial.println(name ? name : "NULL");
    Serial.print("Transferred Size: ");
    Serial.println(transferredSize);
}

void handleRoot() {
    float voltage = readBatteryVoltage();
    float percentage = (voltage - 3.0) / (4.2 - 3.0) * 100.0;
    percentage = constrain(percentage, 0, 100);

    String html = "<html><body>";
    html += "<h1>Battery Percentage: " + String(percentage) + "%</h1>";
    html += "<p>Battery Voltage: " + String(voltage, 2) + " V</p>";
    html += "<p>SD Card Inserted: " + String(digitalRead(SD_EN)) + "</p>";
    html += "<p>Voltage Divider Active: " + String(digitalRead(BATTERY_METER_EN_PIN)) + "</p>";
    html += "<form method='POST' action='/toggle'><button type='submit'>Toggle Divider</button></form>";
    html += "<form method='POST' action='/led'><button type='submit'>Toggle LED</button></form>";
    html += "</body></html>";

    server.send(200, "text/html", html);
}

void handleToggle() {
    digitalWrite(BATTERY_METER_EN_PIN, !digitalRead(BATTERY_METER_EN_PIN));
    server.sendHeader("Location", "/");
    server.send(303); // Redirect to root
}

void handleLed() {
    digitalWrite(USER_LED, !digitalRead(USER_LED));
    server.sendHeader("Location", "/");
    server.send(303); // Redirect to root
}

void setup() {
  Serial.begin(115200);

  // Setup GPIOs
  pinMode(SD_EN, INPUT);
  pinMode(BATTERY_METER_PIN, INPUT);
  pinMode(CHARGE_STATUS_PIN, INPUT);
  pinMode(BATTERY_METER_EN_PIN, OUTPUT);
  pinMode(USER_LED, OUTPUT);
  digitalWrite(USER_LED,1);

  //Is the SD card inserted
  if (0 == digitalRead(SD_EN)){
    Serial.println("SD Card Not inserted!!!");
  }else{
    Serial.println("SD Card Found!");
  }

  // Initialize SD card
  SPI.begin(SD_SCLK, SD_MISO, SD_MOSI, SD_CS);
  if (!SD.begin(SD_CS,SPI)) {
    Serial.println("SD card initialization failed!");
    while (true)
      ;
  }
  Serial.println("SD card initialized.");

  // Read configuration
  readConfig();
  if (strcmp(wifi_mode, "") == 0 || strcmp(ssid, "") == 0 || strcmp(password, "") == 0 || strcmp(ftp_user, "") == 0 || strcmp(ftp_pass, "") == 0) {
    Serial.println("Please fill in the blanks in the config file.");
    while (true)
      ;
  }

  if (strcmp(wifi_mode, "CLIENT") == 0) {
    // Start Wi-Fi
    WiFi.begin(ssid, password);
    Serial.print("Connecting to Wi-Fi");
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.print(".");
    }
    Serial.println("Wi-Fi connected");
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());
  }else{
    WiFi.mode(WIFI_AP);
    WiFi.softAP(ssid, password);
    Serial.println("Wi-Fi AP started.");

    IPAddress IP = WiFi.softAPIP();
    Serial.print("AP IP address: ");
    Serial.println(IP);
  }

  // Start FTP server
  ftpSrv.setCallback(_callback);
  ftpSrv.setTransferCallback(_transferCallback);
  ftpSrv.begin(ftp_user, ftp_pass, " ESP32 MicroSD NAS by SiloCityLabs");//Change the username and password
  Serial.println("FTP server started.");

  // Start Web server
  server.on("/", handleRoot);
  server.on("/toggle", HTTP_POST, handleToggle);
  server.on("/led", HTTP_POST, handleLed);
  server.begin();
  Serial.println("Web server started.");

  //Enable led
  digitalWrite(USER_LED,0);

}

void loop() {
  // Handle FTP server
  ftpSrv.handleFTP();

  // Handle Web server
  server.handleClient();

}

void readConfig() {
  // Check if the config file exists
  if (!SD.exists("/config.txt")) {
    writeExampleConfig();
    Serial.println("Config file not found. Creating a new one.");
    return;
  }

  // Open the config file for reading
  File file = SD.open("/config.txt");
  if (file) {
    while (file.available()) {
      String line = file.readStringUntil('\n');
      line.trim(); // Remove whitespace or newline characters

      Serial.println("Read line: " + line);

      if (line.startsWith("MODE=")) {
        strncpy(wifi_mode, line.substring(5).c_str(), sizeof(wifi_mode) - 1);
        wifi_mode[sizeof(wifi_mode) - 1] = '\0'; // Null-terminate
      } else if (line.startsWith("HOMEWIFI=")) {
        strncpy(ssid, line.substring(9).c_str(), sizeof(ssid) - 1);
        ssid[sizeof(ssid) - 1] = '\0'; // Null-terminate
      } else if (line.startsWith("HOMEPASS=")) {
        strncpy(password, line.substring(9).c_str(), sizeof(password) - 1);
        password[sizeof(password) - 1] = '\0'; // Null-terminate
      } else if (line.startsWith("FTPUSER=")) {
        strncpy(ftp_user, line.substring(8).c_str(), sizeof(ftp_user) - 1);
        ftp_user[sizeof(ftp_user) - 1] = '\0'; // Null-terminate
      } else if (line.startsWith("FTPPASS=")) {
        strncpy(ftp_pass, line.substring(8).c_str(), sizeof(ftp_pass) - 1);
        ftp_pass[sizeof(ftp_pass) - 1] = '\0'; // Null-terminate
      } else {
        Serial.println("Invalid line in config file: " + line);
      }
    }
    file.close();
  } else {
    Serial.println("Error opening config file.");
    return;
  }

  // Log the read values
  Serial.println("===== CONFIGURATION =====");
  Serial.print("Wi-Fi Mode: ");
  Serial.println(wifi_mode);
  Serial.print("Wi-Fi SSID: ");
  Serial.println(ssid);
  Serial.print("Wi-Fi Password: ");
  Serial.println(password);
  Serial.print("FTP User: ");
  Serial.println(ftp_user);
  Serial.print("FTP Password: ");
  Serial.println(ftp_pass);
  Serial.println("=========================");
}


void writeExampleConfig() {
  File file = SD.open("/config.txt", FILE_WRITE);
  if (file) {
    file.println("MODE=CLIENT");
    file.println("HOMEWIFI=WIFINAME");
    file.println("HOMEPASS=PASSWORD");
    file.println("FTPUSER=user");
    file.println("FTPPASS=password");
    file.close();
    Serial.println("Example File written.");
  } else {
    Serial.println("Error opening file.");
  }
}
```

Upload this code to your **XIAO ESP32-C3** or **XIAO ESP32-C6** using the Arduino IDE.

**Step 5: Connecting to the FTP Server**

**For Wi-Fi Client Mode:**

1.  Connect your computer or device to the same Wi-Fi network as the ESP32.
    
2.  Use an FTP client (e.g., FileZilla) to connect:
    
    *   Host: IP address of the ESP32 (shown in Serial Monitor)
        
    *   Username: ftpuser
        
    *   Password: ftppass
        

**For Access Point Mode:**

1.  Connect to the ESP32's Wi-Fi network (YourSSID).
    
2.  Use the IP address of the ESP32 (default is usually 192.168.4.1) in your FTP client.
    

**Step 6: Web Interface**

The ESP32 also includes a web interface for monitoring and controlling the device:

*   Battery Voltage: View the current battery voltage and estimated percentage.
    
*   SD Card Status: Check whether the MicroSD card is inserted.
    
*   GPIO Control: Toggle GPIO states via buttons.
    

Access the interface by entering the ESP32’s IP address in your browser.

**Applications of ESP32 FTP Server**

*   Local NAS: Use the ESP32 as a compact file server for small networks.
    
*   Remote Monitoring: Check device and battery status from anywhere on your network.
    
*   IoT Integration: Serve data files to other devices or systems in your IoT ecosystem.
    

The **ESP32MiniSD** makes it easy to create a fully functional FTP server with additional features like battery monitoring and web control. With this setup, you can extend your IoT projects with robust file-serving capabilities.

For more details and updates, visit the [ESP32MiniSD GitHub repository](https://github.com/SiloCityLabs/esp32-minisd) or check out the [SiloCityLabs shop](https://shop.silocitylabs.com/).