# Copyright 2025 Canonical Ltd.
# SPDX-License-Identifier: AGPL-3.0

BIN="/usr/local/bin"
VERSION="v1.50.0" # last release as of 2025-02-11

if [ "$1" = "--latest" ]; then
  VERSION=$(curl -s "https://api.github.com/repos/bufbuild/buf/releases/latest" | jq -r '.name')
fi

echo "Attempting to install buf version $VERSION"

curl -sSL "https://github.com/bufbuild/buf/releases/download/${VERSION}/buf-$(uname -s)-$(uname -m)" -o "${BIN}/buf" && chmod +x "${BIN}/buf"

if [ $? ]; then
  echo "Installation succedeed"
else
  echo "Something when wrong"
  exit 1
fi
