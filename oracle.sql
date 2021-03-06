create user c2s identified by myPassword quota unlimited on users;
grant connect to c2s;
grant create table to c2s;
grant create procedure to c2s;
grant create sequence to c2s;

drop table c2s.iots_file_master;

create table c2s.IOTS_FILE_MASTER(
   file_id             number(10,0)       not null,
   file_version        number(10,0)       not null,
   classification_text varchar2(255 BYTE) not null,
   state               varchar2(10)       not null,
   ifl_id              number(2,0)        not null,
   file_origin         varchar2(5)        not null,
   checksum            varchar2(64)       not null,
   file_size           number             not null,
   uri_location        varchar2(200)      not null,
   source_filename     varchar2(256)      not null unique,
   filename            varchar2(256)      not null);

--======================================================================================

drop sequence c2s.file_id_seq;
drop sequence c2s.file_version_seq;

CREATE SEQUENCE c2s.file_id_seq
  START WITH 1
  INCREMENT BY 1
  CACHE 10;

CREATE SEQUENCE c2s.file_version_seq
  START WITH 1
  INCREMENT BY 1
  CACHE 10;

--======================================================================================

CREATE OR REPLACE PACKAGE c2s.file_master_interface AS
  function register_file(p_row in iots_file_master%ROWTYPE) return iots_file_master%ROWTYPE;
END file_master_interface;
/

--===================================================

create or replace PACKAGE BODY c2s.file_master_interface AS
   FUNCTION register_file(p_row IN iots_file_master%ROWTYPE) return iots_file_master%ROWTYPE IS
      curval iots_file_master%ROWTYPE;
   BEGIN
      curval := p_row;
      select file_id_seq.nextval      into curval.file_id from dual;
      select file_version_seq.nextval into curval.file_version from dual;
      curval.filename := to_char(curval.file_version) || '.' || curval.source_filename;
      insert into iots_file_master(file_id,
                               file_version,
                               classification_text,
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
                               curval.classification_text,
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

--======================================================================================

create table c2s.iots_file_locations (
  ifl_id             number(2,0),
  tier               varchar2(5),
  disk_type          varchar2(8),
  network            varchar2(16),
  absolute_path_unix varchar2(128),
  absolute_path_win  varchar2(128),
  seclab             number(11,0),
  created            timestamp(6),
  created_by         varchar2(30),
  last_updated       timestamp(6),
  last_updated_by    varchar2(30)
);


--======================================================================================

insert into c2s.IOTS_FILE_LOCATIONS(ifl_id, tier,   disk_type, network,    absolute_path_unix, absolute_path_win,     seclab)
                         values( 1,    'TIER1', 'C2S',     'UNKNOWN',   '/iots/test',     '\\ifmfs1\iots\test', '1000');
                         
insert into c2s.IOTS_FILE_LOCATIONS(ifl_id, tier,   disk_type, network,    absolute_path_unix, absolute_path_win,     seclab)
                         values( 2,    'TIER2', 'C2S',     'UNKNOWN',   '/iots/train',     '\\ifmfs1\iots\train', '1000');
                         
insert into c2s.IOTS_FILE_LOCATIONS(ifl_id, tier,   disk_type, network,    absolute_path_unix, absolute_path_win,     seclab)
                         values( 3,    'TIER3', 'C2S',     'UNKNOWN',   '/iots/prod',         '\\ifmfs1\iots\prod', '1000');
                         
insert into c2s.IOTS_FILE_LOCATIONS(ifl_id, tier,   disk_type, network,    absolute_path_unix, absolute_path_win,     seclab)
                         values( 4,    'TIER4', 'C2S',     'UNKNOWN',   '/c2s/train',         '\\ifmfs1\c2s\train', '1000');
                         
insert into c2s.IOTS_FILE_LOCATIONS(ifl_id, tier,   disk_type, network,    absolute_path_unix, absolute_path_win,     seclab)
                         values( 5,    'TIER5', 'C2S',     'UNKNOWN',   '/c2s/prod',         '\\ifmfs1\c2s\prod', '1000');
                         
commit;


--======================================================================================
--
-- This is used to test the package
--
--declare
--   a_row iots_file_master%ROWTYPE;
--   p_row iots_file_master%ROWTYPE;
--begin
--   p_row.source_filename := 'jims_file2';
--   p_row.classification := '1000';
--   p_row.state := 'PROCESSED';
--   p_row.ifl_id := 1;
--   p_row.file_origin := 'APX';
--   p_row.checksum := '10923847';
--   p_row.file_size := 120;
--   p_row.uri_location := '/iots/prod/somewhere';
--   a_row := file_master_interface.register_file(p_row);
--end;
--/
--======================================================================================
