<script lang="ts">
  import TopBar from "./components/TopBar.svelte";
  import { clipboardHistory, responseHistory } from "./libs/clipboard";
</script>

<main>
  <TopBar />
  <div class="history-main">
    ClipBoard History Given Here
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
  .clip-list {
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    gap: 10px;
    width: 100%;
    margin-top: 45px;
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

  .response-list {
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    gap: 10px;
  }
  .response-item {
    padding: 20px;
    background-color: rgba(255, 255, 255, 0.411);
  }
</style>
