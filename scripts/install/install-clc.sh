#!/bin/bash

set -e

REPO="estifanos-neway/CLC"
VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
INSTALL_DIR="$HOME/.local/bin"
TMP_DIR="$(mktemp -d)"
ARCH=$(uname -m)
OS=$(uname -s)

# Normalize arch
case "$ARCH" in
x86_64) ARCH="x86_64" ;;
arm64 | aarch64) ARCH="arm64" ;;
i386 | i686) ARCH="i386" ;;
*) echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac

# Normalize OS
case "$OS" in
Darwin) PLATFORM="Darwin" ;;
Linux) PLATFORM="Linux" ;;
*) echo "Unsupported OS: $OS" && exit 1 ;;
esac

FILENAME="CLC_${PLATFORM}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILENAME"

echo "Downloading $FILENAME..."
echo "Downloading $URL..."
curl -L "$URL" -o "$TMP_DIR/clc.tar.gz"

echo "Extracting..."
mkdir -p "$TMP_DIR/extract"
tar -xzf "$TMP_DIR/clc.tar.gz" -C "$TMP_DIR/extract"

echo "Installing to $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
mv "$TMP_DIR/extract/CLC" "$INSTALL_DIR/CLC"
chmod +x "$INSTALL_DIR/CLC"

if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
  echo "Adding $INSTALL_DIR to PATH in your shell profile..."
  SHELL_RC="$HOME/.bashrc"
  [[ $SHELL == *zsh ]] && SHELL_RC="$HOME/.zshrc"
  echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >>"$SHELL_RC"
  echo "-> Added to $SHELL_RC"
fi

echo "Done! You may need to run 'source ~/.bashrc' or open a new terminal for CLC to be available."