<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { sendEmailCode } from '$lib/api/auth';
  import type { Profile } from '$lib/stores/data';

  const dispatch = createEventDispatcher();

  export let profile: Profile;


  function isFormValid() {
    if (profile.name && profile.surname && profile.tgId) {
      return 1;
    } else {
      return 0;
    }
  }


</script>


<section class="card">
  <div class="card__container">

    <header class="card__header">
      <h2 class="card__title m2">Регистрация</h2>
      <p class="card__text m6">Заполните поля необходимой информацией</p>
    </header>

    <div class="card__steps steps">
      <div class="steps_step step">
        <div class="step__circle done">1</div>
        <p class="step__text gradient-text">Создание логина</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle gradient-text">2</div>
        <p class="step__text gradient-text">Персональная информация</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle simple">3</div>
        <p class="step__text simple">Подтверждение</p>
      </div>
    </div>
    <form class="card__form form" on:submit|preventDefault={()=> {
      if (isFormValid()) {
        dispatch('next');
        sendEmailCode(profile.email);
      }
      }}>
      <div class="form__name name">
        <label for="" class="name__label">Имя</label>
        <input type="text" class="name__input m7" pattern="^[А-Яа-яЁё]+$" placeholder="Иван" bind:value={profile.name} required>
      </div>
  
      <div class="form__surname surname">
        <label for="" class="surname__label">Фамилия</label>
        <input type="text" class="surname__input m7" pattern="^[А-Яа-яЁё]+$" placeholder="Иванов" bind:value={profile.surname} required>
      </div>

      <div class="form__tgId tgId">
        <label for="" class="tgId__label">Telegram</label>
        <input type="text" class="tgId__input m7" pattern="^@([A-Za-z0-9_]+)$" placeholder="@qwerty" bind:value={profile.tgId} required>
      </div>
  
      <!-- <div class="form__find-team find-team">
        <label for="" class="find-team__label m6">Ищу команду</label>
        <div class="find-team__switch switch">
          <input type="checkbox" id="toggle" class="switch-input"/>
          <label for="toggle" class="switch-label">
            <span class="switch-inner"></span>
            <span class="switch-switch"></span>
          </label>
        </div>
      </div> -->
  
      <div class="form__buttons">
        <button type="button" class="form__button grey-button m6" on:click={()=> dispatch('prev')}>Назад</button>
        <button type="submit" class="form__button gradient-button m6">Далее</button>
      </div>
    </form>
  </div>
</section>






<style>
  .card {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .card__container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    text-align: center;
    flex-direction: column;

    background-color: var(--white);
    border-radius: 28px;
    padding: 36px;
    width: calc(480px - 72px);
    height: calc(607px - 72px);
  }

  .card__title {
    color: var(--black);
  }
  
  .card__text {
    color: var(--dark-grey);
  }

  .steps {
    display: flex;
    margin: 16px 0;
  }

  .step {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;

    width: 30%;
    margin: 0 auto;
  }

  .step__text {
    font-family: "Manrope-regular";
    font-size: 12px;
    line-height: 12px;
    margin-top: 7px;
  }

  .step__circle {
    font-family: "Montserrat";
    font-size: 18px;

    display: flex;
    align-items: center;
    justify-content: center;

    background-color: #fff;
    width: 29px;
    height: 29px;
    border: 2px solid var(--dark-grey);
    border-radius: 50%;
  }
  .done {
    background: var(--gradient);
    width: 31px;
    height: 31px;
    border: none;
    color: var(--white);
  }

  .simple {
    color: var(--dark-grey);
  }

  .card__form {
    width: 100%;
  }

  .name, .surname, .tgId {
    display: flex;
    flex-direction: column;
    text-align: left;
    width: 100%;
  }

  .name__label, .surname__label, .tgId__label {
    font-family: "Manrope-semibold";
    font-size: 18px;
  }

  .name__input, .surname__input, .tgId__input {
    background-color: #F3F3F3;
    border: 0;
    border-radius: 50px;
    padding: 8px 16px;
    margin: 10px 0;
  }


  .find-team {
    display: flex;
    flex-direction: row;
    align-items: center;
    text-align: left;
    margin: 16px;
  }

  .find-team__label {
    margin-right: 16px;
  }

  .grey-button {
    padding: 8px 16px;
    margin-right: 16px;
  }

  .gradient-button {
    padding: 10px 18px;
  }

  .switch {
    position: relative;
  }

  .switch-input {
      display: none;
  }

  .switch-label {
      display: flex;
      align-items: center;
      cursor: pointer;
      width: 42px;
      height: 24px;
      background-color: #ccc;
      border-radius: 15px;
      position: relative;
      transition: background-color 0.3s ease;
  }

  .switch-inner {
      display: flex;
      justify-content: space-between;
      padding: 0 5px;
  }

  .switch-switch {
      position: absolute;
      width: 18px;
      height: 18px;
      background-color: white;
      border-radius: 50%;
      transition: transform 0.3s ease;
      left: 2px;
  }

  .switch-input:checked + .switch-label {
      background-image: var(--gradient); /* Цвет при включенном состоянии */
  }

  .switch-input:checked + .switch-label .switch-switch {
      transform: translateX(20px); /* Перемещение переключателя */
  }
</style>