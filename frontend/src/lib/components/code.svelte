<script lang="ts">
    import '../../app.css';
  export let code = ["", "", "", "", "", ""]; // Массив для хранения введенных цифр
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

<div class="container">
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
    width: 48px;
    height: 56px;
    text-align: center;
    font-size: 20px;
    margin-right: 12px;
    border: 0px solid #ccc;
    border-radius: 16px;
  }
  .input-field:last-child {
    margin-right: 0px;
  }
  .input-field:focus {
    border-color: #007bff;
    outline: none;
  }
</style>