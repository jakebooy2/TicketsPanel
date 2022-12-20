<div class="parent">
    {#if createActive}  
        <CreateModal {guildId} title={"Create New Form"} label={"Form Name"}
            on:close={() => createActive = false} on:confirm={(response) => createForm(response)}/>
    {/if}

    {#if activeFormId !== null}
        <div class="refresh-wrapper">
            <div class="refresh-button" title="Back to Forms" on:click={() => activeFormId = null}>
                <i class="fas fa-arrow-left" />
            </div>
            <div class="page-title-wrapper" style="width: 86%;">
                <div class="page-title">
                    Editing {activeFormTitle}
                </div>
                <div class="title-dot">&nbsp;</div>
            </div>

            {#if editingTitle}
                <RenameFormModal {guildId} title={activeFormTitle} form_id={activeFormId}
                    on:close={() => editingTitle = false} on:confirm={(name) => createForm(name)}/>
            {/if}

            <Button gap icon="fas fa-save" disabled={formLength === 0} on:click={saveInputs}>Save</Button>

            <Button icon="fas fa-edit" on:click={() => editingTitle = true}>Rename</Button>

            <Button icon="fas fa-trash" danger={true} on:click={() => deleteForm(activeFormId)}>Delete</Button>
        </div>

        {#each forms.find(form => form.form_id === activeFormId).inputs as input, i (input)}
            <div class="form-input-row" animate:flip="{{duration: 500}}">
                <FormInputRow data={input} formId={activeFormId}
                            withSaveButton={true} withDeleteButton={true} withDirectionButtons={true}
                            index={i} {formLength}
                            on:delete={() => deleteInput(activeFormId, input)}
                            on:move={(e) => changePosition(activeFormId, input, e.detail.direction)}/>
            </div>
        {/each}

        <div class="new-input" on:click={addInput}>
            <i class="fas fa-plus"></i> Add Input
        </div>

    {:else}
        <div class="page-title-row">
            <div class="page-title-wrapper" style="width: 86%;">
                <div class="page-title">
                    Forms
                </div>
                <div class="title-dot">&nbsp;</div>
            </div>
            {#if activeFormId == null}
                <Button icon="fas fa-plus" on:click={() => createActive = true}>
                    Create Form
                </Button>
            {/if}
        </div>

        {#each forms as form}
            <div class="form-list-item">
                <div class="form-list-name">
                    {form.title}
                </div>
                <Button icon="fas fa-edit" on:click={() => activeFormId = form.form_id}>
                    Manage Form
                </Button>
            </div>
        {/each}
    {/if}
</div>

<svelte:window bind:innerWidth={windowWidth}/>

<script>
    import {notifyError, notifySuccess, nullIfBlank, withLoadingScreen} from '../js/util'
    import Button from "../components/Button.svelte";
    import axios from "axios";
    import {API_URL} from "../js/constants";
    import {setDefaultHeaders} from '../includes/Auth.svelte'
    import FormInputRow from "../components/manage/FormInputRow.svelte";
    import {flip} from "svelte/animate";
    import RenameFormModal from "../components/manage/RenameFormModal.svelte";
    import CreateModal from "../components/manage/CreateModal.svelte";

    export let currentRoute;
    let guildId = currentRoute.namedParams.id;

    let createActive = false;

    let forms = [];
    let toDelete = {};
    let activeFormId = null;
    $: activeFormTitle = activeFormId !== null ? forms.find(f => f.form_id === activeFormId).title : 'Unknown';

    $: formLength = activeFormId !== null ? forms.find(f => f.form_id === activeFormId).inputs.length : 0;

    let editingTitle = false;
    let renamedTitle = "";
    $: activeFormId, reflectTitle();

    function reflectTitle() {
        renamedTitle = activeFormId !== null ? forms.find(f => f.form_id === activeFormId).title : null;
    }

    $: windowWidth = 0;

    function getForm(formId) {
        return forms.find(form => form.form_id === formId);
    }

    async function updateTitle() {
        const res = await axios.patch(`${API_URL}/api/${guildId}/forms/${activeFormId}`, {title: renamedTitle});
        if (res.status !== 200) {
            notifyError('Failed to update form title');
            return;
        }

        editingTitle = false;
        getForm(activeFormId).title = renamedTitle;
        forms = forms;

        notifySuccess('Form title updated');
    }

    async function createForm(title) {
        let data = {
            title: title.detail,
        };

        const res = await axios.post(`${API_URL}/api/${guildId}/forms`, data);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        createActive = false;
        notifySuccess(`Form ${title.detail} has been created`);
        newTitle = '';

        let form = res.data;
        form.inputs = [];

        activeFormId = null; // Error thrown from {#each forms.find} if we don't temporarily set this to null?
        forms = [...forms, form];
        activeFormId = form.form_id;

        addInput();
    }

    async function deleteForm(id) {
        const res = await axios.delete(`${API_URL}/api/${guildId}/forms/${id}`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        notifySuccess(`Form deleted successfully`);

        activeFormId = null;
    }

    function addInput() {
        const form = getForm(activeFormId);
        if (form.inputs.length >= 5) return;

        const input = {
            form_id: activeFormId,
            position: form.inputs.length + 1,
            style: "1",
            label: "",
            placeholder: "",
            required: true,
            is_new: true,
        };

        form.inputs = [...form.inputs, input];
        forms = forms;
    }

    async function editInput(formId, inputId, data) {
        let mapped = {...data, style: parseInt(data.style)};

        const res = await axios.patch(`${API_URL}/api/${guildId}/forms/${formId}/${inputId}`, mapped);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        let form = getForm(formId);
        form.inputs = form.inputs.filter(input => input.id !== inputId);
        form.inputs = [...form.inputs, res.data];

        notifySuccess('Form input updated successfully');
    }

    async function deleteInput(formId, input) {
        let form = getForm(formId);

        let idx = form.inputs.findIndex((i) => i === input);
        form.inputs.splice(idx, 1);
        for (let i = idx; i < form.inputs.length; i++) {
            form.inputs[i].position--;
        }

        forms = forms;

        if (!input.is_new) {
            if (toDelete[formId] === undefined) {
                toDelete[formId] = [];
            }

            toDelete[formId] = [...toDelete[formId], input.id];
        }
    }

    function changePosition(formId, input, direction) {
        const form = getForm(formId);
        let idx = form.inputs.findIndex((i) => i === input);

        let inputs = form.inputs;
        if (direction === "up") {
            [inputs[idx - 1].position, inputs[idx].position] = [inputs[idx].position, inputs[idx - 1].position];
            [inputs[idx - 1], inputs[idx]] = [inputs[idx], inputs[idx - 1]];
        } else if (direction === "down") {
            [inputs[idx + 1].position, inputs[idx].position] = [form.inputs[idx].position, form.inputs[idx + 1].position];
            [inputs[idx + 1], inputs[idx]] = [form.inputs[idx], form.inputs[idx + 1]];
        }

        forms = forms;
    }

    async function saveInputs() {
        const form = getForm(activeFormId);

        const data = {
            "create": form.inputs.filter(i => i.is_new === true)
                .map(i => ({...i, style: parseInt(i.style), placeholder: nullIfBlank(i.placeholder)})),
            "update": form.inputs.filter(i => !i.is_new)
                .map(i => ({...i, style: parseInt(i.style), placeholder: nullIfBlank(i.placeholder)})),
            "delete": toDelete[activeFormId] || [],
        }

        const res = await axios.patch(`${API_URL}/api/${guildId}/forms/${activeFormId}/inputs`, data);
        if (res.status !== 204) {
            notifyError(res.data.error);
            return;
        }

        toDelete = {};

        const formId = activeFormId;
        await loadForms();
        activeFormId = formId;

        notifySuccess('Form updated successfully');
    }

    async function loadForms() {
        const res = await axios.get(`${API_URL}/api/${guildId}/forms`);
        if (res.status !== 200) {
            notifyError(res.data.error);
            return;
        }

        forms = res.data || [];
        forms.flatMap(f => f.inputs).forEach(i => {
            i.style = i.style.toString();
        });

        // if (forms.length > 0) {
        //    activeFormId = forms[0].form_id;
           // no longer want an active form to begin with 
        // }
    }

    withLoadingScreen(async () => {
        setDefaultHeaders();
        await loadForms();
    });
</script>

<style>
    .parent {
        width: 100%;
        height: 100%;
    }

    .page-title-row{
        display: flex;
    }

    :global(button){
        width: 175px;
        padding: 5px !important;
        height: 40px;
    }

    .page-title-wrapper .title-dot{
        width: 75px;
    }

    .refresh-wrapper :global(.danger){
        margin-left: 10px;
    }

    .content {
        display: flex;
        justify-content: space-between;
        width: 100%;
        height: 100%;
        margin-top: 30px;
        margin-bottom: 50px;
    }

    .refresh-wrapper{
        display: flex;
        width: 100%;
    }

    .refresh-button{
        height: 35px;
        width: 45px;
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

    .new-input{
        width: 100%;
        text-align: center;
        margin-top: 5px;
        background-color: var(--fg-color);
        box-shadow: var(--shadow);
        padding: 10px;
        border-radius: 6px;
        cursor: pointer;
    }

    .body-wrapper {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        padding: 0 1% 1% 1%;
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

    .form-input-row{
        background-color: var(--fg-color);
        box-shadow: var(--shadow);
        padding: 10px;
        border-radius: 6px;
        margin-bottom: 10px;
    }

    .section-title {
        font-size: 36px;
        font-weight: bolder !important;
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

    h3 {
        font-size: 28px;
        margin-bottom: 4px;
    }

    hr.fill {
        border-top: 1px solid #777;
        border-bottom: 0;
        border-left: 0;
        border-right: 0;
        flex: 1;
    }

    .row {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 100%;
        height: 100%;
    }

    .form-select-row {
        justify-content: flex-start;
        gap: 12px;
        max-height: 40px;
    }

    .form-name-edit-wrapper {
        justify-content: flex-start;
        gap: 12px;
    }

    .form-name-save-wrapper {
        height: 48px;
        align-self: flex-end;
    }

    .multiselect-super {
        width: 31%;
    }

    .manage {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        width: 100%;
        height: 100%;
        margin-top: 12px;
    }

    .add-input-container {
        display: flex;
        align-items: center;
        gap: 4px;
        width: unset !important;
    }

    .add-input-container > * {
        cursor: pointer;
    }

    .add-input-disabled > *{
        cursor: default !important;
        color: #777 !important;
    }

    #creation-row {
        justify-content: flex-start !important;
    }

    #create-button-wrapper {
        margin-left: 15px;
        height: 40px;
    }

    @media only screen and (max-width: 950px) {
        .manage {
            flex-direction: column;
        }

        .row {
            flex-direction: column;
        }

        #create-button-wrapper {
            margin-left: unset;
        }

        .form-select-row {
            max-height: unset;
            gap: 8px;
        }

        .multiselect-super {
            width: 100%;
        }

        .form-name-edit-wrapper {
            gap: unset;
        }

        .form-name-save-wrapper {
            width: 100%;
        }
    }
</style>
