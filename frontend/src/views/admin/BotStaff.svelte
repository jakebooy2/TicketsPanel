<div class="wrapper">
    <div class="page-title-wrapper">
        <div class="page-title">
            Bot Staff
        </div>
        <div class="title-dot">&nbsp;</div>
    </div>

  <Card footer="{false}" fill="{false}">
    <h4 slot="title">Add staff member</h4>
    <div slot="body" class="full-width body-wrapper">
      <form class="form-wrapper" on:submit|preventDefault={addStaff}>
        <Input label="User ID" placeholder="585576154958921739" bind:value={tempUserId} />
        <Button type="submit">Add</Button>
      </form>
    </div>
  </Card>

  {#each staff as user}
    <div class="user-row">
        <div class="user-info">
            <div class="user-name">{user.username}<span>#{user.discriminator.padStart(4, '0')}</span></div>
            <div class="user-id">{user.id}</div>
        </div>
        <div class="delete-button" on:click={removeStaff(user.id)}>
            Remove
        </div>
    </div>
  {/each}

  {#if staff.length == 0}
    <div class="user-row">
        <div class="user-info">
            <div class="user-name">No staff found</div>
        </div>
    </div>
  {/if}
</div>

<script>
    import {notifyError, withLoadingScreen} from '../../js/util'
    import axios from "axios";
    import Card from '../../components/Card.svelte'
    import {API_URL} from "../../js/constants";
    import {setDefaultHeaders} from '../../includes/Auth.svelte'
    import Button from "../../components/Button.svelte";
    import Input from "../../components/form/Input.svelte";

    setDefaultHeaders()

    let staff = [];
    let tempUserId = "";

    async function addStaff() {
        if (tempUserId.length === 0) {
            return;
        }

        const res = await axios.post(`${API_URL}/api/admin/bot-staff/${tempUserId}`);
        if (res.status !== 204) {
            notifyError(res.data.error);
            return;
        }

        tempUserId = "";
        await loadData(); // TODO: Return user data with response
    }

    async function removeStaff(userId) {
        const res = await axios.delete(`${API_URL}/api/admin/bot-staff/${userId}`);
        if (res.status !== 204) {
            notifyError(res.data.error);
            return;
        }

        staff = staff.filter(user => user.id !== userId);
    }

    async function loadData() {
        const res = await axios.get(`${API_URL}/api/admin/bot-staff`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        staff = res.data;
    }

    withLoadingScreen(async () => {
        await loadData();
    });
</script>

<style>
    .wrapper {
        display: flex;
        flex-direction: column;
        height: 100%;
        width: 100%;
        padding: 0 25px 10px 25px;
    }

    .page-title-wrapper .title-dot{
        width: 100px;
    }

    .content {
        display: flex;
        justify-content: space-around;
        flex-direction: row;
        width: 95%;

        margin-top: 2%;
    }

    @media only screen and (max-width: 900px) {
        .content {
            flex-direction: column;
        }
    }

    .body-wrapper {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        gap: 30px;
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
