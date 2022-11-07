<div class:col-1={col1} class:col-2={col2} class:col-3={col3} class:col-4={col4}>
    <label for="input" class="form-label">{label}</label>
    <div class="color-picker-wrapper">
        <button type="button" class="color-picker-button" style="background: {color}" on:click={toggleOpen}></button>
        <div class="color-picker-box{open ? ' open' : ''}" style="position: absolute;">
            <ChromaPicker bind:color />
        </div>
    </div>
</div>

<style>
    input {
        width: 100%;
    }

    .color-picker-wrapper{
        width: 100%;
    }

    .color-picker-button{
        border: 1px solid transparent;
        border-radius: 4px;
        width: 100%;
        height: 35px;
        cursor: pointer;
        transition: .2s ease-in-out;
    }
    
    .color-picker-button:hover{
        border-color: white;
    }

    .color-picker-box{
        background: var(--bg-color);
        border-radius: 4px;
        padding: 5px 5px 1px 5px;
        display: none;
        box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.5);
        margin-bottom: 15px;
        z-index: 999;
    }

    .color-picker-box.open{
        display: block;
    }
</style>

<script>
    import ChromaPicker from 'svelte-chroma-picker';

    export let color;

    export let label;
    export let disabled = false;

    export let col1 = false;
    export let col2 = false;
    export let col3 = false;
    export let col4 = false;

    let open = false;

    function toggleOpen(){
        open = !open;
    }
</script>
