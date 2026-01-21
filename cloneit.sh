#!/usr/bin/env bash
set -e

REPOS=$(gh repo list \
  --limit 300 \
  --json name,sshUrl,description \
  --jq '.[] | "\(.name)\t\(.sshUrl)\t\(.description // "")"')

SELECTED=$(echo "$REPOS" | fzf --with-nth=1,3)

[ -z "$SELECTED" ] && exit 0

URL=$(echo "$SELECTED" | cut -f2)

git clone "$URL"
