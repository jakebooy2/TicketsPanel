<label class="form-label">{label}</label>
<div class="multiselect-super">
  <Select placeholder="Select..." items={panels} optionIdentifier="panel_id" getOptionLabel={labelMapper}
          getSelectionLabel={labelMapper} bind:selectedValue={selectedRaw}
          on:select={update} on:clear={handleClear} {isMulti} {isSearchable}
          placeholderAlwaysShow={true} />
</div>

<script>
    import Select from 'svelte-select';
    import {onMount} from "svelte";

    export let label;
    export let panels;
    export let selected;
    export let isMulti = true;
    export let isSearchable = false;

    let selectedRaw = isMulti ? panels.filter((p) => selected.includes(p.panel_id)) : selected;

    function labelMapper(panel) {
        return panel.title || "";
    }

    function update() {
        if (selectedRaw === undefined) {
            selectedRaw = [];
        }

        if (isMulti) {
            selected = selectedRaw.map((panel) => panel.panel_id);
        } else {
            if (selectedRaw) {
                selected = selectedRaw.panel_id;
            } else {
                selected = undefined;
            }
        }
    }

    function handleClear() {
        if (isMulti) {
            selected = [];
        } else {
            selected = undefined;
        }
    }

    function applyOverrides() {
        if (isMulti) {
            //selected = [];
            selectedRaw = panels.filter((p) => selected.includes(p.panel_id));
        } else {
            if (selectedRaw) {
                selectedRaw = selectedRaw.panel_id;
            }
        }
    }

    onMount(() => {
        applyOverrides();
    });
</script>

<style>
  :global(.listContainer){
    padding: 5px;
    background: #383844 !important;
    box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1) !important;
  }
  :global(.listContainer) :global(.item){
    border-radius: 4px !important;
    transition: .2s ease-in-out !important;
  }
  :global(.listContainer) :global(.item):hover{
    cursor: pointer;
  }
  :global(.listContainer) :global(.item.hover){
    background: rgba(255, 255, 255, .08) !important;
  }

  .multiselect-super{
    --background: rgba(255, 255, 255, .08);
    --border: none;
    --placeholderColor: white;
    --clearSelectFocusColor: white;

    --padding: 5px;
  }
</style>