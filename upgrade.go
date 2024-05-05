package main

import (
	"fmt"
	"os"
	"os/exec"
)

func installBinary(filepath string) error {
	// Rename the current executable
	currentExec, err := os.Executable()
	if err != nil {
		return err
	}

	backupPath := currentExec + ".bak"
	if err := os.Rename(currentExec, backupPath); err != nil {
		return err
	}

	// Replace the current executable with the new one
	if err := os.Rename(filepath, currentExec); err != nil {
		// Restore the original executable if there's an error
		os.Rename(backupPath, currentExec)
		return err
	}

	// Delete the backup file
	os.Remove(backupPath)

	// Get the path to the currently running executable
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Start a new instance of the application
	cmd := exec.Command(executablePath, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting process:", err)
		return err
	}

	// Exit the current instance
	os.Exit(0)

	return nil
}
