rm -f c2s debug

if [ -z ${ORACLE_SID+x} ]; then
   echo "Sourcing db.env"
   . ~/db.env
else 
   echo "Oracle is setup"
fi

go build .

oracleUser="c2s" oraclePassword="myPassword" oracleHost="db1.protoeffect.cxm" oraclePort="1521" oracleService="orcl.protoeffect.cxm" kafkaConsumerAddr="10.0.0.36:2181" topicName="c2s" zookeeperAddr="10.0.0.36:9092" ./c2s
