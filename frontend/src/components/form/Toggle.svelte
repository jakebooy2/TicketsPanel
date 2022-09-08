<div class:inline>
  {#if label !== undefined}
    <label class="form-label" style="margin-bottom: 0 !important;">{label}</label>
  {/if}
  <Toggle hideLabel
          toggledColor="{toggledColour}"
          untoggledColor="{untoggledColour}"
          bind:toggled={value}/>
</div>

<script>
    import Toggle from "svelte-toggle";

    export let value;
    export let label;

    export let toggledColour = "#995DF3";
    export let untoggledColour = "rgba(255, 255, 255, .15)";
    export let inline = false;
</script>

<style>
  div {
      display: flex;
      flex-direction: column;
  }

  .form-label{
    padding-bottom: 7px !important;
  }

  .inline {
      flex-direction: row !important;
      gap: 4px;
  }
</style>