# c2s
Kafka to Oracle 

This project will monitor a Kafka topic and insert the value part of the Kafka message into an Oracle table.  It uses the Shopify/sarama package to monitor Kafka.

1: Install/setup go 1.10, docker, git

2: Download c2s
	git clone https://github.com/JamesFraley/c2s.git

3: Install Oracle instant client
	sudo rpm -i oracle-instantclient12.2-basic-12.2.0.1.0-1.x86_64.rpm
	sudo rpm -i oracle-instantclient12.2-devel-12.2.0.1.0-1.x86_64.rpm
	sudo rpm -i oracle-instantclient12.2-sqlplus-12.2.0.1.0-1.x86_64.rpm
	
4: Create the ~/db.env
	export PATH=/usr/lib/oracle/12.2/client64/bin:$PATH
	export LD_LIBRARY_PATH=/usr/lib/oracle/12.2/client64/lib/:$LD_LIBRARY_PATH
	export ORACLE_SID=orcl
	export TNS_ADMIN=/usr/lib/oracle/12.2/client64/network
	export LD_LIBRARY_PATH=/usr/lib/oracle/12.2/client64/lib

5: Create the oci8.pc

	cp $GOPATH/src/c2s/oci8.pc /usr/lib64/pkgconfig/oci8.pc
	ln /usr/lib64/pkgconfig/oci8.pc  /usr/share/pkgconfig/oci8.pc


6: Build c2s
	cd $GOPATH/src/c2s
	./build.share

7:  Build the docker image
	cd $GOPATH/src/c2s
	docker build . -t c2s:0.0.1


--------------------------------------------
--
--  Testing starts here
--
--------------------------------------------

A: Ensure an Oracle database is started and accessible.
	-Use EZCONNECT "CONNECT username/password@[//]host[:port][/service_name]"
	-You'll have to look up the service name
	-There may be a couple of errors where the script tries to drop nonexistent objects.
	
	sqlplus c2s/myPassword@db1:1521/orcl.????.???? @oracle.sql
	
B: Start kafka
	-Currently, we use: kafka_2.11-0.10.1.0.tgz
	-Download, unzip and move kafka to /opt/
	-Edit the config file
		cd /opt/kafka_2.11-0.10.1.0/config/
		vi server.properties 
		Uncomment delete.topic.enable=true
	-Start Kafka
		cd $GOPATH/src/c2s
		./startKafka.sh

	-You can use ./cleanKafka.sh to delete the topic and recreate it.

C: Run the program stand alone:
	-You will have to update the following configurations in runC2S.sh:
		oracleHost
		oracleService
	-Crate the data directories
		sudo mkdir -p /c2s/prod/data/cbf/tfrm/2018/03;sudo chmod -R 777 /c2s
		cd /c2s/prod/data/cbf/tfrm/2018/03
		for X in `seq -w 1 31`;do mkdir $X; done; chmod 777 *

	-Start the program
		./runC2S.sh

	-You can push a file to the Kafka topic with:
		./loadKafkaFile.sh

	-Next you check oracle by logging in and executing the following command:
	  select * from c2s.iots_file_master;
	  
D:

docker build . -t c2s:0.0.1

docker run -v /home/fraleyjd/go/src/c2s:/config -v /c2s:/data c2s:0.0.1

REFERENCES:
	https://github.com/mattn/go-oci8
	https://github.com/golang/go/wiki/SQLDrivers
	https://gocodecloud.com/blog/2016/08/09/accessing-an-oracle-db-in-go/
	https://apextips.blogspot.com/2015/09/making-connections-to-oracle-database.html
	https://jsoneditoronline.org/
	https://github.com/paulmach/go.geojson.git
	http://www.orafaq.com/wiki/EZCONNECT
