<script lang="ts">
  import "../../app.css"
  import { page } from '$app/stores';

  const navItems = [
    { name: 'Мой аккаунт', path: '/account' },
    { name: 'Уведомления', path: '/notifications' },
    { name: 'Мои анкеты', path: '/forms'},
    { name: 'Мои команды', path: '/teams' },
    { name: 'Я участвую', path: '/participate' },
    { name: 'Мои достижения', path: '/achievements' }
  ];

  $: currentPath = $page.url.pathname;
</script>

<nav class="sidebar">
  <div class="sidebar-item">
    {#each navItems as item}
    
    <a
      href={"/profile"+item.path}
      class:active={currentPath === item.path}
    ><img src="/profile-nav{item.path}.svg" style="width: 20px; height: 20px;" alt="">
      {item.name}
    </a>
    {/each}
  </div>
  
</nav>

<style>
  .sidebar {
    display: flex;
    flex-direction: column;
    font-family: "Manrope";
    font-size: 20px;
    font-weight: bold;
    color: #fff;
    text-decoration: none;
    margin-left: 20px;
  }

  .sidebar-item {
    display: flex;

    flex-direction: column;

  }

  .sidebar a {
    color: var(--light-grey);
    text-decoration: none;
    font-size: 20px;
    margin: 10px 0;
    transition: font-size 0.3s ease;
  }

  .sidebar a.active {
    font-size: 20px; 
    font-weight: bold;
  }

  .sidebar a:hover {
    color: #fff;
    font-size: 24px; /* Add a hover effect for better interactivity */
  }
</style>
