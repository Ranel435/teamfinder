<script lang="ts">
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
  
    // Поля для хранения состояния
    let username: string | null = null;
    let errorMessage: string = "";
  
    // Функция для отправки данных на сервер
    async function login() {
      try {
        // Отправка запроса на сервер
        const response = await fetch('http://89.111.154.197:8080/api/product/1', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          }
        });
  
        // if (response.ok) {
        //   // Обработка успешного ответа, например, получение и сохранение токена
        //   const data = await response.json();
        //   const token = data.token; // Извлечение токена из ответа
        //   localStorage.setItem('token', token); // Сохранение токена для последующего использования
        // } else {
        //   // Обработка ошибок
        //   errorMessage = ('Ошибка входа. Пожалуйста, повторите попытку через 5 минут.');
        // }
      } catch (error) {
        // console.error('Ошибка при подключении к серверу:', error);
        errorMessage = ('Ошибка подключения. Попробуйте еще раз позже.');
      }
    }
</script>
  
  <!-- HTML-разметка формы -->
<!-- <form on:submit|preventDefault={login}>
<label for="username">Введите ваше имя пользователя:</label>
<input type="text" id="username" bind:value={username} required /> -->

<button type="submit" on:click={() => login()}>Войти</button>

{#if errorMessage}
    <p class="error">{errorMessage}</p>
{/if}
<!-- </form> -->

<style>
    .error {
        color: red;
    }
</style>
  