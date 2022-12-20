<div class="sidebar-container">
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

    export let serverName;
    export let iconUrl;
    export let guildId;
</script>