# 📋 gogcli Project Roadmap & TODO

This document tracks the immediate steps required to fix the codebase and get the project fully "Live," as well as long-term goals for future development.

---

## 🚀 Phase 1: Immediate Fixes (Get it working)

Currently, the project has import mismatches and dependency issues. Follow these steps to resolve them.

- [ ] **Fix Module Identity**: Update `go.mod` to use the correct module path `github.com/jibankumarpanda/gogcli.1`.
- [ ] **Bulk Update Imports**: Systematically replace `github.com/jibankumarpanda/gogcli/` with `github.com/jibankumarpanda/gogcli.1/` across all `.go` files.
- [ ] **Repair Makefile**: Update the `Makefile` to use the correct module paths for `LDFLAGS` and formatting tools.
- [ ] **Dependency Tidy**: Run `go mod tidy` to resolve all external dependencies (Google APIs, Kong, etc.).
- [ ] **Pass Local Tests**: Ensure `go test ./...` passes locally.
- [ ] **Fix GitHub Actions**: Resolve the failing `ci/test(push)` job by pushing the corrected import paths.

---

## 🛠️ Phase 2: Launch & Google Setup (Live Operations)

Once the code builds, you must connect it to Google.

- [ ] **Create Google Cloud Project**: Initialize a project in the [Google Cloud Console](https://console.cloud.google.com/).
- [ ] **Enable Required APIs**:
    - [ ] Gmail API
    - [ ] Google Calendar API
    - [ ] Google Drive API
    - [ ] People API (for Contacts)
- [ ] **Configure OAuth Consent**:
    - [ ] Set "User Type" to External.
    - [ ] **Important**: Add your own email to the "Test Users" list.
- [ ] **Generate Credentials**:
    - [ ] Create "OAuth Client ID" (Type: Desktop App).
    - [ ] Download the secret JSON file.
- [ ] **Initial Auth**:
    - [ ] Run `gog auth credentials <path-to-json>` once to register the app.
    - [ ] Run `gog auth add <your-email>` to log in for the first time.

---

## ⚙️ Phase 3: Workflow & CI/CD Stability

- [ ] **Headless Auth for CI**: Investigate using `GOG_KEYRING_PASSWORD` to allow automated tests to run without interactive prompts.
- [ ] **Automatic Releases**: Configure a GitHub Secret for `GITHUB_TOKEN` to allow `release.yml` to automatically publish binaries when you push a version tag (e.g., `v1.0.0`).
- [ ] **Documentation Audit**: Update all internal docs in `/docs` to reflect the `.1` module name.

---

## 🔮 Phase 4: Future Plans (The Vision)

- [ ] **Extended Service Support**:
    - [ ] **Google Photos**: Add commands to list and download photos.
    - [ ] **YouTube**: Basic channel/video management.
    - [ ] **Google Keep**: Improve support for non-Workspace users if possible.
- [ ] **Enhanced UI/UX**:
    - [ ] Add interactive prompts for complex commands using `bubbletea`.
    - [ ] Progress bars for large Drive uploads/downloads.
- [ ] **Plugin System**: Allow users to write custom scripts or "agents" that use `gogcli` as a backend.
- [ ] **Cross-Platform Installer**: Create an installer for Windows/macOS/Linux instead of just raw binaries.
- [ ] **Email Tracking Dashboard**: Build a small local UI to view tracking stats from the Cloudflare Worker.

---

## 📝 Notes
*   Always run `make fmt` before committing.
*   Keep `credentials.json` secret and never commit it to Git.
*   Use `gog auth list --check` periodically to verify token health.
