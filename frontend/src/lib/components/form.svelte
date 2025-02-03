<script lang="ts">
  // import {type Vacancy} from "$lib/stores/data"
  interface Form {
    role: string;
    hakathon: string;
    skills: string[];
    experience: string;
    how_many_haks_in_this_role: number;
    name: string;
    surname: string;
  }

  let form: Form = {
    role: "Backend developer",
    hakathon: "VTB API",
    skills: ['Python', 'SQL'],
    experience: "3",
    how_many_haks_in_this_role: 5,
    name: "Чин",
    surname: "Азес"
  }
</script>


<div class="card">
  <div class="card__container">
    <header class="card__header">
      <span class="gradient-text m8">В поиске команды</span>
    </header>
    <h2 class="card__title m4">{form.role}</h2>
    <h3 class="card__subtitle">На <a class="card__hackaton gradient-text" href="/">{form.hakathon}</a> &bull; Онлайн, финал - Москва</h3>
    <p class="card__details m7">{form.skills[0]} &bull; {form.skills[1]}</p>
    <hr class="card__divider">
    <!-- <div class="card__archievements">
      <h3 class="m7 gradient-text">Достижения</h3>
    </div> -->
    <span class="card__name m7">{form.name}</span>
    <span class="card__surname m7">{form.surname}</span>
    <p class="card__description m8">С 5 лет в IT, люблю пиццу</p>

    <div class="card__buttons">
        <button class="card__button black-white-button m6">Подробнее</button>
        <button class="card__button gradient-button m6">Пригласить</button>
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

  .card__description {
    margin-bottom: 16px;
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