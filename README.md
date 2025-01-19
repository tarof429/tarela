# Tarela

## Introduction

Tarela is a backup tool. It has clealy documented usage information and provides interactive input. 

## Usage
```sh
$ ./dist/tarela -input /home/taro/Code/ -output /home/taro/dest/ -keep 3
1 files in /home/taro/dest will be removed. Continue? (y/N): y
Removing file: backup_202518010944.sfs
Continue with backup? (y/N): y
✓          
```

## Building

This project makes use of the Make tool. Otherwise you'll need Go.

```sh
$ make
help                           Display this help
build                          Build the code
test                           Run all tests
clean                          Remove the dist directory
```