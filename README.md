# OrionAI — AI Desktop Assistant & Chatbot

Clipr is a cross-platform desktop assistant and chatbot powered by **Go**, **Svelte**, and **Wails**. It integrates the lightweight **Phi-3 Mini** model to provide fast, local AI responses and assistant functionalities.

Go is used to handle **prompt engineering**, **file handling**, and **system-level commands**, while **Svelte** ensures a reactive, efficient, and beautiful GUI. Wails bridges the frontend and backend, compiling everything into a native desktop application.

---

## ⚙️ Technologies Used

- **Go** – Core logic, file I/O, and system integration
- **Svelte** – Lightweight and reactive frontend framework
- **Wails** – Native desktop app framework connecting Go and Svelte
- **Phi-3 Mini** – Embedded AI language model for chatbot responses

---

## ✅ Features

- 🔌 **Prompt Engineering Backend** – Handled in Go for flexible and safe local AI interactions
- 💬 **Chatbot Interface** – Converse with Phi-3 Mini directly from your desktop
- 💾 **Persistent Configuration** – Save and load user settings using JSON
- 🗂️ **Cross-Platform File Handling** – Automatic path detection for config storage (Linux & Windows)
- 🎨 **Transparent UI with Blur** – Native blur effects using Wails + Svelte
- 🧰 **System Command Execution** – Control and interact with your OS through chat commands

---

## 📦 Installation

```bash
git clone https://github.com/yourusername/clipr
cd clipr
wails dev
