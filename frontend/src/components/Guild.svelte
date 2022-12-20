<div class="server-box" on:click={goto(guild.id)}>
    <div class="server-info">
        {#if guild.icon === undefined || guild.icon === ""}
            <div class="invite-icon">
                <i class="fas fa-question"></i>
            </div>
        {:else}
            <img class="guild-icon" src="{getIconUrl()}" alt="Guild Icon"/>
        {/if}
        <div class="server-name">{guild.name}</div>
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