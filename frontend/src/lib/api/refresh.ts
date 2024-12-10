// src/lib/api/refresh.ts
export async function refreshAccessToken(refreshToken: string): Promise<string | null> {
  try {
    const response = await fetch('http://localhost:8090/auth/refresh', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ refresh: refreshToken }),
    });

    if (response.ok) {
      const { access } = await response.json();
      return access;
    }
    return null;
  } catch (error) {
    console.error('Ошибка обновления access токена:', error);
    return null;
  }
}