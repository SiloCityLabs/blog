---
title: Local Failover for DigitalOcean Functions
subtitle: How a local dev tool became a resilient, cost-aware backend
author: Luis Rodriguez
type: post
date: 2026-02-08
categories:
  - Infrastructure
tags:
  - Serverless
  - Failover
  - DigitalOcean
  - TrueNAS
  - Golang
  - Security
---

What started as a simple local development tool for DigitalOcean Functions quietly turned into something much more serious: a production-grade failover system that now runs across home TrueNAS servers owned by the operators of [FilaMeter](https://filameter.com).

Internally, we call it **DOrunner**.

The original goal was mundane—test Go-based serverless functions locally without burning through API calls or waiting on cloud deploys. But as usage grew and traffic patterns became less predictable, that local runner evolved into a core part of our production architecture.

This post explains why that happened, what problem it solves, and how the system works in practice.

## The Real Problem With Serverless

Serverless platforms are operationally attractive but economically fragile. The pay-per-invocation model turns traffic into a direct financial liability.

The most dangerous failure mode isn’t downtime—it’s **cost runaway**.

In traditional infrastructure, abuse usually manifests as degraded performance. In serverless, abuse manifests as a bill.

This is where **Denial-of-Wallet (DoW)** attacks matter. Instead of knocking a service offline, an attacker simply forces it to keep working until it becomes unaffordable. Bot traffic, aggressive scrapers, and AI crawlers make this problem worse, especially as they increasingly ignore conventions like `robots.txt`.

For us, the risks broke down into four concrete issues:

- **Unpredictable spend** from traffic spikes
- **Automated scraping** that provides no user value
- **Hard free-tier limits** that legitimate usage can exceed
- **Vendor dependency** when all execution happens in one cloud

We needed cost ceilings that could not be exceeded—without taking the service offline.

## From Local Runner to Failover Backend

DOrunner already existed as a local execution environment for DigitalOcean Functions. It compiled and ran the same Go code we deployed to the cloud, using the same interfaces and contracts.

That symmetry made the next step obvious: instead of *only* using it locally, we could run it **elsewhere**.

TrueNAS boxes—spread across multiple homes—became the overflow capacity. These machines were already reliable, online 24/7, and under our direct control. Electricity and bandwidth costs were fixed. Compute was effectively “free.” They sit idle at a mere 18MB ram usage.

At that point, DOrunner stopped being a dev tool and became a failover backend.

## Hard Limits, Based on Data

Before routing any traffic, we looked at real usage.

Historical DigitalOcean data showed that we could safely handle about **18,000 function invocations per month** before entering paid tiers. To stay well clear of that boundary, we enforced a **hard cap of 600 invocations per day**.

The key decision:  
**routing happens before limits are exceeded, not after.**

Once the daily budget approaches exhaustion, new requests are diverted automatically to self-hosted runners. No throttling. No degraded service. No surprise invoices.

This is not rate limiting—it’s cost-aware execution placement.

## Traffic Control and Gateway Enforcement

All requests enter through **Crate.cc**, which acts as a gateway in front of both DigitalOcean and the TrueNAS runners.

Crate.cc is responsible for:

- Tracking invocation counts in real time
- Enforcing daily hard limits
- Detecting abusive or automated traffic
- Routing requests to the appropriate backend

From the application’s perspective, there is exactly one HTTPS endpoint. Whether a request runs in the cloud or on a basement server is an internal decision.

Legitimate user traffic gets priority access to cloud execution. Everything else—bots, scrapers, anomalous traffic—gets diverted or blocked before it can consume metered resources.

## Execution Model

Each TrueNAS node runs containerized versions of our DigitalOcean Functions.

Key properties of the runner:

- Identical Go binaries and business logic across environments
- Dynamic build and execution from staging directories
- Automatic dependency resolution
- Binary caching to avoid cold starts

This keeps behavior consistent regardless of where the code runs. A function does not “know” whether it’s executing in DigitalOcean or on a friend’s NAS.

## Distributed, Not Centralized

There is no single failover box.

Each node is independently capable of handling overflow traffic. If one machine goes offline, routing simply shifts to the remaining capacity. The system degrades gracefully instead of catastrophically.

This model spreads operational risk while keeping ownership and control local.

## Results in Production

Since deploying DOrunner for **filameter.com**, the system has delivered measurable results:

- **Zero unexpected cloud charges** for over eight months
- **99.9% uptime** for sync operations
- **Sub-30ms routing overhead** at the gateway
- **2.3× traffic growth** absorbed without cloud cost increases
- **Automated mitigation** of denial-of-wallet attempts

The most important outcome isn’t performance—it’s predictability. Costs are bounded by design.

## When robots.txt Stops Working

AI training crawlers and large-scale scrapers increasingly ignore exclusion signals. Blocking them at the application layer is too late; by then, you’ve already paid.

Crate.cc’s detection layer identifies scraping behavior—request rate anomalies, enumeration patterns, User-Agent spoofing—and intercepts it before it reaches metered execution.

Traffic that insists on existing is redirected to infrastructure we own. It consumes our power, not our budget.

This isn’t ideological. It’s practical.

The 600-request daily budget is reserved for actual users, not training corpora.

## Why This Matters

Academic research is already acknowledging denial-of-wallet attacks as a serious serverless threat, not a theoretical one. Systems that assume “scale equals success” ignore the economic reality of small services.

DOrunner is a pragmatic response:  
use the cloud when it makes sense, and stop using it *before* it becomes a liability.

Serverless doesn’t have to mean surrendering cost control. With hard limits, intelligent routing, and overflow capacity you actually own, you can keep the convenience without the risk.

Sometimes the smartest cloud strategy is knowing exactly when not to use the cloud—and having that decision enforced automatically.