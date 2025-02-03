<script lang="ts">
  import "../../app.css";
  import { page } from '$app/stores';

  const navItems = [
    { name: "Мой аккаунт", path: "/account" },
    { name: "Уведомления", path: "/notifications" },
    { name: "Мои анкеты", path: "/forms" },
    { name: "Мои команды", path: "/teams" },
    { name: "Я участвую", path: "/participate" },
    { name: "Достижения", path: "/achievements" },
  ];
   $: currentPath = $page.url.pathname;
</script>

<nav class="sidenav">
  <div class="sidenav__items">
    {#each navItems as item}
      <div class="sidenav__item">
        <a
          href={"/profile" + item.path}
          class="sidenav__link"
        >
          <img
            src="/profile-nav{item.path}.svg"
            style="width: 20px; height: 20px;"
            alt=""
          />
          {item.name}
        </a>
      </div>
    {/each}
  </div>
  {#if $page.url.pathname != "/profile"+navItems[0].path}
    <div class="line"></div>
  {/if}
</nav>

<style>
  .sidenav {
    display: flex;
    flex-direction: row;
    justify-content: center;
    font-family: "Manrope-bold";
    font-size: 20px;
    color: #fff;
    text-decoration: none;
    margin-left: 20px;
  }

  .sidenav__items {
    display: flex;
    flex-direction: column;
  }

  .sidenav__item{
    margin: 15px 0;
    width: 190px;
    height: 30px;
  }

  .sidenav__link {
    color: var(--light-grey);
    text-decoration: none;
    font-size: 20px;
    transition: font-size 0.3s ease;
  }

  .sidenav__link.active {
    font-size: 24px;
    font-weight: bold;
  }

  .sidenav__link:hover {
    color: #fff;
    font-size: 24px;
  }

  .line {
    margin-left: 36px;
    width: 3px;
    height: 100%;
    background-color: #fff;
  }
</style>
