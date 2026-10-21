---
title: TrueNAS Cron Jobs Worth Setting Up on Every Home Server
author: Luis Rodriguez
type: post
date: 2026-10-21
draft: false
categories: []
tags:
  - TrueNAS
  - homelab
  - docker
  - automation
subtitle: Here are 3 of them
description: "Three useful TrueNAS SCALE cron jobs for keeping dynamic DNS updated, refreshing game servers, and pulling the latest Plex container images."
---

A home server is great until you have to keep logging in to do the same maintenance tasks over and over. I've been running TrueNAS SCALE for a while, and a few simple cron jobs have saved me from having to babysit some of my services.

These are three jobs I actually use: one to keep my domain pointed at my home IP address, one to redeploy my Palworld server each morning, and one to pull updated Plex images monthly.

Not every TrueNAS system needs all three. If you host apps or expose services through a domain, though, they're worth considering.

## Where to add cron jobs in TrueNAS SCALE

In the TrueNAS web interface, go to **System > Advanced Settings > Cron Jobs** and select **Add**. Depending on your TrueNAS version, the wording or layout may differ slightly.

For each job, set a description, enter the command, choose an account with the required permissions, and configure the schedule. My jobs run as `truenas_admin`, which has the permissions needed on my system. You may need to use a different administrator account.

You can also enable or disable the job without deleting it. I recommend testing commands manually before scheduling them, especially commands that restart containers.

## 1. Keep your dynamic DNS updated

**Schedule:** Every day at midnight (`0 0 * * *`)

**Description:** Dynamic DNS update

**Command:**

```bash
wget -q --read-timeout=0.0 --waitretry=5 --tries=400 --background 'https://ipv4.cloudns.net/api/dynamicURL/?q=YOUR_DYNAMIC_DNS_TOKEN'
```

I use [ClouDNS](https://www.cloudns.net/) for dynamic DNS. If your internet provider changes your public IP address, your domain can end up pointing to the old address. The dynamic DNS URL tells ClouDNS to update its record to the public IPv4 address it sees from your request.

Replace `YOUR_DYNAMIC_DNS_TOKEN` with the unique URL or token provided by your DNS provider. Keep the full URL in quotes so the shell doesn't interpret special characters such as `&`.

The `wget` options here are from my existing setup. They make the request quiet, configure retries, and launch it in the background. However, backgrounding the process also means TrueNAS won't reliably know whether the request ultimately succeeded. For a simpler job with an observable exit status, I'd use:

```bash
wget -q --timeout=30 --tries=3 -O /dev/null 'https://ipv4.cloudns.net/api/dynamicURL/?q=YOUR_DYNAMIC_DNS_TOKEN'
```

A daily update is fine if your IP rarely changes, but it can leave your hostname wrong for up to a day after an address change. If that matters, use a shorter interval or a DDNS client that detects changes automatically.

**Important:** Treat your dynamic DNS URL like a password. Don't publish your real token in a blog post, screenshot, or public repository. Also, DDNS only updates DNS. It won't fix a downed reverse proxy, broken port forwarding, or an unavailable service.

## 2. Redeploy a game server on a schedule

**Schedule:** Every day at 4:00 AM (`0 4 * * *`)

**Description:** Palworld redeploy

**Command:**

```bash
midclt call app.redeploy palworld
```

I run a Palworld server as a TrueNAS app. Rather than manually redeploy it, I schedule the operation for early morning when people are less likely to be playing.

`midclt` is the TrueNAS command-line client for calling middleware API methods. The `app.redeploy` method redeploys the named app; in recent TrueNAS releases, the operation stops the app, pulls its configured images, and starts it again.

Change `palworld` to the **actual app name** shown in your TrueNAS installation if you're using this for something else.

A few things to know before enabling it:

- Players will be disconnected during the redeployment.
- Save game data and persistent volumes need to be configured properly. A redeploy is not a backup.
- An image refresh only gets whatever version the app's configured image tag points to. A pinned version won't magically advance to the next release.
- This does **not** replace installing available TrueNAS catalog app upgrades.

A daily redeploy is useful for my particular game-server setup, but I wouldn't blindly schedule daily restarts for databases or other apps that need uninterrupted availability.

## 3. Pull updated Plex images once a month

**Schedule:** Midnight on the first day of each month (`0 0 1 * *`)

**Description:** Plex image update

**Command:**

```bash
midclt call -job app.pull_images plex
```

Plex doesn't need a forced redeploy every morning on my setup, so I use a monthly image pull instead.

The `app.pull_images` method downloads the configured container images for the app. In recent TrueNAS SCALE versions, the default behavior is to **redeploy the app after pulling**, meaning Plex can briefly go offline while the container restarts.

The `-job` flag tells `midclt` to follow the long-running middleware job, which makes it more appropriate for an operation that may take time to finish.

The result depends on the image tags your app uses. If Plex is pinned to a specific image version, pulling it again won't upgrade Plex to a newer version. Similarly, if the TrueNAS catalog has a new app version or configuration migration, pulling Docker images alone won't apply that catalog upgrade.

If you want to download images **without** redeploying, newer TrueNAS API versions expose a `redeploy: false` option. Check the API documentation for your installed TrueNAS release before changing the command.

## All three jobs at a glance

| Job | Command | Schedule |
| --- | --- | --- |
| Dynamic DNS | `wget ... 'https://ipv4.cloudns.net/api/dynamicURL/?q=YOUR_DYNAMIC_DNS_TOKEN'` | Daily, 12:00 AM |
| Palworld redeploy | `midclt call app.redeploy palworld` | Daily, 4:00 AM |
| Plex image pull | `midclt call -job app.pull_images plex` | Monthly, day 1 at 12:00 AM |

These schedules use the TrueNAS server's configured time zone. Also note that the DDNS and Plex jobs both start at midnight on the first day of the month. They do different things, but it's easy to stagger them if you'd rather not overlap scheduled tasks.

## A few recommendations before you automate everything

Cron is handy, but it's also easy to forget that you set something to run months ago. Keep descriptions clear, check job output when troubleshooting, and leave errors visible while testing.

Most importantly, don't confuse maintenance automation with backups. A scheduled redeploy doesn't protect your app configuration or data, and automatically pulling the latest image can introduce a breaking change. Keep backups and have a rollback plan for apps you depend on.

Also remember that a restart may briefly take an app offline. Schedule jobs when nobody is using the service, and make sure nobody is relying on a long-running task inside the container.

## Final thoughts

I like having TrueNAS handle repetitive maintenance instead of remembering to log in and click buttons. Dynamic DNS, scheduled game-server redeployment, and occasional Plex image pulls are three small changes that make my home server easier to manage.

They're not universal defaults, and I wouldn't recommend automatic restarts for every service. But if you run similar apps, these commands provide a good starting point. Set your own schedules, test them, and let TrueNAS handle the repetitive work.

### References

- [TrueNAS: Managing Cron Jobs](https://www.truenas.com/docs/scale/systemsettings/advanced/managecronjobs/)
- [TrueNAS API: app.redeploy](https://api.truenas.com/v26.0/api_methods_app.redeploy.html)
- [TrueNAS API: app.pull_images](https://api.truenas.com/v25.04/api_methods_app.pull_images.html)
- [ClouDNS: Getting Started with Dynamic DNS](https://www.cloudns.net/wiki/article/364/)