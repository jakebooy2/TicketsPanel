<div class="wrapper">
  <div class="page-title">
    Admin
  </div>

  <Card footer="{false}" fill="{false}">
    <h4 slot="title">Bot Staff</h4>
    <div slot="body" class="full-width body-wrapper">
      <form class="form-wrapper" on:submit|preventDefault={addStaff}>
        <Input label="User ID" placeholder="585576154958921739" bind:value={tempUserId} />
        <Button type="submit">Add</Button>
      </form>

      <table>
        <thead>
        <tr>
          <th>Username</th>
          <th></th>
        </tr>
        </thead>
        <tbody>
        {#each staff as user}
          <tr class="bg">
            <td>{user.username}#{user.discriminator.padStart(4, '0')} ({user.id})</td>
            <td>
              <button type="button" class="delete-button" danger on:click={() => removeStaff(user.id)}>
                Delete
              </button>
            </td>
          </tr>
        {/each}
        {#if staff.length == 0}
          <tr class="bg">
            <td>No staff found.</td>
          </tr>
        {/if}
        </tbody>
      </table>
    </div>
  </Card>
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
        align-items: center;
        padding: 0 25px 10px 25px;
    }

    .page-title{
      float: left;
      text-align: left;
      width: 100%;
      font-size: 28px;
      font-weight: bold;
      margin-bottom: 15px;
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
      padding: 5px 15px !important;
      border: none;
      background-color: #dc3545 !important;
      border-color: #dc3545 !important;
      border-radius: 4px;
      color: white;
      transition: .2s ease-in-out !important;
      cursor: pointer !important;
    }

    .delete-button:hover:enabled, .delete-button:active {
      background-color: rgba(11, 11, 11, .2) !important;
      border-color: rgba(11, 11, 11, .2) !important;
    }
</style>
