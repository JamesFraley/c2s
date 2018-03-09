export PATH=/usr/lib/oracle/12.2/client64/bin:$PATH
export LD_LIBRARY_PATH=/usr/lib/oracle/12.2/client64/lib/:$LD_LIBRARY_PATH
export ORACLE_SID=orcl
#export TNS_ADMIN=/tnsadmin
export LD_LIBRARY_PATH=/usr/lib/oracle/12.2/client64/lib
source /config/c2s.config
/c2s
#oracleUser="fraleyjd" oraclePassword="myPassword" oracleHost="10.103.20.161" oraclePort="1521" oracleService="orcl.altamiracorp.com" kafkaConsumerAddr="172.18.0.1:2181" topicName="place" zookeeperAddr="127.0.0.1:9092" ./c2s
