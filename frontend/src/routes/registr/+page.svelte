<script lang="ts">
  let current_form: string = "registration";
</script>

<section class="registr">
  <div class="registr-container">
    <div class="text">
      <h2>Регистрация</h2>
      <p>Заполните поля необходимой информацией</p>
    </div>
    {#if current_form === "registration"}
      <div class="steps">
        <div class="step active">
          <div class="circle">1</div>
          <p>Создание логина</p>
        </div>
        <div class="step">
          <div class="circle">2</div>
          <p>Персональная информация</p>
        </div>
        <div class="step">
          <div class="circle">3</div>
          <p>Подтверждение</p>
        </div>
      </div>
      <div class="form">
        <div class="login">
          <label for="">Логин</label>
          <input class="input" type="text" required />
        </div>
        <div class="email">
          <label for="">Почта</label>
          <input class="input" type="email" required />
        </div>
        <div class="password">
          <label for="">Пароль</label>
          <input type="password" required />
        </div>
        <div class="password">
          <label for="">Повторите пароль</label>
          <input type="password" required />
        </div>
      </div>
      <div class="buttons">
        <a href="/auth">Назад</a>
        <a
          class="login-a"
          href=""
          on:click={() => (current_form = "personal_info")}>Далее</a
        >
      </div>
    {/if}
    {#if current_form === "personal_info"}
      <div class="steps">
        <div class="step active">
          <div class="circle">1</div>
          <p>Создание логина</p>
        </div>
        <div class="step active">
          <div class="circle">2</div>
          <p>Персональная информация</p>
        </div>
        <div class="step">
          <div class="circle">3</div>
          <p>Подтверждение</p>
        </div>
      </div>
      <div class="form">
        <div class="name">
          <label for="">Имя</label>
          <input class="input" type="text" required placeholder="Иван"/>
        </div>
        <div class="surname">
          <label for="">Фамилия</label>
          <input class="input" type="text" required placeholder="Иванов"/>
        </div>
        <div class="tgid">
          <label for="">Телеграм</label>
          <input type="text" required placeholder="@telegram" />
        </div>
      </div>
      <div class="buttons">
        <a href="" on:click={() => current_form = "registration"}>Назад</a>
        <a
          class="login-a"
          href=""
          on:click={() => (current_form = "confirm")}>Далее</a
        >
      </div>
    {/if}
    {#if current_form === "confirm"}
      <div class="steps">
        <div class="step active">
          <div class="circle">1</div>
          <p>Создание логина</p>
        </div>
        <div class="step active">
          <div class="circle">2</div>
          <p>Персональная информация</p>
        </div>
        <div class="step active">
          <div class="circle">3</div>
          <p>Подтверждение</p>
        </div>
      </div>
      <div class="form">
        <div class="email">
          <label for="">Логин</label>
          <input class="email-input" type="email" required />
        </div>
        <div class="email">
          <label for="">Почта</label>
          <input class="email-input" type="email" required />
        </div>
        <div class="password">
          <label for="">Пароль</label>
          <input type="password" required />
        </div>
        <div class="password">
          <label for="">Повторите пароль</label>
          <input type="password" required />
        </div>
      </div>
      <div class="buttons">
        <a href="" on:click={() => current_form = "personal_info"}>Назад</a>
        <a
          class="login-a"
          href=""
          on:click={() => (current_form = "confirm")}>Далее</a
        >
      </div>
    {/if}
  </div>
</section>

<style>
  .registr {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
  }

  .registr-container {
    padding: 50px 36px;
    width: 568px;
    height: auto;
    background-color: #fff;
    border-radius: 32px;
    display: flex;
    flex-direction: column;
    text-align: center;
    align-items: center;
    justify-content: space-between;
  }

  .text {
    margin-bottom: 38px;
  }

  h2 {
    font-family: "Noto Sans Display Variable";
    margin: 0;
    font-size: 48px;
    font-weight: 900px;
    font-style: medium;
  }

  p {
    font-family: "Noto Sans Display Variable";
    margin: 0;
    color: #a2a2a2;
    font-size: 24px;
  }

  .form {
    display: flex;
    flex-direction: column;
    text-align: left;
    justify-content: space-between;
    width: calc(100% - 54px);
    margin-bottom: 38px;
  }
  .email, .login, .name, .surname {
    display: flex;
    flex-direction: column;
    margin-bottom: 20px;
  }
  .input {
    margin-bottom: 36px;
  }
  .password, .tgid {
    display: flex;
    flex-direction: column;
  }

  label {
    font-family: "Noto Sans Display Variable";
    font-size: 24px;
    margin-bottom: 7px;
  }

  input {
    background-color: #f3f3f3;
    border: 0;
    height: 24px;
    border-radius: 50px;
    padding: 8px 16px;
    font-size: 24px;
  }

  .buttons {
    display: flex;
    justify-content: center;
    width: 100%;
  }

  a {
    font-family: "Noto Sans Display Variable";
    font-size: 24px;
    color: #a2a2a2;
    padding: 8px 27px;
    border: 3px solid #c9c9c9;
    border-radius: 50px;
    text-decoration: none;
  }

  .login-a {
    margin-left: 25px;
    display: flex;
    align-items: center;
    color: #fff;
    border: 0;
    background-image: linear-gradient(-120deg, #fba31c, #e5469a);
  }

  .steps {
    display: flex;
    margin-right: 20px;
    margin-bottom: 38px;
  }

  .step {
    width: 30%;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
  }

  .circle {
    width: 44px;
    height: 44px;
    border-radius: 50%;
    background: #fff;
    border: 2px solid #ccc;
    color: #ccc;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 18px;
    font-weight: bold;
    font-family: "Noto Sans Display Variable";
  }

  .step.active .circle {
    border-color: #ff4081;
    color: #ff4081;
    background: #fff;
  }

  .step p {
    font-family: "Noto Sans Display Variable";
    margin-top: 10px;
    font-size: 19px;
    color: #999;
  }

  .step.active p {
    color: #ff4081;
  }

  /* Линия соединения */
  .steps::before {
    content: "";
    position: absolute;
    top: 20px;
    left: 10%;
    width: 80%;
    height: 2px;
    background: #ccc;
    z-index: -1;
  }

  .step.active ~ .step::before {
    background: #ff4081;
  }
</style>
