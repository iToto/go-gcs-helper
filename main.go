package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	PROJECT            = "ENV_GCS_PROJECT"
	SNAPSHOT_BASE_NAME = "ENV_GCS_SNAPSHOT_BASE_NAME"
	VOLUME             = "ENV_GCS_VOLUME"
	ZONE               = "ENV_GCS_ZONE"
)

func main() {
	var validCommands = [...]string{"snapshot"}

	if len(os.Args) != 2 {
		err := fmt.Errorf("Usage: go-gcs-helper [COMMAND]")
		log.Printf(err.Error())
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "snapshot":
		createSnapshot()
	default:
		err := fmt.Errorf(
			"invalid command. List of valid commands: %v",
			validCommands,
		)
		log.Printf(err.Error())
		os.Exit(1)
	}
}

func createSnapshot() {
	// Validate ENV vars
	if "" == os.Getenv(VOLUME) ||
		"" == os.Getenv(PROJECT) ||
		"" == os.Getenv(ZONE) ||
		"" == os.Getenv(SNAPSHOT_BASE_NAME) {
		fmt.Println("Missing Env Vars!")
		os.Exit(1)
	}

	start := time.Now()
	snapshotName := fmt.Sprintf(
		"%s-%d",
		os.Getenv(SNAPSHOT_BASE_NAME),
		start.Unix(),
	)
	arg1 := "beta"
	arg2 := "compute"
	arg3 := fmt.Sprintf("--project=%s", os.Getenv(PROJECT))
	arg4 := "disks"
	arg5 := "snapshot"
	arg6 := os.Getenv(VOLUME)
	arg7 := fmt.Sprintf("--zone=%s", os.Getenv(ZONE))
	arg8 := fmt.Sprintf("--snapshot-names=%s", snapshotName)
	arg9 := "--storage-location=us"

	cmd := exec.Command(
		"gcloud",
		arg1,
		arg2,
		arg3,
		arg4,
		arg5,
		arg6,
		arg7,
		arg8,
		arg9,
	)

	log.Printf("Running GCS Snapshot command and waiting for it to finish...")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Command finished with error: %v", err)
		fmt.Printf("Output: %s", string(out))
	} else {
		log.Printf("Snapshot: %s successfully created!", snapshotName)
	}
}
