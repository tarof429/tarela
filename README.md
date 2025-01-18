# Tarela

## Introduction

Tarela is a backup tool. It has clealy documented usage information and provides interactive input. 

## Usage
```sh
$ ./dist/tarela -input /home/taro/src/ -output /home/taro/dest/ -keep 3
No files need to be removed
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