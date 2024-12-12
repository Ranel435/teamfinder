<script lang="ts">
  // import { rollupVersion } from "vite";
  import "../../../../app.css"
  interface notification {
    type: string;
    hackaton: string;
    role: string;
    date: string;
    team: string;
    profile: string;
  }

  let notifications: notification[] = [
    {
      type: "Приглашение в команду",
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      team: "KoVCheG",
      profile: ""
    },
    {
      type: "Заявка на вступление",
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      team: "KoVCheG",
      profile: "qquerell"
    },
    {
      type: "Приглашение в команду",
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      team: "KoVCheG",
      profile: ""
    },
    {
      type: "Приглашение в команду",
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      team: "KoVCheG",
      profile: ""
    },
    {
      type: "Заявка на вступление",
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      team: "KoVCheG",
      profile: "qquerell"
    },
    {
      type: "Приглашение в команду",
      hackaton: "Ответ на анкету Т1 Москва",
      role: "UI/UX",
      date: "18.12.2024",
      team: "KoVCheG",
      profile: ""
    }
  ]

  function deleteNotification(index: number) {
    notifications = notifications.splice(index);
  }



</script>


<section class="notifications">
    {#if notifications.length === 0} 
    <div class="no-notifications">
      <p class="no-notifications-text">Уведомлений пока нет</p>
    </div>
    {:else}
    <div class="notifications-container">
      <div class="notifications-header">
        <h2>Все уведомления</h2>
        <div class="notifications-header-buttons">
          <button class="gradient-button">Прочитать все</button>
          <button class="button" on:click={() => notifications = []}>Очистить все</button>
        </div>
      </div>
      <ul class="notifications-list">
        {#each notifications as notification, index}
          <div class="notification">
            <div class="notification-container">
              <div class="notification-info">
                <h3>{notification.type}</h3>
                <p>{notification.hackaton} * {notification.role} * {notification.date}</p>
                <p>{notification.team} * {notification.profile}</p>
              </div>
              <button class="delete-button button" on:click={() => notifications = notifications.splice(index)}>Удалить</button>
            </div>
          </div>
        {/each}
      </ul>
  </div>
    {/if}
</section>



<style>

    button {
      font-family: "Manrope";
      transition: all 0.5s ease;
    }

  .notifications {
    margin-left: 50px;
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .no-notifications {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
  }

  .no-notifications-text {
    font-family: "Manrope";
    font-size: 48px;
    color: #fff;
  }

  .notifications-container {
    display: flex;
    flex-direction: column;
    padding: 32px;
  }
  .notifications-header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 40px;
  }

  .notifications-list {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-height: 50vh; /* Фиксированная высота для прокрутки */
    overflow-y: auto; /* Добавление вертикального скроллбара */
    padding-right: 10px; /* Добавляем отступ для скроллбара */
    box-sizing: content-box; /* Учитываем отступ, не влияя на ширину */
  }

  .notifications-header h2{
    font-family: "Manrope";
    font-size: 48px;
    font-weight: 400;
    color: #fff;
  }

  .notifications-header-buttons {
    display: flex;
    justify-content: space-between;
  }

  .button {
    border: 3px solid #fff;
    padding: 10px 16px;
  }

  .notifications-header-buttons button {
    font-size: 18px;
    color: #fff;
    border-radius: 50px;
    background-color: #000;
  }


  .gradient-button {
    padding: 13px 19px;
    border: none;
    margin-right: 10px;
    transition: all 0.5s ease;
  }

  .gradient-button:hover {
    background:linear-gradient(90deg, #ED666E, #FBA31C);
  }

  .notification {
    border: 2px solid #fff;
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .notification-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
  }

  .notification-container h3 {
    font-size: 24px;
    color: #fff;
    font-family: "Manrope";
  }

  .notification-container p {
    font-size: 16px;
    color: #fff;
    font-family: "Manrope";
  }
  .notification:last-child {
    margin-bottom: 0;
  }

  
  
  .button:hover {
    color: black;
    background-color: #fff;
  }
  /* Кнопка удаления */
  .delete-button {
    height: 50px;
    display: none;
    font-size: 18px;
    color: #fff;
    border: 3px solid #fff;
    border-radius: 50px;
    background-color: transparent;
    cursor: pointer;
    margin-right: 16px;
  }



  .notification:hover {
    /* border: 2px var(--gradient) #fff; */
    border-image: var(--pink);
  }

  /* Показываем кнопку при наведении */
  .notification:hover .delete-button {
    display: block;
  }



  /* Стиль для скроллбара */
  .notifications-list::-webkit-scrollbar {
    width: 26px; /* Ширина скроллбара */
  }

  .notifications-list::-webkit-scrollbar-thumb {
    background-color: var(--dark-grey); /* Цвет ползунка */
    border: 2px solid #2a2a2a;
    
    border-radius: 32px; /* Скругление */
  }

  .notifications-list::-webkit-scrollbar-thumb:hover {
    background-color: var(--light-grey); /* Цвет при наведении */
  }
</style>