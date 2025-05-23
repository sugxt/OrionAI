<script>
  // your script goes here
</script>

<div>
  <div class="tools">
    <a href="#">File</a>
    <a href="#">Edit</a>
    <a href="#">View</a>
    <a href="#">Settings</a>
  </div>
</div>

<!-- markup (zero or more items) goes here -->

<style>
  /* your styles go here */
  .tools {
    position: fixed;
    top: 0;
    left: 0;
    display: flex;
    flex-direction: row;
    gap: 20px;
    padding: 8px;
    width: 100%;
    background-color: rgb(0, 11, 26);
  }
  .tools a {
    font-size: 14px;
    text-decoration: none;
    color: #ffffff;
    font-weight: 300;
  }
</style>
