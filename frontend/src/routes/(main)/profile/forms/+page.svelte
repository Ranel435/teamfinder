<script lang="ts">
    interface form {
      hackaton: string;
      role: string;
      date: string;
      response_count: number;
      teams_count: number
  }


  let forms: form[] = [
    {
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4
    },
    {
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4
    },
    {
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4
    },
    {
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4
    },
    {
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4
    },
    {
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      response_count: 2,
      teams_count: 4
    }
  ]
</script>



<section class="forms">
  {#if forms.length === 0}
  <div class="no-forms">
    <p class="no-forms-text">Анкет пока нет</p>
    <div class="no-forms-buttons">
      <button>Создать</button>
      <button>Хакатоны</button>
      <button>Рынок вакансий</button>
    </div>
  </div>
  {:else}
  <div class="forms-container">
    <div class="forms-header">
      <h2>Активные анкеты</h2>
    </div>
    <ul class="forms-list">
      {#each forms as form, index}
      <div class="form">
        <div class="form-container">
          <div class="form-info">
            <h3>{form.hackaton}</h3>
            <p>{form.role} * {form.date}</p>
            <div class="form-info-links">
              <a href="">Отклики: {form.response_count}</a>
              <a href="">Подходящие команды: {form.teams_count}</a>
            </div>
          </div>
          <button class="delete-button" on:click={() => {forms.splice(index, 1); forms=[...forms]; }}> Удалить</button>
        </div>
      </div>
      {/each}
    </ul>
    <div class="forms-buttons">
      <button class="gradient-button">Все команды</button>
      <button class="gradient-button">Создать</button>
      <button class="button">Хакатоны</button>
    </div>
  </div>
  {/if}
</section>


<style>
  button {
    font-family: "Manrope";
    transition: all 0.5s ease;
  }
  .button {
    border: 3px solid #fff;
    padding: 10px 16px;
  }

  .forms {
    height: 100%;
    margin-left: 50px;
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .no-forms {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }

  .no-forms-text {
    font-family: "Manrope";
    font-size: 48px;
    color: #fff;
    margin-bottom: 32px;
  }
  .no-forms-buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .no-forms-buttons button {
    font-family: "Manrope";
    font-size: 24px;
    color: #fff;
    padding: 10px 24px;
    border: 3px solid #fff;
    border-radius: 50px;
    background-color: transparent;
    cursor: pointer;
    margin-right: 16px;
    transition: all 0.5 ease;
  }

  .no-forms-buttons button:last-child {
    margin-right: 0;
  }

  .no-forms-buttons button:hover {
    background-color: #fff;
    color: black;
  }

  .forms-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    padding: 32px;
  }

  .forms-header {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }

  .forms-list {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin: 40px 0;
    max-height: 50vh; /* Фиксированная высота для прокрутки */
    overflow-y: auto; /* Добавление вертикального скроллбара */
    padding-right: 10px; /* Добавляем отступ для скроллбара */
    box-sizing: content-box; /* Учитываем отступ, не влияя на ширину */
  }


  .forms-header h2{
    font-family: "Manrope";
    font-size: 48px;
    font-weight: 400;
    color: #fff;
  }


  .form {
    border: 2px solid #fff;
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .form:last-child {
    margin-bottom: 0;
  }

  .form-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
  }

  .form-container p {
    color: var(--dark-grey);
    font-family: "Manrope";
    font-style: 16px;
  }
  
  .form-container h3 {
    color: #fff;
    font-family: "Manrope";
    font-size: 24px;
  }

  .form-container a {
    font-family: "Manrope";
    font-style: 16px;
    color: #fff;
    text-decoration: none;
  }

  .form-container a:hover {
    text-decoration: underline;
  }

  .form-info-links {
    display: flex;
    flex-direction: row;
  }

  .form-info-links a{
    margin-right: 10px;
  } 

  .gradient-button {
    padding: 13px 19px;
    border: none;
    margin-right: 10px;
    transition: all 0.5s ease;
  }

  .forms-buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    width: 100%;
  }

  .forms-buttons button {
    font-size: 18px;
    color: #fff;
    border-radius: 50px;
    background-color: #000;
    width: 30%;
  }

  .button:hover {
    background-color: #fff;
    color: #000;
  }
  .gradient-button:hover {
    background:linear-gradient(90deg, #ED666E, #FBA31C);
  }
  /* Кнопка удаления */
  .delete-button {
    display: none;
    font-family: "Manrope";
    font-size: 20px;
    color: #fff;
    height: 50px;
    padding: 10px 24px;
    border: 3px solid #fff;
    border-radius: 50px;
    background-color: transparent;
    cursor: pointer;
    margin-right: 16px;
    transition: all 0.5 ease;
  }

  .delete-button:hover {
    color: #000;
    background-color: #fff;
  }


  .form:hover {
    /* border: 2px var(--gradient) #fff; */
    border-image: var(--pink);
  }

  /* Показываем кнопку при наведении */
  .form:hover .delete-button {
    display: block;
  }



  /* Стиль для скроллбара */
  .forms-list::-webkit-scrollbar {
    width: 26px; /* Ширина скроллбара */
  }

  .forms-list::-webkit-scrollbar-thumb {
    background-color: var(--dark-grey); /* Цвет ползунка */
    border: 2px solid #2a2a2a;
    
    border-radius: 32px; /* Скругление */
  }

  .forms-list::-webkit-scrollbar-thumb:hover {
    background-color: var(--light-grey); /* Цвет при наведении */
  }
</style>