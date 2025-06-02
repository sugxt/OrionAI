<script lang="ts">
  import TopBar from "./components/TopBar.svelte";
  import { clipboardHistory, responseHistory } from "./libs/clipboard";
  import { SetMode } from "../wailsjs/go/api/App";

  let currentTab: string = "clipboard";
  let activeMode: string = "sum"; // Track active mode

  let modes = [
    {
      name: "sum",
      label: "Summarize",
    },
    {
      name: "expand",
      label: "Expand",
    },
    {
      name: "para",
      label: "Paraphrase",
    },
  ];

  const handleModeChange = (mode: string) => {
    activeMode = mode;
    SetMode(mode);
  };
</script>

<main>
  <TopBar />
  <div class="container">
    <div class="header">
      <h1>Clipr History</h1>
      <p class="subtitle">Enhance your clipboard with AI-powered modes</p>
    </div>
    <!-- Mode Selection -->
    <div class="mode-section">
      <h2>Processing Mode</h2>
      <div class="mode-buttons">
        {#each modes as mode}
          <button
            class="mode-btn {activeMode === mode.name ? 'active' : ''}"
            on:click={() => handleModeChange(mode.name)}
          >
            <span>{mode.label}</span>
          </button>
        {/each}
      </div>
    </div>

    <!-- Tab Navigation -->
    <div class="tab-navigation">
      <button
        class="tab-btn {currentTab === 'clipboard' ? 'active' : ''}"
        on:click={() => (currentTab = "clipboard")}
      >
        Clipboard History
        <span class="tab-count">{$clipboardHistory.length}</span>
      </button>
      <button
        class="tab-btn {currentTab === 'responses' ? 'active' : ''}"
        on:click={() => (currentTab = "responses")}
      >
        AI Responses
        <span class="tab-count">{$responseHistory.length}</span>
      </button>
    </div>

    <!-- Content Area -->
    <div class="content-area">
      {#if currentTab === "clipboard"}
        <div class="grid-container">
          {#each $clipboardHistory as clip, index}
            <div class="card clipboard-card">
              <div class="card-header">
                <span class="card-number">#{index + 1}</span>
                <div class="card-actions">
                  <button class="action-btn" title="Copy">📋</button>
                  <button class="action-btn" title="Delete">🗑️</button>
                </div>
              </div>
              <div class="card-content">
                {clip}
              </div>
              <div class="card-footer"></div>
            </div>
          {/each}

          {#if $clipboardHistory.length === 0}
            <div class="empty-state">
              <div class="empty-icon">📋</div>
              <h3>No clipboard history yet</h3>
              <p>Start copying text to see it appear here</p>
            </div>
          {/if}
        </div>
      {/if}

      {#if currentTab === "responses"}
        <div class="grid-container">
          {#each $responseHistory as response, index}
            <div class="card response-card">
              <div class="card-header">
                <span class="card-number">#{index + 1}</span>
                <div class="card-actions">
                  <button class="action-btn" title="Copy">📋</button>
                  <button class="action-btn" title="Delete">🗑️</button>
                </div>
              </div>
              <div class="card-content">
                {response}
              </div>
            </div>
          {/each}

          {#if $responseHistory.length === 0}
            <div class="empty-state">
              <div class="empty-icon">🤖</div>
              <h3>No AI responses yet</h3>
              <p>
                Process some clipboard items to see AI-enhanced content here
              </p>
            </div>
          {/if}
        </div>
      {/if}
    </div>
  </div>
</main>
