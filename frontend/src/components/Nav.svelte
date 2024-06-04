<script>
  import { getContext } from 'svelte';
  import { SignIn } from '../../bindings/corporations-stamping/stampservice.js';

  let { userInfo } = getContext('user');

  const signIn = async () => {
    $userInfo = await SignIn();
  };
</script>

<nav class="bg-slate-800 mb-5 text-slate-50">
  <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
    <div class="flex h-16 items-center justify-end">
      {#if $userInfo}
        {$userInfo.Title}
      {:else}
        <button on:click={signIn}>Sign in</button>
      {/if}
    </div>
  </div>
</nav>
