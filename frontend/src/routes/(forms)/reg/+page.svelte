<script lang="ts">
  import "../../../app.css";
  import Code from "$lib/code.svelte";
  import { goto } from "$app/navigation";

  let current_form: string = "registration"; // переменная для отображения текущей формы
  let main_text: string = "Заполните поля необходимой информацией"; // переменная для текста формы
  let main_title: string = "Регистрация"; // переменная для заголовка формы
  let repeatPasswordPlaceholder: string = "повторите пароль"; // переменная для placeholderа поля ввода повтора пароля

  let password: string = "";
  let repeatPassword: string = "";
  let passwordsMatch: boolean = password === repeatPassword;

  let email: string = "";
  let emailCode: string = ""; // Хранит введенный код подтверждения
  let confirmationCode = ["", "", "", "", "", ""]; // цифры кода подтверждения
  
  let name: string = "";
  let surname: string = "";
  let tgid: string = "";
  let lookingForTeam: boolean = false;
  let formError: string = "";

  async function sendEmailCode(email: string) {
    try {
      const response = await fetch('http://localhost:8090/auth/login/email', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email }),
      });
      if (response.ok) {
        const data = await response.json();
        console.log("done");
      } else {
        const errorData = await response.json();
        console.log("error");
      }
    } catch (error) {
      console.error('Ошибка при отправке запроса:', error);
    }
  }

  async function verifyEmailCode(email: string, password: string) {
    console.log(password);
    try {
      const response = await fetch('http://localhost:8090/auth/login/email', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(
          { email: email, password: password}),
      });
      if (response.ok) {
        const data = await response.json();
        console.log("done");
      } else {
        const errorData = await response.json();
        console.log("error");
      }
    } catch (error) {
      console.error('Ошибка при отправке запроса:', error);
    }
  }

</script>


<section class="registr">
  <div class="registr-container {current_form}">
    <!-- <div class="registr-container-header"> -->
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
            <label for="">Логин</label>
            <input type="text" required placeholder="логин"/>
          </div>
          <div class="email">
            <label for="">Почта</label>
            <input type="email" required placeholder="x@email.com" bind:value={email}/>
          </div>
          <div class="password">
            <label for="">Пароль</label>
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
            if (!name || !surname || !tgid) {
              formError = "Пожалуйста, заполните все обязательные поля.";
            } else {
              formError = ""; // Если ошибки нет, сбрасываем сообщение
              sendEmailCode(email); // Отправка кода на почту
              current_form = "confirm"; // Переход на следующий шаг
              main_title = "Подтверждение";
              main_text = "";
            }
          }}>
          <div class="name">
            <label for="name">Имя</label>
            <input id="name" type="text" required placeholder="Иван" bind:value={name}/>
          </div>
          <div class="surname">
            <label for="surname">Фамилия</label>
            <input id="surname" type="text" required placeholder="Иванов" bind:value={surname}/>
          </div>
          <div class="tgid">
            <label for="tgid">Телеграм</label>
            <input id="tgid" type="text" required placeholder="@telegram" bind:value={tgid}/>
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
          
          {#if formError}
            <p class="error-message">{formError}</p>
          {/if}

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
          <p>на указанную почту {email} был выслан 6-ти значный код</p>
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
        <a href="" on:click={() => {current_form = "personal_info"; main_title = "Регистрация"; main_text = "Зполните поля необходимой информацией"}}>Назад</a>
        <a href="" class="next-button" on:click={() => {verifyEmailCode(email, confirmationCode.join('')); current_form = "end"; main_title = "Регистрация завершена"; main_text = "Заполните информацию в профиле для создания анкеты или присоединения к командам"}}>Завершить</a>
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
</section>

<style>
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

  label {
    font-family: "Manrope";
    font-size: 18px;
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

  a, button{
    font-family: "Manrope";
    font-size: 18px;
    color: var(--dark-grey);
    text-decoration: none;

    padding: 10px 18px;
    border: 3px solid var(--light-grey);
    border-radius: 50px;
    
  }
  
  .buttons {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 16px;
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