<div class="parent">
    {#if activeTeam == null}
        {#if createActive}
            <CreateModal {guildId} title={"Create New Team"} label={"Team Name"}
                on:close={() => createActive = false} on:confirm={(response) => createTeam(response)}/>
        {/if}
    <div class="page-title-row">
        <div class="page-title-wrapper" style="width: 86%;">
            <div class="page-title">
                Staff Teams
            </div>
            <div class="title-dot">&nbsp;</div>
        </div>
        <Button icon="fas fa-plus" on:click={() => createActive = true}>
            Create Team
        </Button>
    </div>

    <br /><!-- needed to put it to the left? hacky -->
    <br />
        {#each teams as team}
            <div class="form-list-item">
                <div class="form-list-name">{team.name}</div>
                <Button icon="fas fa-edit" on:click={() => activeTeam = {id: team.id, name: team.name}}>
                    Manage Team
                </Button>
            </div>
        {/each}
    {:else}
    <div class="refresh-wrapper">
        <div class="refresh-button" title="Back to Forms" on:click={() => activeTeam = null}>
            <i class="fas fa-arrow-left" />
        </div>
        <div class="page-title-wrapper" style="width: 86%;">
            <div class="page-title">
                Managing {activeTeam.name}
            </div>
            <div class="title-dot">&nbsp;</div>
        </div>
        {#if activeTeam.id !== 'default'}
        <Button icon="fas fa-trash" danger={true} on:click={() => deleteTeam(activeTeam.id)}>Delete</Button>
        {/if}
    </div>
        <Card header={false} footer={false}>
        <span slot="title">Staff Team Members</span>
        <div slot="body" class="body-wrapper">
            <div class="section">
            <div class="manage">
                <div class="col">
                <h3>Manage Members</h3>

                {#if members.length == 0}
                    <div class="member-row">
                        <center>
                            No members are in this team yet!
                        </center>
                    </div>
                {/if}

                {#each members as member}
                    <div class="member-row">
                        <b>{member.name}</b>
                        <Button danger>Delete</Button>
                    </div>
                {/each}

                </div>

                <div class="col">
                <h3>Add Member</h3>
                <div class="user-select">
                    <div style="display: flex; flex: 1">
                    <UserSelect {guildId} bind:value={selectedUser}/>
                    </div>

                    <div style="margin-left: 10px">
                    <Button type="button" icon="fas fa-plus" disabled={selectedUser === undefined}
                            on:click={addUser}>Add To Team
                    </Button>
                    </div>
                </div>

                <h3>Add Role</h3>
                <div class="user-select">
                    <div style="display: flex; flex: 1">
                    <RoleSelect {guildId} {roles} bind:value={selectedRole}/>
                    </div>

                    <div style="margin-left: 10px">
                    <Button type="button" icon="fas fa-plus" disabled={selectedRole === undefined}
                            on:click={addRole}>Add To Team
                    </Button>
                    </div>
                </div>
                </div>
            </div>
            </div>
        </div>
        </Card>
    {/if}
</div>

<script>
    import Card from "../components/Card.svelte";
    import {notifyError, notifyRatelimit, notifySuccess, withLoadingScreen} from '../js/util'
    import Button from "../components/Button.svelte";
    import axios from "axios";
    import {API_URL} from "../js/constants";
    import {setDefaultHeaders} from '../includes/Auth.svelte'
    import Input from "../components/form/Input.svelte";
    import Select from 'svelte-select';
    import UserSelect from "../components/form/UserSelect.svelte";
    import RoleSelect from "../components/form/RoleSelect.svelte";
    import CreateModal from "../components/manage/CreateModal.svelte";

    export let currentRoute;
    let guildId = currentRoute.namedParams.id;

    let defaultTeam = {id: 'default', name: 'Default'};

    let createActive = false;

    let createName = "";
    let teams = [];
    let roles = [];
    let activeTeam = null;
    $: activeTeam, updateActiveTeam();
    let members = [];

    let selectedUser;
    let selectedRole;

    function labelMapper(team) {
        return team.name;
    }

    async function updateActiveTeam() {
        const res = await axios.get(`${API_URL}/api/${guildId}/team/${activeTeam.id}`);
        if (res.status !== 200) {
            if (res.status === 429) {
                notifyRatelimit();
            } else {
                notifyError(res.data.error);
            }
            return;
        }

        members = res.data;
    }

    async function addUser() {
        const res = await axios.put(`${API_URL}/api/${guildId}/team/${activeTeam.id}/${selectedUser.id}?type=0`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess(`${selectedUser.username}#${selectedUser.discriminator} has been added to the support team ${activeTeam.name}`);

        let entity = {
            id: selectedUser.id,
            type: 0,
            name: `${selectedUser.username}#${selectedUser.discriminator}`
        }
        members = [...members, entity];
        selectedUser = undefined;
    }

    async function addRole() {
        const res = await axios.put(`${API_URL}/api/${guildId}/team/${activeTeam.id}/${selectedRole.id}?type=1`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess(`${selectedRole.name} has been added to the support team ${activeTeam.name}`);

        let entity = {
            id: selectedRole.id,
            type: 1,
            name: selectedRole.name,
        }
        members = [...members, entity];
        selectedRole = undefined;
    }

    async function removeMember(teamId, entity) {
        const res = await axios.delete(`${API_URL}/api/${guildId}/team/${teamId}/${entity.id}?type=${entity.type}`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess(`${entity.name} has been removed from the team`);
        members = members.filter((member) => member.id !== entity.id);
    }

    async function createTeam(name) {
        let data = {
            name: name.detail,
        };

        const res = await axios.post(`${API_URL}/api/${guildId}/team`, data);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess(`Team ${createName} has been created`);
        createName = '';
        teams = [...teams, res.data];
        createActive = false;
    }

    async function deleteTeam(id) {
        const res = await axios.delete(`${API_URL}/api/${guildId}/team/${id}`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess(`Team deleted successfully`);
        activeTeam = defaultTeam;
        teams = teams.filter((team) => team.id !== id);
    }

    async function loadTeams() {
        const res = await axios.get(`${API_URL}/api/${guildId}/team`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        teams = [defaultTeam, ...res.data];
    }

    async function loadRoles() {
        const res = await axios.get(`${API_URL}/api/${guildId}/roles`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        roles = res.data.roles;
    }

    withLoadingScreen(async () => {
        setDefaultHeaders();

        await Promise.all([
            loadTeams(),
            loadRoles()
        ]);

        // await updateActiveTeam(); // Depends on teams
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
        margin-top: 30px;
    }

    .page-title-row{
        display: flex;
    }

    .body-wrapper {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        padding: 1%;
    }

    .section {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
    }

    .section:not(:first-child) {
        margin-top: 2%;
    }

    .section-title {
        font-size: 36px;
        font-weight: bolder !important;
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


    h3 {
        font-size: 28px;
        margin-bottom: 4px;
    }

    .row {
        display: flex;
        flex-direction: row;
        width: 100%;
        height: 100%;
    }

    .manage {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 100%;
        height: 100%;
        margin-top: 12px;
    }

    .col {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 49%;
        height: 100%;
    }

    table.nice > tbody > tr:first-child {
        border-top: 1px solid #dee2e6;
    }

    .user-select {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 100%;
        height: 100%;
        margin-bottom: 1%;
    }

    @media only screen and (max-width: 950px) {
        .manage {
            flex-direction: column;
        }

        .col {
            width: 100%;
        }
    }

    .form-input-row{
        background-color: var(--fg-color);
        box-shadow: var(--shadow);
        padding: 10px;
        border-radius: 6px;
        margin-bottom: 10px;
    }

    .form-list-item{
        background-color: var(--fg-color);
        box-shadow: var(--shadow);
        padding: 10px 15px;
        border-radius: 6px;
        margin-bottom: 7px;
        display: flex;
        align-items: center;
    }

    .form-list-item .form-list-name{
        font-weight: bold;
        font-size: 18px;
        width: 100%;
    }

    .member-row{
        background-color: rgba(255, 255, 255, .04);
        box-shadow: var(--shadow);
        width: 100%;
        display: flex;
        align-items: center;
        padding: 5px 15px;
        border-radius: 6px;
        margin-bottom: 7px;
    }

    .member-row b{
        width: 100%;
    }
</style>
