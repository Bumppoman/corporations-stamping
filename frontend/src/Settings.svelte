<script>

</script>

<main>
  <div class="border-b border-slate-500 bg-slate-800 px-4 py-5 sm:px-6">
    <div class="-ml-4 -mt-2 flex flex-wrap items-center justify-between sm:flex-nowrap">
      <div class="ml-4 mt-2">
        <h3 class="text-base font-semibold leading-6 text-slate-50">Settings</h3>
      </div>
      <div class="ml-4 mt-2 flex-shrink-0">
        <button
          type="button"
          class="relative inline-flex items-center rounded-md bg-sky-600 px-3 py-2 text-sm font-semibold text-slate-50 shadow-sm hover:bg-sky-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-sky-600"
        >Save</button>
      </div>
    </div>
  </div>
  <div class="bg-slate-700 min-h-full">
    <form class="md:col-span-2">
      <div class="px-4 py-6 sm:p-8">
        <div class="grid max-w-2xl grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
          <div class="sm:col-span-4">
            <label for="listURI" class="block text-sm font-medium leading-6 text-slate-50">SharePoint List URI</label>
            <div class="mt-2">
              <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-slate-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-sky-600 sm:max-w-md">
                <input
                  type="text"
                  name="listURI"
                  id="listURI"
                  class="block flex-1 border-0 bg-transparent py-1.5 pl-3 text-slate-50 placeholder:text-slate-300/40 focus:ring-0 sm:text-sm sm:leading-6"
                  placeholder="https://nysemail.sharepoint.com"
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</main>
