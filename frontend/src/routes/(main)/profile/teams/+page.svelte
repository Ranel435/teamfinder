<script lang="ts">
  interface Team {
    name: string;
    
  }
  let teams: any[] = [];
</script>




<section class="teams">
  {#if teams.length === 0}
  <div class="no-teams">
    <h2 class="no-teams__title m1">Команд пока нет</h2>
    <p class="no-teams__text m5">Вступите в команду чтобы учавствовать в хакатоне</p>
    <div class="no-teams__buttons">
      <button class="no-teams__button gradient-button m5">Создать</button>
      <button class="no-teams__button black-white-button m5">Поиск команды</button>
    </div>
  </div>
  {:else}
  <div class="teams__container">
    <div class="teams__as-lead-teams as-lead-teams">
      <header class="as-lead-teams__header">
        <h2 class="as-lead-teams__title m1">Мои команды</h2>
        <button class="as-lead-teams__button gradient-button m5">Создать</button>
      </header>
      <ul class="as-lead-teams__list">
        {#each teams as team, index}

        {/each}
      </ul>
    </div>

    <div class="teams__as-player-teams as-player-teams">
      <header class="as-player-teams__header m1">В составе</header>
    </div>
  </div>
  {/if}
</section>




<style>
  .teams {
    height: 100%;
    margin-left: 50px;
    display: flex;
    flex-direction: column;
  }

  .no-teams {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }

  .no-teams__title {
    color: var(--white);
  }

  .no-teams__text {
    color: var(--white);
    margin: 22px 0 32px;
  }
  .no-teams__buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .no-teams__button {
    padding: 6px 28px;
    margin-right: 16px;
    transition: all 0.5 ease;
  }

  .no-teams__button:last-child {
    margin-right: 0;
  }

/* /////////////////////////////////////////////////////////////////// */


.teams__container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 32px 0 0 32px;
}


</style>
