package tarela

import (
	"fmt"
	"reflect"
	"testing"
)

// Test whether we get:
// squashfs /home/john/src /home/john/dest/backup.sfs -ef /home/john/dest/exclude.txt
func TestGetSquashBackupCommand(t *testing.T) {
	src := "/home/john/src"
	dest := "/home/john/dest/backup.sfs"
	exclude := "/home/john/dest/exclude.txt"

	cmd, args := GetSquashfsBackupCommand(src, dest, exclude)

	fmt.Println(cmd)
	if cmd != SQUASH_BACKUP_COMMAND {
		t.Fail()
	}

	fmt.Println(args)
	if reflect.DeepEqual(args, []string{src, dest, "-ef", exclude}) != true {
		t.Fail()
	}
}

// Test whether we get:
// tar cf  /home/john/dest/backup.tar /home/john/src
func TestGetTarBackupCommand(t *testing.T) {
	src := "/home/john/src"
	dest := "/home/john/dest/backup.tar"

	cmd, args := GetTarBackupCommand(src, dest)

	fmt.Println(cmd)
	if cmd != TAR_BACKUP_COMMAND {
		t.Fail()
	}

	fmt.Println(args)
	if reflect.DeepEqual(args, []string{"cf", dest, src}) != true {
		t.Fail()
	}
}
