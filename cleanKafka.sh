#!/bin/bash

/home/fraleyjd/Downloads/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --list   --zookeeper localhost:2181
/home/fraleyjd/Downloads/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --delete --zookeeper localhost:2181 --topic place
/home/fraleyjd/Downloads/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic place
/home/fraleyjd/Downloads/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --list   --zookeeper localhost:2181
