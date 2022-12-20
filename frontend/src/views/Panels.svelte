{#if editModal}
  <PanelEditModal {guildId} {channels} {roles} {emojis} {teams} {forms} bind:panel={editData}
                    on:close={() => editModal = false} on:confirm={submitEdit}/>
{/if}

{#if multiEditModal}
  <MultiPanelEditModal {guildId} {channels} {panels} data={multiPanelEditData}
                  on:close={() => multiEditModal = false} on:confirm={submitMultiPanelEdit}/>
{/if}

{#if panelToDelete !== null}
  <ConfirmationModal icon="fas fa-trash-can" isDangerous on:cancel={() => panelToDelete = null}
                     on:confirm={() => deletePanel(panelToDelete.panel_id)}>
    <span slot="body">Are you sure you want to delete the panel {panelToDelete.title}?</span>
    <span slot="confirm">Delete</span>
  </ConfirmationModal>
{/if}

{#if multiPanelToDelete !== null}
  <ConfirmationModal icon="fas fa-trash-can" isDangerous on:cancel={() => multiPanelToDelete = null}
                     on:confirm={() => deleteMultiPanel(multiPanelToDelete.id)}>
    <span slot="body">Are you sure you want to delete the multi-panel {multiPanelToDelete.title}?</span>
    <span slot="confirm">Delete</span>
  </ConfirmationModal>
{/if}


<div class="wrapper">
    <div class="page-title-wrapper" style="width: 100%;">
        <div class="page-title">
            Reaction Panels
        </div>
        <div class="title-dot">&nbsp;</div>
    </div>
  {#if activePage === 'PANELS'}
  <Card footer="{false}">
    <span slot="title">Your Reaction Panels</span>
    <div slot="body" class="card-body">
      <div style="float: right;">
        <div class="tag">Your panel quota: <b>{panels.length} / {isPremium ? '∞' : '3'}</b></div>
        <button class="new-button" on:click={() => setActivePage('CREATE_PANEL')}>New Panel</button>
      </div>
      <table style="margin-top: 10px">
        <thead>
        <tr>
          <th>Channel</th>
          <th>Panel Title</th>
          <th class="category-col">Ticket Channel Category</th>
          <th></th>
          <th></th>
          <th></th>
        </tr>
        </thead>
        <tbody>
        {#each panels as panel}
          <tr class="transcript">
            <td class="left">#{channels.find((c) => c.id === panel.channel_id)?.name ?? 'Unknown Channel'}</td>
            <td>{panel.title}</td>
            <td class="category-col">{channels.find((c) => c.id === panel.category_id)?.name ?? 'Unknown Category'}</td>
            <td class="button-row">
              <Button on:click={() => resendPanel(panel.panel_id)}>Resend</Button>
            </td>
            <td class="button-row">
              <Button on:click={() => openEditModal(panel.panel_id)}>Edit</Button>
            </td>
            <td class="button-row right">
              <Button on:click={() => panelToDelete = panel}>Delete</Button>
            </td>
          </tr>
        {/each}
        </tbody>
      </table>
    </div>
  </Card>
  {/if}

  {#if activePage === 'CREATE_PANEL'}
    <div class="back-to-panels" on:click={() => setActivePage('PANELS')}>← back to panels</div>
  <Card footer="{false}">
    <span slot="title">Create Panel</span>

    <div slot="body" class="body-wrapper">
      {#if !$loadingScreen}
        <PanelCreationForm {guildId} {channels} {roles} {emojis} {teams} {forms} {createPanel} bind:data={panelCreateData}/>
      {/if}
    </div>
  </Card>
  {/if}

  {#if activePage === 'PANELS'}
    <br />
  <Card footer="{false}">
    <span slot="title">Your Multi-Panels</span>
    <div slot="body" class="card-body">
      <div style="float: right;">
        <button class="new-button" on:click={() => setActivePage('CREATE_MULTI_PANEL')}>New Multi Panel</button>
      </div>
      <table style="margin-top: 10px">
        <thead>
        <tr>
          <th>Embed Title</th>
          <th></th>
          <th></th>
          <th></th>
        </tr>
        </thead>
        <tbody>
        {#each multiPanels as panel}
          <tr class="transcript">
            <td>{panel.title}</td>
            <td class="button-row">
              <Button on:click={() => resendMultiPanel(panel.id)}>Resend</Button>
            </td>
            <td class="button-row">
              <Button on:click={() => openMultiEditModal(panel.id)}>Edit</Button>
            </td>
            <td class="button-row">
              <Button on:click={() => multiPanelToDelete = panel}>Delete</Button>
            </td>
          </tr>
        {/each}
        </tbody>
      </table>
    </div>
  </Card>
  {/if}

  {#if activePage === 'CREATE_MULTI_PANEL'}
    <div class="back-to-panels" on:click={() => setActivePage('PANELS')}>← back to panels</div>

    <Card footer={false}>
        <span slot="title">Create Multi-Panel</span>
        <div slot="body" class="card-body">
          <p>Note: The panels which you wish to combine into a multi-panel must already exist</p>

          {#if !$loadingScreen}
            <div style="margin-top: 10px">
              <MultiPanelCreationForm {guildId} {channels} {panels} bind:data={multiPanelCreateData}/>
              <div style="display: flex; justify-content: center; margin-top: 2%">
                <Button icon="fas fa-paper-plane" fullWidth={true} on:click={createMultiPanel}>Submit</Button>
              </div>
            </div>
          {/if}
        </div>
      </Card>
  {/if}
</div>
<!--  <div class="col-main">-->
<!--    <div class="row">-->

<!--    </div>-->
<!--    <div class="row">-->

<!--    </div>-->
<!--  </div>-->
<!--  <div class="col-small">-->
<!--    <div class="row">-->
<!--    </div>-->
<!--    <div class="row">-->
<!--      <Card footer={false}>-->
<!--        <span slot="title">Create Multi-Panel</span>-->
<!--        <div slot="body" class="card-body">-->
<!--          <p>Note: The panels which you wish to combine into a multi-panel must already exist</p>-->

<!--          {#if !$loadingScreen}-->
<!--            <div style="margin-top: 10px">-->
<!--              <MultiPanelCreationForm {guildId} {channels} {panels} bind:data={multiPanelCreateData}/>-->
<!--              <div style="display: flex; justify-content: center; margin-top: 2%">-->
<!--                <Button icon="fas fa-paper-plane" fullWidth={true} on:click={createMultiPanel}>Submit</Button>-->
<!--              </div>-->
<!--            </div>-->
<!--          {/if}-->
<!--        </div>-->
<!--      </Card>-->
<!--    </div>-->
<!--  </div>-->
<!--</div>-->

<style>
  .card{
    width: 100%;
    display: block;
  }
  .wrapper {
    height: 100%;
    width: 100%;
  }

  table {
    width: 100%;
  }

  table *{
    border: none !important;
  }

  th {
    text-align: left;
    font-weight: bold;
    border-bottom: 1px solid #dee2e6;
    padding-left: 10px;
  }

  tr {
    border-bottom: 1px solid #dee2e6;
  }

  tr.transcript{
    background: rgba(255, 255, 255, .04);
    margin-top: 5px;
  }

  .button-row{
    width: 1%;
    white-space: nowrap;
    padding: 5px !important;
  }

   .button-row :global(button){
    padding: 5px 10px !important;
  }

  .button-row td{
    padding: 0;
  }

  td {
    padding: 10px 0 10px 10px;
    border-radius: 6px;
  }

  :global(.fill){
    height: unset !important;
  }

  .tag{
    padding: 7px 25px;
    background: rgba(255, 255, 255, .08);
    border-radius: 6px;
    color: white;
    margin: 0 auto 10px auto !important;
    display: inline-block;
    font-family: 'Poppins', sans-serif !important;
  }

  .card-body{
    width: 100%;
  }

  .new-button{
    border: none;
    border-radius: 6px;
    height: 38px;
    background: #873ef5;
    cursor: pointer;
    transition: .2s ease-in-out;
    color: white;
    font-weight: bold;
    text-transform: uppercase;
    padding: 0 15px;
  }
  .new-button:hover{
    background: #8a45f1;
  }
  .back-to-panels{
    font-weight: bold;
    text-transform: uppercase;
    cursor: pointer;
    padding: 3px 10px;
    transition: .2s ease-in-out;
    display: inline-block;
  }
  .back-to-panels:hover{
    color: #873ef5;
  }

  .page-title-wrapper .title-dot{
    width: 200px;
  }
</style>


<script>
    import Card from "../components/Card.svelte";
    import {notifyError, notifySuccess, withLoadingScreen} from "../js/util";
    import {loadingScreen} from "../js/stores";
    import axios from "axios";
    import {API_URL} from "../js/constants";
    import {setDefaultHeaders} from '../includes/Auth.svelte'
    import Button from "../components/Button.svelte";
    import PanelEditModal from "../components/manage/PanelEditModal.svelte";
    import PanelCreationForm from "../components/manage/PanelCreationForm.svelte";
    import MultiPanelCreationForm from '../components/manage/MultiPanelCreationForm.svelte';
    import MultiPanelEditModal from "../components/manage/MultiPanelEditModal.svelte";
    import ConfirmationModal from "../components/ConfirmationModal.svelte";

    export let currentRoute;
    export let params = {};

    setDefaultHeaders()

    let guildId = currentRoute.namedParams.id;

    let channels = [];
    let roles = [];
    let emojis = [];
    let teams = [];
    let forms = [];
    let panels = [];
    let multiPanels = [];
    let isPremium = false;

    let editModal = false;
    let multiEditModal = false;
    let panelToDelete = null;
    let multiPanelToDelete = null;

    let panelCreateData;
    let editData;
    let multiPanelCreateData;
    let multiPanelEditData;

    let activePage = "PANELS";

    function setActivePage(page){
      activePage = page;
    }

    function openEditModal(panelId) {
        editData = panels.find((p) => p.panel_id === panelId);
        editModal = true;
    }

    function openMultiEditModal(id) {
        multiPanelEditData = multiPanels.find((mp) => mp.id === id);
        multiEditModal = true;
    }

    async function resendPanel(panelId) {
        const res = await axios.post(`${API_URL}/api/${guildId}/panels/${panelId}`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess("Panel resent successfully");
    }

    async function deletePanel(panelId) {
        const res = await axios.delete(`${API_URL}/api/${guildId}/panels/${panelId}`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        panels = panels.filter((p) => p.panel_id !== panelId);
        panelToDelete = null;
    }

    async function resendMultiPanel(id) {
        const res = await axios.post(`${API_URL}/api/${guildId}/multipanels/${id}`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess("Multipanel resent successfully")
    }

    async function deleteMultiPanel(id) {
        const res = await axios.delete(`${API_URL}/api/${guildId}/multipanels/${id}`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        multiPanels = multiPanels.filter((p) => p.id !== id);
        multiPanelToDelete = null;
    }

    async function createPanel() {
        let mapped = Object.fromEntries(Object.entries(panelCreateData).map(([k, v]) => {
            if (v === "null") {
                return [k, null];
            } else {
                return [k, v];
            }
        }));

        const res = await axios.post(`${API_URL}/api/${guildId}/panels`, mapped);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        await loadPanels();
        notifySuccess('Panel created successfully');
    }

    async function submitEdit(e) {
        let data = e.detail;

        const res = await axios.patch(`${API_URL}/api/${guildId}/panels/${data.panel_id}`, data);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        } else {
            editModal = false;
            editData = {};
            notifySuccess('Panel updated successfully');
        }

        await loadPanels();
    }

    async function submitMultiPanelEdit(e) {
        let data = e.detail;

        const res = await axios.patch(`${API_URL}/api/${guildId}/multipanels/${data.id}`, data);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        } else {
            multiEditModal = false;
            multiPanelEditData = {};
            notifySuccess('Multi-panel updated successfully');
        }

        await loadPanels();
    }

    async function createMultiPanel() {
        const res = await axios.post(`${API_URL}/api/${guildId}/multipanels`, multiPanelCreateData);
        if (res.status !== 200) {
            notifyError(res.data.error);
        } else {
            notifySuccess('Multi-panel created successfully');
        }

        await loadMultiPanels();
    }

    async function loadPremium() {
        const res = await axios.get(`${API_URL}/api/${guildId}/premium?include_voting=false`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        isPremium = res.data.premium;
    }

    async function loadChannels() {
        const res = await axios.get(`${API_URL}/api/${guildId}/channels`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        channels = res.data;
    }

    async function loadPanels() {
        const res = await axios.get(`${API_URL}/api/${guildId}/panels`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        // convert button_style and form_id to string
        panels = res.data.map((p) => Object.assign({}, p, {
          button_style: p.button_style.toString(),
          form_id: p.form_id === null ? "null" : p.form_id
        }));
    }

    async function loadMultiPanels() {
        const res = await axios.get(`${API_URL}/api/${guildId}/multipanels`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        multiPanels = res.data.data;
    }

    async function loadTeams() {
        const res = await axios.get(`${API_URL}/api/${guildId}/team`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        teams = res.data;
    }

    async function loadRoles() {
        const res = await axios.get(`${API_URL}/api/${guildId}/roles`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        roles = res.data.roles;
    }

    async function loadEmojis() {
        const res = await axios.get(`${API_URL}/api/${guildId}/emojis`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        emojis = res.data;
    }

    async function loadForms() {
        const res = await axios.get(`${API_URL}/api/${guildId}/forms`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        forms = res.data || [];
    }

    withLoadingScreen(async () => {
      await Promise.all([
          loadPremium(),
          loadChannels(),
          loadTeams(),
          loadForms(),
          loadRoles(),
          loadEmojis(),
          loadPanels(),
          loadMultiPanels()
      ]);
    });
</script>
