<script lang="ts">
  import "../../../app.css";
  import Code from "$lib/code.svelte";
  $: current_form = "registration";
  $: main_text = "Заполните поля необходимой информацией";
  $: main_title = "Регистрация";
  let email: string = "";
  let errorMessage: string = "";

  async function sendEmailCode(email: string) {
    try {
      const response = await fetch('http://localhost:8090/auth/login/email', {
        method: 'POST', // Используйте POST для отправки данных
        headers: {
          'Content-Type': 'application/json', // Указываем тип контента
        },
        body: JSON.stringify({ email }), // Преобразуем объект в строку JSON
      });
      console.log('Response status:', response.status);
      console.log('Response headers:', response.headers);
      const responseText = await response.text();
      console.log('Response body:', responseText);

      if (response.ok) {
        const data = await response.json();
        console.log("done");
        //message = data.message; // Сообщение от сервера
      } else {
        const errorData = await response.json();
        console.log(1111);
        //message = errorData.error; // Обработка ошибок
      }
    } catch (error) {
      console.error('Ошибка при отправке запроса:', error);
      console.log(email);
      //message = 'Ошибка подключения. Попробуйте еще раз позже.';
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
    <!-- </div> -->
    
    <!-- <div class="registr-container-main"> -->
      {#if current_form === "registration"}
      <div class="registr-container-main">
        <div class="form">
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
            <input type="password" required placeholder="пароль" />
            <input type="password" class="repeat-password" required placeholder="повторите пароль"/>
          </div>
        </div>
      </div>
      {/if}
      {#if current_form === "personal_info"}
      <div class="registr-container-main">
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
            <input type="text" required placeholder="@telegram" />
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
      </div>
      {/if}
      {#if current_form === "confirm"}
      <div class="registr-container-main">
        <div class="verify-container">
          <p>на указанную почту {email} был выслан 6-ти значный код</p>
          <div class="verify">
            <p style="color: #000">Введите код для продолжения</p>
            <Code />
          </div>
        </div>
      </div>

      {/if}
      <!-- {#if current_form === "end"} 
      <div class="registr-container-main">
        <p>Заполните информацию в профиле для создания анкеты или присоединения
          к командам</p>
      </div>
      {/if} -->
    <!-- </div> -->

    <div class="registr-container-buttons">
      {#if current_form === "registration"}
      <div class="buttons">
        <a href="/forms/auth">Назад</a>
        <a href="" class="next-button" on:click={() => {current_form = "personal_info"}}>Далее</a>
      </div>
      {/if}
      {#if current_form === "personal_info"}
      <div class="buttons">
        <a href="" on:click={() => {current_form = "registration"; main_title = "Регистрация"; main_text = "Зполните поля необходимой информацией"}}>Назад</a>
        <a href="" class="next-button" on:click={() => {sendEmailCode(email); current_form = "confirm"; main_title = "Подтверждение"; main_text = ""}}>Далее</a>
      </div>
      {/if}
      {#if current_form === "confirm"}
      <div class="buttons">
        <a href="" on:click={() => {current_form = "personal_info"; main_title = "Регистрация"; main_text = "Зполните поля необходимой информацией"}}>Назад</a>
      <a href="" class="next-button" on:click={() => {current_form = "end"; main_title = "Регистрация завершена"; main_text = "Заполните информацию в профиле для создания анкеты или присоединения к командам"}}>Завершить</a>
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
    margin: 28px 0;
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

  a {
    font-family: "Manrope";
    font-size: 18px;
    color: var(--dark-grey);
    text-decoration: none;

    padding: 10px 18px;
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