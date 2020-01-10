<template>
  <div id="app">
    <div is="sui-container">
      <Nav />
      <router-view />
    </div>
  </div>
</template>
<script>
import Nav from "@/components/Nav.vue";

export default {
  name: "home",
  components: {
    Nav
  }
};
</script>

<style lang="scss">
@import "@/assets/styles/main.scss";
</style>
