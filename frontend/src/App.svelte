<script lang="ts">
  import TopBar from "./components/TopBar.svelte";
  import { AskOllama } from "../wailsjs/go/api/App";
  let name: string = "";
  let chatHistory: { user: string; bot: string }[] = [];
  let isLoading: boolean = false;
  let isStart: boolean = false;
  let wails = window.runtime;
  async function sendPrompt() {
    if (!name.trim()) return;

    const userMessage = name.trim();
    isLoading = true;
    name = ""; // Clear input immediately

    // Add user message to chat
    chatHistory = [...chatHistory, { user: userMessage, bot: "" }];

    wails.LogPrint(userMessage);
    const res = await AskOllama(userMessage);
    isLoading = false;

    if (res) {
      wails.LogPrint(res);
      // Update the last entry with bot response
      chatHistory[chatHistory.length - 1].bot = res;
      chatHistory = [...chatHistory]; // Trigger reactivity
    }
  }
  function quit(timer: number): void {
    setInterval(() => {
      wails.Quit();
    }, timer);
  }
</script>

<main>
  <TopBar />
  <div class="chat-container">
    <div class="chat-box">
      {#each chatHistory as chat, index}
        <div class="message-group">
          <!-- User message -->
          <div class="message user-message">
            <div class="message-content user-content">
              {chat.user}
            </div>
            <div class="message-time">
              {new Date().toLocaleTimeString([], {
                hour: "2-digit",
                minute: "2-digit",
              })}
            </div>
          </div>
          <!-- Bot message or loading -->
          {#if chat.bot}
            <div class="message bot-message">
              <div class="bot-avatar">🦙</div>
              <div class="message-content bot-content">
                {chat.bot}
              </div>
              <div class="message-time">
                {new Date().toLocaleTimeString([], {
                  hour: "2-digit",
                  minute: "2-digit",
                })}
              </div>
            </div>
          {:else if index === chatHistory.length - 1 && isLoading}
            <div class="message bot-message">
              <div class="bot-avatar">🦙</div>
              <div class="message-content bot-content">
                <div class="typing-indicator">
                  <div class="typing-dot"></div>
                  <div class="typing-dot"></div>
                  <div class="typing-dot"></div>
                </div>
              </div>
            </div>
          {/if}
        </div>
      {/each}
    </div>
  </div>

  <div class="input-container">
    <div class="input-box">
      <input
        placeholder="Type your message..."
        autocomplete="off"
        bind:value={name}
        class="input"
        type="text"
        disabled={isLoading}
        on:keydown={(e) => {
          if (e.key === "Enter" && !isLoading) sendPrompt();
        }}
      />
      <button
        class="send-btn"
        on:click={sendPrompt}
        disabled={isLoading || !name.trim()}
        class:loading={isLoading}
      >
        {#if isLoading}
          <div class="spinner"></div>
        {:else}
          <svg
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="m22 2-7 20-4-9-9-4Z" />
            <path d="M22 2 11 13" />
          </svg>
        {/if}
      </button>
    </div>
    <button class="quit-btn" on:click={() => quit(700)}>
      <svg
        width="16"
        height="16"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
      >
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
      Quit
    </button>
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
      Ubuntu, Cantarell, sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
  }

  main {
    display: flex;
    flex-direction: column;
    height: 100vh;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  }

  .chat-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    max-width: 900px;
    margin: 0 auto;
    padding: 20px;
    overflow: hidden;
  }

  .chat-box {
    flex: 1;
    overflow-y: auto;
    padding: 20px 0;
    display: flex;
    flex-direction: column;
    gap: 24px;
    scroll-behavior: smooth;
  }

  .chat-box::-webkit-scrollbar {
    width: 6px;
  }

  .chat-box::-webkit-scrollbar-track {
    background: rgba(255, 255, 255, 0.1);
    border-radius: 3px;
  }

  .chat-box::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.3);
    border-radius: 3px;
  }

  .chat-box::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.5);
  }

  .message-group {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .message {
    display: flex;
    flex-direction: column;
    max-width: 70%;
    position: relative;
  }

  .user-message {
    align-self: flex-end;
    align-items: flex-end;
  }

  .bot-message {
    align-self: flex-start;
    align-items: flex-start;
    flex-direction: row;
    gap: 12px;
    max-width: 75%;
  }

  .bot-avatar {
    width: 36px;
    height: 36px;
    border-radius: 50%;
    background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 18px;
    flex-shrink: 0;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  }

  .message-content {
    padding: 16px 20px;
    border-radius: 20px;
    position: relative;
    word-wrap: break-word;
    line-height: 1.4;
    font-size: 15px;
  }

  .user-content {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border-bottom-right-radius: 6px;
    box-shadow: 0 2px 12px rgba(102, 126, 234, 0.3);
  }

  .bot-content {
    background: white;
    color: #333;
    border-bottom-left-radius: 6px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
    border: 1px solid rgba(0, 0, 0, 0.05);
  }

  .message-time {
    font-size: 11px;
    color: rgba(255, 255, 255, 0.7);
    margin-top: 4px;
    font-weight: 500;
  }

  .bot-message .message-time {
    color: rgba(0, 0, 0, 0.5);
    margin-left: 48px;
  }

  .typing-indicator {
    display: flex;
    gap: 4px;
    align-items: center;
    padding: 8px 0;
  }

  .typing-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background: #ccc;
    animation: typingAnimation 1.4s infinite ease-in-out;
  }

  .typing-dot:nth-child(1) {
    animation-delay: -0.32s;
  }

  .typing-dot:nth-child(2) {
    animation-delay: -0.16s;
  }

  @keyframes typingAnimation {
    0%,
    80%,
    100% {
      transform: scale(0.8);
      opacity: 0.5;
    }
    40% {
      transform: scale(1);
      opacity: 1;
    }
  }

  .input-container {
    padding: 20px;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(20px);
    border-top: 1px solid rgba(255, 255, 255, 0.2);
  }

  .input-box {
    display: flex;
    gap: 12px;
    max-width: 900px;
    margin: 0 auto;
    align-items: center;
  }

  .input {
    flex: 1;
    padding: 16px 20px;
    border: none;
    border-radius: 25px;
    background: white;
    font-size: 15px;
    outline: none;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
  }

  .input:focus {
    box-shadow: 0 4px 20px rgba(102, 126, 234, 0.3);
    transform: translateY(-1px);
  }

  .input:disabled {
    opacity: 0.7;
    cursor: not-allowed;
  }

  .send-btn {
    width: 48px;
    height: 48px;
    border: none;
    border-radius: 50%;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    box-shadow: 0 2px 12px rgba(102, 126, 234, 0.3);
  }

  .send-btn:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 4px 20px rgba(102, 126, 234, 0.4);
  }

  .send-btn:active {
    transform: translateY(0);
  }

  .send-btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
    transform: none;
  }

  .spinner {
    width: 20px;
    height: 20px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-radius: 50%;
    border-top-color: white;
    animation: spin 1s ease-in-out infinite;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }

  .quit-btn {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 12px 16px;
    border: none;
    border-radius: 20px;
    background: rgba(255, 255, 255, 0.2);
    color: white;
    cursor: pointer;
    font-size: 13px;
    font-weight: 500;
    transition: all 0.3s ease;
    backdrop-filter: blur(10px);
  }

  .quit-btn:hover {
    background: rgba(255, 255, 255, 0.3);
    transform: translateY(-1px);
  }

  /* Mobile responsiveness */
  @media (max-width: 768px) {
    .chat-container {
      padding: 10px;
    }
    .message {
      max-width: 85%;
    }
    .bot-message {
      max-width: 90%;
    }
    .input-container {
      padding: 15px;
    }
    .input-box {
      flex-wrap: wrap;
      gap: 10px;
    }
    .quit-btn {
      order: 1;
      width: 100%;
      justify-content: center;
      margin-top: 10px;
    }
  }
</style>
