<script lang="ts">
  import TopBar from "./components/TopBar.svelte";
  import { onMount } from "svelte";
  import { loadConfig } from "./libs/config";
  import { AskOllama } from "../wailsjs/go/api/App";

  let name: string = "";
  let chatHistory: { user: string; bot: string }[] = [];
  let wails = window.runtime;

  onMount(async () => {
    const loaded = await loadConfig();
    wails.LogPrint(loaded.theme);
  });

  async function sendPrompt() {
    if (!name.trim()) return;
    wails.LogPrint(name);
    const res = await AskOllama(name);
    if (res) {
      wails.LogPrint(res);
      chatHistory = [...chatHistory, { user: name, bot: res }];
      name = ""; // Clear input
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

  <div class="chat-box">
    {#each chatHistory as chat}
      <div class="chat-entry">
        <div class="chat-user"><strong>You:</strong> {chat.user}</div>
        <div class="chat-bot"><strong>Bot:</strong> {chat.bot}</div>
      </div>
    {/each}
  </div>

  <div class="input-box" id="input">
    <input
      placeholder="Enter your prompt"
      autocomplete="off"
      bind:value={name}
      class="input"
      id="name"
      type="text"
      on:keydown={(e) => {
        if (e.key === "Enter") sendPrompt();
      }}
    />
    <button class="btn" on:click={sendPrompt}>Ask the Llama</button>
    <button class="btn" on:click={() => quit(700)}>Quit</button>
  </div>
</main>

<style>
  .chat-box {
    margin-top: 30px;
    color: black;
    width: 80%;
    max-height: 300px;
    overflow-y: auto;
    margin: 1rem auto;
    padding: 1rem;
    background-color: #f3f3f3;
    border-radius: 8px;
    box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
  }

  .chat-entry {
    margin-bottom: 1rem;
  }

  .chat-user,
  .chat-bot {
    margin: 0.25rem 0;
  }
</style>
