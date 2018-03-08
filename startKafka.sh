#!/bin/bash

typeset -i X=`ps -ef | grep kaf | grep -v grep | wc -l`

if (( X>0 )); then
   echo Kafka is already running
   exit
fi

echo Starting Zookeeper
cd /home/fraleyjd/Downloads/kafka_2.11-0.10.1.0
bin/zookeeper-server-start.sh config/zookeeper.properties >/home/fraleyjd/zookeeper.log 2>&1 &
sleep 5

echo Starting Kafka
cd /home/fraleyjd/Downloads/kafka_2.11-0.10.1.0
bin/kafka-server-start.sh config/server.properties >/home/fraleyjd/kafka.log 2>&1 &
sleep 5

echo Creating partition
cd /home/fraleyjd/Downloads/kafka_2.11-0.10.1.0
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic place
sleep 5

echo Starting Kafdrop
cd /home/fraleyjd/projects/go/src/Kafdrop
java -jar ./target/kafdrop-2.0.6.jar --zookeeper.connect=127.0.0.1:2181 >/home/fraleyjd/kafdrop.log 2>&1 &

