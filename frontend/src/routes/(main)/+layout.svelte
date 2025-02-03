<script>
    import { accessToken } from "$lib/stores/authStore";
    import "../../app.css";
</script>
<header class="header-nav">
    <div class="header-nav__logo">
        <a href="/">
            <img class="header-nav__img" src="/teamfinder.svg" alt="teamfinder">
        </a>
    </div>

    <nav class="header-nav__links">
        <ul class="header-nav__list" role="navigation">
            <li><a class="header-nav__link buttonS" href="/">Хакатоны</a></li>
            <li><a class="header-nav__link buttonS" href="/teams">Команды</a></li>
            <li><a class="header-nav__link buttonS" href="/forms">Анкеты</a></li>
            <li><a class="header-nav__link buttonS" href="/community">Сообщества</a></li>
            <li><a class="header-nav__link buttonS" href="/about">О проекте</a></li>
            <li><a href="/profile/notifications"><img src="/profile-nav/notifications.svg" alt="Уведомления"></a></li>
        </ul>
    </nav>
    
    <div class="header-nav__account account">
        <a class="account__link buttonM" href="/auth" >Вход</a>
        <!-- <a class="account__link buttonM" href="/profile/account">Профиль</a> -->
    </div>
</header>
<slot />
<footer class="footer">
    <section class="footer__nav footer-nav">
        <ul class="footer-nav__list" role="navigation">
            <li class="footer-nav__item">
                <img class="footer-nav__img" src="/teamfinder.svg" alt="">
            </li>
            <li class="footer-nav__item">
                <a class="footer-nav__link buttonM" href="/help">Поддержка</a>
            </li>
            <li class="footer-nav__item">
                <a class="footer-nav__link buttonM" href="https://t.me/+WQeYWDOPnvs5yhhY" rel="noopener noreferrer">Сообщество</a>
            </li>
            <li class="footer-nav__item">
                <a class="footer-nav__link buttonM" href="/privacy_policy">Политика конфиденциальности</a>
            </li>
            <li class="footer-nav__item">
                <a class="footer-nav__link buttonM" href="/our_team">Наша команда</a>
            </li>
        </ul>
    </section>
    <hr class="footer__line">
    <section class="footer__bottom bottom">
        <div class="bottom__rights">
            <p>©2024 Все права защищены</p>
        </div>
        <div class="bottom__list">
            <div class="bottom__item">
                <label for="">Сообщество</label>
                <a class="bottom__link" href="https://t.me/+WQeYWDOPnvs5yhhY">
                    <img class="bottom__img" src="/tg-icon.svg" alt="">
                </a>
            </div>
            <!-- <div class="bottom__item">
                <label for="">Код</label>
                <a class="bottom__link" href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">
                    <img class="bottom__img" src="/github-icon.svg" alt="">
                </a>
            </div>
            <div class="bottom__item">
                <label for="">Дизайн</label>
                <a class="bottom__link" href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">
                    <img class="bottom__img" src="/n-icon.svg" alt="">
                </a>
            </div> -->
            <div class="bottom__item">
                <label for="">При поддержке</label>
                <a class="bottom__link" href="https://t.me/itatmisis">
                    <img class="itam-logo" src="/itam.svg" alt="">
                </a>
            </div>
        </div>
    </section>
</footer>


<style>
    
    ul {
        align-items: center;
    }
    .header-nav {
        padding: 0 128px;
    }
    .footer {
        padding: 0 96px;
        min-height: 100%;
    }
/* ///////////////////////////////////////////////////////////////////////////////////////////////////// */

    .header-nav {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding-top: 33px;
    }

    .header-nav__img {
        max-height: 56px;
        max-width: 232px;
    }

    .header-nav__list {
        list-style: none;
        display: flex;
    }

    .header-nav__list li {
        margin: 0 15px;
    }

    .header-nav__link {
        color: white;
        text-decoration: none;
    }

    .header-nav__link:hover {
        text-decoration: underline;
    }

    .account__link {
        color: white;
        max-width: 164px;
        text-decoration: none;
        padding: 3px 50px;
        border: 3px solid;
        border-color: #fff;
        border-radius: 50px;
    }

    .account__link:hover {
        transition: all 0.5s ease;
        color: #000;
        background-color: #fff;
    }

/* //////////////////////////////////////////////////////////////////////////////////// */


    footer {
        display: flex;
        flex-direction: column;
     }

    .footer__nav {
        margin-bottom: 74px;
    }

    .footer-nav__list {
        display: flex;
        justify-content: space-between;
        align-items: center;
        list-style-type: none;
    }

    .footer-nav__img {
        height: 64px;
        width: 233px;
    }

    .footer-nav__link {
        color: var(--dark-grey);
        font-size: 18px;
        font-weight: 500;
        text-decoration: none;
    }

    .footer-nav__link:hover {
        text-decoration: underline;
    }

    .footer__line {
        border-top: 2px solid var(--dark_grey); /* Толщина и цвет линии */
        width: 100%; /* Ширина линии */
    }

    .footer__bottom {
        margin-top: 48px;
        margin-bottom: 56px;
        display: flex;
        align-items: center;
        flex-direction: row;
    }

    .bottom__rights {
        width: 50%;
    }

    .bottom__rights p {
        font-family: "Manrope-semibold";
        font-size: 18px;
        font-weight: 600;
        color: var(--dark-grey);
    }
    
    .bottom__list {
        display: flex;
        flex-direction: row;
        justify-content: right;
        width: 50%;
        padding-right: 128px;
    }

    .bottom__img {
        height: 39px;
        width: 39px;
    }

    .itam-logo {
        width: 136px;
        height: 42px;
    }

    .bottom__item {
        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
        cursor: pointer;
        margin-right: 96px;
    }

    .bottom__item:last-child {
        margin-right: 0;
    }

    .bottom__item label {
        font-family: "Manrope-semibold";
        font-size: 12px;
        color: #fff;
        margin-bottom: 14px;
    }
</style>
