FROM centos:latest
ADD c2s /
COPY ./instantclient/oracle-instantclient12.2-basic-12.2.0.1.0-1.x86_64.rpm /
COPY ./instantclient/oracle-instantclient12.2-devel-12.2.0.1.0-1.x86_64.rpm /
COPY ./instantclient/oracle-instantclient12.2-sqlplus-12.2.0.1.0-1.x86_64.rpm /
COPY c2s.sh /
RUN /usr/bin/mkdir /config
RUN chmod 777 /c2s.sh /config
RUN yum install libaio -y
RUN rpm -i /oracle-instantclient12.2-basic-12.2.0.1.0-1.x86_64.rpm
RUN rpm -i /oracle-instantclient12.2-devel-12.2.0.1.0-1.x86_64.rpm
RUN rpm -i /oracle-instantclient12.2-sqlplus-12.2.0.1.0-1.x86_64.rpm
ENTRYPOINT /c2s.sh
