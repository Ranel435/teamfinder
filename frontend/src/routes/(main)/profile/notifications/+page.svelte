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
        <button on:click={() => notifications = []}>Прочитать все</button>
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
              <button class="delete-button" on:click={() => notifications = notifications.splice(index)}>Удалить</button>
            </div>
          </div>
        {/each}
      </ul>
  </div>
    {/if}
</section>



<style>
  .notifications {
    border: 4px solid #fff;
    border-radius: 32px;
    height: 100%;
    margin-left: 100px;
    display: flex;
    flex-direction: column;
    height: 100%;
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
    align-items: center;
    justify-content: flex-start;
    padding: 32px;
  }
  .notifications-header {
    width: 100%;
    display: flex;
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



  .notification {
    border: 2px solid #fff;
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .notification-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    padding: 16px 24px;
  }

  .notification-container h3, p {
    color: #fff;
    font-family: "Manrope";
  }
  .notification:last-child {
    margin-bottom: 0;
  }

  /* Кнопка удаления */
  .delete-button {
    display: none;
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


  button {
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
  
  button:hover {
    color: black;
    background-color: #fff;
  }
  
</style>