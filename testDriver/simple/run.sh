#!/bin/sh

mv ../../etc/cfg.json ../../etc/cfg.json.bak
cp cfg.json ../../etc/cfg.json
curl http://localhost:8888/api/reload

go run bench.go -t init -p 3306
go run bench.go -t type-test
