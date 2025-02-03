<script lang="ts">
  import { goto } from "$app/navigation";
  import Code from "$lib/components/code.svelte";
  let current_form = "first";
  let main_title = "Восстановление пароля";

  let email: string = "";
  let login: string ="";
</script>


<section class="forget-password">
  <div class="forget-password__container">
    
      {#if current_form === "first"}
      <header class="forget-password__header">
        <h2 class="forget-password__title m2">Восстановление пароля</h2>
      </header>
      <div class="forget-password__main">
        <form class="forget-password__form form"  on:submit|preventDefault={() => {
          if (!email || !login) {
          } else {
            current_form = "second"; 
            main_title = "Подтверждение входа";
          }
        }}>
          <div class="form__email email">
            <label class="email__label m6" for="">Email</label>
            <input class="email__input m7" type="email" required placeholder="x@email.com" bind:value={email}/>
          </div>
          <div class="form__password password">
            <label class="password__label m6" for="">Фамилия</label>
            <input class="password__input m7" type="text" pattern="^[А-Яа-я]+$" required placeholder="Иванов" bind:value={login}/>
          </div>
          <div class="forget-password__buttons">
              <button class="forget-password__button grey-button m6" on:click={() => {current_form = "gg"; main_title = "ГГ("}}>Нет доступа к почте</button>
              <button class="forget-password__button gradient-button m6" type="submit">Далее</button>
          </div>
        </form>
      </div>
      {/if}

      {#if current_form === "second"}
      <header class="verify__header">
        <h2 class="verify__title m2">Восстановление пароля</h2>
      </header>
      <div class="verify">
        <p class="verify__text buttonM">На указанную почту было выслано письмо с инструкцией
          для восстановления пароля</p>
        
      </div>
      <div class="verify__buttons">
        <button class="verify__button grey-button m6" on:click={() => {current_form = "first"; main_title = "Восстановление пароля"}}>Назад</button>
        <button class="verify__button gradient-button m6">Вход</button>
      </div>
      {/if}

      {#if current_form === "gg"}
      <header class="gg__header m2">ГГ(</header>
      <p class="gg__text buttonM">Чтобы восстановить аккаунт, обратитесь в поддержку</p>
      <div class="gg__buttons">
        <a class="gg__button grey-button m6" href="/">Поддержка</a>
        <a class="gg__button grey-button m6" href="/">На главную</a>
      </div>
      {/if}
  </div>
</section>


<style>
  .forget-password {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .forget-password__container {
    display: flex;
    flex-direction: column;
    text-align: center;
    align-items: center;
    justify-content: space-between;

    background-color: #fff;
    border-radius: 32px;
    width: 408px;
    height: 342px;
    padding: 36px 36px;
  }

  .forget-password__header {
    display: flex;
    flex-direction: column;
    text-align: center;
  }

  .forget-password__title {
    margin-bottom: 24px;
  }

  .forget-password__main {
    display: flex;
    flex-direction: column;
    text-align: left;
    justify-content: space-between;
    width: 100%;
    margin-bottom: 18px;
  }

  .form__email, .form__password {
    display: flex;
    flex-direction: column;
  }

  .form__email {
    margin-bottom: 22px;
  }

  .email__input, .password__input {
    background-color: var(--light-grey);
    border: 0;
    border-radius: 50px;
    padding: 10px 16px;
  }

  .email__label, .password__label {
    margin-left: 10px;
    margin-bottom: 8px;
  }

  .forget-password__buttons {
    margin-top: 16px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    width: 100%;
  }

  .grey-button {
    padding: 8px 16px;
    margin-right: 12px;
  }

  .gradient-button {
    padding: 10px 18px;
  }





  .verify {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    text-align: center;
  }

  .verify__text {
    height: 100%;
    justify-content: center;
    align-items: center;
    color: var(--dark-grey);
  }

  .gg__text {
    color: var(--dark-grey);
  }

  .gg__button {
    text-decoration: none;
  }
</style>