<template>
  <div class="home">
    <sui-grid>
        <sui-grid-row>
          <sui-grid-column :computer="2" :tablet="1" class=""></sui-grid-column>
          <sui-grid-column :computer="12" :tablet="14" :mobile="16" class="">
            <div class="ui action left icon input full-width">
              <input
                type="text"
                placeholder="Paste your URL and GoShort"
                class="main-input"
              />
              <div class="ui teal button"><i class="paper plane icon"></i></div>
            </div>
          </sui-grid-column>
          <sui-grid-column :computer="2" :tablet="1" class=""></sui-grid-column>
        </sui-grid-row>
      </sui-grid>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "home",
  components: {}
};
</script>
