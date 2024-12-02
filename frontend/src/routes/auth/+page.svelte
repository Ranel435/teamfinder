<script lang="ts">
  let email: string = "";
  let password: string = "";
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
      <a href="/reg">Регистрация</a>
      <a href="/forget_password">Забыл пароль</a>
      <a class="login-a" href="/">Войти</a>
    </div>
  </div>
</section>

<style>

  .auth {
    display: flex;
    justify-content: center;
    height: 100vh;
    margin: 100px;
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


  .buttons {
    display: flex;
    justify-content: space-between;
    width: 100%;
  }

  a {
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