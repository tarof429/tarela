package main

import (
	"flag"
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

func removeFiles(files []os.DirEntry, remove int) {

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
func getBackupPathName(s string) string {
	now := time.Now()

	outputFileName := fmt.Sprintf("backup_%v.sfs", now.Format("200602010304"))

	return filepath.Join(s, outputFileName)
}

// Generate the command that will be used to create the backup
// It is hardcoded to use mksquashfs, but the output file can be
// whatever we like, useful for testing.
func getBackupCommand(input, outputFilePath, exclude string) (string, []string) {

	args := []string{}

	args = append(args, input)
	args = append(args, outputFilePath)

	if len(exclude) > 0 {
		args = append(args, "-ef")
		args = append(args, exclude)
	}

	return BACKUP_COMMAND, args
}

// Backup the directory to a tar file
func backup(input, outputFilePath, exclude string) {

	fmt.Printf("Continue with backup file creation? (y/N): ")

	var choice string

	fmt.Scanf("%v", &choice)

	if choice == "y" {
		fmt.Printf("Creating %v\n", outputFilePath)

		cmd, args := getBackupCommand(input, outputFilePath, exclude)

		_, err := exec.Command(cmd, args...).Output()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	input := flag.String("input", "", "Input directory")
	output := flag.String("output", "", "Output file")
	keep := flag.Int("keep", 0, "Number of files to keep")
	exclude := flag.String("exclude", "", "Exclude file")

	flag.Parse()

	if len(*input) == 0 || len(*output) == 0 || *keep < 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Make sure we can read the input directory
	if _, err := os.ReadDir(*input); err != nil {
		log.Fatal(err)
	}

	// Make sure we can read the output directory
	outputPath := filepath.Dir(*output)

	files, err := os.ReadDir(outputPath)

	if err != nil {
		log.Fatal(err)
	}

	backupPathName := getBackupPathName(outputPath)

	if err := os.Chdir(outputPath); err != nil {
		log.Fatal(err)
	}

	// Cleanup the backup directory (if needed)
	if *keep > 0 && len(files) >= *keep {
		var choice string

		remove := len(files) - *keep + 1

		fmt.Printf("%v files in %v will be removed. Continue? (y/N): ", remove, outputPath)
		fmt.Scanf("%v", &choice)

		switch choice {
		case "y":
			fmt.Println("Removing files")
			removeFiles(files, remove)
		default:
			fmt.Println("Aborting")
			os.Exit(0)
		}
	} else {
		fmt.Println("No files need to be removed")
	}

	// Perform the backup. This is independent of any cleanup we did previously.
	backup(*input, backupPathName, *exclude)

}
