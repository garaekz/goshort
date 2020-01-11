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
              v-model="url"
              @keyup.enter="saveURL"
            />
            <div class="ui teal button" @click.prevent="saveURL">
              <i class="paper plane icon"></i>
            </div>
          </div>
        </sui-grid-column>
        <sui-grid-column :computer="2" :tablet="1" class=""></sui-grid-column>
      </sui-grid-row>
    </sui-grid>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "home",
  components: {},
  data: () => ({
    url: null
  }),
  methods: {
    async saveURL() {
      await axios
        .post("http://localhost:8080/api/v1/shorten", {
          original_url: this.url
        })
        .then(
          response => {
            console.log(response);
          },
          error => {
            console.log(error);
          }
        );
    }
  }
};
</script>
