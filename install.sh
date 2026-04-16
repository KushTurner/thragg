#!/usr/bin/env bash
set -e

INSTALL_DIR="/usr/local/bin"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Installing thragg..."

chmod +x "$SCRIPT_DIR/bin/thragg"
ln -sf "$SCRIPT_DIR/bin/thragg" "$INSTALL_DIR/thragg"

echo "Done. Run 'thragg init' inside any project to get started."
