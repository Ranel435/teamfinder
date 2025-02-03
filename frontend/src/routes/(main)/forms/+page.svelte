<script lang="ts">
  import { type Profile } from "$lib/stores/data";
  let array: number[] = [1, 2, 3, 4, 5, 6];

  let profiles: Profile[] = [];
</script>



<section class="forms">
  <div class="forms-containner">
    <div class="forms-title">
      <h2>Анкеты</h2>
      <div class="forms-grid">
          {#each array as item}
            <div class="form">
                <h3 class="form-role">Frontend</h3>
                <p class="form-skills">Svelte / TS· Опыт 10+ лет</p>
                <hr>
                <h3 style="margin-bottom: 8px;" class="gradient-text">Достижения</h3>
                {#each profiles as profile}
                {/each}
                <!-- <div class="form-achievements">
                    <p>1 хак</p>
                    <p>0 побед</p>
                    <p>3м 5д в сборе</p>
                </div> -->
                <button class="slot-more_about">Подробнее</button>
            </div>
          {/each}
        </div>
        <div class="more-roles">
            <a class="more" href="/">Другие команды</a>
        </div>
    </div>
</section>


<style>

    .forms {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 128px;
    }

    .forms-grid {
        display: grid;
        grid-template-columns: repeat(3, minmax(auto, 100vw));
    }

    .form {
        display: flex;
        flex: 1;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        text-align: left;

        border: 3px solid #fff;
        border-radius: 24px;
        padding: 16px;
        margin: 0px 36px 36px 0;
        font-family: 'Manrope';
        color: #fff;  
    }

    .form:nth-child(3n) {
        margin-right: 0;
    }


    .form h4 {
        color: var(--pink);
        font-size: 12px;
        font-weight: 500;
    }

    .form-role {
        font-size: 24px;
    }

    .form-skills {
        font-size: 12px;
        color: var(--dark-grey);
    }


    .form-achievements {
        display: flex;
        flex-direction: row;
        justify-content: left;
        width: 100%;
        align-items: center;
        margin-bottom: 21px;
    }

    .form-achievements p {
        font-size: 10px;
        background-color: #9C004A;
        padding: 5px 13px;
        border-radius: 12px;
        margin-right: 10px;
    }

    .slot-more_about {
        font-size: 14px;
        padding: 8px 0;
        width: 100%;
        border: 2px solid;
        border-color: #fff;
        border-radius: 50px;
        background-color: rgba(0, 0, 0, 1);
        cursor: pointer;
    }

    .slot-more_about:hover {
        transition: all 0.7s ease;
        color: #000;
        background-color: #fff;
    }

    hr {
      width: 100%;
    }
</style>