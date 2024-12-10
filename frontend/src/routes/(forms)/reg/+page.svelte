<script lang="ts">
  import Code from "$lib/components/code.svelte"; // компонент ввода кода
  import { goto } from "$app/navigation"; // навигация между страницами

  import { sendEmailCode, verifyEmailCode } from '$lib/api/auth';
  import { saveTokens } from '$lib/stores/authStore';
  import { profile } from '$lib/stores/data';

  let current_form: string = "registration"; // переменная для отображения текущей формы
  let main_text: string = "Заполните поля необходимой информацией"; // переменная для текста формы
  let main_title: string = "Регистрация"; // переменная для заголовка формы
  let repeatPasswordPlaceholder: string = "повторите пароль"; // переменная для placeholderа поля ввода повтора пароля

  let password: string = "";
  let repeatPassword: string = "";
  let passwordsMatch: boolean = password === repeatPassword;

  let emailCode: string = ""; // Хранит введенный код подтверждения
  let confirmationCode: string[] = ["", "", "", "", "", ""]; // цифры кода подтверждения
  
  let lookingForTeam: boolean = false;

  // Отправка email
  async function handleSendEmail() {
    const success = await sendEmailCode($profile.email);
    if (!success) {
      alert('Не удалось отправить код. Попробуйте снова.\n');

    } else {
      alert('Код отправлен');
    }
  }

  // Проверка кода
  async function handleVerifyCode() {
    const tokens = await verifyEmailCode($profile.email, confirmationCode.join(''));
    if (tokens) {
      saveTokens(tokens.access, tokens.refresh);
      current_form = "end";
      main_title = "Регистрация завершена"; 
      main_text = "Заполните информацию в профиле для создания анкеты или присоединения к командам";
      return 1;
    } else {
      alert('Неверный код подтверждения.');
      return 0;
    }
  }





</script>


