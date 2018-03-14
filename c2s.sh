export PATH=/usr/lib/oracle/12.2/client64/bin:$PATH
export LD_LIBRARY_PATH=/usr/lib/oracle/12.2/client64/lib/:$LD_LIBRARY_PATH
source ./config/c2s.config

echo oracleuser=$oracleUser
echo oraclePassword=$oraclePassword
echo oracleHost=$oracleHost
echo oraclePort=$oraclePort
echo oracleServiice=$oracleService
echo kafkaConsumerAddr=$kafkaConsumerAddr
echo zookeeperAddr=$zookeeperAddr
echo topicName=$topicName

oracleUser=$oracleUser \
   oraclePassword=$oraclePassword \
   oracleHost=$oracleHost \
   oraclePort=$oraclePort \
   oracleService=$oracleService \
   kafkaConsumerAddr=$kafkaConsumerAddr \
   topicName=$topicName \
   zookeeperAddr=$zookeeperAddr \
   ./c2sp
