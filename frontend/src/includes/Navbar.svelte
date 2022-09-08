<script>
    import {Navigate} from "svelte-router-spa";
</script>

<div class="navbar">
  <div class="flex-center">
    <img src="https://ticketsbot.net/assets/img/logo-trans.webp" class="logo" />
  </div>

  <div class="links">
    <a href="/">Home</a>
    <a href="https://docs.ticketsbot.net">Documentation</a>
    <a href="https://ticketsbot.net/premium">Premium</a>
  </div>
  <div class="flex-right"></div>
</div>

<style>
  .navbar {
      display: flex;
      flex-direction: row;
      gap: 24px;

      width: 100%;

      background-color: var(--fg-color);
      border-radius: 6px;

      padding: 10px 20px 7px 20px;
  }

  .flex-center {
      display: flex;
      justify-content: center;
      align-items: center;
  }

  .flex-right{
    display: flex;
    justify-content: right;
    align-items: center;
    float: right;
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
      width: 87%;
  }

  a {
      color: white;
      text-decoration: none;
      cursor: pointer;
      text-align: center;
      font-size: 16px;
      transition: .2s ease-in-out;
  }
  a:hover{
    color: var(--primary-color);
  }
</style>
