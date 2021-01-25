<template>
  <div id="app">
    <!-- <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div> -->
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
