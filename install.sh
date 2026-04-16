#!/usr/bin/env bash
set -e

INSTALL_DIR="/usr/local/bin"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Building thragg..."
cd "$SCRIPT_DIR"
go build -o bin/thragg .

echo "Installing..."
ln -sf "$SCRIPT_DIR/bin/thragg" "$INSTALL_DIR/thragg"

echo ""
echo "Done."
echo ""
echo "Next steps:"
echo "  cd into a repo and run 'thragg init' to copy CLAUDE.md and settings into it"
