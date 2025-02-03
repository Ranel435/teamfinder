<script lang="ts">
  import Code from "$lib/components/code.svelte";
  import type { Profile } from "$lib/stores/data";
  import { createEventDispatcher } from 'svelte';
  import { verifyEmailCode } from '$lib/api/auth';
  import { saveTokens } from "$lib/stores/authStore";
  import { goto } from "$app/navigation";

  const dispatch = createEventDispatcher();
  let confirmationCode: string[] = ["", "", "", "", "", ""]; // цифры кода подтверждения

  export let profile: Profile;

  async function handleVerifyCode() {
    console.log('Sending verification data:', {
      email: profile.email,
      code: confirmationCode.join(''),
      password: profile.password
    });

    const tokens = await verifyEmailCode(
      profile.email,
      confirmationCode.join(''),
      profile.name,
      profile.password
    );

    if (tokens) {
      saveTokens(tokens.access, tokens.refresh);
      dispatch('next');
      return true;
    } else {
      confirmationCode= ["", "", "", "", "", ""]; // цифры кода подтверждения
      alert('Неверный код подтверждения.');
      return false;
    }
  }
</script>

<section class="card">
  <div class="card__container">

    <header class="card__header">
      <h2 class="card__title m2">Подтверждение почты</h2>
    </header>

    <div class="card__steps steps">
      <div class="steps_step step">
        <div class="step__circle done">1</div>
        <p class="step__text gradient-text">Создание логина</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle done">2</div>
        <p class="step__text gradient-text">Персональная информация</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle gradient-text">3</div>
        <p class="step__text gradient-text">Подтверждение</p>
      </div>
    </div>
    
    <p class="card__text m6">на указанную почту {profile.email}
      был выслан 6-ти значный код</p>

    <div class="card__code">
      <label for="" class="code__label m6">Введите код для продолжения</label>
      <Code  bind:code={confirmationCode}/>
    </div>

    <div class="card__buttons">
      <button class="card__button card__button-back grey-button m6" on:click={() => dispatch('prev')}>Назад</button>
      <button class="card__button grey-button m6" on:click={() => handleVerifyCode()}>Завершить</button>
      <button class="card__button grey-button m6">Выслать повторно</button>
    </div>
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

  .card__code {
    display: flex;
    flex-direction: column;
  }

  .code__label {
    margin-bottom: 16px;
  }

  .grey-button {
    padding: 8px 16px;
    margin-right: 16px;
  }

  .gradient-button {
    padding: 10px 18px;
  }

  .card__button-back {
    margin-bottom: 16px;
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
      background-image: var(--gradient);
  }

  .switch-input:checked + .switch-label .switch-switch {
      transform: translateX(20px);
  }
</style>