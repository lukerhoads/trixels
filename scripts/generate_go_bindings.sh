#!/bin/bash

DIR=$1

for file in $DIR/*; 
do 
    file_name="${file##*/}"
    file_prefix="${file_name%.*}"
    abigen --abi=$file --pkg=abigen --out=abigen/$file_prefix.go; 
done