// src/lib/stores/authStore.ts
import { writable } from 'svelte/store';

interface User {
  id: number;
  username: string;
  email: string;
}

export const accessToken = writable<string | null>(sessionStorage.getItem('accessToken'));
export const refreshToken = writable<string | null>(localStorage.getItem('refreshToken'));
export const currentUser = writable<User | null>(null);

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
  currentUser.set(null);
  sessionStorage.removeItem('accessToken');
  localStorage.removeItem('refreshToken');
}

export async function fetchUserProfile() {
  const token = sessionStorage.getItem('accessToken');
  if (!token) return null;

  try {
    const response = await fetch('/api/users/me', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });
    if (response.ok) {
      const userData = await response.json();
      currentUser.set(userData);
      return userData;
    }
    return null;
  } catch (error) {
    console.error('Failed to fetch user profile:', error);
    return null;
  }
}
