---
title: How I acchieved 98 seo score by dropping wordpress
type: post
date: 2019-07-15T22:36:25+00:00
comments: false
draft: true
categories:
  - Projects
tags:
  - luis-portfolio
  - maave-portfolio

---



<!--more-->


## Kafka

*

### Kafka Commands

 

**Create Main topic**

 

`/opt/kafka/bin/kafka-topics.sh --create --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --replication-factor 3 --partitions 1 --topic {{topic name}}`

 

**List all topics**

 

`/opt/kafka/bin/kafka-topics.sh --list --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 XXX.XXX.XXX.XXX:2181`

 

**Delete topic**

 

`/opt/kafka/bin/kafka-topics.sh --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --delete --topic {{topic name}}`

 

**Describe Topic**

 

`/opt/kafka/bin/kafka-topics.sh --describe --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --topic {{topic name}}`

 

**Alter Topic**

 

`/opt/kafka/bin/kafka-topics.sh --zookeeper XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181, XXX.XXX.XXX.XXX:2181 --alter --topic {{topic name}} --replication-factor 3`

 

**Console Consumer**

 

/opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server XXX.XXX.XXX.XXX:9092, XXX.XXX.XXX.XXX:9092 XXX.XXX.XXX.XXX:9092 --topic {{topic name}} --from-beginning

 

**Create main stream**

 

```

CREATE STREAM {{topic name}} (api BIGINT, type INT, page VARCHAR, cid VARCHAR, cip VARCHAR, exception INT, extra VARCHAR, date VARCHAR, time BIGINT) WITH (KAFKA_TOPIC='{{topic name}}', VALUE_FORMAT='JSON', TIMESTAMP='time');

```

 

## Zookeeper

*

 

https://gist.github.com/ursuad/e5b8542024a15e4db601f34906b30bb5

https://gist.github.com/sahilsk/a4d9d989fa4763c1abcbaec028f8346e

 

## Other

 

**latency test command**

```
curl -X GET \
  http:// XXX.XXX.XXX.XXX:3000/ip \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: XXX.XXX.XXX.XXX:3000' \
  -H 'X-IP: XXX.XXX.XXX.XXX ' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -s -o /dev/null -w '%{time_total}'
```
