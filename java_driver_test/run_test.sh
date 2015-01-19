#!/bin/sh

cd cm-test
mvn install && mvn assembly:assembly && java -jar target/com.wandoujia.cm-test-1.0-SNAPSHOT.jar
