// Замените хардкод URL на переменную окружения
const API_URL = import.meta.env.PROD 
  ? '/api'  // В production используем относительный путь
  : import.meta.env.VITE_API_URL || 'http://localhost:8090';

export async function sendEmailCode(email: string): Promise<boolean> {
  try {
    const response = await fetch(`${API_URL}/auth/login/email`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include', // Добавляем для работы с куками
      body: JSON.stringify({ email }),
    });
    return response.ok;
  } catch (error) {
    console.log(email);
    console.error('Ошибка отправки email:', error);
    return false;
  }
}

export async function verifyEmailCode(email: string, code: string, login: string, password: string): Promise<{ access: string, refresh: string } | null> {
  try {
    const response = await fetch(`${API_URL}/auth/verify/email`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        email,
        code,
        username: login,
        password
      }),
    });

    if (response.ok) {
      const tokens = await response.json();
      return tokens;
    }
    return null;
  } catch (error) {
    console.error('Ошибка верификации кода:', error);
    return null;
  }
}

// Добавляем функцию для сохранения токенов
function saveTokens(accessToken: string, refreshToken: string) {
  sessionStorage.setItem('accessToken', accessToken);
  localStorage.setItem('refreshToken', refreshToken);
}

export async function login(email: string, password: string) {
  try {
    const response = await fetch(`${API_URL}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        email,
        password
      }),
    });

    if (response.ok) {
      const data = await response.json();
      saveTokens(data.access_token, data.refresh_token);
      return true;
    }
    return false;
  } catch (error) {
    console.error('Ошибка входа:', error);
    return false;
  }
}