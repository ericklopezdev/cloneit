#!/bin/bash
set -e

echo "Installing CloneIt..."

# Detect OS and arch
os=$(uname -s | tr '[:upper:]' '[:lower:]')
case $os in
linux) os="linux" ;;
darwin) os="darwin" ;;
mingw* | msys*) os="windows" ;;
*)
  echo "Unsupported OS: $os"
  exit 1
  ;;
esac

arch=$(uname -m)
case $arch in
x86_64) arch="x86_64" ;;
arm64 | aarch64) arch="arm64" ;;
*)
  echo "Unsupported arch: $arch"
  exit 1
  ;;
esac

# Get latest release
repo="ericklopezdev/cloneit" # Replace with actual owner/repo
api_url="https://api.github.com/repos/$repo/releases/latest"
release_data=$(curl -s $api_url)

# Find asset URL
if [ "$os" = "windows" ]; then
  pattern="cloneit_windows_${arch}.zip"
  ext="zip"
else
  pattern="cloneit_${os}_${arch}.tar.gz"
  ext="tar.gz"
fi

asset_url=$(echo "$release_data" | jq -r ".assets[] | select(.name | test(\"$pattern\")) | .browser_download_url")

if [ -z "$asset_url" ]; then
  echo "No matching asset found for $os $arch"
  exit 1
fi

# Download and extract
temp_dir=$(mktemp -d)
cd "$temp_dir"

echo "Downloading $asset_url"
curl -L -o archive.$ext "$asset_url"

if [ "$ext" = "zip" ]; then
  unzip archive.zip
else
  tar xzf archive.tar.gz
fi

# Find binary
binary=$(find . -name "cloneit*" -type f -executable | head -1)
if [ -z "$binary" ]; then
  echo "Binary not found"
  exit 1
fi

# Install
install_dir="/usr/local/bin"
if ! sudo -n true 2>/dev/null; then
  install_dir="$HOME/.local/bin"
  mkdir -p "$install_dir"
fi

echo "Installing to $install_dir"
if [ "$install_dir" = "/usr/local/bin" ]; then
  sudo cp "$binary" "$install_dir/cloneit"
else
  cp "$binary" "$install_dir/cloneit"
fi

# Add to PATH if needed
if ! command -v cloneit >/dev/null 2>&1; then
  echo "Add $install_dir to your PATH"
  echo "export PATH=\"$install_dir:\$PATH\"" >>~/.bashrc
  echo "Restart your shell or run: source ~/.bashrc"
fi

echo "CloneIt installed successfully!"
echo "Make sure gh and fzf are installed."
echo "Run: cloneit"
