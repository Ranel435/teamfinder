<script lang="ts">
  import { goto } from "$app/navigation";
  import { saveTokens } from "$lib/stores/authStore";
  import { login } from '$lib/api/auth';

  let email: string = "";
  let password: string = "";
  let n_email: string = "qquerell.fw@gmail.com";
  let n_password: string = "12341234";

  async function handleLogin(event: Event) {
    event.preventDefault();

    if (!email || !password) {
      alert("Все поля обязательны");
      return;
    } 

    const success = await login(email, password);
    if (success) {
      goto("/profile");
      //лучше перекидывать на главную или еще лучше на хаки
    } else {
      alert('Неверный email или пароль');
    }
  }
</script>

<section class="auth">
  <div class="auth__container">
    <h2 class="auth__title m2">Вход</h2>
    <form class="auth__form" on:submit={handleLogin}>
      <div class="auth__email email">
        <label class="email__label m6" for="email">Почта</label>
        <input class="email__input m7" type="email" id="email" bind:value={email} placeholder="x@email.com" required>
      </div>
      <div class="auth__password password">
        <label class="password__label m6" for="password">Пароль</label>
        <input class="password__input m7" type="password" bind:value={password} placeholder="пароль" minlength="8" required>
      </div>
      <div class="auth__buttons">
        <div class="auth__first-buttons">
          <a class="auth__button grey-button m6" href="/reg">Регистрация</a>
          <a class="auth__button grey-button m6" href="/forget_password">Забыл пароль</a>
        </div>
        <div class="auth__second-buttons">
          <a class="auth__button grey-button m6" href="/">На главную</a>
          <button class="auth__button gradient-button m6" type="submit">Войти</button>
        </div>
      </div>
    </form>
  </div>
</section>


<style>
  .auth {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .auth__container {
    display: flex;
    flex-direction: column;
    text-align: center;
    align-items: center;
    justify-content: space-between;

    padding: 36px;
    width: 408px;
    height: 342px;
    background-color: #fff;
    border-radius: 32px;
  }

  .auth__form {
    display: flex;
    flex-direction: column;
    text-align: left;
    justify-content: space-between;
    width: calc(100% - 32px);
  }

  .auth__email {
    display: flex;
    flex-direction: column;
    margin-bottom: 16px;
  }
  
  .auth__password {
    display: flex;
    flex-direction: column;
  }

  .email__label, .password__label {
    margin-bottom: 8px;
  }

  .email__input, .password__input {
    background-color: #F3F3F3;
    border: 0;
    border-radius: 50px;
    padding: 8px 16px;
  }

  .auth__buttons {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    margin-top: 38px;
  }

  .auth__button {
    padding: 8px 16px;
    margin-right: 20px;
    text-decoration: none;
  }

  .auth__first-buttons, .auth__second-buttons {
    display: flex;
    justify-content: space-between;
  }

  .auth__second-buttons {
    margin-top: 12px;
  }
</style>