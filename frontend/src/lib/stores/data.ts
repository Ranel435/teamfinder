import { writable } from 'svelte/store';

export interface User {
  login: string;
  name: string;
  surname: string;
  tgid: string;
  email: string;
  about: string;

  software: string;
  skills: string;
  qualities: string;

  edu: string;
  edu_group: string;
  extra_edu: string;
}


export interface Hackathon {
  id: number; // идентификатор хакатона
  name: string; // название хакатона
  description: string; // описание хакатона
  startDate: string; // дата начала хакатона
  endDate: string; // дата окончания хакатона
}

export interface Profile { // Анкета
  id: number; // идентификатор анкеты
  userId: number; // идентификатор пользователя
  name: string; // имя пользователя
  surname: string; 
  academicGroup: string; // академическая группа
  telegramHandle: string; // никнейм в Telegram
  desiredRole: string; // желаемая роль
  skills: string[]; // навыки пользователя
  aboutMe: string; // информация о пользователе
  achievements: string[]; // достижения пользователя
  status: string; // статус профиля
}

export interface Team {
  id: string;
  name: string;
  about: string;
  teammates: Profile[];
}


export const profile = writable<User>({
  login: "",
  name: "",
  surname: "",
  tgid: "",
  email: "",
  about: "",

  software: "",
  skills: "",
  qualities: "",

  edu: "",
  edu_group: "",
  extra_edu: ""
});


