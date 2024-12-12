<script lang="ts">
  import { type Profile } from "$lib/stores/data";
  let array: number[] = [1, 2, 3, 4, 5, 6];

  let profiles: Profile[] = [
    {
    id: 1,
    userId: 101,
    name: "Алексей",
    surname: "Иванов",
    academicGroup: "Группа A1",
    telegramHandle: "@alexey_ivanov",
    desiredRole: "Разработчик",
    skills: ["JavaScript", "TypeScript", "React"],
    aboutMe: "Я увлекаюсь программированием и разработкой веб-приложений.",
    achievements: ["Участие в хакатоне 2023", "Сертификат по JavaScript"],
    status: "Активен"
  },
  {
    id: 2,
    userId: 102,
    name: "Мария",
    surname: "Петрова",
    academicGroup: "Группа B2",
    telegramHandle: "@maria_petrovna",
    desiredRole: "Дизайнер",
    skills: ["Photoshop", "Figma", "UI/UX"],
    aboutMe: "Я дизайнер с опытом работы в различных проектах.",
    achievements: ["Победа в конкурсе дизайна", "Стажировка в крупной компании"],
    status: "Активен"
  },
  {
    id: 3,
    userId: 103,
    name: "Иван",
    surname: "Сидоров",
    academicGroup: "Группа C3",
    telegramHandle: "@ivan_sidorov",
    desiredRole: "Аналитик",
    skills: ["Анализ данных", "SQL", "Python"],
    aboutMe: "Интересуюсь анализом данных и машинным обучением.",
    achievements: ["Курс по Data Science", "Проект по анализу данных"],
    status: "Активен"
  },
  {
    id: 4,
    userId: 104,
    name: "Елена",
    surname: "Кузнецова",
    academicGroup: "Группа D4",
    telegramHandle: "@elena_kuznetsova",
    desiredRole: "Менеджер проекта",
    skills: ["Управление проектами", "Командная работа"],
    aboutMe: "Опытный менеджер проектов с успешными кейсами.",
    achievements: ["Успешное завершение проекта XYZ"],
    status: "Активен"
  },
  {
    id: 5,
    userId: 105,
    name: "Дмитрий",
    surname: "Смирнов",
    academicGroup: "Группа E5",
    telegramHandle: "@dmitry_smirnov",
    desiredRole: "Тестировщик",
    skills: ["Тестирование ПО", "Автоматизация тестирования"],
    aboutMe: "Специализируюсь на тестировании программного обеспечения.",
    achievements: ["Сертификат ISTQB", "Опыт работы в Agile-команде"],
    status: "Активен"
  },
  {
    id: 6,
    userId: 106,
    name: "Ольга",
    surname: "Федорова",
    academicGroup: "Группа F6",
    telegramHandle: "@olga_fedorova",
    desiredRole: "Frontend-разработчик",
    skills: ["HTML", "CSS", "JavaScript"],
    aboutMe: "Frontend-разработчик с опытом работы над сложными проектами.",
    achievements: ["Создание нескольких веб-сайтов"],
    status: "Активен"
  },
  {
    id: 7,
    userId: 107,
    name: "Артем",
    surname: "Морозов",
    academicGroup: "Группа G7",
    telegramHandle: "@artem_morozov",
    desiredRole: "Backend-разработчик",
    skills: ["Node.js", "Express", "MongoDB"],
    aboutMe: "Backend-разработчик с опытом работы с RESTful API.",
    achievements: ["Разработка нескольких серверных приложений"],
    status: "Активен"
  },
  {
   id :8,
   userId :108,
   name :"Наталья" ,
   surname :"Лебедева" ,
   academicGroup :"Группа H8" ,
   telegramHandle :"@natalya_lebedyeva" ,
   desiredRole :"Системный администратор" ,
   skills :["Linux" ,"Networking" ,"Security"] ,
   aboutMe :"Опытный системный администратор с навыками управления серверами." ,
   achievements :["Сертификат CCNA" ,"Участие в проекте по миграции серверов"] ,
   status :"Активен"
},
{
   id :9,
   userId :109,
   name :"Сергей" ,
   surname :"Павлов" ,
   academicGroup :"Группа I9" ,
   telegramHandle :"@sergey_pavlov" ,
   desiredRole :"Data Scientist" ,
   skills :["Python" ,"Machine Learning" ,"Statistics"] ,
   aboutMe :"Студент, увлеченный машинным обучением и анализом данных." ,
   achievements :["Курс по машинному обучению" ,"Участие в соревновании Kaggle"] ,
   status :"Активен"
}

  ];
</script>



<section class="forms">
  <div class="forms-containner">
    <div class="forms-title">
      <h2>Анкеты</h2>
      <div class="forms-grid">
          {#each profiles as profile}
            <div class="form">
                <h3 class="form-role">{profile.desiredRole}</h3>
                <p class="form-skills">{profile.skills}</p>
                <hr>
                <h3 style="margin-bottom: 8px;" class="gradient-text">Достижения</h3>
                <div class="form-achievements">
                    {#each profile.achievements as ach}
                      <p>{ach}</p>
                    {/each}
                </div>
                <div class="form-fio">
                  <p>{profile.name}</p>
                  <p>{profile.surname}</p>
                </div>
                <div class="form-about">
                  <p>{profile.aboutMe}</p>
                </div>
                <div class="form-buttons">
                  <button class="slot-more_about">В профиль</button>
                  <button class="slot-more_about">Пригласить</button>
                </div>
            </div>
          {/each}
        </div>
        <div class="more-roles">
            <a class="more" href="/">Другие команды</a>
        </div>
    </div>
</section>


<style>

    .forms {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 128px;
    }

    .forms-grid {
        display: grid;
        grid-template-columns: repeat(3, minmax(auto, 100vw));
    }

    .form {
        display: flex;
        flex: 1;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        text-align: left;

        border: 3px solid #fff;
        border-radius: 24px;
        padding: 16px;
        margin: 0px 36px 36px 0;
        font-family: 'Manrope';
        color: #fff;  
    }

    .form:nth-child(3n) {
        margin-right: 0;
    }


    .form h4 {
        color: var(--pink);
        font-size: 12px;
        font-weight: 500;
    }

    .form-role {
        font-size: 24px;
    }

    .form-skills {
        font-size: 12px;
        color: var(--dark-grey);
    }


    .form-achievements {
        display: flex;
        flex-direction: row;
        justify-content: left;
        width: 100%;
        align-items: center;
        margin-bottom: 8px;
    }

    .form-achievements p {
        font-size: 10px;
        background-color: #9C004A;
        padding: 5px 13px;
        border-radius: 12px;
        margin-right: 10px;
    }

    .slot-more_about {
        font-size: 14px;
        padding: 8px 0;
        width: 100%;
        border: 2px solid;
        border-color: #fff;
        border-radius: 50px;
        background-color: rgba(0, 0, 0, 1);
        cursor: pointer;
    }

    .slot-more_about:hover {
        transition: all 0.7s ease;
        color: #000;
        background-color: #fff;
    }

    hr {
      width: 100%;
    }

    .form-fio {
      display: flex;
      flex-direction: row;
      margin-bottom: 8px;
    }

    .form-fio p {
      margin-right: 8px;
    }

    .form-about p {
      font-family: "Manrope";
      font-size: 16px;
      color: var(--dark-grey);
    }

    .form-about {
      margin-bottom: 8px;
    }

    .form-buttons {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
    }

    .form-buttons button{
      padding: 8px 16px;
      color:#fff;
      margin-right: 10px;
    }
</style>