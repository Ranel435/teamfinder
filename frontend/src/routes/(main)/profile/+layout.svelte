<script>
  import Sidebar from '$lib/components/sidenav.svelte';
  import "../../../app.css"
</script>

<div class="app">
  <Sidebar />
  <div style="width: 1px; height: 100%; background-color: #fff;"></div>
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