package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBackupCommand(t *testing.T) {
	src := "/home/john/src"
	dest := "/home/john/dest/backup.sfs"
	exclude := "/home/john/dest/exclude.txt"

	cmd, args := getBackupCommand(src, dest, exclude)

	fmt.Println(cmd)
	if cmd != "mksquashfs" {
		t.Fail()
	}

	fmt.Println(args)
	if reflect.DeepEqual(args, []string{src, dest, "-ef", exclude}) != true {
		t.Fail()
	}

}
