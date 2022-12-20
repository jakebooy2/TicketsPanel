<div class="wrapper">
  {#if imageUrl !== null}
    <img src={imageUrl} class="logo" bind:this={logo} on:error={useDefaultLogo}/>
  {:else}
    <img src="/assets/img/grey.png" class="logo"/>
  {/if}
  <div class="details">
    <div class="title-row">
      <span class="title">{name}</span>
      {#if builtIn} 
        <Badge colour="#0c8f43">Built-In</Badge>
      {/if}
      {#if added}
        <Badge>Active</Badge>
      {/if}
      {#if guildCount !== undefined}
        <Badge>
          <div class="guild-count">
            <i class="fas fa-server"></i>
            {guildCount}
          </div>
        </Badge>
      {/if}
    </div>

    {#if showAuthor && author}
      <div class="author">
        <a href="https://discord.com/users/{author.id}" class="link" style="gap: 4px">
          <img src="https://cdn.discordapp.com/avatars/{author.id}/{author.avatar}.webp" class="author-avatar"
               alt="Author avatar" on:error={(e) => handleAvatarError(e, author.discriminator)}/>
          <b>{author.username}#{author.discriminator}</b>
        </a>
      </div>
    {:else if showAuthor}
      <div class="author">
        <a href="https://discord.com/users/{ownerId}" class="link" style="gap: 4px">
          <img src="https://cdn.discordapp.com/embed/avatars/0.png" class="author-avatar"
               alt="Author avatar" />
          <b>Unknown User</b>
        </a>
      </div>
    {/if}

    <span class="description">
      <slot name="description"></slot>
    </span>

    {#if !hideLinks}
      <div class="links">
        {#if builtIn}
          <a href="{viewLink}" target="_blank" class="integration-link single">View</a>
        {:else if added}
          <Navigate to="/manage/{guildId}/integrations/view/{integrationId}" styles="integration-link">View</Navigate>
          <Navigate to="/manage/{guildId}/integrations/manage/{integrationId}" styles="integration-link">Configure</Navigate>
          <a href="#" class="integration-link" on:click={() => dispatch("remove", {})}>Remove</a>
        {:else}
          {#if owned}
            <Navigate to="/manage/{guildId}/integrations/view/{integrationId}" styles="integration-link">Preview</Navigate>
            <Navigate to="/manage/{guildId}/integrations/configure/{integrationId}" styles="integration-link">Configure</Navigate>
            <Navigate to="/manage/{guildId}/integrations/activate/{integrationId}" styles="integration-link">Add to server</Navigate>
        {:else}
            <Navigate to="/manage/{guildId}/integrations/view/{integrationId}" styles="integration-link double">View</Navigate>
            <Navigate to="/manage/{guildId}/integrations/activate/{integrationId}" styles="integration-link double">Add to server</Navigate>
          {/if}
        {/if}
      </div>
    {/if}
  </div>
</div>

<script>
    import Badge from "../Badge.svelte";
    import {Navigate} from "svelte-router-spa";
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    export let guildId;

    export let integrationId;
    export let name;
    export let imageUrl;
    export let ownerId;
    export let owned = false;
    export let guildCount;
    export let author;
    export let showAuthor = false;

    export let added = false;
    export let builtIn = false;
    export let hideLinks = false;
    export let viewLink;

    let logo;

    function useDefaultLogo() {
        logo.src = "/assets/img/grey.png";
    }

    function handleAvatarError(ev, discriminator) {
        const src = `https://cdn.discordapp.com/embed/avatars/${(discriminator || 0) % 5}.png`;
        if (ev.target.src === src) { // Setting onerror to null does not work with svelte
            return;
        }

        ev.target.src = src;
    }
</script>

<style>
    .wrapper {
        display: flex;
        flex-direction: column;
        border-radius: 10px;

        background-color: var(--fg-color) !important;
        box-shadow: var(--shadow);
        transition: all .3s ease-in-out;

        height: 100%;
    }

    .logo {
        width: 100%;
        border-radius: 10px 10px 0 0;
        min-height: 150px;
        height: 150px;
        max-height: 150px;
        object-fit: cover;
    }

    .details {
        display: flex;
        flex-direction: column;
        padding: 10px 20px 10px 20px;
        height: 100%;
    }

    .title-row {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;
    }

    .title {
        font-size: 18px;
        font-weight: bolder;
        word-break: break-all;
    }

    .description {
        flex: 1;
        font-size: 14px;
        color: #b9bbbe;
    }

    .guild-count {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 4px;
    }

    .links {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-around;
        gap: 4px;
    }

    .author-avatar{
        margin-top: 5px;
    }

    .link{
        color: white;
        margin-top: -5px;
    }

    .author {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 4px;
        color: white;
    }

    .author-avatar {
        height: 24px;
        border-radius: 50%;
    }

    :global(.integration-link){
        background-color: var(--primary-color);
        margin-top: 10px;
        text-align: center;
        text-decoration: none;
        color: white;
        border-radius: 6px;
        padding: 5px 10px;
    }

    :global(.single){
        width: 100%;
        display: block;
    }

    :global(.double){
        width: 50%;
        display: block;
    }
</style>