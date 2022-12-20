<div class="server-box" on:click={invite}>
    <div class="server-info">
        <div class="invite-icon">
            <i class="fas fa-plus fa-2x guild-icon-fa"></i>
        </div>
        <div class="server-name">Invite to your server</div>
    </div>
</div>

<style>
    :global(.invite-icon){
        width: 52px;
        height: 52px;
        box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1);
        border-radius: 100%;
        background-color: var(--bg-color);
        text-align: center;
    }

    :global(.invite-icon) :global(i){
        line-height: 52px;
        font-size: 20px;
    }
</style>

<script>
    function invite() {
        window.location.href = `https://invite.ticketsbot.net`;
    }
</script>
