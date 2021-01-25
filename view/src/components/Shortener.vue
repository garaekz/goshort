<template>
  <div class="flex flex-wrap items-stretch py-12 relative md:w-44r w-full">
    <input type="text" class="flex-shrink flex-grow flex-auto leading-normal w-px flex-1 border h-12 border-grey-light rounded px-3 relative focus:border-none" placeholder="Recipient's username">
    <div class="flex -mr-px">
      <button class="flex items-center bg-urbano-green leading-normal bg-grey-lighter rounded rounded-l-none border border-l-0 border-grey-light px-3 whitespace-no-wrap text-white h-12 text-sm">
        <svg class="h-4 fill-current text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M476 3.2L12.5 270.6c-18.1 10.4-15.8 35.6 2.2 43.2L121 358.4l287.3-253.2c5.5-4.9 13.3 2.6 8.6 8.3L176 407v80.5c0 23.6 28.5 32.9 42.5 15.8L282 426l124.6 52.2c14.2 6 30.4-2.9 33-18.2l72-432C515 7.8 493.3-6.8 476 3.2z"/></svg>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Shortener',
};
</script>

<style scoped lang="scss">

</style>
