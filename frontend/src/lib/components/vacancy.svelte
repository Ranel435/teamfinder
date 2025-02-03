<script lang="ts">
  // import {type Vacancy} from "$lib/stores/data"
  interface Vacancy {
    role: string;
    hakathon: string;
    skills: string[];
    experience: string;
    team: string;
    teammatesRoles: string[];
  }

  let vacancy: Vacancy = {
    role: "Backend developer",
    hakathon: "VTB API",
    skills: ['Python', 'SQL'],
    experience: "3",
    team: "Kovcheg at MISIS",
    teammatesRoles: ["Backend developer", "Frontend developer", "Designer"]
  }
</script>


<div class="card">
  <div class="card__container">
    <header class="card__header">
      <span class="gradient-text m8">Ищем</span>
      <span class="gradient-text m8">Осталось дней: 2</span>
    </header>
    <h2 class="card__title m4">{vacancy.role}</h2>
    <h3 class="card__subtitle">На <a class="card__hackaton gradient-text" href="/">{vacancy.hakathon}</a> &bull; Онлайн, финал - Москва</h3>
    <p class="card__details m7">{vacancy.skills[0]} &bull; {vacancy.skills[1]}</p>
    <hr class="card__divider">
    <p class="card__team">В команду <a class="team-name gradient-text" href="/">{vacancy.team}</a></p>
    <p class="card__stack">Состав: {#each vacancy.teammatesRoles as teammate, index} <span class="card__teammate m8">&bull; {teammate} </span>{/each}</p>
  
    <div class="card__buttons">
        <button class="card__button black-white-button m6">Подробнее</button>
        <button class="card__button gradient-button m6">Откликнуться</button>
    </div>
  </div>
</div>
<style>
  .card {
      width: 350px;
      background-color: transparent;
      border-radius: 20px;
      padding: 20px;
      color: var(--white);
      border: 2px solid var(--white);
  }

  .card__header {
      display: flex;
      justify-content: space-between;
      font-size: 18px;
      font-weight: bold;
  }

  .card__hackaton {
    text-decoration: underline;
  }

  .card__subtitle {
    font-family: "Manrope-semibold";
    font-size: 18px;
  }

  .card__details {
      color: #B3B3B3;
      margin-bottom: 16px;
  }

  .card__team {
    color: var(--dark-grey);
    font-family: "Manrope-semibold";
    font-size: 18px;
    margin-bottom: 10px;
  }

  .card__stack {
    color: var(--dark-grey);
    font-family: "Manrope-semibold";
    font-size: 18px;
    margin-bottom: 25px;
  }

  .card__teammate {
    color: var(--white);
    font-family: "Manrope-semibold";
    font-size: 18px;
  }
  .team-name {
    text-decoration: underline;
  }

  .card__buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .card__button {
    padding: 9px 26px;
  }
</style>