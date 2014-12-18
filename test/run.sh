#!/bin/sh
rm ./test &> /dev/null
go build -o ./test

./test -t init

mv ../etc/cfg.json ../etc/cfg.json.bak

cp cfg.json ../etc/cfg.json
curl http://127.0.0.1:8888/api/reload
./test -t type-test

mv ../etc/cfg.json.bak ../etc/cfg.json
