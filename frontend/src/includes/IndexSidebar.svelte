<div class="sidebar">
  <div class="server-container">
    <img src="https://ticketsbot.net/assets/img/logo-trans.webp" class="server-icon" />
    <span class="server-name">Tickets Bot</span>
  </div>

  <div class="content">
    <Navigate to="/" styles="sidebar-link">
        Servers
    </Navigate>
    <Navigate to="/whitelabel" styles="sidebar-link">
      Whitelabel
    </Navigate>
    <Navigate to="/admin/bot-staff" styles="sidebar-link">
      Admin
    </Navigate>
  </div>
</div>

<style>
  .sidebar {
      display: flex;
      flex-direction: column;

      border-radius: 6px;
      background: linear-gradient(71.3deg, #D87BFF 0%, #995DF3 100%);
      width: 250px !important;
      height: 100%;
  }

  .server-container {
      display: flex;
      flex-direction: row;
      align-items: center;
      gap: 10px;
      width: 250px !important;

      margin: 20px 30px 0 30px;
      font-weight: bold;
  }

  .server-icon {
      width: 48px;
      height: 48px;
      border-radius: 50%;
  }

  .content {
      border-radius: 5px;
      margin: 20px 10px;
  }
</style>

<script>
    import { Navigate } from "svelte-router-spa";
</script>