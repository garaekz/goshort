<template>
  <div id="app">
    <Navbar />
    <router-view/>
  </div>
</template>

<script>
// @ is an alias to /src
import Navbar from '@/components/Navbar.vue';

export default {
  name: 'Home',
  components: {
    Navbar,
  },
};
</script>

<style lang="scss">
@import './assets/sass/index.scss';
</style>
