import { writable } from 'svelte/store';

export interface Profile {
  profileId: string;
  name: string;
  surname: string;
  tgId: string; // @айди
  email: string; // ***@**.**
  password: string;
  about: string; // до 300 символов
  portfolio: string; // ссылка
  university: string;
  academicGroup: string;
  extraEducation: string;
}

export interface Hakathon {
  hakathonId: string;
  name: string;
  startDate: string;
  endDate: string;
  online: number;
  place: string;
  about: string;
}

export interface Form { // анкета
  profileId: string;
  hakathonId: string;
  formId: string;
  role: string;
  skills: string[];
  experience: string;
}

export interface Team {
  teamId: string;
  name: string;
  teammatesProfiles: Profile[];
  teammatesForms: Form[];
}

export interface Vacancy { // вакансия
  teamId: string;
  hakathonId: string;
  vacancyId: string;
  role: string;
  skills: string[];
}


// export interface User {
//   login: string;
//   name: string;
//   surname: string;
//   tgid: string;
//   email: string;
//   about: string;

//   software: string;
//   skills: string;
//   qualities: string;

//   edu: string;
//   edu_group: string;
//   extra_edu: string;
// }


// export interface Hackathon {
//   id: number; // идентификатор хакатона
//   name: string; // название хакатона
//   description: string; // описание хакатона
//   startDate: string; // дата начала хакатона
//   endDate: string; // дата окончания хакатона
// }

// export interface Profile { // Анкета
//   id: string; // идентификатор анкеты
//   userId: number; // идентификатор пользователя
//   name: string; // имя пользователя
//   surname: string; 
//   academicGroup: string; // академическая группа
//   telegramHandle: string; // никнейм в Telegram
//   desiredRole: string; // желаемая роль
//   skills: string[]; // навыки пользователя
//   aboutMe: string; // информация о пользователе
//   achievements: string[]; // достижения пользователя
//   status: number; // статус профиля
// }

// export const profile = writable<User>({
//   login: "",
//   name: "",
//   surname: "",
//   tgid: "",
//   email: "",
//   about: "",
//   software: "",
//   skills: "",
//   qualities: "",
//   edu: "",
//   edu_group: "",
//   extra_edu: ""
// });








