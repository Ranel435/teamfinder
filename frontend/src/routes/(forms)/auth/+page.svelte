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
  <div class="auth-container">
    <h2>Вход</h2>
    <form class="form" on:submit={handleLogin}>
      <div class="email">
        <label for="email">Почта</label>
        <input type="email" id="email" bind:value={email} placeholder="x@email.com" required>
      </div>
      <div class="password">
        <label for="password">Пароль</label>
        <input type="password" bind:value={password} placeholder="пароль" minlength="8" required>
      </div>
      <div class="all-buttons">
        <div class="buttons">
          <button class="button" type="button" on:click={() => goto("/reg")}>Регистрация</button>
          <button class="button" type="button" on:click={() => goto("/forget_password")}>Забыл пароль</button>
        </div>
        <div class="buttons second-buttons">
          <button class="button" type="button" on:click={() => goto('/')}>Назад</button>
          <button class="login-a" type="submit">Войти</button>
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

  .auth-container {
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

  h2 {
    font-family: 'Manrope';
    margin: 0;
    font-size: 36px;
    font-style: normal;
    font-weight: 100;
  }
  .form {
    display: flex;
    flex-direction: column;
    text-align: left;
    justify-content: space-between;
    width: calc(100% - 32px);
  }
  .email {
    display: flex;
    flex-direction: column;
    margin-bottom: 16px;
  }
  .password {
    display: flex;
    flex-direction: column;
  }

  label {
    font-family: 'Manrope';
    font-size: 18px;
    margin-bottom: 8px;
  }

  input {
    font-family: 'Manrope';
    font-size: 16px;
    background-color: #F3F3F3;
    border: 0;
    border-radius: 50px;
    padding: 8px 16px;
  }

  .all-buttons {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    margin-top: 38px;
  }

  .buttons {
    display: flex;
    justify-content: space-between;
  }

  .second-buttons {
    margin-top: 20px;
  }

  button {
    background-color: #fff;
    font-family: 'Manrope';
    font-size: 18px;
    text-decoration: none;
    color: var(--dark-grey);
    padding:8px 16px;
    border: 3px solid var(--dark-grey);
    border-radius: 50px;
    margin-right: 20px;
    cursor: pointer;
    transition: all 0.5s ease;
  }

  .login-a{
    padding: 11px 22px;
    display: flex;
    align-items: center;
    color: #fff;
    border: 0;
    background-image:  linear-gradient(120deg, #E5469A, #FBA31C);
    transition: all 0.5s ease;
  }


  .button:hover {
    background-color: var(--light-grey);
    color: var(--dark-grey);
  }

  .login-a:hover {
    background-image:  linear-gradient(180deg, #E5469A, #FBA31C);

  }

</style>