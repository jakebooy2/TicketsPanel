<div class="badge" style="--badge-background-color: {colour}">
  <slot></slot>
</div>

<script>
  export let colour = '#995df3';
</script>

<style>
    .badge {
        display: flex;
        align-items: center;

        background-color: var(--badge-background-color, #995df3);
        border-radius: 4px;
        font-size: 14px;
        padding: 2px 10px;
        margin-left: 7px;
    }
</style>