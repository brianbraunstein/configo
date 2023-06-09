#!/usr/bin/env bash
set -euo pipefail

staging_dir="$(mktemp -dp $RULEDIR)"

# Build up directory.
mkdir -p $staging_dir/configo/DEBIAN/
envsubst < packaging/deb/deb_control.envsubst > "$staging_dir/configo/DEBIAN/control"
mkdir -p $staging_dir/configo/usr/local/bin/
cp $cli_path $staging_dir/configo/usr/local/bin/configo

# Build the deb package from the directory.
(cd "$staging_dir"; dpkg-deb --build --root-owner-group ./configo)
mv "$staging_dir/configo.deb" "$OUT_FILE"

rm -rf "$staging_dir"

