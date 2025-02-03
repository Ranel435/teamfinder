<!-- <script lang="ts">
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
      <button class="gradient-button" >Добавить</button>
    </div>
    <ul class="achievements-list">
      
      {#each achievements as achievement, index}
      <div class="achievement">
        <div class="achievement-container">
          <div class="achievement-container-text">
            <p class="gradient-text">{achievement.place} место * {achievement.hackaton}</p>
            <p style="color: #fff;">В составе команды: {achievement.team}</p>
            <p style="color: #fff;">Дата: {achievement.date}</p>
          </div>
          <div class="achievement-container-buttons">
            <button class="delete-button">Редактировать</button>
            <button class="delete-button" on:click={() => {achievements.splice(index, 1); achievements=[...achievements];}}>Удалить</button>
          </div>
        </div>
      </div>
      {/each}
    </ul>
  </div>
  {/if}
</section>

<style>
  button {
    font-family: "Manrope";
    transition: all 0.5s ease;
  }

  .achievements {
    height: 100%;
    margin-left: 50px;
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

  .achievements-container {
    display: flex;
    flex-direction: column;
    padding: 32px;
  }
  .achievements-header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 40px;
  }

  .achievements-list {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-height: 50vh; /* Фиксированная высота для прокрутки */
    overflow-y: auto; /* Добавление вертикального скроллбара */
    padding-right: 10px; /* Добавляем отступ для скроллбара */
    box-sizing: content-box; /* Учитываем отступ, не влияя на ширину */
  }

  .achievements-header h2{
    font-family: "Manrope";
    font-size: 48px;
    font-weight: 400;
    color: #fff;
  }

  .achievements-header button {
    font-family: "Manrope";
    padding: 13px 19px;
    transition: all 0.5s ease;
    font-size: 18px;
    border: none;
    border-radius: 50px;
    text-align: center;
  }

  .achievements-header button:hover {
    background:linear-gradient(90deg, #ED666E, #FBA31C);
  }


  .achievement {
    border: 2px solid var(--pink);
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .achievement-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
  }

  .achievement-container-text {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .achievement-container-buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .achievement-container p {
    font-family: "Manrope";
  }

  .gradient-text {
    font-size: 20px;
  }
  .achievement:last-child {
    margin-bottom: 0;
  }


  .delete-button {
    /* height: 50px; */
    padding: 10px 16px;
    display: none;
    font-size: 18px;
    color: #fff;
    border: 3px solid #fff;
    border-radius: 50px;
    background-color: transparent;
    cursor: pointer;
    margin-right: 16px;
  }

  .delete-button:hover {
    color: #000;
    background-color: #fff;
  }

  .achievement:hover .delete-button {
    display: block;
  }




  /* Стиль для скроллбара */
  .achievements-list::-webkit-scrollbar {
    width: 26px; /* Ширина скроллбара */
  }

  .achievements-list::-webkit-scrollbar-thumb {
    background-color: var(--dark-grey); /* Цвет ползунка */
    border: 2px solid #2a2a2a;
    
    border-radius: 32px; /* Скругление */
  }

  .achievements-list::-webkit-scrollbar-thumb:hover {
    background-color: var(--light-grey); /* Цвет при наведении */
  }
</style> -->



<section class="archievements">
  <h2 class="archievement__title m3">Страница в разработке</h2>
  <p class="archievements__text m5">Скоро будет готово</p>
  <div class="archievements__buttons">
    <a href="/profile/ac2zcount" class="archievements__button gradient-button m6">Назад</a>
    <a href="/" class="archievements__button black-white-button m6">На главную</a>
  </div>
</section>


<style>
  .archievements {
    color: var(--white);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }

  .archievements__text {
    color: var(--grey);
  }

  .archievements__buttons {
    margin-top: 16px;
  }
  .archievements__button {
    padding: 6px 26px;
  }
</style>