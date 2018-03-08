#!/bin/bash

typeset -i X=`ps -ef | grep kaf | grep -v grep | grep -v kafkaStatus | wc -l`
echo "There are " $X " instances of Kafka running"
