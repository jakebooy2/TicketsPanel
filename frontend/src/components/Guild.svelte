<div class="guild-badge" on:click={goto(guild.id)}>
  <div class="guild-icon-bg">
    {#if guild.icon === undefined || guild.icon === ""}
      <i class="fas fa-question guild-icon-fa"></i>
    {:else}
      <img class="guild-icon" src="{getIconUrl()}" alt="Guild Icon"/>
    {/if}
  </div>

  <div>
    <span class="guild-name">
      {guild.name}
    </span>
  </div>
</div>

<script>
    import axios from 'axios';
    import {API_URL} from "../js/constants";
    import {notifyError} from "../js/util";

    export let guild;

    function isAnimated() {
        if (guild.icon === undefined || guild.icon === "") {
            return false;
        } else {
            return guild.icon.startsWith('a_')
        }
    }

    function getIconUrl() {
        if (isAnimated()) {
            return `https:\/\/cdn.discordapp.com/icons/${guild.id}/${guild.icon}.gif?size=256`
        } else {
            return `https:\/\/cdn.discordapp.com/icons/${guild.id}/${guild.icon}.webp?size=256`
        }
    }

    async function goto(guildId) {
        const permissionLevel = await getPermissionLevel(guildId);
        if (permissionLevel === 2) {
            window.location.href = `/manage/${guildId}/settings`;
        } else {
            window.location.href = `/manage/${guildId}/transcripts`;
        }
    }

    async function getPermissionLevel(guildId) {
        const res = await axios.get(`${API_URL}/user/permissionlevel?guild=${guildId}`);
        if (res.status !== 200 || !res.data.success) {
            notifyError(res.data.error);
            return;
        }

        return res.data.permission_level;
    }
</script>

<style>
    :global(.guild-badge) {
        display: flex;
        align-items: center;
        background-color: rgba(255, 255, 255, .06);
        border-radius: 6px;
        cursor: pointer;
        padding: 10px 0;
        margin-bottom: 5px;
    }

    @media (max-width: 950px) {
        :global(.guild-badge) {
            width: 100%;
        }
    }

    :global(.guild-icon-bg) {
        height: 52px;
        width: 52px;
        background-color: #272727;
        border-radius: 6px;
        margin-left: 10px;
        transition: .2s ease-in-out;
    }

    :global(.guild-icon) {
        height: 52px;
        width: 52px;
        border-radius: 50%;
    }

    :global(.guild-icon-fa) {
        border-radius: 50%;
        color: white;
        font-size: 32px !important;
        width: 52px;
        height: 52px;
        text-align: center;
        margin-top: 10px;
    }

    :global(.guild-name) {
        color: white !important;
        padding-left: 10px;
    }

    :global(.guild-badge):hover :global(.guild-icon-bg){
      background: #995DF3;
    }
</style>
