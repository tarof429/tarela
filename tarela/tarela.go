package tarela

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	BACKUP_COMMAND = "mksquashfs"
)

func RemoveFiles(files []os.DirEntry, remove int) {

	for _, f := range files[0:remove] {
		fmt.Printf("Removing file: %v\n", f.Name())
		if err := os.Remove(f.Name()); err != nil {
			fmt.Printf("Unable to remove %v\n", f.Name())
			log.Fatal(err)
		}
	}
}

// Generate a backup path name. Precision is only down to the minute.
// In real-life scenarios, this may not be a problem; however, note that
// if the backup path name already exists, it will be overwritten!
func GetBackupPathName(s string) string {
	now := time.Now()

	outputFileName := fmt.Sprintf("backup_%v.sfs", now.Format("200602010304"))

	return filepath.Join(s, outputFileName)
}

// Generate the command that will be used to create the backup
// It is hardcoded to use mksquashfs, but the output file can be
// whatever we like, useful for testing.
func GetBackupCommand(input, outputFilePath, exclude string) (string, []string) {

	args := []string{}

	args = append(args, input)
	args = append(args, outputFilePath)

	if len(exclude) > 0 {
		args = append(args, "-ef")
		args = append(args, exclude)
	}

	return BACKUP_COMMAND, args
}

// Backup the directory to a sfs file
func Backup(input, outputFilePath, exclude string) {
	cmd, args := GetBackupCommand(input, outputFilePath, exclude)

	_, err := exec.Command(cmd, args...).Output()

	if err != nil {
		log.Fatal(err)
	}
}
