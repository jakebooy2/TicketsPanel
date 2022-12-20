<div class="sidebar-container">
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
</div>

<style>
  .sidebar {
      border-radius: 6px;
      background: var(--fg-color);
      width: 275px !important;
      padding: 1px 0;
      box-shadow: var(--shadow);
  }

  .server-container {
      display: flex;
      flex-direction: row;
      align-items: center;
      gap: 10px;

      font-weight: bold;

      margin: 10px;
      padding: 5px 10px;
      border-radius: 4px;

      background: linear-gradient(33.3deg, #873ef5 0%, #995DF3 100%);
      box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1);
      margin-bottom: -5px;
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