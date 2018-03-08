#!/bin/bash

if [[ $# < 1 ]]; then
   echo No file identified!
   exit
fi

cd ./data/

rm -f tmp.json tmp2.json input_file.json

FPAHERE=`shuf -i 1000000-9999999 -n 1`
 SNHERE=`shuf -i 1000000-9999999 -n 1`

sed -e s/FPAHERE/$FPAHERE/g $1 > tmp.json
sed -e s/SNHERE/$SNHERE/g   tmp.json > tmp2.json
cat tmp2.json | tr -d '\n' > input_file.json

/home/fraleyjd/Downloads/kafka_2.11-0.10.1.0/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic place < input_file.json

rm -f tmp.json tmp2.json input_file.json

cd ..
