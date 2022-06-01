#!/bin/bash

DIR=$1
TARGET_DIR=$2

mkdir -p $TARGET_DIR

for file in $DIR/abi/*; 
do 
    file_name="${file##*/}"
    file_prefix="${file_name%.*}"
    mkdir -p $TARGET_DIR/$file_prefix
    abigen --abi=$file --bin=$DIR/bin/$file_prefix.bin --pkg=$file_prefix --out=abigen/$file_prefix/$file_prefix.go; 
done