# CLC - Command Line Chat

**CLC** is a command-line tool that uses generative AI to help you automate common system-level tasks by describing your intent in natural language.

---

## ✨ Features

- Accepts natural-language prompts (e.g., "increase brightness", "create a zip of all PDFs")
- Detects your OS and architecture automatically
- Friendly error handling and output
- Lightweight and fast

---

## 🛠️ Installation

### 📦 macOS & Linux

Open a terminal and run:

```bash
bash <(curl -sSL https://raw.githubusercontent.com/estifanos-neway/CLC/main/scripts/install/install-clc.sh)
````

> This script will:
>
> * Detect your architecture (x86\_64 or arm64)
> * Download the correct release binary
> * Extract and move the `CLC` binary to `~/.local/bin`
> * Add it to your shell's `PATH` if necessary

**Note:** You may need to start a new terminal session or run `source ~/.profile`, `source ~/.bashrc`, or `source ~/.zshrc` depending on your shell, for the CLC command to become available.

### 🪟 Windows (PowerShell)

Open PowerShell **as Administrator** and run:

```powershell
iwr -useb https://raw.githubusercontent.com/estifanos-neway/CLC/main/scripts/install/install-clc.ps1 | iex
```

> This script will:
>
> * Detect your CPU architecture (x86, x64, arm64)
> * Download and unzip the correct release
> * Place `CLC.exe` in your local programs folder
> * Add it to your PATH if not already present

**Note:** You may need to restart your terminal or system for the PATH changes to take effect.

---

## 🚀 Usage

Once installed, run `CLC` followed by a natural-language prompt:

```bash
CLC create a backup of all .txt files in this folder
```

Available flags:

* `-s`: Skip executing the generated script (preview only)
* `-k`: Keep the generated script file

Example:

```bash
CLC -s "restart my wifi adapter"
```

---

## 🧱 Structure

The core logic lives in:

* `main.go`: CLI entry point and flag handling
* `internal/api/clc`: Prompt handling and response parsing
* `internal/api/gemini`: Gemini API integration
* `scripts/install/`: Cross-platform install scripts

---

## 📦 Releases

Pre-built binaries are available for:

* **macOS**: `CLC_Darwin_arm64.tar.gz`, `CLC_Darwin_x86_64.tar.gz`
* **Linux**: `CLC_Linux_arm64.tar.gz`, `CLC_Linux_i386.tar.gz`, `CLC_Linux_x86_64.tar.gz`
* **Windows**: `CLC_Windows_arm64.zip`, `CLC_Windows_i386.zip`, `CLC_Windows_x86_64.zip`

Find them on the [Releases page](https://github.com/estifanos-neway/CLC/releases).

---

## 🧪 Example Prompts

* `CLC increase screen brightness`
* `CLC create a tar.gz of all PNGs`
* `CLC kill any process using port 8080`
* `CLC set system volume to 50%`

---

## 🧩 Requirements

* No dependencies (compiled Go binary)
* Optional: `curl`, `tar`, or `powershell` for installation

---

## 🙋 Support

Found a bug? Have a feature request? Open an issue or PR on [GitHub](https://github.com/estifanos-neway/CLC/issues).

---

## 📝 License

MIT License