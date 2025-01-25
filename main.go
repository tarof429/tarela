package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/theckman/yacspin"

	"tarela/tarela"
)

const (
	INPUT_FLAG             string = "Input directory"
	OUTPUT_FLAG            string = "Output file"
	KEEP_FLAG              string = "Number of files to keep"
	EXCLUDE_FLAG           string = "Exclude file"
	COMPRESSION_TYPE       string = "What compression to use"
	SPINNER_MESSAGE        string = "Running backup"
	SPINNER_GRAPHIC        int    = 35
	SPINNER_STOP_CHARACTER string = "✓"
	SPINNER_STOP_COLOR     string = "fgGreen"
)

func main() {
	input := flag.String("input", "", INPUT_FLAG)
	output := flag.String("output", "", OUTPUT_FLAG)
	keep := flag.Int("keep", 0, KEEP_FLAG)
	exclude := flag.String("exclude", "", EXCLUDE_FLAG)
	comp := flag.String("comp", "sfs", COMPRESSION_TYPE)

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

	var backupPathName string

	if *comp == "sfs" {
		backupPathName = tarela.GetSquashfsBackupPathName(outputPath)
	} else if *comp == "tar" {
		backupPathName = tarela.GetTarBackupPathName(outputPath)
	} else {
		flag.Usage()
		os.Exit(1)
	}

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
			tarela.RemoveFiles(files, remove)
		default:
			fmt.Println("Aborting")
			os.Exit(0)
		}
	} else {
		fmt.Println("No files need to be removed")
	}

	fmt.Printf("Continue with backup? (y/N): ")

	var choice string

	fmt.Scanf("%v", &choice)

	if choice == "y" {

		// Perform the backup. This is independent of any cleanup we did previously.
		cfg := yacspin.Config{
			Frequency:       200 * time.Millisecond,
			CharSet:         yacspin.CharSets[SPINNER_GRAPHIC],
			Suffix:          " ",
			SuffixAutoColon: true,
			Message:         SPINNER_MESSAGE,
			StopCharacter:   SPINNER_STOP_CHARACTER,
			StopColors:      []string{SPINNER_STOP_COLOR},
		}
		var spinner *yacspin.Spinner

		spinner, err = yacspin.New(cfg)

		if err != nil {
			os.Exit(1)
		}

		err = spinner.Start()

		if err != nil {
			os.Exit(1)
		}

		spinner.Message(fmt.Sprintf("Creating %v...", backupPathName))

		// Create the backup
		if *comp == "sfs" {
			tarela.BackupSquashfs(*input, backupPathName, *exclude)
		} else if *comp == "tar" {
			tarela.BackupTar(*input, backupPathName)
		} else {
			os.Exit(1)
		}
		err = spinner.Stop()

		if err != nil {
			os.Exit(1)
		}
	}

}
