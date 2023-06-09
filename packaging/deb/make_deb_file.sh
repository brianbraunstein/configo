#!/usr/bin/env bash
set -euo pipefail

staging_dir="$(mktemp -dp $RULEDIR)"

cp -R packaging/deb/archive_template "$staging_dir/configo"
mkdir -p $staging_dir/configo/usr/local/bin/
cp $cli_path $staging_dir/configo/usr/local/bin/configo
(cd "$staging_dir"; dpkg-deb --build --root-owner-group ./configo)
mv "$staging_dir/configo.deb" "$OUT_FILE"

rm -rf "$staging_dir"

