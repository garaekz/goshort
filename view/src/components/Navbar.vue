<template>
  <header class="flex items-center justify-center p-4 pb-0 pb-4 bg-urbano-menubg">
    <!-- Global navigation -->
    <nav>
      <ul class="list-reset flex items-center">
        <li class="ml-4">
          <router-link class="text-urbano-menulinks uppercase text-sm no-underline py-2 border-none p-0" :to="'/'">
            Home
          </router-link>
        </li>
        <li class="ml-4">
          <a class="text-urbano-menulinks uppercase text-sm no-underline py-2 border-none p-0" href="https://github.com/garaekz/goshort/">
            Github
          </a>
        </li>
        <!-- <li class="ml-4">
          <a class="text-urbano-menulinks uppercase text-sm no-underline py-2 border-none p-0" href="#">
            Features
          </a>
        </li>
        <li class="ml-4">
          <a class="text-urbano-menulinks uppercase text-sm no-underline py-2 border-none p-0" href="#">
            Pricing
          </a>
        </li> -->
      </ul>
    </nav>
    <!-- END Global navigation -->
  </header>
</template>

<script>

export default {
  name: 'Navbar',
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
nav {
  font-family: Montserrat, Helvetica, Arial, sans-serif;
  font-weight: 500;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
</style>
