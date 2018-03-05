#!/bin/bash


if [[ $1 = "erase" ]]; then
   echo Deleting
   rm -Rf /tmp/zookeeper /tmp/kafka-logs 2>/dev/null
else
   echo keeping
fi

cd /home/fraleyjd/Downloads/kafka_2.11-0.10.1.0
bin/zookeeper-server-start.sh config/zookeeper.properties >/home/fraleyjd/zookeeper.log 2>&1 &
sleep 10

cd /home/fraleyjd/Downloads/kafka_2.11-0.10.1.0
bin/kafka-server-start.sh config/server.properties >/home/fraleyjd/kafka 2>&1 &
sleep 10

cd /home/fraleyjd/Downloads/kafka_2.11-0.10.1.0
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic place

sleep 10
cd /home/fraleyjd/projects/go/src/Kafdrop
java -jar ./target/kafdrop-2.0.6.jar --zookeeper.connect=127.0.0.1:2181 >/home/fraleyjd/kafdrop.log 2>&1 &

