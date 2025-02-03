<script lang="ts">
  interface form {
    hackaton: string;
    role: string;
    date: string;
    response_count: number;
    teams_count: number;
  }

  let forms: form[] = [
    {
      hackaton: "Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4,
    },
    {
      hackaton: "Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4,
    },
    {
      hackaton: "Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4,
    },
    {
      hackaton: "Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4,
    },
    {
      hackaton: "Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4,
    },
    {
      hackaton: "Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4,
    },
  ];
</script>

<section class="forms">
  {#if forms.length === 0}
    <div class="forms__no-forms no-forms">
      <p class="no-forms__title m1">Пока нет анкет</p>
      <p class="no-forms__text m5">Создайте анкету, чтобы искать команду</p>
      <div class="no-forms__buttons">
        <button class="no-forms__button gradient-button m5">Создать</button>
        <a class="no-forms__button black-white-button m5" href="/">Хакатоны</a>
        <a class="no-forms__button black-white-button m5" href="/"
          >Рынок вакансий</a
        >
      </div>
    </div>
  {:else}
    <div class="forms__container">
      <header class="forms__header">
        <h2 class="forms__title m1">Активные анкеты</h2>
        <div class="forms__buttons">
          <button class="forms__button gradient-button button__create m5"
            >Создать</button
          >
          <button class="forms__button black-white-button button__hakathons m5"
            >Хакатоны</button
          >
        </div>
      </header>
      <ul class="forms__list">
        {#each forms as form, index}
          <div class="forms__form form">
            <div class="form__container">
              <div class="form__info">
                <h3 class="form__hackaton m5">{form.hackaton}</h3>
                <span class="form__role m7">{form.role} • </span><span
                  class="form__date m7">{form.date}</span
                >
                <div class="form__links">
                  <a class="form__link link__responses" href="/"
                    >Отклики: {form.response_count}
                    <img src="/link.svg" alt="link" /></a
                  >
                  <a class="form__link link__teams" href="/"
                    >Подходящие команды: {form.teams_count}
                    <img src="/link.svg" alt="link" /></a
                  >
                </div>
              </div>
            </div>
          </div>
        {/each}
      </ul>
    </div>
  {/if}
</section>

<style>
  .forms {
    height: 100%;
    margin-left: 28px;
    display: flex;
    flex-direction: column;
  }

  .forms__no-forms {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }

  .no-forms__title {
    color: var(--white);
  }

  .no-forms__text {
    color: var(--white);
    margin: 22px 0 32px;
  }
  .no-forms__buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .no-forms__button {
    padding: 6px 28px;
    margin-right: 16px;
    text-decoration: none;
  }

  .no-forms__button:last-child {
    margin-right: 0;
  }

  /* ////////////////////////////////////////////////////////////////////////////////// */

  .forms__container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    padding: 32px 0 0 0;
  }

  .forms__header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
  }

  .forms__buttons {
    display: flex;
    flex-direction: row;
  }

  .forms__button {
    padding: 5px 70px;
  }

  .button__create {
    margin-right: 32px;
  }

  .forms__title {
    width: 100%;
    display: flex;
    justify-content: space-between;
    color: var(--white);
  }

  .forms__list {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin: 32px;
    /* max-height: 50vh;
    overflow-y: auto;
    padding-right: 10px;
    box-sizing: content-box; */
  }

  .forms__form {
    border: 2px solid #fff;
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .forms__form:last-child {
    margin-bottom: 0;
  }

  .form__container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
  }

  .form__hackaton {
    color: var(--white);
  }

  .form__role {
    color: var(--grey);
  }

  .form__date {
    color: var(--yellow);
  }

  .form__links {
    display: flex;
    flex-direction: row;
  }

  .form__link {
    font-family: "Manrope-semibold";
    font-size: 16px;
    margin-right: 10px;
    color: var(--white);
    text-decoration: none;
  }

  .form__link:hover {
    text-decoration: underline;
  }

  .form:hover {
    border-image: var(--pink);
  }

  .form:hover .form__delete-button {
    display: block;
  }

  /* .forms__list::-webkit-scrollbar {
    width: 28px;
  }
  .forms__list::-webkit-scrollbar-thumb {
    background-color: var(--dark-grey);
    border: 5px solid transparent;
    border-radius: 32px;
    background-clip: padding-box;
  }
  .forms__list::-webkit-scrollbar-thumb:hover {
    background-color: var(--light-grey);
  }
  .forms__list::-webkit-scrollbar-track {
    background-color: transparent;
    border-radius: 32px;
    border: 2px solid #fff;
  } */
</style>
