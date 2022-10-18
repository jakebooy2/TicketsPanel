<Head/>

<div class="outer">
  <div class="wrapper">
    <ServerSidebar {guildId} serverName="My Server" iconUrl="https://miro.medium.com/fit/c/176/176/0*jTQUf7SQznjDASXn.png"/>
    <div class="main">
      <Navbar />
      <div class="inner" class:dropdown={$dropdown}>
        <LoadingScreen/>
        <div class="content-container" class:hide={$loadingScreen}>
          <Route {currentRoute} {params}/>
        </div>
        <NotifyModal/>
      </div>
    </div>
  </div>
</div>

<style>
    body {
        padding: 10px !important;
        font-family: 'Poppins', sans-serif !important;
    }

    .outer {
        padding: 25px;
        box-sizing: border-box; /* important */
        height: 100%;
        width: 100%;
    }

    .wrapper {
        display: flex;
        flex-direction: row;
        width: 100%;
        height: 100%;

        margin: 0 !important;
        padding: 0 !important;
        gap: 10px;
    }

    .main {
        display: flex;
        flex-direction: column;
        gap: 4vh;

        padding: 0 20px;
        width: 100%;
    }

    .inner {
        display: flex;
        padding-bottom: 0px !important;
    }

    .content-container {
        display: flex;
        width: 100%;
        height: 100%;


        border-radius: 6px;
    }

    .hide {
        visibility: hidden;
    }

    :global(.page-title){
      float: left;
      text-align: left;
      width: 100%;
      font-size: 28px;
      font-weight: bold;
      margin-bottom: 15px;
    }
</style>

<script>
    import Head from '../includes/Head.svelte'
    import LoadingScreen from '../includes/LoadingScreen.svelte'
    import NotifyModal from '../includes/NotifyModal.svelte'
    import Navbar from '../includes/Navbar.svelte'

    import {Route} from 'svelte-router-spa'
    import {dropdown, loadingScreen} from '../js/stores'
    import {notifyError, withLoadingScreen} from '../js/util';
    import axios from "axios";
    import {API_URL} from "../js/constants";
    import {setDefaultHeaders} from '../includes/Auth.svelte'
    import ServerSidebar from "../includes/ServerSidebar.svelte";
import Guild from '../components/Guild.svelte';

    export let currentRoute;
    export let params = {};

    let guildId = currentRoute.namedParams.id;
    let permissionLevel = 0;

    setDefaultHeaders();

    async function loadPermissionLevel() {
        const res = await axios.get(`${API_URL}/user/permissionlevel?guild=${guildId}`);
        if (res.status !== 200 || !res.data.success) {
            notifyError(res.data.error);
            return;
        }

        permissionLevel = res.data.permission_level;
    }

    withLoadingScreen(async () => {
        await loadPermissionLevel();
    });
</script>