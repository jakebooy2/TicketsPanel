<script>
    import {Navigate} from "svelte-router-spa";
</script>

<div class="navbar">
  <div class="flex-center">
    <img src="/assets/img/logo-trans.png" class="logo" />
  </div>

  <div class="links">
    <a href="/">Home</a>
    <a href="https://docs.ticketsbot.net">Documentation</a>
    <a href="https://ticketsbot.net/premium">Premium</a>
  </div>
</div>

<style>
  .navbar {
      display: flex;
      flex-direction: row;
      gap: 24px;

      width: 100%;

      background-color: var(--fg-color);
      border-radius: 10px;

      padding: 5px 20px;
  }

  .flex-center {
      display: flex;
      justify-content: center;
      align-items: center;
  }

  .logo {
      width: 48px;
      height: 48px;
  }

  .links {
      display: flex;
      flex-direction: row;
      align-items: center;
      gap: 20px;
  }

  a {
      color: white;
      text-decoration: none;
      cursor: pointer;
      text-align: center;

      font-size: 20px;
      font-weight: bold;
  }
</style>
