# gogcli Project Guide

This document explains how this project is structured, how the workflow works, and lists useful commands for testing and development.

---

## 1. Project Purpose
**gogcli** is a Command Line Interface (CLI) built in Go that allows you to interact with Google Workspace (Gmail, Drive, Calendar, etc.) directly from your terminal.

## 2. Folder Structure
- **`.github/workflows/`**: Automation scripts for Testing (CI) and Publishing (Release).
- **`cmd/gog/`**: The main entry point. `main.go` starts the application.
- **`docs/`**: The website source code (Landing page, CSS, and documentation).
- **`internal/`** (Currently Missing): The core logic and "engine" of the tool.

---

## 3. Current Workflow (Web & Docs)

If you are updating the website or documentation:

1. **Edit**: Modify files in the `docs/` folder (e.g., `index.html`).
2. **Local Test**: Open `docs/index.html` in your browser to see changes.
3. **Commit & Push**:
   ```powershell
   git add .
   git commit -m "Brief description of your change"
   git push origin main
   ```
4. **Deploy**: GitHub will automatically update your live site within 1-2 minutes.

---

## 4. Useful Terminal Commands (Testing)

These are the types of commands used to investigate and verify the project state.

### File & Directory Investigation
- **List all files**: `git ls-files`
- **Search for a specific file**: `Get-ChildItem -Filter "index.html" -Recurse`
- **Check current Git status**: `git status`

### Searching for Text
- **Find a string in the docs**: `Select-String -Pattern "github.com" -Path .\docs\*.html`
- **Find all Go imports**: `Select-String -Pattern "import" -Path .\cmd\gog\*.go`

### Diagnostics & Research (The "Antigravity" Commands)
These are specific commands I used to help you manage the project today:

- **List every file in the project (Hidden & Visible)**:
  `cmd /c "dir /b /s"`
- **List just the folders to see structure**:
  `cmd /c "dir /b /ad"`
- **Find every line where "github.com" is mentioned (to check URLs)**:
  `grep -r "github.com" . --exclude-dir=.git`
- **Find specifically where your username is used**:
  `grep -r "jibankumarpanda" . --exclude-dir=.git`
- **Verify if a specific file exists anywhere**:
  `cmd /c "dir /s /b go.mod"`
- **See the commit history with file names changed**:
  `git log -n 5 --oneline --name-only`
- **Compare changes between two specific versions (commits)**:
  `git diff HEAD~1 HEAD --summary`

### Repository Management
- **Check Remote URL**: `git remote -v`
- **Show recent commits**: `git log --oneline -n 10`

---

## 5. Future Go Commands (When logic is restored)
When the full source code is present, you will use these:
- **Run Tests**: `go test ./...`
- **Build the App**: `go build -o gog.exe ./cmd/gog`
- **Run the App**: `./gog.exe --help`
- **Format Code**: `go fmt ./...`

---

## 6. Important Links
- **GitHub Repository**: [jibankumarpanda/gogcli.1](https://github.com/jibankumarpanda/gogcli.1)
- **Live Documentation**: [gogcli.sh](https://gogcli.sh/) (or your GitHub Pages link)
