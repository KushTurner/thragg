#!/usr/bin/env bash
set -e

INSTALL_DIR="/usr/local/bin"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Installing thragg CLI..."

chmod +x "$SCRIPT_DIR/bin/thragg"
ln -sf "$SCRIPT_DIR/bin/thragg" "$INSTALL_DIR/thragg"

echo ""
echo "Done."
echo ""
echo "Next steps:"
echo "  1. cd into a project repo"
echo "  2. Run 'thragg init' to configure it"
echo "  3. Run 'thragg apply' to copy CLAUDE.md and settings into the repo"
