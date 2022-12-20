<form on:submit|preventDefault={forwardCreate} class="input-form">
    <div class="input-row-container">
        <div class="input-row">
            <Input  label="Label" bind:value={data.label} placeholder="Name of the field" />
            <Textarea  label="Placeholder" bind:value={data.placeholder} minHeight="90px"
                    placeholder="Placeholder text for the field, just like this text" />
        </div>
        <div class="sidebar-row">
            {#if windowWidth > 950}
                {#if withDirectionButtons}
                    <div class="form-movement-container">
                        {#if withDeleteButton}
                            <form on:submit|preventDefault={forwardDelete} class="button-form">
                                <Button icon="fas fa-trash" danger={true}>Delete</Button>
                            </form>
                        {/if}
                        <div class="form-movement">
                            <form on:submit|preventDefault={() => forwardMove("down")} class="button-form">
                                <Button disabled={index >= formLength - 1}>
                                <i class="fas fa-chevron-down"></i>
                                </Button>
                            </form>
                            <form on:submit|preventDefault={() => forwardMove("up")} class="button-form">
                                <Button disabled={index === 0}>
                                <i class="fas fa-chevron-up"></i>
                                </Button>
                            </form>
                        </div>
                    </div>
                {/if}
                {#if withDeleteButton && !withDirectionButtons}
                <form on:submit|preventDefault={forwardDelete} class="button-form">
                    <Button icon="fas fa-trash" danger={true}>Delete</Button>
                </form>
                {/if}
            {/if}

            <div class="row">
                <Dropdown col1={true} label="Style" bind:value={data.style}>
                  <option value=1 selected>Short</option>
                  <option value=2>Paragraph</option>
                </Dropdown>
              </div>
              <div class="row toggle-box-edit" style="width: 100%;">
                <ToggleBox label="Required" bind:value={data.required}/>
              </div>
        </div>
    </div>
    
  {#if windowWidth <= 950}
    <div class="col-1">
      {#if withDirectionButtons}
        <div class="row">
          <div class="col-2-force">
            <form on:submit|preventDefault={() => forwardMove("down")} class="button-form">
              <Button fullWidth={true} disabled={index >= formLength - 1}>
                <i class="fas fa-chevron-down"></i>
              </Button>
            </form>
          </div>
          <div class="col-2-force">
            <form on:submit|preventDefault={() => forwardMove("up")} class="button-form">
              <Button fullWidth={true} disabled={index === 0}>
                <i class="fas fa-chevron-up"></i>
              </Button>
            </form>
          </div>
        </div>
      {/if}
      <div class="row">
        <div class="col-2-force">
          {#if withDeleteButton}
            <form on:submit|preventDefault={forwardDelete} class="button-form">
              <Button icon="fas fa-trash" danger={true}>Delete</Button>
            </form>
          {/if}
        </div>
      </div>
    </div>
  {/if}
</form>

<svelte:window bind:innerWidth={windowWidth} />

<script>
  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();

  import Input from "../form/Input.svelte";
  import Dropdown from "../form/Dropdown.svelte";
  import Button from "../Button.svelte";
  import Textarea from "../form/Textarea.svelte";
  import Checkbox from "../form/Checkbox.svelte";
  import ToggleBox from '../form/ToggleBox.svelte';

  export let withCreateButton = false;
  export let withDeleteButton = false;
  export let withDirectionButtons = false;
  export let disabled = false;

  export let index;
  export let formLength;

  export let data = {};

  $: windowWidth = 0;

  function forwardCreate() {
    dispatch('create', data);
  }

  function forwardDelete() {
    dispatch('delete', {});
  }

  function forwardMove(direction) {
    dispatch('move', {direction: direction});
  }
</script>

<style>
    .input-form {
        display: flex;
        flex-direction: column;
        width: 100%;
    }

    .input-row-container{
        display: flex;
        gap: 20px;
    }

    .input-row{
        display: flex;
        gap: 5px;
        flex-direction: column;
        width: 100%;
    }

    .sidebar-row{
        display: flex;
        gap: 5px;
        flex-direction: column;
        width: 25%;
    }

    .form-movement-container{
        display: flex;
        gap: 10px;
        margin-top: 22px;
    }

    .toggle-box-edit :global(.toggle-box){
        width: 100%;
        height: 43px;
    }

    .toggle-box-edit :global(.toggle-box) :global(.form-label){
        width: 100%;
    }

    .form-movement{
        display: flex;
        gap: 2px;
    }

    .form-movement :global(button){
        width: 50px;
    }

    .row {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 100%;
        height: 100%;
    }

    .sub-row {
        display: flex;
        flex-direction: row;
    }

    .button-form {
        display: flex;
        flex-direction: column;
        justify-content: flex-end;
        padding-bottom: 0.5em;
    }

    .buttons-row > :not(:last-child) {
        margin-right: 10px;
    }

    @media only screen and (max-width: 950px) {
        .settings-row {
            flex-direction: column-reverse !important;
        }

        .button-form {
            width: 100%;
        }
    }

    :global(.form-input){
        width: 100% !important;
    }


</style>