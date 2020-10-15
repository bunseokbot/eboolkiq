#!/bin/zsh

for f in **/*.proto
do
    protoc -I . -I /usr/local/include \
        --go_out=plugins=grpc:. \
        --go_opt=paths=source_relative \
        "${f}"
done