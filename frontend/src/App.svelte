<script lang="ts">
  import TopBar from "./components/TopBar.svelte";
  import UserConfig from "./components/UserConfig.svelte";
  import { LoadConfig, DeleteConfig, GetClipBoardHistory } from "../wailsjs/go/api/App";
  import { AskOllama } from "../wailsjs/go/api/App";
  import { onMount } from "svelte";
  let message: string = "";
  let isConfigWindow = false;
  let chatHistory: { user: string; bot: string }[] = [];
  let isLoading: boolean = false;
  let wails = window.runtime;

  let currentMode = "";

  const modes = [
    {
      name: "General",
      mode: "general",
    },
    {
      name: "Tasks",
      mode: "tasks",
    },
  ];

  onMount(() => {
    const getDetails = async () => {
      const res = await LoadConfig();
      // wails.LogPrint(res);
      if (res == "") {
        isConfigWindow = true;
      }
    };
    getDetails();
  });

  async function sendPrompt() {
    if (!message.trim()) return;

    const userMessage = message.trim();
    isLoading = true;
    message = ""; // Clear input immediately

    // Add user message to chat
    chatHistory = [...chatHistory, { user: userMessage, bot: "" }];

    wails.LogPrint(userMessage);
    if (currentMode == "") {
      currentMode = "general";
    }
    const res = await AskOllama(userMessage, currentMode);
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
  {#if isConfigWindow}
    <UserConfig />
  {/if}
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
    <div class="mode-select">
      {#each modes as mode}
        <div class="mode-details">
          <button
            class="{mode.mode == currentMode ? 'active':''}"
            on:click={() => {
              currentMode = ""
              currentMode = mode.mode;
              wails.LogPrint(currentMode)
            }}
            >{mode.name}
          </button>
        </div>
      {/each}
    </div>
    <div class="input-box">
      <input
        placeholder="Type your message..."
        autocomplete="off"
        bind:value={message}
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
        disabled={isLoading || !message.trim()}
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


