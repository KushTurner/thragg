package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const configFile = ".thragg.json"

type Repo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Config struct {
	Project string `json:"project"`
	Repos   []Repo `json:"repos"`
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		cmdInit()
	case "apply":
		cmdApply()
	case "list":
		cmdList()
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Usage: thragg <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init     Set up thragg for the current project")
	fmt.Println("  apply    Copy CLAUDE.md and .claude/settings.json into each configured repo")
	fmt.Println("  list     Show configured repos for this project")
}

func cmdInit() {
	if _, err := os.Stat(configFile); err == nil {
		fmt.Print(".thragg.json already exists. Overwrite? [y/N] ")
		var confirm string
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) != "y" {
			return
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Setting up thragg for this project.")
	fmt.Println()
	fmt.Print("Project name: ")
	scanner.Scan()
	project := strings.TrimSpace(scanner.Text())

	var repos []Repo
	for {
		fmt.Println()
		fmt.Print("Repo name (leave blank to finish): ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())
		if name == "" {
			break
		}

		fmt.Printf("Path to %s: ", name)
		scanner.Scan()
		path := strings.TrimSpace(scanner.Text())
		path = expandHome(path)

		repos = append(repos, Repo{Name: name, Path: path})
	}

	if len(repos) == 0 {
		fmt.Println("No repos added. Exiting.")
		os.Exit(1)
	}

	cfg := Config{Project: project, Repos: repos}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "error writing config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("Created %s\n", configFile)
}

func cmdApply() {
	cfg := loadConfig()
	thraggDir := thraggRoot()

	for _, repo := range cfg.Repos {
		path := expandHome(repo.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("Warning: %s does not exist, skipping.\n", path)
			continue
		}

		fmt.Printf("Applying to %s...\n", path)

		copyFile(
			filepath.Join(thraggDir, "CLAUDE.md"),
			filepath.Join(path, "CLAUDE.md"),
		)

		if err := os.MkdirAll(filepath.Join(path, ".claude"), 0755); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		copyFile(
			filepath.Join(thraggDir, ".claude", "settings.json"),
			filepath.Join(path, ".claude", "settings.json"),
		)

		copyDir(
			filepath.Join(thraggDir, "rules"),
			filepath.Join(path, "rules"),
		)
	}

	fmt.Println()
	fmt.Println("Done.")
}

func cmdList() {
	cfg := loadConfig()
	fmt.Printf("Project: %s\n\n", cfg.Project)
	for _, r := range cfg.Repos {
		fmt.Printf("  %s: %s\n", r.Name, r.Path)
	}
}

func loadConfig() Config {
	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "No .thragg.json found. Run: thragg init")
		os.Exit(1)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing config: %v\n", err)
		os.Exit(1)
	}
	return cfg
}

func thraggRoot() string {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving executable: %v\n", err)
		os.Exit(1)
	}
	// Resolve symlinks so we get the real path
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving symlink: %v\n", err)
		os.Exit(1)
	}
	// bin/thragg -> bin -> thragg root
	return filepath.Dir(filepath.Dir(exe))
}

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}

func copyFile(src, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", src, err)
		os.Exit(1)
	}
	if err := os.WriteFile(dst, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "error writing %s: %v\n", dst, err)
		os.Exit(1)
	}
}

func copyDir(src, dst string) {
	if err := os.MkdirAll(dst, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "error creating %s: %v\n", dst, err)
		os.Exit(1)
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", src, err)
		os.Exit(1)
	}
	for _, entry := range entries {
		copyFile(
			filepath.Join(src, entry.Name()),
			filepath.Join(dst, entry.Name()),
		)
	}
}
