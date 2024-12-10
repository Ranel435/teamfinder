// src/lib/stores/authStore.ts
import { writable } from 'svelte/store';

export const accessToken = writable<string | null>(sessionStorage.getItem('accessToken'));
export const refreshToken = writable<string | null>(localStorage.getItem('refreshToken'));

// Утилиты
export function saveTokens(access: string, refresh: string) {
  accessToken.set(access);
  refreshToken.set(refresh);
  sessionStorage.setItem('accessToken', access);
  localStorage.setItem('refreshToken', refresh);
}

export function clearTokens() {
  accessToken.set(null);
  refreshToken.set(null);
  sessionStorage.removeItem('accessToken');
  localStorage.removeItem('refreshToken');
}
