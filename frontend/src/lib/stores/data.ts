import { writable } from 'svelte/store';

export interface Profile {
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


export interface Form {
  id: string;
  name: string;
  surname: string;
  tgid: string;
  role: string;
  skills: string;
  about: string;
  archivements: string;
  status: number; // если 1 -> ищет команду, если 0 -> не ищет команду
}


export interface Team {
  id: string;
  name: string;
  about: string;
  teammates: Form[];
}


export const profile = writable<Profile>({
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