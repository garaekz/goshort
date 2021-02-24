<template>
  <div class="home pt-36">
    <div class="flex items-center justify-center">
      <img class="h-32" alt="GoShort Logo" src="../assets/goshort.svg">
    </div>
    <div class="mt-12 flex items-center justify-center bg-gray-600 bg-opacity-20 pt-6 pb-6">
      <Shortener/>
    </div>
    <div class="w-full flex justify-center pt-2">
      Previous <span class="pl-2 uppercase font-extrabold">Shorts</span>
    </div>
    <div class="w-full flex justify-center pb-8">
      <svg class="h-4" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" viewBox="0 0 350 350" xml:space="preserve">
      <g id="icon" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: none; fill-rule: nonzero; opacity: 1;" transform="translate(-1.9444444444444287 -1.9444444444444287) scale(3.89 3.89)" >
        <polygon points="0,38.92 2.83,36.08 45,78.25 87.17,36.08 90,38.92 45,83.92 " style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(0,0,0); fill-rule: nonzero; opacity: 1;" transform="  matrix(1 0 0 1 0 0) "/>
        <polygon points="0,8.92 2.83,6.08 45,48.25 87.17,6.08 90,8.92 45,53.92 " style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(0,0,0); fill-rule: nonzero; opacity: 1;" transform="  matrix(1 0 0 1 0 0) "/>
      </g>
      </svg>
    </div>
    <div class="w-full flex justify-center">
      <ShortedList/>
    </div>
    <div class="py-48">
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Shortener from '@/components/Shortener.vue';
import ShortedList from '@/components/ShortedList.vue';
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'Home',
  components: {
    Shortener,
    ShortedList,
  },
  computed: {
    ...mapGetters(['shorteds']),
  },
  mounted() {
  },
  methods: {
    ...mapActions({
      getSavedShorteds: 'GET_SAVED_SHORTEDS',
    }),
  },
};
</script>
<style scoped lang="scss">
.home {
  background: url('../assets/goshortbg.svg') no-repeat center center / cover;
}
</style>
