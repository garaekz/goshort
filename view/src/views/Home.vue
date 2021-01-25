<template>
  <div class="home pt-36">
    <div class="flex items-center justify-center">
      <img class="h-24 " alt="GoShort Logo" src="../assets/goshort.svg">
    </div>
    <div class="mt-12 flex items-center justify-center bg-gray-600 bg-opacity-20 px-8">
      <Shortener/>
    </div>
    <div class="py-48">
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Shortener from '@/components/Shortener.vue';

export default {
  name: 'Home',
  components: {
    Shortener,
  },
};
</script>
<style scoped lang="scss">
.home {
  background: url('../assets/goshortbg.svg') no-repeat center center / cover;
}
</style>
