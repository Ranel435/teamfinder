import { writable } from 'svelte/store';

const isBrowser = typeof window !== 'undefined';

interface User {
  id: number;
  username: string;
  email: string;
}

export const currentUser = writable<User | null>(null);

export const accessToken = writable<string | null>(
  isBrowser ? sessionStorage.getItem('accessToken') : null
);
export const refreshToken = writable<string | null>(
  isBrowser ? localStorage.getItem('refreshToken') : null
);

// Utilities
export function saveTokens(access: string, refresh: string) {
  if (isBrowser) {
    sessionStorage.setItem('accessToken', access);
    localStorage.setItem('refreshToken', refresh);
    accessToken.set(access);
    refreshToken.set(refresh);
  }
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
