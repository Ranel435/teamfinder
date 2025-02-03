<script lang="ts">
  import type { Profile } from '$lib/stores/data';
  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();

  export let profile: Profile;

  let repeatedPassword: string = "";
  let termsAccepted: boolean = false;

  let showPassword: boolean = false;
  let showRepeatPassword: boolean = false;

  // Функция для проверки валидности формы
  function isFormValid() {
    if (repeatedPassword !== profile.password) {
      alert('Пароли не совпадают');
      repeatedPassword = "";
      termsAccepted = false;
      return 0;
    } 
    if (!termsAccepted) {
      alert('Согласитесь с правилами и политикой конфиденциальности для дальнейшей регистрации');
      return 0;
    }
    if (profile.email && profile.password && profile.password === repeatedPassword && termsAccepted) {
      return 1;
    };
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
        <div class="step__circle gradient-text">1</div>
        <p class="step__text gradient-text">Создание логина</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle simple">2</div>
        <p class="step__text simple">Персональная информация</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle simple">3</div>
        <p class="step__text simple">Подтверждение</p>
      </div>
    </div>
    <form class="card__form form" on:submit|preventDefault={()=> {
      if (isFormValid()) {
        dispatch('next');
      }
      }}>
      <div class="form__email email">
        <label for="email" class="email__label">Почта</label>
        <input type="email" class="email__input m7" placeholder="email@email.com" bind:value={profile.email} 
        required>
      </div>
  
      <div class="form__passwod password">
        <label for="password" class="password__label">Пароль</label>
        {#if showPassword}
        <input type="text" minlength="8" class="password__input m7" placeholder="Пароль" bind:value={profile.password} required>
        {:else}
        <input type="password" minlength="8" class="password__input m7" placeholder="Пароль" bind:value={profile.password} required>
        {/if}
        <input type="password" minlength="8" class="password__second-input m7" placeholder="Повторите пароль" bind:value={repeatedPassword} required>
      </div>
  
      <div class="form__terms terms">
        <label for="" class="terms__label m6">Согласен с <a href="" class="terms__link">правилами</a> сервиса 
          и <a href="" class="terms__link">политикой конфиденциальности</a> </label>
        <div class="terms__switch switch">
          <input type="checkbox" id="toggle" class="switch-input" bind:checked={termsAccepted}/>
          <label for="toggle" class="switch-label">
            <span class="switch-inner"></span>
            <span class="switch-switch"></span>
          </label>
        </div>
      </div>
  
      <div class="form__buttons">
        <a href="/auth" class="form__button grey-button m6">Назад</a>
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

  .simple {
    color: var(--dark-grey);
  }

  .card__form {
    width: 100%;
  }

  .email, .password {
    display: flex;
    flex-direction: column;
    text-align: left;
    width: 100%;
  }

  .email__label, .password__label {
    font-family: "Manrope-semibold";
    font-size: 18px;
  }

  .email__input, .password__input, .password__second-input {
    background-color: #F3F3F3;
    border: 0;
    border-radius: 50px;
    padding: 8px 16px;
    margin: 10px 0;
  }


  .terms {
    display: flex;
    flex-direction: row;
    align-items: center;
    text-align: left;
    margin: 16px;
  }

  .terms__link {
    color: #000;
  }

  .grey-button {
    padding: 8px 16px;
    margin-right: 16px;
  }

  .gradient-button {
    padding: 10px 18px;
  }

  .form__button {
    text-decoration: none;
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