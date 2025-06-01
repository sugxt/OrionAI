<script lang="ts">
  import TopBar from "./components/TopBar.svelte";
  import { clipboardHistory, responseHistory } from "./libs/clipboard";

  var currentMode:string = "summarize"
  const setMode = (mode: string) => {
    //Making a function incase there is more to handle here
    currentMode = mode;
  };
</script>

<main>
  <TopBar />
  <div class="history-main">
    <h1>Clipr History</h1>
    <div class="mode-select">
      <h2>Select Modes</h2>
      <div class="button-list">
        <button on:click={()=> {
          setMode("sum")
        }}>Summarize</button>
        <button on:click={()=> {
          setMode("expand")
        }}>Expand</button>
      </div>
    </div>
    <ul class="clip-list">
      {#each $clipboardHistory as clip}
        <li class="clip-item">{clip}</li>
      {/each}
    </ul>
    <ul class="response-list">
      {#each $responseHistory as res}
        <li class="response-item">
          {res}
        </li>
      {/each}
    </ul>
  </div>
</main>

<style>
  .history-main {
    margin-top: 35px;
    display: flex;
    flex-direction: column;
    align-items: start;
    justify-content: start;
  }
  .history-main h1 {
    font-weight: 600;
    font-size: x-large;
    margin-left: 20px;
  }
  .clip-list {
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    gap: 10px;
    width: 100%;
  }

  .clip-item {
    list-style-type: none;
    text-decoration: none;
    background-color: rgba(255, 255, 255, 0.411);
    padding: 20px;
    width: 25%;
    border-radius: 10px;
    border: 2px solid rgba(255, 255, 255, 0.452);
    transition: cubic-bezier();
    transition-duration: 200ms;
  }
  .clip-item:hover {
    scale: 1.03;
  }

  .mode-select {
    padding: 20px;
    display: flex;
    flex-direction: column;
    align-items: start;
    margin-left: 20px;
    display: flex;
    flex-direction: column;
    margin: 0;
    gap: 10px;
  }

  .mode-select h2 {
    font-weight: 500;
    padding:0;
    margin: 0;
    font-size: 16px;
  }

  .button-list {
    display: flex;
    flex-direction: row;
    align-items: start;
    gap: 10px;
  }

  .button-list button{
    width: auto;
    padding: 10px;
    font-size: 14px;
    font-weight: 600;
    border-radius: 12px;
    border: none;
    background-color: white;
  }

  .response-list {
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    gap: 10px;
  }
  .response-item {
    list-style-type: none;
    padding: 20px;
    border-radius: 10px;
    width: 25%;
    background-color: rgba(255, 255, 255, 0.411);
  }
</style>
