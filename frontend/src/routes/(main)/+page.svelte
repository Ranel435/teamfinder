<script lang="ts">
    import "../../app.css";
    import {type Form} from "$lib/stores/data"

    let roles: string[] = ["Backend", "Frontend", "Fullstack", "Analyst", "Product-Manager", "Design", "IOS-Developer", "ML", "DevOps"]
    let array: number[] = [1, 2, 3, 4, 5, 6]
    let count_teams: number = 14;


    let vacancy: Form = {
        id: "14874231-8512489-y3",
        name: "Max",
        surname: "Uskov",
        tgid: "@qquerell",
        role: "Frontend",
        skills: "SvelteKit / TS",
        about: "Senior++ html dev ceo",
        archivements: "Shorthack 44th place",
        status: 1
    }
</script>


<main>
    <section class="main">
        <div class="text">
            <p>Команда найдется для каждого</p>
        </div>
        <div class="main-text">
            <img src="Начни Хакатонить!.svg" alt="">
        </div>
        <div class="main-buttons">
            <a  class="main-button" href="/forms">Смотреть анкеты</a>
            <a class="main-button" href="">Найти команду</a>
            <a class="main-button" href="">Создать анкету</a>
        </div> 
    </section>
    <section class="popular-roles">
        <h2 class="title">Популярные роли</h2>
        <div class="roles-grid">
            {#each roles as role, index}
            <button class="role {role}">
                <div class="role-name">{role}</div>
                <div class="role-teams">{count_teams} команд</div>
              </button>
            {/each}
        </div>
        <div class="more-roles">
            <a class="more" href="/">Еще роли</a>
        </div>
    </section>

    <section class="teams_slots">
        <h2 class="teams_slots-title">Поиск сокомандников</h2>
        <div class="slots-grid">
            {#each array as item}
            <div class="slots-grid-slot">
                <h4>Ищем</h4>
                <h3 class="slot-role">Роль</h3>
                <p class="slot-skills">Стек/навыки · Опыт</p>
                <hr>
                <h3 class="slot-team_name">Название команды</h3>
                <div class="slot-players">
                    <p class="slot-players-player">Игрок 1</p>
                    <p class="slot-players-player">Игрок 2</p>
                    <p class="slot-players-player">Игрок 3</p>
                    <p class="slot-players-player">Игрок 4</p>
                </div>
                <div class="slot-features">
                    <p>1 хак</p>
                    <p>0 побед</p>
                    <p>3м 5д в сборе</p>
                </div>
                <button class="slot-more_about">Подробнее</button>
            </div>
            {/each}
        </div>
        <div class="more-roles">
            <a class="more" href="/">Другие команды</a>
        </div>
    </section>
</main>



<style>
    * {
        color: #fff;
    }

    hr {
        width: 100%;
    }

    img {
        width: 815px;
        height: 258px;
    }

    p {
        font-family: "Igra";
        font-size: 18px;
        line-height: 130%;
        letter-spacing: 5%;
        color: #fff;
    }

    main {
        /* margin: 256px 0; */
        /* background-image: url(".static/PinkEllipse.svg"), url("/static/YellowEllipse.svg"), url("/static/Star.svg"); 
        background-repeat: no-repeat, no-repeat, no-repeat; 
        background-position: center center, center center, center center; 
        background-size: cover; */
        margin: 100px 80px;
    }

    .main {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }

    .main-text {
        text-align: center;
        margin: 10px 0 40px;
    }

    .main-buttons {
        margin-bottom: 128px;
    }
    
    

    /* main buttons */
    .main-button {
        font-family: 'Manrope';
        color: #fff;
        text-decoration: none;

        max-width: 164px;
        border: 3px solid;
        border-color: #fff;
        border-radius: 50px;

        font-size: 18px;
        margin-right: 20px;
        padding: 5px 16px;
    }

    .main-button:hover {
        transition: all 0.7s ease;
        color: #000;
        background-color: #fff;
    }

    .popular-roles {
        display: flex;
        flex-wrap: wrap;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        text-align: center;
        margin-top: 0;
    }

    .title, .teams_slots-title {
        font-family: "Igra";
        font-weight: normal;
        color: #fff;
        margin-bottom: 80px;
        font-size: 56px;
    }

    .roles-grid {
        display: grid;
        grid-template-columns: repeat(3, minmax(auto, 100vw));
    }

    .role {
        position: relative; 
        overflow: hidden; 
        display: flex;
        flex: 1;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        text-align: left;
        border: 4px solid #fff;
        border-radius: 24px;
        padding: 10px;
        margin: 0px 36px 36px 0;
        background: rgba(0, 0, 0,0);
        cursor: pointer;
        font-family: 'Manrope';   
        color: #fff;  
        transition: transform 0.2s, border-color 0.2s;
    }
    .role:nth-child(3n) {
        margin-right: 0;
    }

    .role::before {
        content: ""; 
        position: absolute;
        top: 0;
        left: 0;
        width: 5%; 
        height: 100%;
        background-color: rgba(0, 31, 184, 0.5);
        z-index: -1;
        transition: width 0.3s ease;
    }

    .role:hover::before {
        width: 100%; /* Заполняем всю кнопку */
    }

    .role-name {
        font-weight: medium;
        margin-bottom: 10px;
        padding-left: 19px;
        font-size: 27px;
    }

    .role-teams {
        padding-left: 19px;
        padding-bottom: 10px;
        font-size: 18px;
    }

    .more-roles {
        display: flex;
        justify-content: right;
        margin-top: 20px;
        transition: color 0.2s;
        text-decoration: none;
        width: 100%;
    }

    .more {
        font-family: "Igra";
        font-size: 24px;
        color: #fff;
        text-decoration: none;
    }

    .more:hover {
        text-decoration: underline;
    }

    .Backend::before {
        background-color: rgba(0, 31, 184, 0.3);
    }
    .Frontend::before {
        background-color: rgba(255, 77, 0, 0.3);
    }
    .Fullstack::before {
        background-color: rgba(255, 39, 39, 0.3);
    }
    .Analyst::before {
        background-color: rgba(0, 197, 3, 0.3);
    }
    .ML::before {
        background-color: rgba(0, 90, 12, 0.3);
    }
    .DevOps::before {
        background-color: rgba(153, 0, 255, 0.3);
    }
    .Design::before {
        background-color: rgba(14, 179, 255, 0.3);
    }
    .IOS-Developer::before {
        background-color: rgba(160, 0, 83, 0.3);
    }
    .Product-Manager::before {
        background-color: rgba(255, 225, 0, 0.3);
    }

    .teams_slots {
        display: flex;
        flex-direction: column;
        align-items: center;

        margin-top: 128px;
    }

    .slots-grid {
        display: grid;
        grid-template-columns: repeat(3, minmax(auto, 100vw));
    }

    .slots-grid-slot {
        display: flex;
        flex: 1;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        text-align: left;

        border: 3px solid #fff;
        border-radius: 24px;
        padding: 16px;
        margin: 0px 36px 36px 0;
        font-family: 'Manrope';
        color: #fff;  
    }

    .slots-grid-slot:nth-child(3n) {
        margin-right: 0;
    }


    .slots-grid-slot h4 {
        color: var(--pink);
        font-size: 12px;
        font-weight: 500;
    }

    .slot-role {
        font-size: 24px;
    }

    .slot-skills {
        font-size: 12px;
        color: var(--dark-grey);
    }

    .slot-team_name {
        font-size: 16px;
        font-weight: 500;
        color: var(--pink);
        background-image: var(--gradient); /* Укажите градиент */
        -webkit-background-clip: text; /* Обрезка фона по тексту (для WebKit-браузеров) */
        background-clip: text; /* Обрезка фона по тексту (для других браузеров) */
        color: transparent; /* Делаем цвет текста прозрачным */
    }

    .slot-players {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        margin: 8px 0 8px;
    }

    .slot-players-player {
        font-family: 'Manrope';
        font-weight: 300;
        font-size: 14px;
        margin: 0px 8px 8px 0;
    }

    .slot-features {
        display: flex;
        flex-direction: row;
        justify-content: left;
        width: 100%;
        align-items: center;
        margin-bottom: 21px;
    }

    .slot-features p {
        font-size: 10px;
        background-color: #9C004A;
        padding: 5px 13px;
        border-radius: 12px;
        margin-right: 10px;
    }

    .slot-more_about {
        font-size: 14px;
        padding: 8px 0;
        width: 100%;
        border: 2px solid;
        border-color: #fff;
        border-radius: 50px;
        background-color: rgba(0, 0, 0, 1);
        cursor: pointer;
    }

    .slot-more_about:hover {
        transition: all 0.7s ease;
        color: #000;
        background-color: #fff;
    }
</style>