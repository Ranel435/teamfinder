<script>
  import Sidebar from '$lib/components/sidenav.svelte';
  import "../../../app.css"
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
    margin: 100px 84px;
  }

  main {
    flex: 1;
  }
</style>