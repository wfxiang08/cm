-- http://kimbriggs.com/computers/computer-notes/mysql-notes/mysql-data-types-50.file
drop database test_type;
create database test_type;

use test_type;

create table test_int (id INT, data INT);
insert into test_int values (0, 1);

create table test_tinyint (id INT, data TINYINT);
insert into test_tinyint values (0, 1);

create table test_bigint (id INT, data BIGINT);
insert into test_int values (0, 1);

create table test_double (id INT, data DOUBLE);
insert into test_double values (0, 1.5);

create table test_float (id INT, data FLOAT);
insert into test_float values (0, 1.5);

create table test_bool (id INT, data BOOL);
insert into test_bool values (0, TRUE);

create table test_varchar (id INT, data varchar(100));
insert into test_varchar values (0, "hello");

create table test_char (id INT, data varchar(1024));
insert into test_char values (0, "hello");

create table test_text (id INT, data text);
insert into test_text values (0, "hello");

create table test_longtext (id INT, data LONGTEXT);
insert into test_longtext values (0, "hello");

create table test_tinytext (id INT, data TINYTEXT);
insert into test_tinytext values (0, "hello");

create table test_binary (id INT, data BINARY);
insert into test_binary values (0, "hello");

create table test_varbinary (id INT, data VARBINARY(255));
insert into test_varbinary values (0, "hello");

create table test_blob (id INT, data BLOB);
insert into test_blob values (0, "hello");

create table test_tinyblob (id INT, data TINYBLOB);
insert into test_tinyblob values (0, "hello");

create table test_longblob (id INT, data LONGBLOB);
insert into test_longblob values (0, "hello");

create table test_timestamp (id INT, data TIMESTAMP);
insert into test_timestamp values (0, NOW());

create table test_datetime (id INT, data DATETIME);
insert into test_datetime values (0, NOW());

create table test_date (id INT, data DATE);
insert into test_date values (0, NOW());

create table test_enum (id INT, data enum("a", "b", "c"));
insert into test_enum values (0, "a");

create table test_set (id INT, data set("a", "b", "c"));
insert into test_set values (0, "a");
