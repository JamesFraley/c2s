--
create table t(
  file_id number,
  file_name varchar2(200));

create or replace procedure testInsert(file_id IN NUMBER, file_name IN varchar2) as
begin
  insert into t(file_id, file_name) values (file_id, file_name);
  commit;
end;
/


======================================================================================

drop table fraleyjd.iots_file_master;

create table IOTS_FILE_MASTER(
   file_id             number(10,0)     not null,
   file_version        number(10,0)     not null,
   classification      varchar2(4 byte) not null,
   state               varchar2(10)     not null,
   ifl_id              number(2,0)      not null,
   file_origin         varchar2(5)      not null,
   checksum            varchar2(64)     not null,
   file_size           number           not null,
   uri_location        varchar2(200)    not null,
   source_filename     varchar2(256)    not null unique,
   filename            varchar2(256)    not null);

======================================================================================

drop sequence fraleyjd.file_id_seq;
drop sequence fraleyjd.file_version_seq;

CREATE SEQUENCE file_id_seq
  START WITH 1
  INCREMENT BY 1
  CACHE 10;

CREATE SEQUENCE file_version_seq
  START WITH 1
  INCREMENT BY 1
  CACHE 10;

======================================================================================

CREATE OR REPLACE PACKAGE file_master_interface AS
  function register_file(p_row in iots_file_master%ROWTYPE) return iots_file_master%ROWTYPE;
END file_master_interface;
/

===================================================

create or replace PACKAGE BODY file_master_interface AS
   FUNCTION register_file(p_row IN iots_file_master%ROWTYPE) return iots_file_master%ROWTYPE IS
      curval iots_file_master%ROWTYPE;
   BEGIN
      curval := p_row;
      select file_id_seq.nextval      into curval.file_id from dual;
      select file_version_seq.nextval into curval.file_version from dual;
      curval.filename := to_char(curval.file_version) || '.' || curval.source_filename;
      insert into iots_file_master(file_id,
                                  file_version,
                                  classification,
                                 state,
                               ifl_id,
                               file_origin,
                               checksum,
                               file_size,
                               uri_location,
                               source_filename,
                               filename)
                        values(curval.file_id,
                               curval.file_version,
                               curval.classification,
                               curval.state,
                               curval.ifl_id,
                               curval.file_origin,
                               curval.checksum,
                               curval.file_size,
                               curval.uri_location,
                               curval.source_filename,
                               curval.filename);
      return(curval);
   END register_file;
END  file_master_interface;
/

======================================================================================

declare
   a_row iots_file_master%ROWTYPE;
   p_row iots_file_master%ROWTYPE;
begin
   p_row.source_filename := 'jims_file2';
   p_row.classification := '1000';
   p_row.state := 'PROCESSED';
   p_row.ifl_id := 1;
   p_row.file_origin := 'APX';
   p_row.checksum := '10923847';
   p_row.file_size := 120;
   p_row.uri_location := '/iots/prod/somewhere';
   a_row := file_master_interface.register_file(p_row);
end;
/

