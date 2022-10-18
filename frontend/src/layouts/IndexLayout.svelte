<Head/>
<div class="outer">
  <div class="wrapper">
    <IndexSidebar {name} {avatar} {isWhitelabel} {isAdmin} />
    <div class="main">
      <Navbar />
      <div class="inner">
        <LoadingScreen/>
        <div class="content-container" class:hide={$loadingScreen}>
          <Route {currentRoute} {params}/>
        </div>
        <NotifyModal/>
      </div>
    </div>
  </div>
</div>

<script>
    import {Route} from 'svelte-router-spa'
    import Head from '../includes/Head.svelte'
    import Sidebar from '../includes/Sidebar.svelte'
    import LoadingScreen from '../includes/LoadingScreen.svelte'
    import NotifyModal from '../includes/NotifyModal.svelte'
    import axios from "axios";
    import {API_URL} from '../js/constants'
    import {notifyError} from '../js/util'
    import {loadingScreen} from "../js/stores"
    import {redirectLogin, setDefaultHeaders} from '../includes/Auth.svelte'
    import IndexSidebar from "../includes/IndexSidebar.svelte";
    import Navbar from "../includes/Navbar.svelte";

    export let currentRoute;
    export let params = {};

    setDefaultHeaders()

    let name;
    let avatar;

    let isWhitelabel = false;
    let isAdmin = false;

    async function loadData() {
        const res = await axios.get(`${API_URL}/api/session`);
        if (res.status !== 200) {
            if (res.data.auth === true) {
                redirectLogin();
            }

            notifyError(res.data.error);
            return;
        }

        name = res.data.username;
        avatar = res.data.avatar;
        isWhitelabel = res.data.whitelabel;
        isAdmin = res.data.admin;
    }

    loadData();
</script>

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
    /*body {*/
    /*    padding: 0 !important;*/
    /*    font-family: 'Poppins', sans-serif;*/
    /*}*/

    /*.wrapper {*/
    /*    display: flex;*/
    /*    width: 100%;*/
    /*    height: 100%;*/
    /*    margin: 0 !important;*/
    /*    padding: 0 !important;*/
    /*}*/

    /*.super-container {*/
    /*    display: flex;*/
    /*    width: 100%;*/
    /*    height: 100%;*/
    /*}*/

    /*.content-container {*/
    /*    display: flex;*/
    /*    width: 100%;*/
    /*    height: 100%;*/
    /*}*/

    /*.hide {*/
    /*    visibility: hidden;*/
    /*}*/

    /*@media (max-width: 950px) {*/
    /*    .wrapper {*/
    /*        flex-direction: column;*/
    /*    }*/
    /*}*/
</style>