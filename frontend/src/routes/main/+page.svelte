<script lang="ts">
    import "../../app.css";
    let token: string = "";
    let auth: string = "/auth";

    let free_agents: string = "/free_agents";
    let find_team: string = "/find_team";
    let create_form: string = "/create-form";

    let count_teams: number = 14;

    if (!token) {
        free_agents = "/auth";
        find_team = "/auth";
        create_form = "/auth";
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
            <a  class="main-button" href={free_agents}>Свободные агенты</a>
            <a class="main-button" href={find_team}>Найти команду</a>
            <a class="main-button" href={create_form}>Создать анкету</a>
        </div> 
    </section>
    <section class="popular-roles">
        <h2 class="title">Популярные роли</h2>
        <div class="roles-grid">
          <button class="role backend">
            <div class="role-name">Backend</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role frontend">
            <div class="role-name">Frontend</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role fullstack">
            <div class="role-name">Fullstack dev</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role analytics">
            <div class="role-name">Аналитик данных</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role pm">
            <div class="role-name">Product manager</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role ui">
            <div class="role-name">UX / UI дизайнер</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role mobile">
            <div class="role-name">IOS / Android dev</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role ml">
            <div class="role-name">ML engineer</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
          <button class="role devops">
            <div class="role-name">DevOps инженер</div>
            <div class="role-teams">{count_teams} команд</div>
          </button>
        </div>
        <div class="more-roles">
            <a class="more" href="/">Еще роли</a>
        </div>
    </section>
</main>



<style>
    /* main {
        background-image: url("teamfinder/frontend/static/PinkEllipse.svg"), url("/static/YellowEllipse.svg"), url("/static/Star.svg"); 
        background-repeat: no-repeat, no-repeat, no-repeat; 
        background-position: center center, center center, center center; 
        background-size: cover;
    } */
    @media (max-width: 1440px) {
        img {
            width: 815px;
            height: 258px;
        }

        .main {
            margin-top: 0;
            margin: 128px 0;
        }

        .popular-roles {
            margin-top: 0;
        }
    }

    .main {
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
    
    p {
        font-family: "Igra";
        font-size: 24px;
        line-height: 130%;
        letter-spacing: 5%;
        color: #fff;
    }

    /* main buttons */
    .main-button {
        font-family: 'Manrope';
        font-size: 24px;
        color: #fff;
        text-decoration: none;

        max-width: 164px;
        border: 3px solid;
        border-color: #fff;
        border-radius: 50px;

        margin: 0 18px;
        padding: 8px 50px;
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
        margin-top: 198px;
    }

    .title {
        font-family: "Igra";
        font-size: 48px;
        font-weight: normal;
        color: #fff;
        margin-bottom: 80px;
        padding: 0;
    }

    /* grid container */
    .roles-grid {
        display: grid;
        grid-template-columns: repeat(3, minmax(auto, 440px));
    }

    /* role card */
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
        font-size: 36px;
        font-weight: medium;
        margin-bottom: 10px;
        padding-left: 19px;
    }

    .role-teams {
        font-size: 24px;
        padding-left: 19px;
        padding-bottom: 10px;
    }

    /* more roles link */
    .more-roles {
        display: flex;
        justify-content: right;
        margin-top: 20px;
        cursor: pointer;
        transition: color 0.2s;
        text-decoration: none;
        width: 80%;
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

    /* colors for roles */
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
    .ml::before {
        background-color: rgba(0, 90, 12, 0.3);
    }
    .devops::before {
        background-color: rgba(153, 0, 255, 0.3);
    }
    .ui::before {
        background-color: rgba(14, 179, 255, 0.3);
    }
    .mobile::before {
        background-color: rgba(160, 0, 83, 0.3);
    }
    .pm::before {
        background-color: rgba(255, 225, 0, 0.3);
    }




</style>