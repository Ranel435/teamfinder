<script lang="ts">
  import { goto } from "$app/navigation";
  import "../../../app.css"
  let email: string = "";
  let password: string = "";
</script>

<section class="auth">
  <div class="auth-container">
    <h2>Вход</h2>
    <form class="form">
      <div class="email">
        <label for="email">Почта</label>
        <input type="email" id="email" bind:value={email} placeholder="x@email.com" required>
      </div>
      <div class="password">
        <label for="password">Пароль</label>
        <input type="password" bind:value={password} placeholder="пароль" required>
      </div>
    </form>
    <div class="all-buttons">
      <div class="buttons">
        <a class="button" href="/reg">Регистрация</a>
        <a href="/forms/forget_password">Забыл пароль</a>
      </div>
      <div class="buttons second-buttons">
        <button class="button" on:click={() => goto('/')}>Назад</button>
        <a class="login-a" href="/">Войти</a>
      </div>
    </div>
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
    padding: 36px;
    width: 480px;
    height: 414px;
    background-color: #fff;
    border-radius: 32px;
    display: flex;
    flex-direction: column;
    text-align: center;
    align-items: center;
    justify-content: space-between;
    
  }

  h2 {
    font-family: 'Manrope';
    margin: 0;
    font-size: 48px;
    font-style: normal;
    font-weight: 100;
  }
  .form {
    display: flex;
    flex-direction: column;
    text-align: left;
    justify-content: space-between;
    width: calc(100% - 54px);
  }
  .email {
    display: flex;
    flex-direction: column;
    margin-bottom: 20px;
  }
  .password {
    display: flex;
    flex-direction: column;
  }

  label {
    font-family: 'Manrope';
    font-size: 24px;
    margin-bottom: 12px;
  }

  input {
    font-family: 'Manrope';
    font-size: 20px;
    background-color: #F3F3F3;
    border: 0;
    border-radius: 50px;
    height: 24px;
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
    /* margin-bottom: 30px; */
  }

  .second-buttons {
    margin-top: 20px;
  }

  .button {
    margin-right: 20px;
  }

  a, button {
    font-family: 'Manrope';
    font-size: 24px;
    text-decoration: none;
    color: #A2A2A2;
    padding:8px 27px;
    border: 3px solid #C9C9C9;
    border-radius: 50px;
  }

  .login-a{
    display: flex;
    align-items: center;
    color: #fff;
    border: 0;
    background-image: linear-gradient(-120deg, #FBA31C, #E5469A);
  }
</style>