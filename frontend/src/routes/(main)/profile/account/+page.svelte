<script lang="ts">
  import "../../../../app.css"
  import { profile } from "$lib/stores/data";
  let isEditing = false; // Статус редактирования


  
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
  <div class="profile-inf">
    <div class="profile-inf-container">
      <img style="width: 300px; height: 300px" src="/profile.png" alt="">
      <div class="main_text">
        <p>Мой профиль</p>
        <div class="main_text-other">
          <p>Роль</p>
          <p>Стек</p>
        </div>
      </div>
      <input type="text" placeholder="Фамилия" bind:value={profile.surname}
      disabled = {!isEditing}>
      <input type="text" placeholder="Имя" bind:value={profile.name}
      disabled = {!isEditing}>
      <input type="text" pattern="^@([A-Za-z0-9_]+)$" placeholder="@tgid" bind:value={profile.tgid}
      disabled = {!isEditing}>
      <input type="email" placeholder="почта@mail.ru" bind:value={profile.email}
      disabled = {!isEditing}>
      <textarea rows="2" placeholder="Обо мне" bind:value={profile.about}
      disabled = {!isEditing}></textarea>
      {#if isEditing}
        <button class="save" on:click={saveData}>Сохранить</button>
      {:else}
        <button class="edit" on:click={toggleEdit}>Редактировать</button>
      {/if}
    </div>
  </div>
  <div class="profile-other_inf">
    <div class="profile-skills">
      <div class="profile-skills-container">
        <h2 class="header">Навыки</h2>
        <div class="form-item">
          <label for="">Владение софтом</label>
          <input type="text" placeholder="Владение софтом" disabled = {!isEditing} bind:value={$profile.software}>
        </div>
        <div class="form-item">
          <label for="">Полный стэк</label>
          <input type="text" placeholder="Полный стек" disabled = {!isEditing} bind:value={$profile.skills}>
        </div>
        <div class="form-item">
          <label for="">Качества</label>
          <input type="text" placeholder="Качества" disabled = {!isEditing} bind:value={$profile.qualities}>
        </div>
      </div>
    </div>

    <div class="profile-edu">
      <div class="profile-edu-container">
        <h2 class="profile-edu-header">Образование</h2>
        <div class="form-item">
          <label for="">Учебное заведение</label>
          <input type="text" placeholder="Учебное заведение" disabled = {!isEditing} bind:value={$profile.edu}>
        </div>
        <div class="form-item">
          <label for="">Учебная группа</label>
          <input type="text" placeholder="Учебная группа" disabled = {!isEditing} bind:value={$profile.edu_group}>
        </div>
        <div class="toggle-switch">
          <label class="toggle-label" for="educationToggle">Уже имею образование</label>
          <input type="checkbox" id="educationToggle" class="toggle-checkbox" disabled = {!isEditing} >
          <label class="toggle-slider" for="educationToggle"></label>
        </div>
        <div class="form-item">
          <label for="">Дополнительное образование</label>
          <input type="text" placeholder="Дополнительное образование" disabled = {!isEditing} bind:value={$profile.extra_edu}>
        </div>
      </div>
    </div>
  </div>  
</section>


<style>
  h2 {
    color: var(--pink);
    font-family: "Manrope";
    margin-bottom: 4px;
  }
  label {
    font-family: "Manrope";
    font-size: 18px;
    color: #fff;
    text-align: left; /* Выравнивание текста влево */
    display: block; /* Чтобы каждая метка занимала отдельную строку */
    margin-bottom: 4px;
  }
  textarea {
    resize: none; /* Убираем возможность изменять размер */
  }
  input {
    width: 100%;
    box-sizing: border-box
  }
  p {
    color: #fff;
  }

  .profile {
      display: flex;
      flex-direction: row;
      align-items: stretch; /* Растягивает дочерние элементы по высоте */
      justify-content: space-between;
      padding-left: 40px;
    }


  .profile-inf, .profile-skills, .profile-edu {
    border: 4px solid #fff;
    border-radius: 32px;
    
  }

  .profile-other_inf {
    justify-content: space-between;
  }

  .profile-inf-container, .profile-skills-container, .profile-edu-container {
    padding: 14px 36px;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: left;
  }



  .profile-inf, .profile-other_inf {
    width: 48%;
    display: flex;
    flex-direction: column; /* Выравнивание содержимого */
    min-height: auto; /* Автоматическая высота, если фиксированная не требуется */
  }

  .form-item {
    width: 100%;
    text-align: left;
  }

  .profile-other_inf input {
    padding: 8px 16px;
    background-color: var(--light-grey);
    border-radius: 27px;
    font-family: "Manrope";
    font-size: 16px;
    margin-bottom: 4px;
  }

  .profile-inf input, textarea {
    width: 100%; /* Растягиваем на всю ширину контейнера */
    padding: 0 0; /* Отступ сверху и снизу */
    background: none; /* Убираем фон */
    border: none; /* Убираем рамки */
    border-bottom: 1px solid #fff; /* Тонкая нижняя линия */
    color: #fff; /* Белый текст */
    font-family: "Manrope";
    font-size: 18px; /* Размер текста */
    outline: none; /* Убираем обводку при фокусе */
    transition: border-color 0.3s ease; /* Анимация при изменении цвета */
    margin-bottom: 16px;
  }

  .main_text {
    font-family: "Manrope";
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: start;
    width: 100%;
  }

  .main_text-other {
    font-family: "Manrope";
    display: flex;
    flex-direction: column;
  }

  .edit {
    font-family: "Manrope";
    color: #fff;
    font-size: 20px;
    background-color: transparent;
    border: 2px solid #fff;
    border-radius: 50px;
    padding: 12px 24px;
  }

  input:disabled, textarea:disabled {
    color: #aaa;
  }
  .edit, .save {
    font-family: "Manrope";
    color: #fff;
    font-size: 20px;
    background-color: transparent;
    border: 2px solid #fff;
    border-radius: 50px;
    padding: 12px 24px;
    cursor: pointer;
  }
  .edit:hover, .save:hover {
    background-color: #fff;
    color: #000;
  }






.toggle-switch {
  display: flex;
  align-items: center;
  text-align: left;
  width: 100%;
  gap: 10px; /* Расстояние между текстом и переключателем */
  margin-bottom: 4px;

}

/* Стиль текста */
.toggle-label {
  font-family: "Manrope";
  font-size: 18px;
  color: #fff;
  cursor: pointer;
  text-align: left;
}

/* Прячем стандартный чекбокс */
.toggle-checkbox {
  display: none;
}

/* Стиль переключателя */
.toggle-slider {
  width: 40px;
  height: 20px;
  background-color: #ccc; /* Цвет выключенного переключателя */
  border-radius: 20px;
  position: relative;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

/* Круглая кнопка внутри переключателя */
.toggle-slider::before {
  content: "";
  width: 16px;
  height: 16px;
  background-color: #fff; /* Цвет кнопки */
  border-radius: 50%;
  position: absolute;
  top: 2px;
  left: 2px;
  transition: transform 0.3s ease;
}

/* Когда переключатель включен */
.toggle-checkbox:checked + .toggle-slider {
  background-image: var(--gradient); /* Цвет включенного переключателя */
}

/* Смещаем круглую кнопку вправо */
.toggle-checkbox:checked + .toggle-slider::before {
  transform: translateX(20px);
}
</style>