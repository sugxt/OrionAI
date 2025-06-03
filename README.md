# Clipr — Clipboard Assistant

Clipr is a cross-platform clipboard assistant and chatbot powered by **Go**, **Svelte**, and **Wails**. It integrates the lightweight **Phi-3 Mini** model to provide fast, local AI responses and assistant functionalities.

Go is used to handle **prompt engineering**, **file handling**, and **clipboard access**, while **Svelte** ensures a reactive, efficient, and beautiful GUI. Wails bridges the frontend and backend, compiling everything into a native desktop application.

---

## ⚙️ Technologies Used

- **Go** – Core logic, file I/O, and system integration
- **Svelte** – Lightweight and reactive frontend framework
- **Wails** – Native desktop app framework connecting Go and Svelte
- **Phi-3 Mini** – Embedded AI language model for chatbot responses

---

## 📦 Installation

```bash
git clone https://github.com/yourusername/clipr
cd clipr
wails dev
