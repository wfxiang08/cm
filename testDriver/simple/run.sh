#!/bin/sh

mv ../../etc/cfg.json ../../etc/cfg.json.bak
cp cfg.json ../../etc/cfg.json

go run bench.go -t init -p 3306


curl http://localhost:8888/api/reload
go run bench.go -t type-test
