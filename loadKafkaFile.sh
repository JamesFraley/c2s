#!/bin/bash

if [[ $# < 1 ]]; then
   inFile=catalog-record-formatted.json
else
   inFile=$1
fi

cd ./data/

rm -f tmp.json tmp2.json input_file.json

FPAHERE=`shuf -i 1000000-9999999 -n 1`
 SNHERE=`shuf -i 1000000-9999999 -n 1`

filePath=/c2s/prod/data/cbf/tfrm/2018/03/08
outFile=20180308_0953_$FPAHERE_$SNHERE.txt

sed -e s/FNHERE/$outFile/g    $inFile  > tmp.json
sed -e s,PATHHERE,$filePath,g tmp.json > tmp2.json
cat tmp2.json | tr -d '\n'             > input_file.json

echo `date +"%F %T"` > $filePath/$outFile

/home/fraleyjd/Downloads/kafka_2.11-0.10.1.0/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic place < input_file.json

#rm -f tmp.json tmp2.json input_file.json

cd ..
