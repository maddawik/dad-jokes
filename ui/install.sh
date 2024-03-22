#!/bin/bash

VERSION=$(gum choose linux-arm-64 linux-armv7 linux-x64 macos-arm64 macos-x64 windows-arm64.exe windows-x64.exe)

echo "Downloding $VERSION tailwindcli binary"
curl -OL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$VERSION
chmod +x tailwindcss-$VERSION
mv tailwindcss-$VERSION tailwindcss
