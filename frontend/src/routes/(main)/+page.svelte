 <script lang="ts">
    import "../../app.css";
    import {type Form} from "$lib/stores/data";

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
        <div class="main__content">
            <p class="main__text">Команда найдется для каждого</p>
            <img class="main__image" src="Начни Хакатонить!.svg" alt="Начни хакатонить!">
        </div>
        <div class="main__buttons">
            <a class="main__button" href="/forms">Смотреть анкеты</a>
            <a class="main__button" href="/teams">Найти команду</a>
            <a class="main__button" href="/create_form">Создать анкету</a>
        </div>
    </section>
    <section class="popular-roles">
        <h2 class="popular-roles__title">Популярные роли</h2>
        <div class="popular-roles__roles">
            {#each roles as role, index}
            <button class="popular-roles__role role {role}">
                <p class="role__name">{role}</p>
                <p class="role__teams">{count_teams} команд</p>
            </button>
            {/each}
        </div>
        <div class="more">
            <a class="more__link" href="/">Еще роли</a>
        </div>
    </section>

    <section class="teams">
        <h2 class="teams__title">Поиск сокомандников</h2>
        <div class="teams__slots">
            {#each array as item}
            <div class="teams__slot slot">
                <h3 class="slot__title">Ищем</h3>
                <h4 class="slot__role">Роль</h4>
                <p class="slot__skills">Стек/навыки · Опыт</p>
                <hr>
                <h3 class="slot__team-name gradient-text">Название команды</h3>
                <div class="slot__players">
                    <p class="slot__player">Игрок 1</p>
                    <p class="slot__player">Игрок 2</p>
                    <p class="slot__player">Игрок 3</p>
                    <p class="slot__player">Игрок 4</p>
                </div>
                <div class="slot__features">
                    <p class="slot__feature">1 хак</p>
                    <p class="slot__feature">0 побед</p>
                    <p class="slot__feature">3м 5д в сборе</p>
                </div>
                <button class="slot__more-about">Подробнее</button>
            </div>
            {/each}
        </div>
        <div class="more">
            <a class="more__link" href="/">Другие команды</a>
        </div>
    </section>
</main>



<style>
    p, h2, h4 {
        color: #fff;
    }

    hr {
        width: 100%;
    }

    main {
        margin: 100px 80px;
    }


/* /////////////////////////////////////////////////////////////////////     */
    
    .main {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }

    .main__content {
        text-align: center;
        margin: 10px 0 40px;
    }

    .main__text {
        font-family: "Igra";
        color: #fff;
    }

    .main__image {
        width: 815px;
        height: 258px;
    }

    .main__buttons {
        margin-bottom: 128px;
    }

    .main__button {
        font-family: 'Manrope';
        color: #fff;
        text-decoration: none;

        border: 3px solid;
        border-color: #fff;
        border-radius: 50px;

        font-size: 18px;
        margin-right: 20px;
        padding: 5px 16px;
    }

    .main__button:hover {
        transition: all 0.7s ease;
        color: #000;
        background-color: #fff;
    }

/* ////////////////////////////////////////////////////////////////////////// */

    .popular-roles {
        display: flex;
        flex-wrap: wrap;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        text-align: center;
        margin-top: 0;
    }

    .popular-roles__title, .teams__title {
        font-family: "Igra";
        font-weight: normal;
        color: #fff;
        margin-bottom: 80px;
        font-size: 56px;
    }

    .popular-roles__roles {
        display: grid;
        grid-template-columns: repeat(3, minmax(auto, 100vw));
    }

    .popular-roles__role {
        position: relative; 
        overflow: hidden; 
        display: flex;
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
        transition: transform 0.2s, border-color 0.2s;
    }

    .popular-roles__role:nth-child(3n) {
        margin-right: 0;
    }

    .popular-roles__role::before {
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

    .popular-roles__role:hover::before {
        width: 100%;
    }

    .role__name {
        margin-bottom: 10px;
        padding-left: 19px;
        font-size: 27px;
    }

    .role__teams {
        padding-left: 19px;
        padding-bottom: 10px;
        font-size: 18px;
    }

    .more {
        display: flex;
        justify-content: right;
        margin-top: 20px;
        transition: color 0.2s;
        text-decoration: none;
        width: 100%;
    }

    .more__link {
        font-family: "Igra";
        font-size: 24px;
        color: #fff;
        text-decoration: none;
    }

    .more__link:hover {
        text-decoration: underline;
    }

    .Backend::before {background-color: rgba(0, 31, 184, 0.3);}
    .Frontend::before {background-color: rgba(255, 77, 0, 0.3);}
    .Fullstack::before {background-color: rgba(255, 39, 39, 0.3);}
    .Analyst::before {background-color: rgba(0, 197, 3, 0.3);}
    .ML::before {background-color: rgba(0, 90, 12, 0.3);}
    .DevOps::before {background-color: rgba(153, 0, 255, 0.3);}
    .Design::before {background-color: rgba(14, 179, 255, 0.3);}
    .IOS-Developer::before {background-color: rgba(160, 0, 83, 0.3);}
    .Product-Manager::before {background-color: rgba(255, 225, 0, 0.3);}

/* /////////////////////////////////////////////////////////////////////// */

    .teams {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 128px;
    }

    .teams__slots {
        display: grid;
        grid-template-columns: repeat(3, minmax(auto, 100vw));
    }

    .teams__slot {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        text-align: left;

        border: 3px solid #fff;
        border-radius: 24px;
        padding: 16px;
        margin: 0px 36px 36px 0;
        font-family: 'Manrope';
    }

    .teams__slot:nth-child(3n) {
        margin-right: 0;
    }

    .slot__title {
        color: var(--pink);
        font-size: 12px;
        font-weight: 500;
    }

    .slot__role {
        font-size: 24px;
    }

    .slot__skills {
        font-size: 12px;
        color: var(--dark-grey);
    }

    .slot__team-name {
        font-size: 16px;
        font-weight: 500;
    }

    .slot__players {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        margin: 8px 0 8px;
    }

    .slot__player {
        font-size: 14px;
        margin: 0px 8px 8px 0;
    }

    .slot__features {
        display: flex;
        flex-direction: row;
        justify-content: left;
        width: 100%;
        align-items: center;
        margin-bottom: 21px;
    }

    .slot__feature {
        font-size: 10px;
        background-color: var(--burgundy);
        padding: 5px 13px;
        border-radius: 12px;
        margin-right: 10px;
    }

    .slot__more-about {
        font-size: 14px;
        color: #fff;
        padding: 8px 0;
        width: 100%;
        border: 2px solid;
        border-color: #fff;
        border-radius: 50px;
        background-color: rgba(0, 0, 0, 1);
        cursor: pointer;
    }

    .slot__more-about:hover {
        transition: all 0.7s ease;
        color: #000;
        background-color: #fff;
    }

</style>