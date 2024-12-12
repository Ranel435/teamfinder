<script lang="ts">
  let teams: any[] = [];
</script>




<section class="teams">
  {#if teams.length === 0}
  <div class="no-teams">
    <p class="no-teams-text">Команд пока нет</p>
    <div class="no-forms-buttons">
      <button>Создать</button>
      <button>Поиск</button>
    </div>
  </div>
  {:else}
  <div class="teams-container">
    <div class="teams-header">
      <h2>Команды</h2>
    </div>
    <ul class="teams-list">
      {#each teams as team, index}
      <div class="team"teams>
        <div class="team-container">
          <!-- <div class="form-info">
          <h3>{form.hackaton}</h3>
          <p>{form.role} * {form.date}</p>
          <div class="form-info-links">
            <a href="">Отклики: {form.response_count}</a>
            <a href="">Подходящие команды: {form.teams_count}</a>
          </div>
          
          </div>
          <button class="delete-button" on:click={() => { forms.splice(index, 1); forms = [...forms]; }}> Удалить</button> -->
        </div>
      </div>
      {/each}
    </ul>
    <div class="teams-buttons">
      <button>Все команды</button>
      <button>Создать</button>
      <button>Хакатоны</button>
    </div>
  </div>
  {/if}
</section>
