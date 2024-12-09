<script>
  import Sidebar from '$lib/sidenav.svelte';
</script>

<div class="app">
  <Sidebar />
  <main>
    <slot />
  </main>
</div>

<style>
  .app {
    display: flex;
    margin: 128px 0;
  }

  main {
    flex: 1;
  }
</style>