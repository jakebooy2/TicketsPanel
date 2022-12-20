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
  <div class="flex-right">
    <div class="user-info">
      <img src="https://cdn.discordapp.com/avatars/601612991422791726/3164a94600ddc0fca86def35e00756b4.webp?size=512" />
    </div>
    <div class="user-name">
      Jake S.
    </div>
    <div class="logout">
      <i class="fa-solid fa-right-from-bracket"></i>
    </div>
  </div>
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
      box-shadow: var(--shadow);
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
    margin-top: -3px;
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
      margin-top: -3px;
  }

  a {
      color: white;
      text-decoration: none;
      cursor: pointer;
      text-align: center;
      font-size: 15px;
      transition: .2s ease-in-out;
  }
  a:hover{
    color: var(--primary-color);
  }

  .user-info{
    display: flex;
    align-items: center;
  }

  .user-info img{
    width: 44px;
    height: 44px;
    border-radius: 100%;
  }

  .user-info:hover + .user-name,
  .user-name:hover{
    background:  rgba(255, 255, 255, .08);
  }

  .user-name{
    border-radius: 25px;
    background: rgba(255, 255, 255, .04);
    padding: 10px 15px;
    overflow: hidden;
    min-width: 90px;
    margin-left: 5px;
    cursor: default;
    transition: .2s ease-in-out;
    margin-top: 2px;
  }

  .logout{
    width: 44px;
    height: 44px;
    border-radius: 100%;
    margin-left: 5px;
    text-align: center;
    background: rgba(255, 255, 255, .04);
    cursor: pointer;
    transition: .2s ease-in-out;
  }

  .logout:hover{
    background: #dc3545 !important;
  }

  .logout i{
    font-size: 20px;
    line-height: 46px;
  }
</style>
