FROM centos:latest
COPY c2s /c2sp
COPY c2s.sh /
COPY ./instantclient/oracle-instantclient12.2-basic-12.2.0.1.0-1.x86_64.rpm /
COPY ./instantclient/oracle-instantclient12.2-devel-12.2.0.1.0-1.x86_64.rpm /
COPY ./instantclient/oracle-instantclient12.2-sqlplus-12.2.0.1.0-1.x86_64.rpm /
RUN /usr/bin/mkdir /config
RUN /usr/bin/mkdir /c2s
RUN chmod 777 /c2s.sh /config /c2s
RUN yum install libaio -y
RUN rpm -i /oracle-instantclient12.2-basic-12.2.0.1.0-1.x86_64.rpm
RUN rpm -i /oracle-instantclient12.2-devel-12.2.0.1.0-1.x86_64.rpm
RUN rpm -i /oracle-instantclient12.2-sqlplus-12.2.0.1.0-1.x86_64.rpm
ENTRYPOINT /bin/sh
#ENTRYPOINT /c2s.sh
