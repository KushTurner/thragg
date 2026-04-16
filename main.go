package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "apply":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: thragg apply <repo-path> [repo-path ...]")
			os.Exit(1)
		}
		cmdApply(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Usage: thragg <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  apply <repo-path> [repo-path ...]   Copy CLAUDE.md, settings.json, and rules/ into each repo")
}

func cmdApply(paths []string) {
	thraggDir := thraggRoot()

	for _, path := range paths {
		path = expandHome(path)
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

func thraggRoot() string {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving executable: %v\n", err)
		os.Exit(1)
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving symlink: %v\n", err)
		os.Exit(1)
	}
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
