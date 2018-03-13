#!/bin/bash

/opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --list   --zookeeper localhost:2181
/opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --delete --zookeeper localhost:2181 --topic c2s
/opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 10 --topic c2s
/opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --list   --zookeeper localhost:2181
