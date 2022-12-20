<div class="content" transition:fade>

    <div class="refresh-wrapper">
        <div class="refresh-button" title="Refresh Servers">
            <i class="fas fa-sync" />
        </div>
        <div class="page-title-wrapper" style="width: 100%;">
            <div class="page-title">
                Select a server
            </div>
            <div class="title-dot">&nbsp;</div>
        </div>
    </div>

    <div class="server-list">
        <InviteBadge />

        {#each guilds as guild}
            <Guild guild={guild}/>
        {/each}
    </div>
</div>

<script>
    import axios from 'axios';
    import {fade} from 'svelte/transition';
    import {notifyError, withLoadingScreen} from '../js/util'
    import {setDefaultHeaders} from '../includes/Auth.svelte'
    import {API_URL} from "../js/constants.js";
    import Guild from '../components/Guild.svelte'
    import Card from '../components/Card.svelte'
    import InviteBadge from '../components/InviteBadge.svelte'
    import Button from '../components/Button.svelte'

    setDefaultHeaders();

    export let guilds = [];

    async function loadData() {
        const res = await axios.get(`${API_URL}/user/guilds`);
        guilds = res.data;
    }

    async function refreshGuilds() {
        await withLoadingScreen(async () => {
            const res = await axios.post(`${API_URL}/user/guilds/reload`);
            if (res.status !== 200) {
                notifyError(res.data.error);
                return;
            }

            if (!res.data.success && res.data['reauthenticate_required'] === true) {
                window.location.href = "/login";
                return;
            }

            await loadData();
        });
    }

    withLoadingScreen(async () => {
        await loadData();
    });
</script>

<style>
    .content {
        /* display: flex; */
        height: 100%;
        width: 100%;
        /* justify-content: center; */
    }

    .refresh-wrapper{
        display: flex;
        width: 100%;
    }

    .refresh-button{
        height: 35px;
        width: 35px;
        /* background-color: blue; */
        border-radius: 8px;
        text-align: center;
        margin-top: 10px;
        margin-left: 10px;
        background-color: rgba(255, 255, 255, .06);
        cursor: pointer;
        transition: .2s ease-in-out;
    }

    .refresh-button:hover{
        background-color: var(--primary-color);
    }

    .refresh-button i{
        line-height: 35px;
        font-size: 20px;
    }

    .server-wrapper{
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        gap: 10px;
        margin-top: 15px;
    }

    .server-list :global(.server-box){
        width: 100%;
        background-color: var(--fg-color);

        padding: 10px 20px;
        box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1);
        border-radius: 6px;
        margin-bottom: 10px;
        border: 1px solid transparent;
        cursor: pointer;
        transition: .2s ease-in-out;
    }

    .server-list :global(.server-box):hover{
        border-color: var(--primary-color);
        box-shadow: 0 6px 6px rgba(135, 62, 245, .1), 0 0 0 1px rgba(135, 62, 245, .1);
    }

    .server-list :global(.server-box) :global(.server-info){
        display: flex;
        gap: 20px;
        align-items: center;
    }

    .server-list :global(.server-box) :global(.server-info) :global(img){
        width: 52px;
        height: 52px;
        box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1);
        border-radius: 100%;
    }

    .server-list :global(.server-box) :global(.server-info) :global(.server-name){
        font-size: 17px;
    }

    .page-title-wrapper .title-dot{
        width: 175px;
    }
    
    .card-wrapper {
        display: block;
        width: 75%;
        margin-top: 5%;
    }

    #guild-container {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-evenly;
    }

    #refresh-container {
        display: flex;
        justify-content: center;

        margin: 10px 0;
        color: white;
    }

    @media (max-width: 576px) {
        .card-wrapper {
            width: 100%;
        }
    }
</style>
