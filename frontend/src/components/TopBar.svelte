<script>
  // your script goes here
  import { DeleteConfig, ClearHistory } from "../../wailsjs/go/api/App";
  let toggleSettings = false;
  async function resetConfig() {
    const res = await DeleteConfig();
    if (res == "deleted") {
      window.runtime.LogPrint(res);
      window.runtime.WindowReload();
      toggleSettings != toggleSettings;
    }
  }
</script>

<div class="top-bar">
  <div class="settings">
    <button
      on:click={() => {
        toggleSettings = !toggleSettings;
      }}>Settings</button
    >
  </div>
  {#if toggleSettings}
    <div class="settings-box">
      <h1>Settings</h1>
      <button
        on:click={() => {
          ClearHistory("");
        }}>Clear Response History</button
      >
      <button
        on:click={() => {
          ClearHistory("clipboard");
        }}
      ></button>
    </div>
  {/if}
</div>

<!-- markup (zero or more items) goes here -->

<style>
  /* your styles go here */
  .settings {
    padding: 10px;
    position: fixed;
    top: 0;
    right: 0;
    z-index: 10;
  }
  .settings button {
    padding: 10px;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(20px);
    border: 2px solid rgba(255, 255, 255, 0.2);
    border-radius: 8px;
    color: white;
  }
  .settings-box {
    background-color: rgba(255, 255, 255, 0.185);
    width: 200px;
    position: fixed;
    top: 10px;
    z-index: 0;
    right: 10px;
    backdrop-filter: blur(10px);
    border-radius: 10px;
  }
</style>
