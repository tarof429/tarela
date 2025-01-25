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
	SQUASH_BACKUP_COMMAND = "mksquashfs"
	TAR_BACKUP_COMMAND    = "tar"
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

// Generate a backup path name for a tar file.
func GetTarBackupPathName(s string) string {
	now := time.Now()

	outputFileName := fmt.Sprintf("backup_%v.tar", now.Format("200602010304"))

	return filepath.Join(s, outputFileName)
}

// Generate a backup path name. Precision is only down to the minute.
// In real-life scenarios, this may not be a problem; however, note that
// if the backup path name already exists, it will be overwritten!
func GetSquashfsBackupPathName(s string) string {
	now := time.Now()

	outputFileName := fmt.Sprintf("backup_%v.sfs", now.Format("200602010304"))

	return filepath.Join(s, outputFileName)
}

// Generate the command that will be used to create the backup
// It is hardcoded to use mksquashfs, but the output file can be
// whatever we like, useful for testing.
func GetTarBackupCommand(input, outputFilePath string) (string, []string) {

	args := []string{}

	args = append(args, "cf")
	args = append(args, outputFilePath)
	args = append(args, input)

	return TAR_BACKUP_COMMAND, args
}

// Generate the command that will be used to create the backup
// It is hardcoded to use mksquashfs, but the output file can be
// whatever we like, useful for testing.
func GetSquashfsBackupCommand(input, outputFilePath, exclude string) (string, []string) {

	args := []string{}

	args = append(args, input)
	args = append(args, outputFilePath)

	if len(exclude) > 0 {
		args = append(args, "-ef")
		args = append(args, exclude)
	}

	return SQUASH_BACKUP_COMMAND, args
}

// BackupSquashfs the directory to a sfs file
func BackupSquashfs(input, outputFilePath, exclude string) {
	cmd, args := GetSquashfsBackupCommand(input, outputFilePath, exclude)

	_, err := exec.Command(cmd, args...).Output()

	if err != nil {
		log.Fatal(err)
	}
}

// BackupSquashfs the directory to a sfs file
func BackupTar(input, outputFilePath string) {
	cmd, args := GetTarBackupCommand(input, outputFilePath)

	_, err := exec.Command(cmd, args...).Output()

	if err != nil {
		log.Fatal(err)
	}
}
