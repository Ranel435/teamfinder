<script lang="ts">
    import '../app.css';
  let code = ["", "", "", "", "", ""]; // Массив для хранения введенных цифр
  let inputs: HTMLInputElement[] = []; // Масиив для ссылок на элементы input
  function handleInput(event: Event, index: number) {
    const target = event.target as HTMLInputElement;
    const value = target.value;
    // Проверка на ввод только цифр
    if (/^\d*$/.test(value)) {
      code[index] = value;

      // Переключение фокуса на следующее поле
      if (value && index < code.length - 1) {
        inputs[index + 1].focus();
      }
    } else {
      code[index] = "";
    }
    console.log(typeof value);
  }

  function handleKeyDown(event: KeyboardEvent, index: number) {
    if (event.key === "Backspace" && !code[index] && index > 0) {
      inputs[index - 1].focus();
    }
  }
</script>

<div>
  {#each code as digit, index}
    <input
      bind:this={inputs[index]}
      class="input-field"
      maxlength="1"
      bind:value={code[index]}
      on:input={(event) => handleInput(event, index)}
      on:keydown={(event) => handleKeyDown(event, index)}
    />
  {/each}
</div>

<style>
  .input-field {
    background-color: var(--light-grey);
    width: 52px;
    height: 60px;
    text-align: center;
    font-size: 24px;
    /* margin: 0 5px; */
    margin-right: 12px;
    border: 0px solid #ccc;
    border-radius: 36px;
  }
  .input-field:last-child {
    margin-right: 0px;
  }
  .input-field:focus {
    border-color: #007bff;
    outline: none;
  }
</style>


<!-- 
<script lang="ts">
  let code: string[] = ['', '', '', '', '', '']; // Массив для хранения введенных цифр
  let inputs: HTMLInputElement[] = []; // Массив для ссылок на элементы input

  function handleInput(event: Event, index: number) {
    const target = event.target as HTMLInputElement; // Приводим к типу HTMLInputElement
    const value = target.value;

    // Проверка на ввод только цифр
    if (/^\d*$/.test(value)) {
      code[index] = value; // Сохраняем введенное значение

      // Переключение фокуса на следующее поле
      if (value && index < code.length - 1) {
        inputs[index + 1].focus(); // Устанавливаем фокус на следующее поле
      }
    } else {
      code[index] = ''; // Очищаем поле, если введено не число
    }
  }

  function handleKeyDown(event: KeyboardEvent, index: number) {
    if (event.key === 'Backspace' && !code[index] && index > 0) {
      inputs[index - 1].focus(); // Устанавливаем фокус на предыдущее поле
    }
  }
</script>

<style>
  .input-field {
    width: 40px;
    height: 40px;
    text-align: center;
    font-size: 24px;
    margin: 0 5px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }

  .input-field:focus {
    border-color: #007BFF;
    outline: none;
  }
</style>

<div>
  {#each code as digit, index}
    <input
      bind:this={inputs[index]} 
      class="input-field"
      maxlength="1"
      bind:value={code[index]}
      on:input={event => handleInput(event, index)}
      on:keydown={event => handleKeyDown(event, index)}
    />
  {/each}
</div> -->