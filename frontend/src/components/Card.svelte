<script>
    export let footer = true;
    export let fill = true;
    export let footerRight = false;
    export let dropdown = false;
    export let ref;

    let dropdownActive = false;
</script>

<div class="card" class:fill>
  <div class="card-header" class:dropdown on:click={() => dropdownActive = dropdown && !dropdownActive}>
    <h4 class="card-title">
      <slot name="title">
        No Title :(
      </slot>
    </h4>
  </div>
  <div class="card-body" class:dropdown class:dropdownActive class:dropdownInactive={dropdown && !dropdownActive} {ref}>
    <div class="inner" class:dropdown>
      <slot name="body">
        No Content :(
      </slot>
    </div>
  </div>

  {#if footer}
    <div class="card-footer">
      <div class="footer-content" class:footerRight>
        <slot name="footer" />
      </div>
    </div>
  {/if}
</div>

<style>
    .card {
        display: flex;
        flex-direction: column;
        width: 100%;
        border-radius: 5px;
        box-shadow: 0 4px 4px rgba(0, 0, 0, 0.25);
        transition: all .3s ease-in-out;
        margin-bottom: 0 !important;
        background-color: var(--fg-color);
    }

    .fill {
        height: 100%;
    }

    .card-title {
        color: white;
        font-size: 20px;
        font-weight: bold !important;

        font-family: 'Poppins', sans-serif;
        padding: 10px 15px;
        margin: 0;
    }

    .card-header {
        display: flex;

        margin: 7px;
        background-color: rgba(255, 255, 255, .06);
        border-radius: 4px;
    }

    .card-header.dropdown {
        cursor: pointer;
        user-select: none;
    }

    .card-body {
        display: flex;
        flex: 1;
        min-height: 75px;

        color: white;
        margin: 10px 10px;
    }

    .inner {
        display: flex;
        height: 100%;
        width: 100%;
    }

    .inner.dropdown {
        position: absolute;
    }

    .card-body.dropdown {
        position: relative;
        transition: min-height .3s ease-in-out, margin-top .3s ease-in-out, margin-bottom .3s ease-in-out;
    }

    .card-body.dropdownInactive {
        height: 0;
        visibility: hidden;

        margin: 0;
        flex: unset;
        min-height: 0 !important;
    }

    .card-body.dropdownActive {
        visibility: visible;
        min-height: auto;
        overflow: hidden;
    }

    .card-footer {
        display: flex;
        color: white;
        border-top: 1px solid rgba(0, 0, 0, .125);
        padding: 10px 20px;
    }

    .footer-content {
        display: flex;
        align-items: center;
        height: 100%;
        width: 100%;
    }

    .footerRight {
        flex-direction: row-reverse;
    }

    :global(div [slot=footer]) {
        display: flex;
        flex-direction: row;
    }

    .inner > * {
        width: 100%;
    }
</style>