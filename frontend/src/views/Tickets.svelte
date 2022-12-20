<div class="parent">
    <div class="refresh-wrapper">
        <div class="page-title-wrapper" style="width: 86%;">
            <div class="page-title">
                Tickets
            </div>
            <div class="title-dot">&nbsp;</div>
        </div>
    </div>

    <Card footer={false}>
      <span slot="title">Open Tickets</span>
      <div slot="body" class="body-wrapper">
        <table>
          <thead>
          <tr class="bg">
            <th>ID</th>
            <th>Panel</th>
            <th>User</th>
            <th>View</th>
          </tr>
          </thead>
          <tbody>
          {#each tickets as ticket}
            <tr>
              <td>{ticket.id}</td>
              <td>{ticket.panel_title}</td>
              {#if ticket.user !== undefined}
                <td>{ticket.user.username}#{ticket.user.discriminator}</td>
              {:else}
                <td>Unknown</td>
              {/if}
              <td>
                <Navigate to="/manage/{guildId}/tickets/view/{ticket.id}" styles="link">
                  <Button type="button">View</Button>
                </Navigate>
              </td>
            </tr>
          {/each}
          </tbody>
        </table>
      </div>
    </Card>
</div>

<script>
    import Card from "../components/Card.svelte";
    import {notifyError, notifySuccess, withLoadingScreen} from '../js/util'
    import axios from "axios";
    import {API_URL} from "../js/constants";
    import {setDefaultHeaders} from '../includes/Auth.svelte'
    import Button from "../components/Button.svelte";
    import {Navigate} from 'svelte-router-spa';

    export let currentRoute;
    let guildId = currentRoute.namedParams.id;

    let tickets = [];

    async function loadTickets() {
        const res = await axios.get(`${API_URL}/api/${guildId}/tickets`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        tickets = res.data;
    }

    withLoadingScreen(async () => {
        setDefaultHeaders();
        await loadTickets();
    });
</script>

<style>
    .parent {
        width: 100%;
        height: 100%;
    }

    .content {
        width: 100%;
        height: 100%;
    }

    .body-wrapper {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
    }

    th{
      text-align: left !important;
      padding: 5px;
    }

    .form-wrapper {
        display: flex;
        flex-direction: column;
        width: 100%;
    }

    td, th {
    border-radius: 4px;
  }

    tr.bg{
      background: rgba(255, 255, 255, .04);
      border-radius: 4px !important;
      padding: 5px 15px;
      border: none !important;
    }

    tr.bg td{
      border-radius: 4px;
      padding: 5px 10px;
    }

    .button-row{
      width: 1%;
      white-space: nowrap;
      background: #dc3545 !important;
      padding: 0 !important;
    }

    :global(.link){
        text-decoration: none;
    }

    .button-row td{
      padding: 0;
    }

    .delete-button{
      padding: 7px 15px !important;
      height: 37px;
      margin-top: 5px;
      border: none;
      background-color: #dc3545 !important;
      border-color: #dc3545 !important;
      border-radius: 4px;
      color: white;
      transition: .2s ease-in-out !important;
      cursor: pointer !important;
      float: right;
    }

    .delete-button:hover:enabled, .delete-button:active {
      background-color: rgba(11, 11, 11, .2) !important;
      border-color: rgba(11, 11, 11, .2) !important;
    }

    .user-row{        
        display: flex;
        gap: 10px;
        margin-top: 15px; 
        
        background-color: var(--fg-color);
        border-radius: 6px;
        padding: 10px 20px;
        box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1);
    }

    .user-row .user-info{
        width: 100%;
    }

    .user-row .user-info .user-name{
        font-weight: bold;
        font-size: 18px;
    }

    .user-row .user-info span{
        font-weight: normal;
        color: rgba(255, 255, 255, .6);
    }

    .user-row .user-id{
        color: rgba(255, 255, 255, .5);
        margin-top: -2px;
        font-size: 14px;
    }
</style>
