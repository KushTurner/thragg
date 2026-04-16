package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		cmdInit()
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Usage: thragg <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init   Copy CLAUDE.md, settings.json, and rules/ into the current directory")
}

func cmdInit() {
	thraggDir := thraggRoot()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting current directory: %v\n", err)
		os.Exit(1)
	}

	copyFile(
		filepath.Join(thraggDir, "CLAUDE.md"),
		filepath.Join(cwd, "CLAUDE.md"),
	)

	if err := os.MkdirAll(filepath.Join(cwd, ".claude"), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	copyFile(
		filepath.Join(thraggDir, ".claude", "settings.json"),
		filepath.Join(cwd, ".claude", "settings.json"),
	)

	copyDir(
		filepath.Join(thraggDir, "rules"),
		filepath.Join(cwd, "rules"),
	)

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
