<script lang="ts">
  import Code from "$lib/components/code.svelte";
  let current_form = "first";
  let main_title = "Восстановление пароля";
</script>


<section class="forget_password">
  <div class="forget_password-container">
    <div class="forget_password-container-header">
      <h2>{main_title}</h2>
    </div>
    <div class="forget_password-container-main">
      {#if current_form === "first"}
        <div class="form">
          <div class="email">
            <label for="">Email</label>
            <input type="text" required placeholder="x@email.com"/>
          </div>
          <div class="password">
            <label for="">Логин</label>
            <input type="text" required placeholder="логин"/>
          </div>
        </div>
      {/if}

      {#if current_form === "second"}
      <div class="verify-container">
        <p>на указанную почту x@email.com был выслан 6-ти значный код</p>
        <div class="verify">
          <p style="color: #000">Введите код для продолжения</p>
          <Code />
        </div>
      </div>
      {/if}

      {#if current_form === "gg"}
      <div class="gg-container">
        <p>Чтобы восстановить аккаунт, обратитесь в поддержку</p>
      </div>
      {/if}
    </div>

    <div class="forget_password-container-buttons">
      {#if current_form === "first"}
      <div class="buttons">
        <a href="" on:click={() => {current_form = "gg"; main_title = "ГГ("}}>Нет доступа к почте</a>
        <a href="" class="next-button" on:click={() => {current_form = "second"; main_title = "Подтверждение входа"}}>Далее</a>
      </div>
      {/if}

      {#if current_form === "second"}
      <div class="buttons">
        <a href="" on:click={() => {current_form = "first"; main_title = "Восстановление пароля"}}>Назад</a>
        <a href="/" class="next-button">Завершить</a>
      </div>
      <div class="repeat-button">
        <a href="" class="repeat-button">Выслать повторно</a>
      </div>
      {/if}

      {#if current_form === "gg"}
      <div class="buttons">
        <a href="/">Поддержка</a>
        <a href="" style="margin-left: 25px;">На главную</a>
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
    width: 568px;
    height: 408px;
    padding: 50px 36px;
  }

  .forget_password-container-header {
    display: flex;
    flex-direction: column;
    text-align: center;
  }

  .forget_password-container-header h2 {
    font-family: "Manrope";
    font-size: 48px;
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
    width: calc(100% - 54px);
    margin-bottom: 38px;
  }

  .email, .password {
    display: flex;
    flex-direction: column;
  }

  .email {
    margin-bottom: 36px;
  }

  input {
    background-color: var(--light-grey);
    border: 0;
    height: 24px;
    border-radius: 50px;
    padding: 7px 16px;
    font-size: 20px;
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
  .forget_password-container-buttons {
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
    align-items: center;
    color: #fff;
    border: 0;
    background-image: var(--gradient);
  }

  .repeat-button {
    margin-top: 36px;
  }

  .gg-container p{
    font-family: "Manrope";
    font-size: 24px;
    color: var(--dark-grey);
    text-align: center;
  }
</style>