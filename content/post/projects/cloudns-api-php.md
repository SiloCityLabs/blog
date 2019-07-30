---
title: ClouDNS API Library
author: Luis Rodriguez
type: post
date: 2015-01-16T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - php
  - wrapper
  - api
  - dns

---
This is a wrapper library for the ClouDNS API to be used by PHP applications. It tries to ease the integration of the API into your applications by handling all interactions with API and providing a simple interface to interact with. The entire projects [source code](https://gitlab.com/silo-city-labs-llc/cloudns-api-php) is posted on gitlab.

<!--more-->

To begin using the Library, the cloudns.php must be included in your application.

```php
require_once('/path/to/library/cloudns.php');
```

An instance of the ClouDNS must be created to interact with the library. This Object is the gateway to all interactions with the library. The API password obtained from the [ClouDNS](https://www.cloudns.net/api-settings/) must be passed into the ClouDNS by calling set_options.

```php
$cloudns = new ClouDNS();
$cloudns->set_options(array('auth-id' => '999','auth-password' => 'some_password'));
```

Functions
-------------------------------

<table width="100%">
	<tr>
		<th>Function</th>
		<th>Description</th>
	</tr>
	<tr>
		<td><code>detect_ip()</code></td>
		<td>Determine our IP address</td>
	</tr>
	<tr>
		<td><code>list_name_servers()</code></td>
		<td>Get a list with available domain name servers.</td>
	</tr>
	<tr>
		<td><code>list_zones(page,rows,search[optional])</code></td>
		<td>Gets a paginated list with zones you have or zone names matching a keyword.</td>
	</tr>
	<tr>
		<td><code>list_zone_stats()</code></td>
		<td>Gets the number of the zones you have and the zone limit of your customer plan. Reverse zones are included.</td>
	</tr>
	<tr>
		<td><code>delete_domain_zone(domain)</code></td>
		<td>This function is available only for slave zones, master zones and cloud/bulk domains. Works with reverse zones too.</td>
	</tr>
</table>

Examples
-------------------------------
Numerous examples have been provided in the [repository's examples folder](https://gitlab.com/silo-city-labs-llc/cloudns-api-php/tree/master/examples). The examples demonstrate how to accomplish most actions possible in the library. You are encouraged to look at these examples to learn the best practices for using the library.

