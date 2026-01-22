# CloneIt

A cross-platform CLI tool to interactively clone GitHub repositories using the GitHub CLI (`gh`) and `fzf` for selection.

## Features

- Lists all your GitHub repositories (public and private, if authenticated)
- Interactive selection with fuzzy search
- Clones the selected repository via SSH

## Prerequisites

- [GitHub CLI (`gh`)](https://cli.github.com/) installed and authenticated
- [`fzf`](https://github.com/junegunn/fzf) installed for interactive selection
- `jq` for JSON processing

## Usage

1. Ensure you are authenticated with GitHub CLI: `gh auth login`
2. Run: `cloneit`
3. Use fuzzy search to select a repository
4. The tool will clone the selected repository to the current directory

## Installation

### Prerequisites

- [GitHub CLI (`gh`)](https://cli.github.com/) installed and authenticated
- [`fzf`](https://github.com/junegunn/fzf) installed for interactive selection
- `jq` for JSON processing (for install script)

### Option 1: Auto-install (Recommended)

Run the install script:

```bash
curl -fsSL https://raw.githubusercontent.com/ericklopezdev/cloneit/main/install.sh | bash
```

This will download the latest release binary for your platform and install it to `/usr/local/bin` or `~/.local/bin`.

### Option 2: Manual Install

1. Go to [Releases](https://github.com/YOUR_USERNAME/cloneit/releases) and download the appropriate archive for your OS and architecture.
2. Extract the archive.
3. Move the `cloneit` binary to a directory in your PATH (e.g., `/usr/local/bin` or `~/bin`).

### For Windows Users

Download the `.zip` file from releases, extract, and add `cloneit.exe` to your PATH.

![](/screenshots/test.png)
