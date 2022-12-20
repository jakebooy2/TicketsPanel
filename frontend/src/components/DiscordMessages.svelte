<div>
<Card footer={false}>
    <span slot="title">#ticket-{ticketId}</span>
    <div slot="body" class="body-wrapper">
        <div class="section discord-chat">
            <div class="chat-container" bind:this={container}>
                <div>
                    {#each messages as message}
                    <div class="discord-message">
                        <img src={message.author.bot ? 'https://ticketsbot.net/assets/img/logo-trans.webp' : 'https://cdn.discordapp.com/avatars/' + message.author.id + '/' + message.author.avatar + '.jpg?size=64'} class="discord-image" />
                        
                        <div class="discord-message-body">
                            <div class="discord-username">{message.author.username}</div>
                            <div class="discord-message-content">{message.content}</div>
                        </div>
                    </div>
                    {/each}
                </div>
            </div>
            <div class="input-container">
                <div class="input-container-wrapper">
                    <form on:submit|preventDefault={sendMessage}>
                        <input type="text" class="message-input" bind:value={sendContent} disabled={!isPremium}
                            placeholder="{isPremium ? `Message #ticket-${ticketId}` : 'Premium users can receive messages in real-time and respond to tickets through the dashboard'}">
                        <button type="submit" class="context-button">
                            <i class="fa fa-paper-plane" aria-hidden="true"></i>
                        </button>
                    </form>
                    <div class="close-button" on:click={openCloseTicket}>
                        Close
                    </div>
                </div>
            </div>
        </div>     
    </div>
</Card>
</div>
<script>
    import Card from "../components/Card.svelte";
    import {createEventDispatcher} from "svelte";

    export let ticketId;
    export let isPremium = false;
    export let messages = [];
    export let container;

    const dispatch = createEventDispatcher();
    let sendContent = '';
    let loadMessages = '';

    function sendMessage() {
        dispatch('send', sendContent);
        sendContent = '';
    }

    function openCloseTicket(){
        dispatch('closemodal', {});
    }
</script>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Catamaran:wght@300;400;500;600;700;800&display=swap');

    .discord-container {
        display: flex;
        flex-direction: column;

        background-color: #2e3136;
        border-radius: 4px;
        height: 80vh;
        max-height: 100vh;
        margin: 0;
        padding: 0;
        font-family: 'Catamaran', sans-serif !important;
        display: none;
    }

    .channel-header {
        display: flex;
        align-items: center;

        background-color: #1e2124;
        height: 5vh;
        width: 100%;
        border-radius: 4px 4px 0 0;
        position: relative;

        text-align: center;
    }

    #channel-name {
        color: white;
        font-weight: bold;
        padding-left: 20px;
    }

    #message-container {
        display: flex;
        flex-direction: column;
        flex: 1;

        position: relative;
        overflow-y: scroll;
        overflow-wrap: break-word;
    }

    .message {
        color: white !important;
        padding-left: 20px;
    }

    #message-container:last-child {
        margin-bottom: 5px;
    }

    .input-container{
        padding: 7px 14px;
        position: relative;
        display: block;
    }

    .input-container-wrapper{
        display: flex;
        width: 100%;
        gap: 10px;
    }

    .input-container-wrapper form{
        width: 100%;
        position: relative;
        display: block;
    }

    .message-input {
        display: absolute;

        font-size: 16px;
        padding: 10px 15px;
        border-radius: 6px;

        border: none !important;
        background-color: rgba(255, 255, 255, .06) !important;
        color: white !important;
        width: 100%;
        outline: none;
        transition: box-shadow .2s ease-in-out;
    }

    .message-input:focus, .message-input:focus-visible {
        outline-width: 0;
        box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1);
    }

    .context-button{
        position: absolute;
        right: 5px;
        top: 1px;
        height: 45px;
        width: 45px;
        text-align: center;
        cursor: pointer;
        background: none !important;
        padding: 0 !important;
        border: none !important;
        color: white !important;
    }

    .context-button .fa{
        line-height: 45px;
        font-size: 20px;
        transition: .2s ease-in-out;
    }

    .context-button:hover .fa{
        color: var(--primary-color);
    }

    .close-button{
        height: 45px;
        line-height: 45px;
        background-color: #dc3545 !important;
        padding: 0 15px;
        margin-top: 2px;
        border-radius: 6px;
        box-shadow: 0 6px 6px rgba(10, 10, 10, .1), 0 0 0 1px rgba(10, 10, 10, .1);
        font-weight: bold;
        cursor: pointer;
        transition: .2s ease-in-out;
    }

    .close-button:hover{
        background-color: #c82e3e !important;
    }

    .discord-chat{
        width: 100%;
    }

    .chat-container{
        max-height: 600px;
        overflow-y: auto;
        display: flex;
        flex-direction: column-reverse;
    }

    .discord-message{
        padding: 7px 14px;
        border-radius: 6px;
        display: flex;
        width: 100%;
        flex-shrink: 0;
        flex-grow: 1;
        transition: .2s ease-in-out;
        margin-bottom: 5px;
        cursor: default !important;
    }

    .discord-message:hover{
        background-color: rgba(255, 255, 255, .04);
    }

    .discord-image{
        width: 42px;
        height: 42px;
        border-radius: 100%;
        background-color: rgba(11, 11, 11, .2);
        margin-right: 15px;
    }

    .discord-username{
        font-weight: bold;
    }

    .discord-message-content{
        color: rgba(255, 255, 255, .7);
        font-size: 14px;
    }
</style>
