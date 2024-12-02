<script>

    import "../app.css";

</script>
<header class="navbar">
    <div class="logo">
        <a href="/">
            <img class="logo-img" src="teamfinder.svg" alt="Логотип компании">
        </a>
    </div>

    <nav class="nav-links">
        <ul>
            <li><a href="hakatons">Хакатоны</a></li>
            <li><a href="teams">Команды</a></li>
            <li><a href="profiles">Анкеты</a></li>
            <li><a href="community">Сообщества</a></li>
            <li><a href="about">О проекте</a></li>
        </ul>
    </nav>

    <div class="login-button">
        <a href="/auth" class="login-button">Вход</a>
    </div>
</header>
<slot />
<footer>
    <div class="footer-nav">
        <ul>
            <li class="footer-nav-item">
                <img src="teamfinder.svg" alt="">
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
                <a href="">
                    <img src="tg-icon.svg" alt="">
                </a>
            </div>
            <div class="footer-bottom-links-item">
                <label for="">Код</label>
                <a href="">
                    <img src="github-icon.svg" alt="">
                </a>
            </div>
            <div class="footer-bottom-links-item">
                <label for="">Дизайн</label>
                <a href="">
                    <img src="n-icon.svg" alt="">
                </a>
            </div>
            <div class="footer-bottom-links-item">
                <label for="">При поддержке</label>
                <a href="https://t.me/itatmisis">
                    <img style="height: 42px; width: 136px" class="itam-logo" src="itam.svg" alt="">
                </a>
            </div>
        </div>
    </div>
</footer>


<style>
    .navbar {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 18px 0;
    }

    a, button {
        font-family: 'Manrope';
    }

    .logo-img {
        max-height: 56px;
        max-width: 232px;
    }

    ul {
        align-items: center;
    }

    /* navigation */
    .nav-links ul {
        list-style: none;
        display: flex;
        padding: 0;
    }

    .nav-links li {
        margin: 0 15px;
    }

    .nav-links a {
        color: white;
        text-decoration: none;
        font-size: 24px;
    }

    .nav-links a:hover {
        text-decoration: underline;
    }

    /* login */
    .login-button a {
        color: white;
        max-width: 164px;
        text-decoration: none;
        padding: 3px 50px;
        border: 3px solid;
        border-color: #fff;
        border-radius: 50px;
        font-size: 24px;
    }
    .login-button a:hover {
        transition: all 0.5s ease;
        color: #000;
        background-color: #fff;
    }




    footer {
        margin-top: 128px;
        display: flex;
        flex-direction: column;
        /* margin-top: 256px; */
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
        /* justify-content: space-between; */
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

    .footer-bottom-links img {
        height: 52px;
        width: 52px;
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

</style>
