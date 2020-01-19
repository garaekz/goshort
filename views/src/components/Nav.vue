<template>
  <sui-menu secondary>
    <sui-menu-menu>
      <sui-menu-item>
        <router-link is="sui-button" to="/" color="teal">Home</router-link>
      </sui-menu-item>
    </sui-menu-menu>
    <sui-menu-menu position="right">
      <!-- <sui-menu-item>
        Will uncomment when the login's ready
        <sui-button disabled color="teal" content="Login / Register" />
      </sui-menu-item> -->
      <sui-menu-item>
        <a
          href="https://github.com/garaekz/goshort"
          class="ui icon button teal"
        >
          <i class="github icon"></i>
          Github
        </a>
      </sui-menu-item>
    </sui-menu-menu>
  </sui-menu>
</template>

<script>
export default {
  name: "Nav"
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
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
a {
  color: #42b983;
}
</style>
