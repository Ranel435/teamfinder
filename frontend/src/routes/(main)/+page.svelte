<script lang="ts">
  import "../../app.css";
  import { type Profile } from "$lib/stores/data";
  import Vacancy from "$lib/components/vacancy.svelte";
  import Form from "$lib/components/form.svelte";

  let roles: string[] = [
    "Backend",
    "Frontend",
    "Fullstack",
    "Data-Analyst",
    "Product-Manager",
    "Design",
    "IOS-Developer",
    "ML",
    "DevOps",
  ];
  let array: number[] = [1, 2, 3, 4, 5, 6];
  let count_teams: number = 14;

  import { onMount } from "svelte";
  let vacanciesScrollContainer: HTMLElement;
  let formsScrollContainer: HTMLElement;


  function scrollLeft(scrollContainer: HTMLElement) {
    scrollContainer?.scrollBy({ left: -300, behavior: "smooth" });
  }

  function scrollRight(scrollContainer: HTMLElement) {
    scrollContainer?.scrollBy({ left: 300, behavior: "smooth" });
  }

  onMount(() => {
    // Центрируе прокрутку
    if (vacanciesScrollContainer) {
      const containerWidth = vacanciesScrollContainer.scrollWidth;
      const viewportWidth = vacanciesScrollContainer.clientWidth;
      vacanciesScrollContainer.scrollLeft = (containerWidth - viewportWidth) / 3;
    }

    if (formsScrollContainer) {
      const containerWidth = formsScrollContainer.scrollWidth;
      const viewportWidth = formsScrollContainer.clientWidth;
      formsScrollContainer.scrollLeft = (containerWidth - viewportWidth) / 3;
    }
  });
</script>

