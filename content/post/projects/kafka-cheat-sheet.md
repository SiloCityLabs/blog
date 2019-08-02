---
title: Kafka Cheat Sheet
author: Luis Rodriguez
type: post
date: 2019-08-01T08:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - kafka
  - zookeeper
  - cheat-sheet
  - commands
  - cli
  - terminal
  - shell
  - sh

---

Throughout my time working with kafka I have developed great software to use with it. But while the software did many things there are other things I still resorted to this cheat sheet for. I have put together this list of commands which I will keep updated here as I find new links and commands that help. My kafka installtion was under `/opt/kafka` and you will need to change this to your path you installed kafka too. You can also create a symlink and just copy paste commands as they are after changing ip's.

<!--more-->

## Kafka

### Kafka Commands

**Create Main topic**

`/opt/kafka/bin/kafka-topics.sh --create --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --replication-factor 3 --partitions 1 --topic {{topic name}}`

**List all topics**

`/opt/kafka/bin/kafka-topics.sh --list --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181`

**Delete topic**

`/opt/kafka/bin/kafka-topics.sh --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --delete --topic {{topic name}}`

**Describe Topic**

`/opt/kafka/bin/kafka-topics.sh --describe --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --topic {{topic name}}`

**Alter Topic**

`/opt/kafka/bin/kafka-topics.sh --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --alter --topic {{topic name}} --replication-factor 3`

**Console Consumer**

`/opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server XXX.XXX.XXX.XXX:9092, XXX.XXX.XXX.XXX:9092 --topic {{topic name}} --from-beginning`

**Create main stream**

```SQL
CREATE STREAM {{topic name}} (api BIGINT, type INT, page VARCHAR, cid VARCHAR, cip VARCHAR, exception INT, extra VARCHAR, date VARCHAR, time BIGINT) WITH (KAFKA_TOPIC='{{topic name}}', VALUE_FORMAT='JSON', TIMESTAMP='time');
```

## Zookeeper

https://gist.github.com/ursuad/e5b8542024a15e4db601f34906b30bb5
https://gist.github.com/sahilsk/a4d9d989fa4763c1abcbaec028f8346e

## Other

**latency test command**

```sh
curl -X GET \
  http:// XXX.XXX.XXX.XXX:3000/ip \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'X-IP: XXX.XXX.XXX.XXX ' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -s -o /dev/null -w '%{time_total}'
```
