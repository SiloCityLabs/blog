---
title: Sending SMS via the Nexmo SMS gateway.
type: post
date: 2015-09-15T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - php
  - sms
  - gateway

---

A class written to simplify the use of the nexmo api gateway for SMS messaging. This class was a fork from another repository with new gateway endpoints added and adding an entirely new service for automated phone calls with a tts engine. [Download the file here](https://gitlab.com/silo-city-labs-llc/Nexmo-PHP-lib) or look at the code below. 

<!--more-->

### Quick Examples

1) Sending an SMS

```
$sms = new NexmoMessage('account_key', 'account_secret');
$sms->sendText( '+447234567890', 'MyApp', 'Hello world!' );
```


2) Recieving SMS 

```
$sms = new NexmoMessage('account_key', 'account_secret');
if ($sms->inboundText()) {
    $sms->reply('You said: ' . $sms->text);
}
```


3) Recieving a message receipt

```
$receipt = new NexmoReceipt();
if ($receipt->exists()) {
    switch ($receipt->status) {
        case $receipt::STATUS_DELIVERED:
            // The message was delivered to the handset!
            break;
        
        case $receipt::STATUS_FAILED:
        case $receipt::STATUS_EXPIRED:
            // The message failed to be delivered
            break;
    }
}
```


4) List purchased numbers on your account

```
$account = new NexmoAccount('account_key', 'account_secret');
$numbers = $account->numbersList();
```
