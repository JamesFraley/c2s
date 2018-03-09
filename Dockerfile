FROM centos:latest
ADD c2s /
COPY ./instantclient/oracle-instantclient12.2-basic-12.2.0.1.0-1.x86_64.rpm /
COPY ./instantclient/oracle-instantclient12.2-devel-12.2.0.1.0-1.x86_64.rpm /
COPY ./instantclient/oracle-instantclient12.2-sqlplus-12.2.0.1.0-1.x86_64.rpm /
COPY db.env /
RUN chmod 777 /db.env
RUN yum install libaio -y
RUN rpm -i /oracle-instantclient12.2-basic-12.2.0.1.0-1.x86_64.rpm
RUN rpm -i /oracle-instantclient12.2-devel-12.2.0.1.0-1.x86_64.rpm
RUN rpm -i /oracle-instantclient12.2-sqlplus-12.2.0.1.0-1.x86_64.rpm
RUN mkdir /tnsadmin
COPY tnsnames.ora /tnsadmin
#ENV PATH=/usr/lib/oracle/12.2/client64/bin:$PATH
#ENV LD_LIBRARY_PATH=/usr/lib/oracle/12.2/client64/lib/:$LD_LIBRARY_PATH
#ENV ORACLE_SID=orcl
#ENV TNS_ADMIN=/tnsadmin
#ENV LD_LIBRARY_PATH=/usr/lib/oracle/12.2/client64/lib
WORKDIR /
ENTRYPOINT /db.env
