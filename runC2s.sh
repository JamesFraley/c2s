rm -f c2s debug

if [ -z ${ORACLE_SID+x} ]; then
   echo "Sourcing db.env"
   . ~/db.env
else 
   echo "Oracle is setup"
fi

go build .

oracleUser="fraleyjd" oraclePassword="myPassword" oracleHost="10.103.20.161" oraclePort="1521" oracleService="orcl.altamiracorp.com" kafkaConsumerAddr="127.0.0.1:2181" topicName="place" zookeeperAddr="127.0.0.1:9092" ./c2s
#kafkaConsumerAddr="127.0.0.1:2181" topicName="place" zookeeperAddr="127.0.0.1:9092" oracleUser="fraleyjd" oraclePassword="myPassword" oracleService="orcl" ./c2s
