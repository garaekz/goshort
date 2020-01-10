<template>
  <sui-menu secondary>
    <sui-menu-item active link name="Home">Home</sui-menu-item>

    <sui-menu-menu position="right">
      <sui-menu-item>
        <sui-button color="teal" content="Login / Register" />
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
