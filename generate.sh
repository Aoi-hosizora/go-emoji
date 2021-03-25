#!/bin/bash

# build
cd _generate || exit
go build -o generate.out
success=$?
cd ..

# generate
if [[ success -eq 0 ]]; then
  ./_generate/generate.out -pkg "emoji" -o "emoji.go"
fi