<main>
  <section class="main">
    <div class="main__content">
      <p class="main__text h4">Команда найдется для каждого</p>
      <img
        class="main__image"
        src="Начни Хакатонить!.svg"
        alt="Начни хакатонить!"
      />
    </div>
    <div class="main__buttons">
      <a class="main__button buttonM" href="/teams">Найти команду</a>
      <a class="main__button buttonM" href="/forms">Смотреть анкеты</a>
      <a class="main__button buttonM" href="/">Что такое хакатоны?</a>
    </div>
  </section>

  <section class="popular-roles">
    <h2 class="popular-roles__title">Популярные роли</h2>
    <div class="popular-roles__roles">
      {#each roles as role, index}
        <button class="popular-roles__role role {role}">
          <p class="role__name m4">{role}</p>
          <div class="role__container">
            <p class="role__teams m6">{count_teams} вакансий</p>
            <p class="role__forms m6">{count_teams} анкет</p>
          </div>
        </button>
      {/each}
    </div>
    <div class="more">
      <a class="more__link h3" href="/">Еще роли</a>
    </div>
  </section>

  <section class="vacancies">
    <h2 class="vacancies__title">Свободные слоты</h2>
    <div bind:this={vacanciesScrollContainer} class="vacancies__container">
      {#each array as item}
        <div class="vacancy">
          <Vacancy />
        </div>
      {/each}
    </div>
    <div class="vacancies__more">
        <div class="vacancies__buttons">
            <button class="vacancies__button" on:click={() => scrollLeft(vacanciesScrollContainer)}><img src="/left.svg" alt=""></button>
            <button class="vacancies__button" on:click={() => scrollRight(vacanciesScrollContainer)}><img src="/right.svg" alt=""></button>
          </div>
          <a href="/" class="more__link h3">Смотреть все</a>
    </div>
  </section>

  <section class="forms">
    <h2 class="forms__title">Активные анкеты</h2>
    <div bind:this={formsScrollContainer} class="forms__container">
      {#each array as item}
        <div class="form">
          <Form />
        </div>
      {/each}
    </div>
    <div class="forms__more">
        <div class="forms__buttons">
          <button class="forms__button" on:click={() => scrollLeft(formsScrollContainer)}><img src="/left.svg" alt=""></button>
          <button class="forms__button" on:click={() => scrollRight(formsScrollContainer)}><img src="/right.svg" alt=""></button>
          </div>
          <a href="/" class="more__link h3">Смотреть все</a>
    </div>
  </section>
</main>

<style>
  p,
  h2 {
    color: #fff;
  }

  main {
    margin: 120px 0px;
  }

  /* /////////////////////////////////////////////////////////////////////     */
  .main {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }

  .main__content {
    text-align: center;
    margin: 10px 0 40px;
  }

  .main__text {
    font-family: "Igra";
    color: #fff;
  }

  .main__image {
    width: 815px;
    height: 258px;
  }

  .main__buttons {
    margin-bottom: 128px;
  }

  .main__button {
    color: #fff;
    text-decoration: none;

    border: 3px solid;
    border-color: #fff;
    border-radius: 50px;

    margin-right: 20px;
    padding: 7px 18px;
  }

  .main__button:hover {
    transition: all 0.7s ease;
    color: #000;
    background-color: #fff;
  }

  /* ////////////////////////////////////////////////////////////////////////// */

  .popular-roles {
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    margin: 0 175px 120px;
  }

  .popular-roles__title,
  .teams__title {
    margin-bottom: 80px;
  }

  .popular-roles__roles {
    display: grid;
    grid-template-columns: repeat(3, minmax(auto, 100vw));
  }

  .popular-roles__role {
    position: relative;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    text-align: left;

    border: 4px solid #fff;
    border-radius: 24px;
    padding: 10px;
    margin: 0px 36px 36px 0;

    background: rgba(0, 0, 0, 0);
  }

  .popular-roles__role:nth-child(3n) {
    margin-right: 0;
  }

  .popular-roles__role::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 5%;
    height: 100%;
    background-color: rgba(0, 31, 184, 0.5);
    z-index: -1;
    transition: width 0.3s ease;
  }

  .popular-roles__role:hover::before {
    width: 100%;
  }

  .role__container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 95%;
  }

  .role__name {
    margin-bottom: 10px;
    padding-left: 19px;
  }

  .role__teams,
  .role__forms {
    padding-left: 19px;
    padding-bottom: 10px;
  }

  .more {
    display: flex;
    justify-content: right;
    margin-top: 20px;
    transition: color 0.2s;
    text-decoration: none;
    width: 100%;
  }

  .more__link {
    color: #fff;
    text-decoration: none;
  }

  .more__link:hover {
    text-decoration: underline;
  }

  .Backend::before {
    background-color: rgba(0, 31, 184, 0.3);
  }
  .Frontend::before {
    background-color: rgba(255, 77, 0, 0.3);
  }
  .Fullstack::before {
    background-color: rgba(255, 39, 39, 0.3);
  }
  .Data-Analyst::before {
    background-color: rgba(0, 197, 3, 0.3);
  }
  .ML::before {
    background-color: rgba(0, 90, 12, 0.3);
  }
  .DevOps::before {
    background-color: rgba(153, 0, 255, 0.3);
  }
  .Design::before {
    background-color: rgba(14, 179, 255, 0.3);
  }
  .IOS-Developer::before {
    background-color: rgba(160, 0, 83, 0.3);
  }
  .Product-Manager::before {
    background-color: rgba(255, 225, 0, 0.3);
  }

  /* //////////////////////////////////////////////////////////////////////// */

  .vacancies {
    display: flex;
    flex-direction: column;
    text-align: center;
  }
  
  .vacancies__container {
    display: flex;
    justify-content: left;
    overflow-x: auto;
    scroll-behavior: smooth;
    gap: 16px;
    padding: 10px;
    scrollbar-width: none;
    margin: 36px 0;
    /* width: 100vw; */
  }

  .vacancies__container::-webkit-scrollbar {
    display: none;
  }

  .vacancy {
    min-width: 350px;
    flex-shrink: 0;
    text-align: left;
  }

  .vacancies__buttons {
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-right: 10px;
  }

  .vacancies__button {
    background: none;
    border: none;
  }

  .vacancies__more {
    display: flex;
    justify-content: right;
    align-items: center;
    margin-top: 20px;
    transition: color 0.2s;
    text-decoration: none;
    margin-right: 175px;
  }

  .more__title {
    color: var(--white);
  }


  /* /////////////////////////////////////////////////////////////////////// */

  .forms {
    display: flex;
    flex-direction: column;
    text-align: center;
    margin-top: 84px;
  }
  
  .forms__container {
    display: flex;
    justify-content: left;
    overflow-x: auto;
    scroll-behavior: smooth;
    gap: 16px;
    padding: 10px;
    scrollbar-width: none;
    margin: 36px 0;
    /* width: 100vw; */
  }

  .forms__container::-webkit-scrollbar {
    display: none;
  }

  .form {
    min-width: 350px;
    flex-shrink: 0;
    text-align: left;
  }

  .forms__buttons {
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-right: 10px;
  }

  .forms__button {
    background: none;
    border: none;
  }


  .forms__more {
    display: flex;
    justify-content: right;
    align-items: center;
    margin-top: 20px;
    transition: color 0.2s;
    text-decoration: none;
    margin-right: 175px;
  }

  .more__title {
    color: var(--white);
  }

  /* //////////////////////////////////////////////////////////////////////////////// */


</style>
