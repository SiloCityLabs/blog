---
title: An SMTP black hole for unit tests
author: Luis Rodriguez
type: post
date: 2026-10-15
draft: false
categories:
  - Projects
tags:
  - smtp
  - golang
  - email
  - testing
  - fakeemaildomain
subtitle: "fakeemaildomain.com accepts any address, answers the way you ask, and throws the message away"
---

A test that sends mail needs somewhere for the message to land. [fakeemaildomain.com](https://fakeemaildomain.com) is that place: a public mail exchanger that accepts the envelope, replies on purpose, and discards the body. The site has no signup, no inbox, and no forwarding.

<!--more-->

Invent any local part. `ada@fakeemaildomain.com` answers `250 OK`, the message is counted, and the bytes are dropped. The address is the whole API. A tag after `+` chooses the reply for the SMTP command that carries it. A numeric subdomain does the same refusal on `RCPT TO`, which is the form you want when the software under test builds the recipient and you would rather keep the local part clean.

## Pick the reply

| Address | What the exchanger does |
| --- | --- |
| `ada@fakeemaildomain.com` | `250 OK`, then the message is discarded |
| `ada+550@fakeemaildomain.com` | That SMTP code, any value from 100 to 599, on the command that carries the address |
| `jon@550.fakeemaildomain.com` | The same refusal, on `RCPT TO`. This needs a wildcard MX |
| `ada+5.1.1@fakeemaildomain.com` | `550 5.1.1 User unknown` |

Enhanced status codes follow the usual SMTP class. `2.x.x` replies `250`, `4.x.x` replies `450`, and `5.x.x` replies `550`. Two mailbox-full cases are the exceptions the RFCs already special-case: `4.2.2` is `452`, and `5.2.2` and `5.2.3` are `552`. A handful of codes also get a real reason string (`5.1.1` User unknown, `5.1.2` Host unknown, `5.2.2` Mailbox full, `5.7.1` Delivery not authorized, and a few neighbors). Anything else comes back as `Requested action not taken`.

## Choose the moment

The same tag can name the command. That matters when the failure has to happen before `DATA`, or only after the client has already sent the body.

| Address | When it fails |
| --- | --- |
| `ada+mail-550@fakeemaildomain.com` | Rejects `MAIL FROM` |
| `ada+rcpt-550@fakeemaildomain.com` | Rejects the recipient |
| `ada+data-552@fakeemaildomain.com` | Accepts the envelope. After the final dot, `552` |

If several recipients on one transaction ask for a `DATA` result, the first one sticks. Later recipients are still accepted. The reply after the dot is the first data code the server saw.

## Slow, rude, and greylisted

Some tests are about the client, not the status code. These tags stall the session, cut it, or pretend to greylist.

| Address | Behavior |
| --- | --- |
| `ada+delay-2s@fakeemaildomain.com` | Pauses, then replies. The wait is capped at 30 seconds |
| `ada+drop@fakeemaildomain.com` | Closes the connection |
| `ada+hang@fakeemaildomain.com` | Sends nothing, until the client closes or two minutes pass |
| `ada+greylist@fakeemaildomain.com` | `450 Try again later` once for that sender, then `250 OK` |

Greylisting is remembered in memory for the life of the process, keyed by the sender address. The table holds 8,192 senders. When it fills, it is cleared and the next delivery from a sender looks like a first attempt again. A restart does the same. That is the behavior to assert in a test: one temporary failure, then acceptance, for as long as the process stays up and the table has room.

`EHLO` advertises `PIPELINING`, `8BITMIME`, `SIZE` (50 MiB), `SMTPUTF8`, and `STARTTLS`. `STARTTLS` uses a certificate generated for the process, so a client that verifies the chain should expect that handshake to fail on purpose. A message larger than 50 MiB is refused with `552` after the dot.

The homepage includes a small dialogue box. Type an address and it prints the SMTP conversation that address will produce, which is a faster way to check a tag than opening a session.

## Point your own domain at it

The exchanger can sit behind a name you already use in tests. Alias a mail host, aim the MX record at that alias, and plus-addresses on your domain work immediately. A wildcard MX lets numeric subdomains such as `user@550.yourdomain.com` reach it too.

```
mail.example.com.   CNAME   fakeemaildomain.com.
example.com.        MX 10   mail.example.com.
*.example.com.      MX 10   mail.example.com.
```

Replace `example.com` with the domain the tests send to. The CNAME is the piece that keeps the setup movable: if the exchanger's address changes, mail follows the name.

## One process, two listeners

The public site and the mail exchanger are the same Go program. SMTP is accepted on port 25. HTTP and HTTPS serve the page, and HTTPS gets its certificate from Let's Encrypt through the HTTP listener. The HTML is embedded in the binary, so a deploy is the binary plus an OpenRC service.

`GET /count` returns `{"messages":N}`, the number of messages discarded since the process started. The page polls that and shows the total. Nothing else is stored. There is no queue to drain and no mailbox to empty between test runs.

The exchanger is free to use. If it saves you a test mailbox, there is a [PayPal](https://paypal.me/ldrrp) and a [Cash App](https://cash.app/$ldrrp) on the site.
