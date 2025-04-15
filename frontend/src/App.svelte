<script lang="ts">
  import { model } from "../wailsjs/go/models";
  import { GetNews } from "../wailsjs/go/main/App";
  import { onMount } from "svelte";

  let newsArr: Array<model.NewsSite> = [];
  let loadingNews = false;

  onMount(() => {
    getNews();
  });

  function getNews() {
    loadingNews = true;
    newsArr = [];
    GetNews()
      .then((result) => (newsArr = result))
      .finally(() => (loadingNews = false));
  }

  function changeVisibilityNews(newsIndex: number) {
    newsArr[newsIndex].Visible = !newsArr[newsIndex].Visible;
  }
</script>

<main>
  <div class="p-4">
    {#if newsArr.length > 0}
      {#each newsArr as newsSite}
        <button
          on:click={() => changeVisibilityNews(newsArr.indexOf(newsSite))}
          class="bg-transparent hover:bg-neutral-950 w-full h-full border-b border-l border-r border-gray-500 rounded-b-lg pl-2 pb-4 mb-3 pt-2 active:bg-neutral-800"
        >
          <div class="border-white text-white flex text-left text-6xl">
            {newsSite.SiteName}
          </div>
        </button>
        <div>
          <div class="grid grid-cols-[26vw_40vw]">
            {#if newsSite.Visible}
              {#each newsSite.News as news}
                <div class="flex-row">
                  <img
                    src={news.ImageLink}
                    alt=""
                    class="object-cover h-60 w-96"
                  />
                </div>
                <div class="text-left">
                  <div class="text-white text-3xl">
                    {news.Headline}
                  </div>
                  <div class="text-gray-400 pt-2 text-2xl">
                    {news.SubHeadling}
                  </div>
                </div>
                <div
                  class="col-span-2 border-b border-gray-500 ml-4 mr-4 pt-2 mb-2"
                ></div>
              {/each}
            {/if}
          </div>
        </div>
      {/each}
    {/if}
    {#if loadingNews}
      <div>
        <div class="flex justify-center items-center">
          <svg
            class="animate-[spin_2000ms_infinite] size-10 text-white fill-current"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 512 512"
            ><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path
              d="M304 48a48 48 0 1 0 -96 0 48 48 0 1 0 96 0zm0 416a48 48 0 1 0 -96 0 48 48 0 1 0 96 0zM48 304a48 48 0 1 0 0-96 48 48 0 1 0 0 96zm464-48a48 48 0 1 0 -96 0 48 48 0 1 0 96 0zM142.9 437A48 48 0 1 0 75 369.1 48 48 0 1 0 142.9 437zm0-294.2A48 48 0 1 0 75 75a48 48 0 1 0 67.9 67.9zM369.1 437A48 48 0 1 0 437 369.1 48 48 0 1 0 369.1 437z"
            /></svg
          >
        </div>
      </div>
    {/if}
  </div>
</main>
