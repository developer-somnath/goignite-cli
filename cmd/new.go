package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project name]",
	Short: "Create a new GoIgnite project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		createProjectStructure(projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createProjectStructure(projectName string) {
	folders := []string{
		"app/controllers",
		"app/models",
		"config",
		"routes",
		"public",
		"resources/views",
	}

	err := os.Mkdir(projectName, 0755)
	if err != nil {
		log.Fatalf("Could not create project folder: %v", err)
	}

	fmt.Printf("Creating new GoIgnite project: %s\n", projectName)

	for _, folder := range folders {
		path := filepath.Join(projectName, folder)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("Failed to create folder %s: %v", path, err)
		}
	}

	createFile(filepath.Join(projectName, "main.go"), mainGoContent(projectName))
	fmt.Println("Project successfully created!")
}

func createFile(path, content string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", path, err)
	}
	defer f.Close()
	f.WriteString(content)
}

func mainGoContent(projectName string) string {
	return fmt.Sprintf(`package main

	import "fmt"

	func main() {
		fmt.Println("Welcome to %s!")
	}
	`, projectName)
}
