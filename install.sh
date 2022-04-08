#!/bin/bash
dir='tmp'
if [ ! -d "$dir" ]
then
    echo 'tmp folder does not exist, cloning to tmp .....'
    git clone https://github.com/go-swagger/go-swagger "$dir" 
fi
cd "$dir"
go install ./cmd/swagger