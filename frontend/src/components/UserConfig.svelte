<script>
  import { SaveConfig } from "../../wailsjs/go/api/App";
  let isSet;
  let name;
  let location;
  async function setConfig() {
    const userObj = {
      name,
      location,
    firstLaunch: false,
    };
    await SaveConfig(JSON.stringify(userObj));
    window.runtime.WindowReload()
  }
</script>

<!-- markup (zero or more items) goes here -->
<main class="setup-page">
  <div class="input-box">
    <h1>Hello, Please Enter your Details to Continue</h1>
    <input type="text" class="intro-input" bind:value={name} placeholder="Enter your name" />
    <input
      type="text"
      class="intro-input"
      bind:value={location}
      placeholder="Enter your location"
    />
    <button on:click={setConfig}>Save Config</button>
  </div>
</main>

<style>

  .input-box {
    display: flex;
    flex-direction: column;
    gap: 10px;
    width: 500px;
    align-items: center;
    justify-content: center;
    width: 100%;
  }
  /* your styles go here */
  .setup-page {
    height: 100%;
    background: #000000;
  }
  .intro-input {
    height: 20px;
    border-radius: 10px;
    border: none;
    padding: 10px;
  }
  
</style>
