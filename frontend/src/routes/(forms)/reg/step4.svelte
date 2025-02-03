<script>
  import Code from "$lib/components/code.svelte";


</script>


<section class="card">
  <div class="card__container">

    <header class="card__header">
      <h2 class="card__title m2">Регистрация завершена</h2>
    </header>

    <p class="card__text m6">Заполните информацию в профиле для создания анкеты или присоединения
      к командам</p>

    <div class="card__steps steps">
      <div class="steps_step step">
        <div class="step__circle done">1</div>
        <p class="step__text gradient-text">Создание логина</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle done">2</div>
        <p class="step__text gradient-text">Персональная информация</p>
      </div>
      <div class="steps__step step">
        <div class="step__circle done">3</div>
        <p class="step__text gradient-text">Подтверждение</p>
      </div>
    </div>

    <div class="card__buttons">
      <a class="card__button grey-button m6" href="/">На главную</a>
      <a class="card__button gradient-button m6" href="/profile/account">К настройкам профиля</a>
    </div>
  </div>
</section>






<style>
  .card {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .card__container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    text-align: center;
    flex-direction: column;

    background-color: var(--white);
    border-radius: 28px;
    padding: 36px;
    width: calc(480px - 72px);
    height: calc(414px - 72px);
  }

  .card__title {
    color: var(--black);
  }
  
  .card__text {
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

  .step__text {
    font-family: "Manrope-regular";
    font-size: 12px;
    line-height: 12px;
    margin-top: 7px;
  }

  .step__circle {
    font-family: "Montserrat";
    font-size: 18px;

    display: flex;
    align-items: center;
    justify-content: center;

    background-color: #fff;
    width: 29px;
    height: 29px;
    border: 2px solid var(--dark-grey);
    border-radius: 50%;
  }
  .done {
    background: var(--gradient);
    width: 31px;
    height: 31px;
    border: none;
    color: var(--white);
  }

  .simple {
    color: var(--dark-grey);
  }

  .card__button {
    text-decoration: none;
  }

  .grey-button {
    padding: 8px 16px;
    margin-right: 16px;
  }

  .gradient-button {
    padding: 10px 18px;
  }

</style>