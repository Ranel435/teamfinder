<script>

    import { accessToken } from "$lib/stores/authStore";
import "../../app.css";

</script>
<header class="navbar">
    <div class="navbar-logo">
        <a href="/">
            <img class="/navbar-logo-img" src="/teamfinder.svg" alt="Логотип компании">
        </a>
    </div>

    <nav class="navbar-links">
        <ul>
            <li><a class="navbar-links-link" href="/hackatons">Хакатоны</a></li>
            <li><a class="navbar-links-link" href="/teams">Команды</a></li>
            <li><a class="navbar-links-link" href="/forms">Анкеты</a></li>
            <li><a class="navbar-links-link" href="/community">Сообщества</a></li>
            <li><a class="navbar-links-link" href="/about">О проекте</a></li>
            <li><a href="/profile/notifications"><img src="/profile-nav/notifications.svg" alt=""></a></li>
        </ul>
    </nav>
    {#if !accessToken}
    <div class="navbar-login-button">
        <a href="/auth" >Вход</a>
    </div>
    {:else}
    <div class="navbar-login-button">
        <a href="/profile/account">Профиль</a>
    </div>
    {/if}
</header>
<slot />
<footer>
    <div class="footer-nav">
        <ul>
            <li class="footer-nav-item">
                <img src="/teamfinder.svg" alt="">
            </li>
            <li class="footer-nav-item">
                <a href="/help">Поддержка</a>
            </li>
            <li class="footer-nav-item">
                <a href="https://t.me/+WQeYWDOPnvs5yhhY">Сообщество</a>
            </li>
            <li class="footer-nav-item">
                <a href="/privacy_policy">Политика конфиденциальности</a>
            </li>
            <li class="footer-nav-item">
                <a href="/our_team">Наша команда</a>
            </li>
        </ul>
    </div>
    <hr class="footer-line">
    <div class="footer-bottom">
        <div class="rights">
            <p>©2024 Все права защищены</p>
        </div>
        <div class="footer-bottom-links">
            <div class="footer-bottom-links-item">
                <label for="">Сообщество</label>
                <a href="https://t.me/+WQeYWDOPnvs5yhhY">
                    <img class="footer-bottom-links-item-logo" src="/tg-icon.svg" alt="">
                </a>
            </div>
            <div class="footer-bottom-links-item">
                <label for="">Код</label>
                <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">
                    <img class="footer-bottom-links-item-logo" src="/github-icon.svg" alt="">
                </a>
            </div>
            <div class="footer-bottom-links-item">
                <label for="">Дизайн</label>
                <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">
                    <img class="footer-bottom-links-item-logo" src="/n-icon.svg" alt="">
                </a>
            </div>
            <div class="footer-bottom-links-item">
                <label for="">При поддержке</label>
                <a href="https://t.me/itatmisis">
                    <img class="itam-logo" src="/itam.svg" alt="">
                </a>
            </div>
        </div>
    </div>
</footer>


<style>
    
    * {
        font-family: "Manrope";
    }
    ul {
        align-items: center;
    }


    .navbar {
        display: flex;
        justify-content: space-between;
        align-items: center;
        height: 94px;
    }

    .navbar-logo-img {
        max-height: 56px;
        max-width: 232px;
    }

    .navbar-links ul {
        list-style: none;
        display: flex;
        padding: 0;
    }

    .navbar-links li {
        margin: 0 15px;
    }

    .navbar-links-link {
        color: white;
        text-decoration: none;
        font-size: 24px;
    }

    .navbar-links-link:hover {
        text-decoration: underline;
    }

    .navbar-login-button a {
        color: white;
        max-width: 164px;
        text-decoration: none;
        padding: 3px 50px;
        border: 3px solid;
        border-color: #fff;
        border-radius: 50px;
        font-size: 20px;
    }

    .navbar-login-button a:hover {
        transition: all 0.5s ease;
        color: #000;
        background-color: #fff;
    }




    footer {
        display: flex;
        flex-direction: column;
     }

    .footer-nav {
        margin-bottom: 74px;
    }

    .footer-nav ul{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        padding: 0;
        list-style-type: none;
    }

    .footer-nav-item img {
        height: 64px;
        width: 233px;
    }

    .footer-nav-item a {
        color: var(--dark-grey);
        font-size: 24px;
        text-decoration: none;
    }

    .footer-nav-item a:hover {
        text-decoration: underline;
    }

    .footer-line {
        border-top: 2px solid var(--dark_grey); /* Толщина и цвет линии */
        width: 100%; /* Ширина линии */
    }

    .footer-bottom {
        margin-top: 48px;
        margin-bottom: 56px;
        display: flex;
        align-items: center;
        flex-direction: row;
    }

    .footer-bottom p {
        font-family: "Manrope";
        font-size: 24px;
        color: var(--dark-grey);
    }

    .rights {
        width: 50%;
    }
    
    .footer-bottom-links {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 50%;
        padding-right: 128px;
    }

    .footer-bottom-links-item-logo {
        height: 52px;
        width: 52px;
    }

    .itam-logo {
        width: 136px;
        height: 42px;
    }

    .footer-bottom-links-item {
        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
        cursor: pointer;
    }

    .footer-bottom-links-item label {
        font-family: "Manrope";
        font-size: 16px;
        color: #fff;
        margin-bottom: 14px;
    } 

    @media (max-width: 1440px) {
        .navbar {
            padding: 0 0;
            height: 80px;
        }

        .navbar-logo-img {
            height: 42px;
            width: 153px;
        }

        .navbar-links li {
            margin: 0 0;
            margin-right: 40px;
        }

        .navbar-links li:last-child {
            margin: 0 0;
            margin-right: 0px;
        }

        .navbar-links-link {
            font-size: 16px;
        }


        .footer-nav-item a {
            font-size: 18px;
        }

        .navbar-login-button a{
            font-size: 20px;
            padding: 5px 50px;
        }


        .rights p {
            font-size: 18px;
        }

        .footer-bottom-links-item label {
            font-size: 12px;
        }

        .footer-bottom-links-item-logo {
            height: 39px;
            width: 39px;
        }

        .itam-logo {
            width: 102px;
            height: 32px;
        }

    }
</style>
