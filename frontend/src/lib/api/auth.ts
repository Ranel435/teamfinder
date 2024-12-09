export async function sendEmailCode(email: string): Promise<boolean> {
  try {
    const response = await fetch('/api/auth/login/email', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    });
    return response.ok;
  } catch (error) {
    console.error('Ошибка отправки email:', error);
    return false;
  }
}

export async function verifyEmailCode(email: string, code: string): Promise<{ access: string; refresh: string } | null> {
  try {
    const response = await fetch('/api/auth/verify/email', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, code }),
    });

    if (response.ok) {
      return response.json();
    }
    return null;
  } catch (error) {
    console.error('Ошибка верификации кода:', error);
    return null;
  }
}