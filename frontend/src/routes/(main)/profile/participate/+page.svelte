<script lang="ts">
  interface participation {
    hackaton: string;
    team: string;
    date: string;
    location: string;
    role: string;
    online: number;
    isCapitan: boolean;
  }


  let participations: participation[] = [
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov",
      role: "Frontend",
      online: 1,
      isCapitan: true,
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov",
      role: "Frontend",
      online: 1,
      isCapitan: false,
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov",
      role: "Frontend",
      online: 1,
      isCapitan: true,
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov",
      role: "Frontend",
      online: 1,
      isCapitan: false,
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov",
      role: "Frontend",
      online: 1,
      isCapitan: true,
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov",
      role: "Frontend",
      online: 1,
      isCapitan: false,
    }
  ];
</script>



<section class="participations">
  {#if participations.length === 0}
  <div class="no-participations">
    <h2 class="no-participations__title m1">Вы нигде не участвуете</h2>
    <p class="no-participations__text m5">Создайте анкету или присоединитесь к команде, заяка отобразиться здесь</p>
    <div class="no-participations__buttons">
      <button class="no-participations__button gradient-button m5">Хакатоны</button>
      <button class="no-participations__button black-white-button m5">Поиск команды</button>
      <button class="no-participations__button black-white-button m5">Создать анкету</button>
    </div>
  </div>
  {:else}
  <div class="participations__container">
    <header class="participations__header">
      <h2 class="participations__title m1">Я участвую</h2>
    </header>
    <ul class="participations__list">
      {#each participations as participation, index}
      <div class="participation">
        <div class="participation__container">
          <div class="participation__info">
            <div class="paticipation__main-info">
              <a class="participation__hakathon m4" href="/">{participation.hackaton} <img src="/link.svg" alt="">  </a>
              <span class="participation__team m4">   • В составе: {participation.team}</span>
            </div>
            <div class="participation__buttons">
              <button class="participation__button grey-button m6"></button>
              <button class="participation__button grey-button m6"></button>
            </div>
          </div>
          <div class="class__paticipation__other-info">
            <span class="participation__online"></span>
            <span class="participation__date m7">{participation.date}</span>
            <span class="participation__role m7"> • {participation.role}</span>
          </div>
        </div>
      </div>
      {/each}
    </ul>
  </div>
  {/if}
</section>

<style>
  .participations {
    height: 100%;
    margin-left: 28px;
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .no-participations {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 100%;
  }

  .no-participations__title {
    color: var(--white);
  }

  .no-participations__text {
    color: var(--dark-grey);
    margin: 22px 0 32px;
  }

  .no-participations__buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .no-participations__button {
    padding: 10px 24px;
    margin-right: 16px;
  }

  .no-participations-buttons button:last-child {
    margin-right: 0;
  }

/* //////////////////////////////////////////////////////////////////   */

  .participations__container {
    display: flex;
    flex-direction: column;
    padding: 32px 0 0 0;
  }
  .participations__header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 40px;
  }

  .participations__title {
    color: #fff;
  }

  .participations__list {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .participation {
    border: 2px solid var(--pink);
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .participation__container {
    display: flex;
    flex-direction: column;
    align-items: left;
    padding: 16px 24px;
  }

  .participation__info {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .participation__hakathon {
    color: var(--white);
    text-decoration: none;
  }

  .participation__team  {
    color: var(--white);
  }

  .participation__date, .participation__role {
    color: var(--white);
  }

  .participation__buttons {
    
  }
</style>