<div class="sidebar">
  <div class="server-container">
    <img src="{iconUrl}" class="server-icon" />
    <span class="server-name">{serverName}</span>
  </div>

  <div class="content">
    <!-- isAdmin var needed for checks -->

      <Navigate to="/" styles="sidebar-link">
          ← Back to Servers
      </Navigate>

    <Navigate to="/manage/{guildId}/settings" styles="sidebar-link">
        Settings
    </Navigate>

    <Navigate to="/manage/{guildId}/transcripts" styles="sidebar-link">
        Transcripts
    </Navigate>

    <Navigate to="/manage/{guildId}/panels" styles="sidebar-link">
        Reaction Panels
    </Navigate>

    <Navigate to="/manage/{guildId}/forms" styles="sidebar-link">
        Forms
    </Navigate>

    <Navigate to="/manage/{guildId}/teams" styles="sidebar-link">
        Staff Teams
    </Navigate>

    <Navigate to="/manage/{guildId}/integrations" styles="sidebar-link">
        Integrations

        <div class="new-badge">New!</div>
    </Navigate>

    <Navigate to="/manage/{guildId}/tickets" styles="sidebar-link">
        Tickets
    </Navigate>

    <Navigate to="/manage/{guildId}/blacklist" styles="sidebar-link">
        Blacklist
    </Navigate>

    <Navigate to="/manage/{guildId}/tags" styles="sidebar-link">
        Tags
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

    export let serverName;
    export let iconUrl;
    export let guildId;
</script>