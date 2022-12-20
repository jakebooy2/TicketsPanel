<div class:inline class="toggle-box">
    
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
  .toggle-box{
    display: flex;
    gap: 15px;
    align-items: center;
    background-color: rgba(255, 255, 255, .08) !important;
    padding: 7px 10px;
    border-radius: 4px;
    margin-top: -5px;
  }

   .form-label{
        color: rgba(255, 255, 255, .8);
    }
</style> 