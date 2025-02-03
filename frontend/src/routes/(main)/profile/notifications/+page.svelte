<script lang="ts">
  // import { rollupVersion } from "vite";
  import "../../../../app.css";
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
  ];

  function deleteNotification(index: number) {
    notifications = notifications.splice(index);
  }
</script>

<section class="notifications">
  {#if notifications.length === 0}
    <div class="no-notifications">
      <h1 class="no-notifications__title m1">Уведомлений пока нет</h1>
      <a class="no-notifications__button gradient-button m5" href="/"
        >На главную</a
      >
    </div>
  {:else}
    <div class="notifications__container">
      <header class="notifications__header">
        <h2 class="notifications__title m1">Все уведомления</h2>
        <div class="notifications__buttons">
          <button class="notifications__button button__read gradient-button m5"
            >Прочитать все</button
          >
          <button
            class="notifications__button button__clear black-white-button m5"
            on:click={() => (notifications = [])}>Очистить все</button
          >
        </div>
      </header>
      <ul class="notifications__list">
        {#each notifications as notification, index}
          <div class="notification">
            <div class="notification__container">
              <div class="notification__info">
                <h3 class="notification__type m5">{notification.type}</h3>
                <div class="notification__info">
                  <span class="notification__hakathon m7"
                    >{notification.hackaton} •
                  </span>
                  <span class="notification__role m7"
                    >{notification.role} •
                  </span>
                  <span class="notification__date m7">{notification.date}</span>
                </div>
                <span class="notification__team m7"
                  >Команда: {notification.team}</span
                >
              </div>
            </div>
          </div>
        {/each}
      </ul>
    </div>
  {/if}
</section>

<style>
  .notifications {
    margin-left: 28px;
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .no-notifications {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    height: 100%;
  }

  .no-notifications__title {
    color: var(--white);
  }

  .no-notifications__button {
    padding: 12px 32px;
    text-decoration: none;
  }

  /* //////////////////////////////////////////////////////////////////////////////////////////////// */

  .notifications__container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    padding: 32px 0 0 0;
  }

  .notifications__header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
  }

  .notifications__title {
    color: var(--white);
  }

  .notifications__buttons {
    display: flex;
    flex-direction: row;
  }

  .notifications__button {
    padding: 5px 36px;
  }

  .button__read {
    margin-right: 32px;
  }

  .notifications__list {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin: 32px;
    /* max-height: 50vh;
    overflow-y: auto;
    padding-right: 10px;
    box-sizing: content-box; */
  }

  /* ///////////////////////////////////////////////////////////////////////// */

  .notification {
    border: 2px solid #fff;
    border-radius: 24px;
    margin-bottom: 16px;
  }

  .notification__container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
  }

  .notification__type {
    color: var(--white);
  }

  .notification__team,
  .notification__hakathon,
  .notification__date,
  .notification__role {
    color: var(--grey);
  }

  .notification:last-child {
    margin-bottom: 0;
  }

  /* .notifications__list::-webkit-scrollbar {
    width: 28px;
  }

  .notifications__list::-webkit-scrollbar-thumb {
    background-color: var(--dark-grey);
    border: 5px solid transparent;
    border-radius: 32px;
    background-clip: padding-box;
  }

  .notifications__list::-webkit-scrollbar-thumb:hover {
    background-color: var(--light-grey);
  }

  .notifications__list::-webkit-scrollbar-track {
    background-color: transparent;
    border-radius: 32px;
    border: 2px solid #fff;
  } */
</style>
