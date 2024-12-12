<script lang="ts">
  import { goto } from "$app/navigation";
  import Code from "$lib/components/code.svelte";
  let current_form = "first";
  let main_title = "Восстановление пароля";

  let email: string = "";
  let login: string ="";
</script>


<section class="forget_password" >
  <div class="forget_password-container">
    <div class="forget_password-container-header">
      <h2>{main_title}</h2>
    </div>
    <div class="forget_password-container-main">
      {#if current_form === "first"}
        <form class="form"  on:submit|preventDefault={() => {
          if (!email || !login) {
          } else {
            current_form = "second"; // Переход на следующий шаг
            main_title = "Подтверждение входа";
          }
        }}>
          <div class="email">
            <label for="">Email</label>
            <input type="email" required placeholder="x@email.com" bind:value={email}/>
          </div>
          <div class="password">
            <label for="">Логин</label>
            <input type="text" pattern="^[A-Za-z]+$" required placeholder="логин" bind:value={login}/>
          </div>
          <div class="forget_password-container-buttons">
            <div class="buttons">
              <button class="button" on:click={() => {current_form = "gg"; main_title = "ГГ("}}>Нет доступа к почте</button>
              <button class="next-button" type="submit">Далее</button>
            </div>
          </div>
        </form>
        
      {/if}

      {#if current_form === "second"}
      <div class="verify-container">
        <p>на указанную почту x@email.com был выслан 6-ти значный код</p>
        <div class="verify">
          <p style="color: #000">Введите код для продолжения</p>
          <Code />
        </div>
      </div>
      <div class="forget_password-container-buttons">
      <div class="buttons">
        <button class="button" on:click={() => {current_form = "first"; main_title = "Восстановление пароля"}}>Назад</button>
        <button on:click={() => goto("/")} class="next-button">Завершить</button>
      </div>
      <div class="repeat-button">
        <button class="button">Выслать повторно</button>
      </div>
      </div>
      {/if}

      {#if current_form === "gg"}
      
      <div class="gg-container">
        <p>Чтобы восстановить аккаунт, обратитесь в поддержку</p>
      </div>
      <div class="forget_password-container-buttons">

      <div class="buttons">
        <button class="button" on:click={() => goto("/")}>Поддержка</button>
        <button style="margin-left: 25px;">На главную</button>
      </div>
      </div>
      {/if}
    </div>
  </div>
</section>


<style>
  .forget_password {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .forget_password-container {
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

  .forget_password-container-header {
    display: flex;
    flex-direction: column;
    text-align: center;
  }

  .forget_password-container-header h2 {
    font-family: "Manrope";
    font-size: 36px;
    font-weight: normal;
    line-height: 48px;
    margin: 0;
    margin-bottom: 24px;
  }

  /* main */
  .forget_password-container-main {
    display: flex;
    flex-direction: column;
    text-align: left;
    justify-content: space-between;
    width: 100%;
    margin-bottom: 18px;
  }

  .email, .password {
    display: flex;
    flex-direction: column;
  }

  .email {
    margin-bottom: 22px;
  }

  input {
    background-color: var(--light-grey);
    border: 0;
    border-radius: 50px;
    padding: 10px 16px;
    font-size: 16px;
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
    margin-bottom: 16px;
  }

    /* buttons */
  .forget_password-container-buttons {
    margin-top: 16px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
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

  .next-button {
    padding:11px 19px;
    margin-left: 25px;
    align-items: center;
    color: #fff;
    border: 0;
    background-image: var(--gradient);
  }

  .repeat-button {
    margin-top: 16px;
  }

  .gg-container p{
    font-family: "Manrope";
    font-size: 24px;
    color: var(--dark-grey);
    text-align: center;
  }
</style>