<section class="registr">
  <div class="registr-container {current_form}">
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

      {#if current_form === "registration"}
      <div class="registr-container-main">
        <form class="form" 
          on:submit|preventDefault={() => {
          if (password !== repeatPassword) {
            repeatPassword = "";
            repeatPasswordPlaceholder = "Пароли не совпадают";
            return; // Прерываем обработку, если пароли не совпадают
          }
          repeatPassword = "";
          repeatPasswordPlaceholder = "повторите пароль"
          current_form = "personal_info"; // Переходим на следующий шаг
          }}>
          <div class="login">
            <label for="">Логин (только латинские буквы)</label>
            <input type="text" pattern="^[A-Za-z]+$" required placeholder="логин" bind:value={$profile.login}/>
          </div>
          <div class="email">
            <label for="">Почта</label>
            <input type="email" required placeholder="x@email.com" bind:value={$profile.email}/>
          </div>
          <div class="password">
            <label for="">Пароль  (от 8 символов)</label>
            <input type="password" minlength="8" required placeholder="пароль" bind:value={password}/>
            <input type="password" class="repeat-password" required placeholder={repeatPasswordPlaceholder} bind:value={repeatPassword}/>
          </div>
          <div class="buttons">
            <button type="button" on:click={() => {goto("/auth")}}>Назад</button>
            <button type="submit" class="next-button">Далее</button>
          </div>
        </form>
        
      </div>
      {/if}
      {#if current_form === "personal_info"}
      <div class="registr-container-main">
        <form class="form" 
          on:submit|preventDefault={() => {
            if (!$profile.name || !$profile.surname || !$profile.tgid) {
            } else {
              sendEmailCode($profile.email); // Отправка кода на почту
              current_form = "confirm"; // Переход на следующий шаг
              main_title = "Подтверждение";
              main_text = "";
            }
          }}>
          <div class="name">
            <label for="name">Имя</label>
            <input id="name" pattern="^[A-Za-zА-Яа-яЁё]+$" type="text" required placeholder="Иван" bind:value={$profile.name}/>
          </div>
          <div class="surname">
            <label for="surname">Фамилия</label>
            <input id="surname" pattern="^[A-Za-zА-Яа-яЁё]+$" type="text" required placeholder="Иванов" bind:value={$profile.surname}/>
          </div>
          <div class="tgid">
            <label for="tgid">Телеграм</label>
            <input id="tgid" type="text" required pattern="^@([A-Za-z0-9_]+)$" placeholder="@telegram" bind:value={$profile.tgid}/>
          </div>
          <div class="looking-for-team">
            <label for="toggle">Ищу команду</label>
            <div class="switch">
              <input type="checkbox" id="toggle" class="switch-input" bind:checked={lookingForTeam} />
              <label for="toggle" class="switch-label">
                <span class="switch-inner"></span>
                <span class="switch-switch"></span>
              </label>
            </div>
          </div>

          <div class="buttons">
            <button type="button" on:click={() => {
                current_form = "registration"; 
                main_title = "Регистрация"; 
                main_text = "Заполните поля необходимой информацией";
              }}
            >Назад</button>
            <button type="submit" class="next-button">Далее</button>
          </div>
        </form>
      </div>
      {/if}
      {#if current_form === "confirm"}
      <div class="registr-container-main">
        <div class="verify-container">
          <p>на указанную почту {$profile.email} был выслан 6-ти значный код</p>
          <div class="verify">
            <p style="color: #000">Введите код для продолжения</p>
            <Code bind:code={confirmationCode} />
          </div>
        </div>
      </div>
      {/if}


    <div class="registr-container-buttons">
      {#if current_form === "confirm"}
      <div class="buttons">
        <button on:click={() => {current_form = "personal_info"; main_title = "Регистрация"; main_text = "Зполните поля необходимой информацией"}}>Назад</button>
        <button class="next-button" on:click={handleVerifyCode}>Завершить</button>
      </div>
      <div class="repeat-button">
        <button class="repeat-button" on:click={handleSendEmail}>Выслать повторно</button>
      </div>
      {/if}
      {#if current_form === "end"}
      <div class="buttons">
        <button on:click={() => goto("/profile/account")}>К настройкам профиля</button>
        <button on:click={() => goto("/")} class="next-button">На главную</button>
      </div>
      {/if}
    </div>
</section>

<style>
  label {
    font-family: "Manrope";
    font-size: 18px;
    margin-left: 10px;
    margin-bottom: 8px;
  }

  /* header */
  .registr {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .registr-container {
    display: flex;
    flex-direction: column;
    text-align: center;
    align-items: center;
    justify-content: space-between;

    background-color: #fff;
    border-radius: 28px;
    width: 408px;
    height: 552px;
    padding: 36px 36px;
  }

  .registr-container-header {
    display: flex;
    flex-direction: column;
    text-align: center;
  }

  .registr-container-header-text h2 {
    font-family: "Manrope";
    font-size: 36px;
    font-weight: normal;
    margin-bottom: 10px;
  }

  .registr-container-header-text p {
    font-family: "Manrope";
    font-size: 18px;
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

  .step-text {
    font-family: "Manrope";
    font-size: 12px;
    line-height: 12px;
    color: var(--dark-grey);
    margin-top: 7px;
  }

  .step-circle {
    font-family: "Montserrat";
    font-size: 18px;
    color: var(--dark-grey);

    display: flex;
    align-items: center;
    justify-content: center;

    background-color: #fff;
    width: 29px;
    height: 29px;
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
    height: 33px;
    width: 33px;
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
    width: calc(100%);
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
    margin-bottom: 12px;
  }

  .name, .surname, .tgid {
    margin-bottom: 12px;
  }

  input {
    background-color: var(--light-grey);
    border: 0;
    border-radius: 50px;
    padding: 10px 16px;
    font-size: 16px;
  }


  

  .repeat-password {
    margin-top: 15px;
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
    font-size: 18px;
    margin-bottom: 32px;
  }

  /* buttons */
  .registr-container-buttons {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 100%;
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
    /* margin-right: 20px; */
    cursor: pointer;
    transition: all 0.5s ease;
  }
  
  .button:hover {
    background-color: var(--light-grey);
    color: var(--dark-grey);
  }


  .buttons {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 16px;
  }

  .next-button {
    padding: 11px 19px;
    margin-left: 25px;
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

  .end {
    width: 408px;
    height: 342px;
  }
</style>