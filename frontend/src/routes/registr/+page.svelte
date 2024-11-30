<script lang="ts">
  import "../../app.css";
  import Code from "$lib/code.svelte";
  $: current_form = "registration";
  $: main_text = "Заполните пол необходимой информацией";
  $: main_title = "Регистрация";
  if (current_form === "confirm") {
    main_title = "Подтверждение почты";
    main_text = "";
  }
  if (current_form === "end") {
    main_title = "Регистрация завершена";
    main_text = "";
  }
</script>

<section class="registr">
  <div class="registr-container">
    <div class="text">
      <h2>{main_title}</h2>
      <p>{main_text}</p>
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
          <input class="input" type="text" required placeholder="логин" />
        </div>
        <div class="email">
          <label for="">Почта</label>
          <input
            class="input"
            type="email"
            required
            placeholder="x@email.com"
          />
        </div>
        <div class="password">
          <label for="">Пароль</label>
          <input type="password" required placeholder="пароль" />
          <input
            type="password"
            class="repeat-password"
            required
            placeholder="повторите пароль"
          />
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
        <div class="step done">
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
          <input class="input" type="text" required placeholder="Иван" />
        </div>
        <div class="surname">
          <label for="">Фамилия</label>
          <input class="input" type="text" required placeholder="Иванов" />
        </div>
        <div class="tgid">
          <label for="">Телеграм</label>
          <input type="text" required placeholder="@telegram" />
        </div>
      </div>
      <div class="buttons">
        <a href="" on:click={() => (current_form = "registration")}>Назад</a>
        <a class="login-a" href="" on:click={() => (current_form = "confirm")}
          >Далее</a
        >
      </div>
    {/if}
    {#if current_form === "confirm"}
      <div class="steps">
        <div class="step done">
          <div class="circle">1</div>
          <p>Создание логина</p>
        </div>
        <div class="step done">
          <div class="circle">2</div>
          <p>Персональная информация</p>
        </div>
        <div class="step active">
          <div class="circle">3</div>
          <p>Подтверждение</p>
        </div>
      </div>
      <div class="verify-container">
        <p>на указанную почту x@email.com был выслан 6-ти значный код</p>
        <div class="verify">
          <p>Введите код для продолжения</p>
          <Code />
        </div>
      </div>

      <div class="buttons">
        <a href="" on:click={() => (current_form = "personal_info")}>Назад</a>
        <a class="login-a" href="" on:click={() => (current_form = "end")}
          >Далее</a
        >
      </div>
    {/if}
    {#if current_form === "end"}
      <div class="steps">
        <div class="step done">
          <div class="circle">1</div>
          <p>Создание логина</p>
        </div>
        <div class="step done">
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
        <a href="" on:click={() => (current_form = "personal_info")}>Назад</a>
        <a class="login-a" href="/" on:click={() => (current_form = "")}
          >На главную</a
        >
      </div>
    {/if}
  </div>
</section>

<style lang="css">
  .registr {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
  }

  .registr-container {
    padding: 50px 36px;
    width: 568px;
    height: 700px;
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
    color: var(--dark-grey);
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
  .email,
  .login,
  .name,
  .surname {
    display: flex;
    flex-direction: column;
  }
  .input {
    margin-bottom: 36px;
  }
  .password,
  .tgid {
    display: flex;
    flex-direction: column;
  }

  label {
    font-family: "Noto Sans Display Variable";
    font-size: 24px;
    margin-left: 10px;
    margin-bottom: 8px;
  }

  input {
    background-color: var(--light-grey);
    border: 0;
    height: 24px;
    border-radius: 50px;
    padding: 7px 16px;
    font-size: 20px;
  }

  .repeat-password {
    margin-top: 15px;
  }

  .buttons {
    display: flex;
    justify-content: center;
    width: 100%;
  }

  a {
    font-family: "Noto Sans Display Variable";
    font-size: 24px;
    color: var(--dark-grey);
    padding: 8px 27px;
    border: 3px solid var(--light-grey);
    border-radius: 50px;
    text-decoration: none;
  }

  .login-a {
    margin-left: 25px;
    display: flex;
    align-items: center;
    color: #fff;
    border: 0;
    background-image: var(--gradient);
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
    width: 42px;
    height: 42px;
    border-radius: 50%;
    background: #fff;
    border: 2px solid var(--dark-grey);
    color: var(--dark-grey);
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 24px;
    font-family: "Montserrat";
    font-weight: 7002;
  }

  .step p {
    font-family: "Noto Sans Display Variable";
    margin-top: 10px;
    font-size: 16px;
    color: var(--dark-grey);
  }
  .active .circle {
    color: var(--pink);
  }
  .active p {
    color: var(--pink);
  }

  .done .circle {
    color: #fff;
    background-image: var(--gradient);
    border: 0;
    height: 48px;
    width: 48px;
  }
  .done p {
    color: var(--pink);
  }

  .verify-container p {
    margin-bottom: 36px;
  }
</style>
