<div class="sidebar">
  <div class="server-container">
    <img src="{iconUrl}" class="server-icon" />
    <span class="server-name">{serverName}</span>
  </div>

  <div class="content">

  </div>
</div>

<style>
  .sidebar {
      display: flex;
      flex-direction: column;

      border-radius: 5px;
      background: linear-gradient(71.3deg, #D87BFF 0%, #995DF3 100%);
      min-width: 250px;
      height: 100%;
  }

  .server-container {
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;
      gap: 10px;

      margin-top: 30px;
  }

  .server-icon {
      width: 48px;
      height: 48px;
      border-radius: 50%;
  }

  .content {
      display: flex;
      height: 100%;
      background: linear-gradient(rgba(0, 0, 0, 0.3), rgba(0, 0, 0, 0.3));
      border-radius: 5px;
      margin: 20px 10px;
  }
</style>

<script>
  export let serverName;
  export let iconUrl;
</script>