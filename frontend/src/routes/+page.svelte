

<section class="main-section">
    <!-- <img src="../static/Ellipse 33.svg" class = "background" alt=""> -->
    <div class="text">
        <p>Команда найдется для каждого</p>
    </div>
    <div class="main-text">
        <img src="Начни Хакатонить!.svg" alt="">
    </div>
    <div class="main-buttons">
        <a  class="main-button" href="">Свободные агенты</a>
        <a class="main-button" href="">Найти команду</a>
        <a class="main-button" href="">Создать анкету</a>
    </div> 
</section>
<section class="popular-roles">
    <h2 class="title">Популярные роли</h2>
    <div class="roles-grid">
      <button class="role backend">
        <div class="role-name">Backend</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role frontend">
        <div class="role-name">Frontend</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role fullstack">
        <div class="role-name">Fullstack dev</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role analytics">
        <div class="role-name">Аналитик данных</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role pm">
        <div class="role-name">Product manager</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role ui/ux">
        <div class="role-name">UX / UI дизайнер</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role ios">
        <div class="role-name">IOS / Android dev</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role ml">
        <div class="role-name">ML engineer</div>
        <div class="role-teams">14 команд</div>
      </button>
      <button class="role devops">
        <div class="role-name">DevOps инженер</div>
        <div class="role-teams">14 команд</div>
      </button>
    </div>
    <div class="more-roles">
        <a class="more" href="/">Еще роли</a>
    </div>
</section>


<style>
    * {
        color: #fff;
    }
    .main-section {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        margin-top: 147px;
    }
    .main-text {
        text-align: center;
        margin: 72px 0 72px;
    }
    .main-buttons {
        margin-bottom: 128px;
    }
    /* .gradient-text {
        background-image: linear-gradient(90deg, #E5469A, #FBA31C);
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
    } */
    
    p {
        font-family: 'Igra Sans Variable';
        font-size: 24px;
        line-height: 130%;
        letter-spacing: 5%;
        padding: 0;
        margin: 0;
    }

    .main-button {
        font-family: 'Noto Sans Display Variable';
        max-width: 164px;
        text-decoration: none;
        padding: 3px 50px;
        border: 3px solid;
        border-color: #fff;
        border-radius: 50px;
        font-size: 24px;
        margin:0 18px;
    }
    .main-button:hover {
        transition: all 0.5s ease; /* Применить плавный переход ко всем свойствам */
        color: #000;
        background-color: #fff;
        /* background-color: #45a049; */
    }



    /* General styles */
    .popular-roles {
        display: flex;
        flex-wrap: wrap;
        flex-direction: column;
        justify-content: center;
        margin-top: 198px;
        align-items: center;
        text-align: center;
    }

    /* Title styles */
    .title {
        font-size: 48px;
        margin-bottom: 80px;
        font-family: 'Igra Sans Variable';
        letter-spacing: -3%;
        font: 400;
        padding: 0;
    }

    /* Grid container */
    .roles-grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        /* grid-auto-rows: 140px;  */
    }

    /* Role card styles */
    .role {
        position: relative; /* Обязателен для правильного позиционирования псевдоэлемента */
        overflow: hidden; /* Обрезает "выходящий" псевдоэлемент */

        display: flex;
        flex: 1; /* Указывает, что все блоки занимают равное пространство */
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        text-align: left;

        border: 4.55px solid #fff;
        border-radius: 24px;

        padding: 10px;
        margin: 0px 36px 36px 0;

        background: rgba(0, 0, 0,0);
        cursor: pointer;
        font-family: 'Noto Sans Display Variable';
        
        width: 440px;
        height: 140px;
        
        transition: transform 0.2s, border-color 0.2s;
    }
    .role:nth-child(3n) {
        margin-right: 0;
    }

    /* Начальное состояние псевдоэлемента (скрыт) */
    .role::before {
        content: ""; /* Обязателен для псевдоэлементов */
        position: absolute;
        top: 0;
        left: 0;
        width: 5%; /* Начальная ширина — 0 */
        height: 100%; /* Полная высота кнопки */
        background-color: rgba(0, 31, 184, 0.5);
        z-index: -1; /* Отправляем за текст кнопки */
        transition: width 0.3s ease; /* Анимация ширины */
    }

    .backend::before {
        background-color: rgba(0, 31, 184, 0.3);
    }
    .frontend::before {
        background-color: rgba(255, 77, 0, 0.3);
    }
    .fullstack::before {
        background-color: rgba(255, 39, 39, 0.3);
    }
    .analytics::before {
        background-color: rgba(0, 197, 3, 0.3);
    }
    .fullstack::before {
        background-color: rgba(255, 39, 39, 0.3);
    }
    .fullstack::before {
        background-color: rgba(255, 39, 39, 0.3);
    }
    .fullstack::before {
        background-color: rgba(255, 39, 39, 0.3);
    }
    .fullstack::before {
        background-color: rgba(255, 39, 39, 0.3);
    }
    .fullstack::before {
        background-color: rgba(255, 39, 39, 0.3);
    }

    .role:hover::before {
        width: 100%; /* Заполняем всю кнопку */
    }


    /* Role name styles */
    .role-name {
        font-size: 36px;
        font-weight: medium;
        margin-bottom: 10px;
        padding-left: 19px;
    }

    /* Teams count styles */
    .role-teams {
        font-size: 24px;
        padding-left: 19px;
        padding-bottom: 10px;
    }

    /* More roles link */
    .more-roles {
        margin-top: 20px;
        font-size: 24px;
        cursor: pointer;
        transition: color 0.2s;
        text-decoration: none;
    }
    .more:hover {
        text-decoration: underline;
    }
</style>