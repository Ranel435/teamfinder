<script lang="ts">
  interface achievement {
    hackaton: string;
    place: number;
    team: string;
    date: string;
    location: string;
  }


  let achievements: achievement[] = [
    {
      hackaton: "VTB API",
      place: 2,
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      place: 2,
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      place: 2,
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      place: 2,
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      place: 2,
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      place: 2,
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    }
  ];
</script>



<section class="achievements">
  {#if achievements.length === 0}
  <div class="no-achievements">
    <p class="no-achievements-title">Достижения не добавлены</p>
    <p class="no-achievements-text">Добавьте победу или участие в хакатоне
      чтобы пользователи больше знали о вашем опыте</p>
    <div class="no-achievements-buttons">
      <button>Добавить</button>
    </div>
  </div>
  {:else}
  <div class="achievements-container">
    <div class="achievements-header">
      <h2>Все достижения</h2>
    </div>
    <ul class="achievements-list">
      {#each achievements as achievement, index}
      <div class="achievement">
        <p>{achievement.place} * {achievement.hackaton}</p>
        <p>В составе команды: {achievement.date}</p>
      </div>
      {/each}
    </ul>
    <div class="achievements-buttons">
      <button>Все команды</button>
      <button>Создать</button>
      <button>Хакатоны</button>
    </div>
  </div>
  {/if}
</section>

<style>
  .achievements {
    height: 100%;
    margin-left: 100px;
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .no-achievements {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 100%;
  }

  .no-achievements-title {
    font-family: "Manrope";
    font-size: 48px;
    color: #fff;
    margin-bottom: 32px;
  }

  .no-achievements-text {
    font-family: "Manrope";
    font-size: 24px;
    color: var(--dark-grey);
    margin-bottom: 32px;
  }

  .no-achievements-buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .no-achievements-buttons button {
    font-family: "Manrope";
    font-size: 24px;
    color: #fff;
    padding: 10px 24px;
    border: 3px solid #fff;
    border-radius: 50px;
    background-color: transparent;
    cursor: pointer;
    margin-right: 16px;
    transition: all 0.5 ease;
  }

  .no-achievements-buttons button:last-child {
    margin-right: 0;
  }

  .no-achievements-buttons button:hover {
    background-color: #fff;
    color: black;
  }
</style>