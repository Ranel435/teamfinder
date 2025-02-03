<script lang="ts">
  import "../../../../app.css"
  import { type Profile} from "$lib/stores/data";
  let isEditing = false; // Статус редактирования


  let profile: Profile = {
    profileId: "",
    name: "",
    surname: "",
    tgId: "",
    email: "",
    about: "",
    portfolio: "",
    university: "",
    academicGroup: "",
    extraEducation: "",
    password: ""
  };
  // Функция переключения режима
  const toggleEdit = () => {
    isEditing = !isEditing;
  };

  // Функция сохранения данных
  const saveData = () => {
    console.log("Сохраненные данные:", profile);
    toggleEdit(); // Завершаем редактирование
  };
</script>


<section class="profile">
  <div class="profile__inf inf">
    <div class="inf__container">
      <img class="int__img" style="width: 300px; height: 300px" src="/profile.png" alt="">
      <div class="inf__header">
        <h4 class="inf__title">Мой профиль</h4>
      </div>
      <input type="text" placeholder="Фамилия" bind:value={profile.surname}
      disabled = {!isEditing}>
      <input type="text" placeholder="Имя" bind:value={profile.name}
      disabled = {!isEditing}>
      <input type="text" pattern="^@([A-Za-z0-9_]+)$" placeholder="@tgid" bind:value={profile.tgId}
      disabled = {!isEditing}>
      <input type="email" placeholder="почта@mail.ru" bind:value={profile.email}
      disabled = {!isEditing}>
      <textarea maxlength=300 placeholder="Обо мне" bind:value={profile.about}
      disabled = {!isEditing}></textarea>
      {#if isEditing}
        <button class="save gradient-button buttonM" on:click={saveData}>Сохранить</button>
      {:else}
        <button class="edit black-white-button buttonM" on:click={toggleEdit}>Редактировать</button>
      {/if}
    </div>
  </div>

  <div class="profile__other-inf">
    <div class="profile__skills skills">
      <div class="skills__container">
        <h2 class="skills__header gradient-text m3">Навыки</h2>
        <div class="form-item">
          <label class="m6" for="">Владение софтом</label>
          <input class="m8" type="text" placeholder="Владение софтом" disabled = {!isEditing}>
        </div>
        <div class="form-item">
          <label class="m6" for="">Полный стэк</label>
          <input class="m8" type="text" placeholder="Полный стек" disabled = {!isEditing}>
        </div>
        <div class="form-item">
          <label class="m6" for="">Качества</label>
          <input class="m8" type="text" placeholder="Качества" disabled = {!isEditing}>
        </div>
      </div>
    </div>
    
    <div class="profile__edu edu">
      <div class="edu__container">
        <h2 class="edu__header gradient-text m3">Образование</h2>
        <div class="form-item">
          <label class="m6" for="">Учебное заведение</label>
          <input class="m8" type="text" placeholder="Учебное заведение" disabled = {!isEditing} bind:value={profile.university}>
        </div>
        <div class="form-item">
          <label class="m6" for="">Учебная группа</label>
          <input class="m8" type="text" placeholder="Учебная группа" disabled = {!isEditing} bind:value={profile.academicGroup}>
        </div>
        <div class="toggle-switch">
          <label class="toggle-label m6" for="educationToggle">Уже имею образование</label>
          <input type="checkbox" id="educationToggle" class="toggle-checkbox" disabled = {!isEditing} >
          <label class="toggle-slider" for="educationToggle"></label>
        </div>
        <div class="form-item">
          <label class="m6" for="">Дополнительное образование</label>
          <input class="m8" type="text" placeholder="Дополнительное образование" disabled = {!isEditing} bind:value={profile.extraEducation}>
        </div>
      </div>
    </div>
  </div>  
</section>


<style>
  .profile__other-inf label {
    color: #fff;
    text-align: left; /* Выравнивание текста влево */
    display: block; /* Чтобы каждая метка занимала отдельную строку */
    margin-bottom: 4px;
  }
  textarea {
    font-size: 16px;
    line-height: 1.5;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    resize: none;
    overflow: hidden;
    min-height: 50px; /* Минимальная высота */
    max-height: 300px; /* Максимальная высота */
  }
  input {
    width: 100%;
    box-sizing: border-box
  }

  .profile {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      padding-left: 40px;
  }

  .profile__inf {
    border: 4px solid #fff;
    border-radius: 32px;
  }

  .profile__inf {
    width: 48%;
    display: flex;
    flex-direction: column; /* Выравнивание содержимого */
  }

  .profile__inf input, textarea {
    width: 100%; /* Растягиваем на всю ширину контейнера */
    padding: 0 0; /* Отступ сверху и снизу */
    background: none; /* Убираем фон */
    border: none; /* Убираем рамки */
    color: #fff; /* Белый текст */
    font-family: "Manrope-medium";
    font-size: 18px; /* Размер текста */
    margin-bottom: 16px;
    border-bottom: 1px solid #fff; /* Тонкая нижняя линия */
  }

  .inf__container {
    padding: 14px 36px;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: left;
  }

  .inf__header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: start;
    width: 100%;
    margin-bottom: 28px;
  }

  .inf__title {
    font-family: "Manrope-medium";
    font-size: 20px;
    color: var(--white);
  }

  .save {
    padding: 12px 24px;
  }

  .edit {
    padding: 9px 24px;
  }
  
  .profile__other-inf {
    width: 48%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    min-height: auto;
  }

  .profile__other-inf input {
    padding: 8px 16px;
    background-color: var(--light-grey);
    border-radius: 27px;
    margin-bottom: 4px;
    outline: none;
    font-family: "Manrope-medium";
    font-size: 12px;
    letter-spacing: -3%;
    border: 0;
  }
   
  .profile__skills, .profile__edu {
    border: 4px solid #fff;
    border-radius: 32px;
    height: auto;
  }

  .skills__container, .edu__container {
    padding: 6px 36px 0 36px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: left;
  }

  .form-item {
    width: 100%;
    text-align: left;
    margin: 0 0 16px 0;
  }

  input:disabled, textarea:disabled {
    color: var(--light-grey);
  }

.toggle-switch {
  display: flex;
  align-items: center;
  text-align: left;
  width: 100%;
  gap: 10px;
  margin-bottom: 4px;

}

.toggle-label {
  color: #fff;
  cursor: pointer;
  text-align: left;
}

.toggle-checkbox {
  display: none;
}

.toggle-slider {
  width: 40px;
  height: 20px;
  background-color: #ccc; 
  border-radius: 20px;
  position: relative;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.toggle-slider::before {
  content: "";
  width: 16px;
  height: 16px;
  background-color: #fff;
  border-radius: 50%;
  position: absolute;
  top: 2px;
  left: 2px;
  transition: transform 0.3s ease;
}

.toggle-checkbox:checked + .toggle-slider {
  background-image: var(--gradient); 
}

.toggle-checkbox:checked + .toggle-slider::before {
  transform: translateX(20px);
}
</style>