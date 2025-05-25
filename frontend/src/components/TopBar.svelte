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
    margin-top: 10px;
    margin-left: 8px;
    margin-right: 8px;
    position: fixed;
    top: 0;
    left: 0;
    display: flex;
    flex-direction: row;
    gap: 20px;
    padding: 10px;
    padding-left: 25px;
    padding-right: 25px;
    border-radius: 40px;
    width:max-content;
    background: linear-gradient(200deg, #5395b9 0%, #96457f 100%);
  }
  .tools a {
    font-size: 14px;
    text-decoration: none;
    color: #ffffff;
    font-weight: 400;
  }
</style>
