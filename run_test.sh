#!/bin/sh
go test -v ./test
cd ./java_driver_test
sh ./run_test.sh
