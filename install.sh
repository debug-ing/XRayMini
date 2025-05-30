#!/bin/bash

set -e

# Detect system architecture
detect_arch() {
  ARCH=$(uname -m)
  case "$ARCH" in
    x86_64) echo "amd64" ;;
    aarch64) echo "arm64" ;;
    armv7l) echo "armv7" ;;
    *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
  esac
}

# Install xray-core
install_xray() {
  echo "[*] Installing xray-core..."
  ARCH=$(detect_arch)
  TMP_DIR=$(mktemp -d)
  cd "$TMP_DIR"
  curl -s https://api.github.com/repos/XTLS/Xray-core/releases/latest | grep "browser_download_url" | grep "linux-amd64.zip" | cut -d '"' -f 4
  LATEST=$(curl -s https://api.github.com/repos/XTLS/Xray-core/releases/latest \
    | grep "browser_download_url" \
    | grep "linux-${ARCH}.zip" \
    | cut -d '"' -f 4)

  wget "$LATEST" -O xray.zip
  unzip xray.zip xray
  install xray /usr/local/bin/xray
  chmod +x /usr/local/bin/xray
  echo "[+] xray-core installed to /usr/local/bin/xray"

  cd /
  rm -rf "$TMP_DIR"
}

# Install tun2proxy
install_tun2proxy() {
  sudo apt install -y unzip wget
  wget https://github.com/tun2proxy/tun2proxy/releases/latest/download/tun2proxy-x86_64-unknown-linux-musl.zip
  unzip tun2proxy-x86_64-unknown-linux-musl.zip
  sudo mv ./tun2proxy-bin  /usr/local/bin/
}

# Main entry
echo "=== Installing xray-core and tun2proxy ==="
bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install
install_tun2proxy
echo "=== Installation completed ==="
