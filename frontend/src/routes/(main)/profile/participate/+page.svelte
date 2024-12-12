<script lang="ts">
  interface participation {
    hackaton: string;
    team: string;
    date: string;
    location: string;
  }


  let participations: participation[] = [
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    },
    {
      hackaton: "VTB API",
      team: "Kovcheg",
      date: "12.12.24",
      location: "Moskov"
    }
  ];
</script>



<section class="participations">
  {#if participations.length === 0}
  <div class="no-participations">
    <p class="no-participations-title">Вы нигде не участвуете</p>
    <p class="no-participations-text">Заявите об участии в хакатоне,
      заявка отобразится здесь</p>
    <div class="no-participations-buttons">
      <button class="gradient-button">Добавить</button>
    </div>
  </div>
  {:else}
  <div class="participations-container">
    <div class="participations-header">
      <h2>Я участвую</h2>
      <button class="gradient-button" >Добавить</button>
    </div>
    <ul class="participations-list">
      {#each participations as participation, index}
      <div class="participation">
        <div class="participation-container">
          <div class="participation-container-text">
            <a class="gradient-text" href="">{participation.hackaton}</a>
            <p>В составе: {participation.team}</p>
            <p>Дата: {participation.date}</p>
          </div>
          <div class="participation-container-buttons">
            <button class="delete-button">Команда</button>
            <button class="delete-button" on:click={() => {participations.splice(index, 1); participations=[...participations];}}>Отказаться от участия</button>
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

  .participations {
    height: 100%;
    margin-left: 50px;
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

  .no-participations-title {
    font-family: "Manrope";
    font-size: 48px;
    color: #fff;
    margin-bottom: 32px;
  }

  .no-participations-text {
    font-family: "Manrope";
    font-size: 24px;
    color: var(--dark-grey);
    margin-bottom: 32px;
  }

  .no-participations-buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .no-participations-buttons button {
    font-family: "Manrope";
    font-size: 24px;
    padding: 10px 24px;
    border-radius: 50px;
    cursor: pointer;
    margin-right: 16px;
    transition: all 0.5 ease;
  }

  .no-participations-buttons button:last-child {
    margin-right: 0;
  }


  .participations-container {
    display: flex;
    flex-direction: column;
    padding: 32px;
  }
  .participations-header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 40px;
  }

  .participations-list {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-height: 50vh; /* Фиксированная высота для прокрутки */
    overflow-y: auto; /* Добавление вертикального скроллбара */
    padding-right: 10px; /* Добавляем отступ для скроллбара */
    box-sizing: content-box; /* Учитываем отступ, не влияя на ширину */
  }

  .participations-header h2{
    font-family: "Manrope";
    font-size: 48px;
    font-weight: 400;
    color: #fff;
  }

  .participations-header button {
    font-family: "Manrope";
    padding: 13px 19px;
    transition: all 0.5s ease;
    font-size: 18px;
    border: none;
    border-radius: 50px;
    text-align: center;
  }

  .participations-header button:hover {
    background:linear-gradient(90deg, #ED666E, #FBA31C);
  }


  .participation {
    border: 2px solid var(--pink);
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .participation-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
  }

  .participation-container-text {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .participation-container-buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .participation-container p {
    font-family: "Manrope";
    color: #fff;
  }
  .participation-container a {
    font-size: 20px;
    font-family: "Manrope";
    text-decoration: none;
  }

  .participation-container a:hover {
    text-decoration: underline;
  }

  .gradient-text {
    font-size: 20px;
  }
  .participation:last-child {
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

  .participation:hover .delete-button {
    display: block;
  }




  /* Стиль для скроллбара */
  .participations-list::-webkit-scrollbar {
    width: 26px; /* Ширина скроллбара */
  }

  .participations-list::-webkit-scrollbar-thumb {
    background-color: var(--dark-grey); /* Цвет ползунка */
    border: 2px solid #2a2a2a;
    
    border-radius: 32px; /* Скругление */
  }

  .participations-list::-webkit-scrollbar-thumb:hover {
    background-color: var(--light-grey); /* Цвет при наведении */
  }
</style>