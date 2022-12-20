<button on:click isTrigger="1" class:gap class:fullWidth class:danger class:iconOnly {disabled} {type}>
  {#if icon !== undefined}
    <i class="{icon}"></i>
  {/if}
  {#if !iconOnly}
    <span class="content">
      <slot/>
    </span>
  {/if}
</button>

<script>
    export let icon;
    export let fullWidth = false;
    export let disabled = false;
    export let type = "submit";
    export let danger = false;
    export let iconOnly = false;
    export let gap = false;
</script>

<style>
    button {
        display: flex;

        justify-content: center;
        align-items: center;

        text-align: center;

        color: white;
        background-color: #873ef5;
        border: none;
        border-radius: .25rem;
        margin: 0;
        padding: 10px 15px;

        cursor: pointer;
        transition: background-color 150ms ease-in-out, border-color 150ms ease-in-out;
        box-shadow: 0 4px 4px rgb(0 0 0 / 25%);
    }

    button:active, button:hover:enabled {
        background-color: #8a45f1;
        border-color: #8a45f1;
    }

    button:disabled {
        background-color: rgba(255, 255, 255, .2);
        cursor: default;
        opacity: .4;
    }

    .gap{
        margin-right: 10px;
    }

    .content {
        margin-left: 5px;
        margin-right: 5px;
    }

    .fullWidth {
        width: 100%;
    }

    .danger {
        background-color: #dc3545 !important;
        border-color: #dc3545 !important;
    }

    .danger:hover:enabled, .danger:active {
        background-color: #c32232 !important;
        border-color: #c32232 !important;
    }

    .iconOnly {
        width: 40px;
        height: 40px;
    }
</style>