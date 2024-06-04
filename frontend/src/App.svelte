<script>
  import { getContext, setContext } from 'svelte';
  import { writable } from 'svelte/store';

  import FileList from './components/FileList.svelte';
  import Nav from './components/Nav.svelte';

  setContext('user', {
    userInfo: writable(null),
  });

  const { userInfo } = getContext('user');
</script>

<main class="bg-slate-700">
  <Nav />

  <div class="mx-auto px-5">
    {#if $userInfo}
      <FileList />
    {:else}
      <div class="flex items-center justify-center min-h-screen text-slate-50 w-full">
        <div>Please sign in to use this application.</div>
      </div>
    {/if}
  </div>
</main>

