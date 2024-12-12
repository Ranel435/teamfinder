import { writable } from 'svelte/store';

const isBrowser = typeof window !== 'undefined';

export const accessToken = writable<string | null>(
  isBrowser ? sessionStorage.getItem('accessToken') : null
);
export const refreshToken = writable<string | null>(
  isBrowser ? localStorage.getItem('refreshToken') : null
);

// Utilities
export function saveTokens(access_token: string, refresh_token: string) {
  accessToken.set(access_token);
  refreshToken.set(refresh_token);

  if (isBrowser) {
    sessionStorage.setItem('accessToken', access_token);
    localStorage.setItem('refreshToken', refresh_token);
  }
}

export function clearTokens() {
  accessToken.set(null);
  refreshToken.set(null);

  if (isBrowser) {
    sessionStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
  }
}
