package errors

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// ============================================================
// WRAPPING ERRORS — always use %w, never %v or %s
// ============================================================

// BAD: breaks the error chain — errors.Is / errors.As won't work
func wrapBad(err error) error {
	return fmt.Errorf("something failed: %v", err) // %v loses the chain
}

// GOOD: preserves the error chain
func wrapGood(err error) error {
	return fmt.Errorf("something failed: %w", err) // %w keeps the chain
}

// ============================================================
// NEVER IGNORE ERRORS
// ============================================================

func ignoreExample() {
	// BAD
	_ = os.Remove("file.txt")

	// GOOD
	if err := os.Remove("file.txt"); err != nil {
		// handle or propagate
	}
}

// ============================================================
// LOG AND RETURN — anti-pattern, handle at one layer only
// ============================================================

// BAD: logs and returns — the same error appears in logs multiple times
func readFileBad(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println("readFile failed:", err) // log here
		return "", err                       // AND return — caller will log again
	}
	return string(data), nil
}

// GOOD: wrap with context, let the top level handle it once
func readFileGood(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", filename, err)
	}
	return string(data), nil
}

// ============================================================
// CHECKING ERRORS — errors.Is / errors.As, not ==
// ============================================================

var ErrNotFound = errors.New("not found")

func checkErrors(err error) {
	// BAD: direct comparison breaks with wrapped errors
	if err == ErrNotFound {
		// won't match if err is fmt.Errorf("looking up user: %w", ErrNotFound)
	}

	// GOOD: errors.Is unwraps the chain automatically
	if errors.Is(err, ErrNotFound) {
		// works even if ErrNotFound is buried in a chain
	}
}

// ============================================================
// CUSTOM ERROR TYPES — implement Unwrap() to support the chain
// ============================================================

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

// Extracting a custom error type from the chain
func handleValidation(err error) {
	var ve ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("field %s: %s\n", ve.Field, ve.Message)
	}
}

// ============================================================
// NO PANIC IN APPLICATION CODE — return errors instead
// ============================================================

// BAD: callers cannot recover, deferred functions are skipped
func parseBad(data []byte) Result {
	result, err := parse(data)
	if err != nil {
		panic(err)
	}
	return result
}

// GOOD: return the error and let the caller decide
func parseGood(data []byte) (Result, error) {
	result, err := parse(data)
	if err != nil {
		return Result{}, fmt.Errorf("parsing data: %w", err)
	}
	return result, nil
}

// ============================================================
// NO log.Fatal OUTSIDE main — skips deferred functions
// ============================================================

// BAD
func helperBad() {
	if err := doSomething(); err != nil {
		log.Fatal(err) // exits immediately, no cleanup
	}
}

// GOOD
func helperGood() error {
	if err := doSomething(); err != nil {
		return fmt.Errorf("doing something: %w", err)
	}
	return nil
}

// Stubs to make the file compile as a reference
type Result struct{}

func parse(_ []byte) (Result, error)  { return Result{}, nil }
func doSomething() error              { return nil }
