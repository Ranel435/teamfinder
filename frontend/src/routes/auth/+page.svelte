<script lang="ts">
  let email: string = "";
  let password: string = "";


  function login(email: string, password: string) {

  }
</script>




<section class="auth">
  <div class="auth-container">
    <h2>Вход</h2>
    <div class="form">
      <div class="email">
        <label for="">Почта или логин</label>
        <input type="email" bind:value={email} placeholder="логин">
      </div>
      <div class="password">
        <label for="">Пароль</label>
        <input type="password" bind:value={password} placeholder="пароль">
      </div>
    </div>
    <div class="buttons">
      <a href="/registr">Регистрация</a>
      <a href="/">Забыл пароль</a>
      <a class="login-a" href="/" on:click={() => login}>Войти</a>
    </div>
  </div>
</section>



<style>

  .auth {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
  }

  .auth-container {
    padding: 72px 36px;
    width: 568px;
    height: 408px;
    background-color: #fff;
    border-radius: 32px;
    display: flex;
    flex-direction: column;
    text-align: center;
    align-items: center;
    justify-content: space-between;
    
  }

  h2 {
    font-family: 'Noto Sans Display Variable';
    margin: 0;
    font-size: 48px;
    font-weight: 900px;
    font-style: medium;
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
    font-family: 'Noto Sans Display Variable';
    font-size: 24px;
    margin-bottom: 12px;
  }

  input {
    font-family: 'Noto Sans Display Variable';
    background-color: #F3F3F3;
    border: 0;
    height: 24px;
    border-radius: 50px;
    padding: 8px 16px;
    font-size: 20px;
  }


  .buttons {
    display: flex;
    justify-content: space-between;
    width: 100%;
  }

  a {
    font-family: 'Noto Sans Display Variable';
    font-size: 24px;
    color: #A2A2A2;
    padding:8px 27px;
    border: 3px solid #C9C9C9;
    border-radius: 50px;
    text-decoration: none;
  }

  .login-a{
    display: flex;
    align-items: center;
    color: #fff;
    border: 0;
    background-image: linear-gradient(-120deg, #FBA31C, #E5469A);
  }
</style>