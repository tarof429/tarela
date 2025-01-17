# Tarela

## Introduction

Tarela is a backup tool. It has clealy documented usage information and provides interactive input. 

## Usage
```sh
$ ./dist/tarela 
Usage of ./dist/tarela:
  -input string
    	Input directory
  -keep int
    	Number of files to keep
  -output string
    	Output file
$ ./dist/tarela  -input /home/taro/Documents/ -output /home/taro/dest/ -keep 1
1 files in /home/taro/dest will be removed. Continue? (y/N): y
Removing files
Removing file: backup_202517011018.tar
Continue with backup file creation? (y/N): y
Creating /home/taro/dest/backup_202517011035.tar
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