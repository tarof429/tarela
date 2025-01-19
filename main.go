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

	backupPathName := tarela.GetBackupPathName(outputPath)

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
			CharSet:         yacspin.CharSets[tarela.SPINNER_GRAPHIC],
			Suffix:          " ",
			SuffixAutoColon: true,
			Message:         "Running backup",
			StopCharacter:   "✓",
			StopColors:      []string{"fgGreen"},
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
		tarela.Backup(*input, backupPathName, *exclude)

		err = spinner.Stop()

		if err != nil {
			os.Exit(1)
		}
	}

}
