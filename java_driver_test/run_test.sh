#!/bin/sh

cd cm-test
mysql -e "DROP DATABASE IF EXISTS test; CREATE DATABASE test;CREATE TABLE test.tbl_test (id INT, data VARCHAR(255));" -uroot -P 3306 --host 127.0.0.1
mvn install && mvn assembly:assembly && java -jar target/com.wandoujia.cm-test-1.0-SNAPSHOT.jar
