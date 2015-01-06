#!/bin/sh

for line in $(find $2 -name "*.json"); do
    FILENAME=${line##*/}
    GOFILENAME=${FILENAME%.*}.go
    go run gen.go $1 $line | gofmt > $GOFILENAME
done

