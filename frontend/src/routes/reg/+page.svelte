<script lang="ts">
  import "../../app.css";
  import Code from "$lib/code.svelte";
  $: current_form = "registration";
  $: main_text = "Заполните поля необходимой информацией";
  $: main_title = "Регистрация";
</script>


<section class="registr">
  <div class="registr-container {current_form}">
    <div class="registr-container-header">
      <div class="registr-container-header-text">
        <h2>{main_title}</h2>
        <p>{main_text}</p>
      </div>
      {#if current_form === "registration"}
      <div class="steps">
        <div class="step active">
          <div class="step-circle">1</div>
          <p class="step-text">Создание логина</p>
        </div>
        <div class="step">
          <div class="step-circle">2</div>
          <p class="step-text">Персональная информация</p>
        </div>
        <div class="step">
          <div class="step-circle">3</div>
          <p class="step-text">Подтверждение</p>
        </div>
      </div>
      {/if}
      {#if current_form === "personal_info"}
      <div class="steps">
        <div class="step done">
          <div class="step-circle">1</div>
          <p class="step-text">Создание логина</p>
        </div>
        <div class="step active">
          <div class="step-circle">2</div>
          <p class="step-text">Персональная информация</p>
        </div>
        <div class="step">
          <div class="step-circle">3</div>
          <p class="step-text">Подтверждение</p>
        </div>
      </div>
      {/if}
      {#if current_form === "confirm"}
      <div class="steps">
        <div class="step done">
          <div class="step-circle">1</div>
          <p class="step-text">Создание логина</p>
        </div>
        <div class="step done">
          <div class="step-circle">2</div>
          <p class="step-text">Персональная информация</p>
        </div>
        <div class="step active">
          <div class="step-circle">3</div>
          <p class="step-text">Подтверждение</p>
        </div>
      </div>
      {/if}
      {#if current_form === "end"}
      <div class="steps">
        <div class="step done">
          <div class="step-circle">1</div>
          <p class="step-text">Создание логина</p>
        </div>
        <div class="step done">
          <div class="step-circle">2</div>
          <p class="step-text">Персональная информация</p>
        </div>
        <div class="step done">
          <div class="step-circle">3</div>
          <p class="step-text">Подтверждение</p>
        </div>
      </div>
      {/if}
    </div>
    
    <div class="registr-container-main">
      {#if current_form === "registration"}
      <div class="form">
        <div class="login">
          <label for="">Логин</label>
          <input type="text" required placeholder="логин"/>
        </div>
        <div class="email">
          <label for="">Почта</label>
          <input type="email" required placeholder="x@email.com"/>
        </div>
        <div class="password">
          <label for="">Пароль</label>
          <input type="password" required placeholder="пароль" />
          <input type="password" class="repeat-password" required placeholder="повторите пароль"/>
        </div>
      </div>
      {/if}
      {#if current_form === "personal_info"}
      <div class="form">
        <div class="name">
          <label for="">Имя</label>
          <input type="text" required placeholder="Иван"/>
        </div>
        <div class="surname">
          <label for="">Фамилия</label>
          <input type="text" required placeholder="Иванов"/>
        </div>
        <div class="tgid">
          <label for="">Телеграм</label>
          <input type="text" required placeholder="@telegram"/>
        </div>
        <div class="looking-for-team">
          <label for="">Ищу команду</label>
          <div class="switch">
            <input type="checkbox" id="toggle" class="switch-input">
            <label for="toggle" class="switch-label">
                <span class="switch-inner"></span>
                <span class="switch-switch"></span>
            </label>
        </div>
        </div>
      </div>
      {/if}
      {#if current_form === "confirm"}
      <div class="verify-container">
        <p>на указанную почту x@email.com был выслан 6-ти значный код</p>
        <div class="verify">
          <p style="color: #000">Введите код для продолжения</p>
          <Code />
        </div>
      </div>
      {/if}
      {#if current_form === "end"}
      {/if}
    </div>

    <div class="registr-container-buttons">
      {#if current_form === "registration"}
      <div class="buttons">
        <a href="/auth">Назад</a>
        <a href="" class="next-button" on:click={() => {current_form = "personal_info"}}>Далее</a>
      </div>
      {/if}
      {#if current_form === "personal_info"}
      <div class="buttons">
        <a href="" on:click={() => {current_form = "registration"; main_title = "Регистрация"; main_text = "Зполните поля необходимой информацией"}}>Назад</a>
        <a href="" class="next-button" on:click={() => {current_form = "confirm"; main_title = "Подтверждение"; main_text = ""}}>Далее</a>
      </div>
      {/if}
      {#if current_form === "confirm"}
      <div class="buttons">
        <a href="" on:click={() => {current_form = "personal_info"; main_title = "Регистрация"; main_text = "Зполните поля необходимой информацией"}}>Назад</a>
      <a href="" class="next-button" on:click={() => {current_form = "end"; main_title = "Регистрация завершена"; main_text = "Заполните информацию в профиле для создания анкеты или присоединенияк командам"}}>Завершить</a>
      </div>
      <div class="repeat-button">
        <a href="" class="repeat-button">Выслать повторно</a>
      </div>
      {/if}
      {#if current_form === "end"}
      <div class="buttons">
        <a href="">К настройкам профиля</a>
        <a href="/" class="next-button">На главную</a>
      </div>
      {/if}
    </div>
  </div>
</section>

<style>
  /* header */
  .registr {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    margin: 100px 0;
  }

  .registr-container {
    display: flex;
    flex-direction: column;
    text-align: center;
    align-items: center;
    justify-content: space-between;

    background-color: #fff;
    border-radius: 32px;
    width: 568px;
    height: 700px;
    padding: 50px 36px;
  }

  .registr-container-header {
    display: flex;
    flex-direction: column;
    text-align: center;
  }

  .registr-container-header-text h2 {
    font-family: "Manrope";
    font-size: 48px;
    font-weight: normal;
    line-height: 48px;
    margin: 0;
    margin-bottom: 24px;
  }

  .registr-container-header-text p {
    font-family: "Manrope";
    font-size: 24px;
    color: var(--dark-grey);

    margin: 0;
  }

  .steps {
    display: flex;
    margin: 38px 0;
  }

  .step {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;

    width: 30%;
    margin: 0 auto;
  }

  .step-text {
    font-family: "Manrope";
    font-size: 16px;
    color: var(--dark-grey);
    margin-top: 10px;
  }

  .step-circle {
    font-family: "Montserrat";
    font-size: 24px;
    color: var(--dark-grey);

    display: flex;
    align-items: center;
    justify-content: center;

    background-color: #fff;
    width: 42px;
    height: 42px;
    border: 2px solid var(--dark-grey);
    border-radius: 50%;
  }

  .active .step-circle {
    color: var(--pink);
  }

  .active p {
    color: var(--pink);
  }
  
  .done .step-circle {
    color: #fff;
    background-image: var(--gradient);
    border: 0;
    height: 48px;
    width: 48px;
  }

  .done p {
    color: var(--pink);
  }

  /* main */
  .registr-container-main {
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
  .surname,
  .password,
  .tgid {
    display: flex;
    flex-direction: column;
  }

  .looking-for-team {
    display: flex;
  }

  .login, .email {
    margin-bottom: 36px;
  }

  .name, .surname, .tgid {
    margin-bottom: 24px;
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

  label {
    font-family: "Manrope";
    font-size: 24px;
    margin-left: 10px;
    margin-bottom: 8px;
  }

  .verify-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
  }

  .verify-container p {
    font-family: "Manrope";
    color: var(--dark-grey);
    font-size: 24px;
    margin-bottom: 32px;
  }

  /* buttons */
  .registr-container-buttons {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 100%;
  }

  a {
    font-family: "Manrope";
    font-size: 24px;
    color: var(--dark-grey);
    text-decoration: none;

    padding: 8px 27px;
    border: 3px solid var(--light-grey);
    border-radius: 50px;
    
  }
  
  .next-button {
    margin-left: 25px;
    /* display: flex; */
    align-items: center;
    color: #fff;
    border: 0;
    background-image: var(--gradient);
  }

  .repeat-button {
    margin-top: 36px;
  }

  /* switcher */
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
      width: 56px;
      height: 32px;
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
      width: 24px;
      height: 24px;
      background-color: white;
      border-radius: 50%;
      transition: transform 0.3s ease;
      left: 2px;
  }

  .switch-input:checked + .switch-label {
      background-image: var(--gradient); /* Цвет при включенном состоянии */
  }

  .switch-input:checked + .switch-label .switch-switch {
      transform: translateX(28px); /* Перемещение переключателя */
  }

  .end {
    width: 568px;
    height: 408px;
  }
</style>