<template>
  <div class="home">
    <sui-grid>
      <sui-grid-row>
        <sui-grid-column :computer="2" :tablet="1" class=""></sui-grid-column>
        <sui-grid-column centered :computer="12" :tablet="14" :mobile="16">
          <div class="ui one column stackable center aligned page grid logo">
            <img :src="require('@/assets/goshort.svg')" alt="" />
          </div>
          <div
            v-if="shorted"
            class="ui one column stackable center aligned page grid shortened-res"
          >
            <div>
              <h1 v-clipboard:copy="shorted" @click="copyToClipboard()">
                <transition name="fade">
                  <sui-icon
                    class="copy-icon"
                    name="copy outline"
                    v-if="!copying"
                  />
                  <sui-icon class="check-icon" name="check" v-else />
                </transition>
                {{ shorted }}
              </h1>
            </div>
          </div>
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
          <sui-message
            v-if="error"
            icon="times circle"
            error
            header="Your URL is invalid"
            content="Please provide a valid URL."
          />
        </sui-grid-column>
        <sui-grid-column :computer="2" :tablet="1" class=""></sui-grid-column>
      </sui-grid-row>
    </sui-grid>
    <sui-grid stackable :columns="2" class="qr-grid">
      <sui-grid-row>
        <sui-grid-column
          class="qr-grid-column"
          v-for="(s, index) in shorteds"
          :key="index"
        >
          <sui-card class="fluid">
            <sui-card-content>
              <sui-card-description>
                <sui-item-group>
                  <sui-item>
                    <qr-code
                      :text="urlify(s.code, true)"
                      :size="200"
                      class="image"
                    ></qr-code>
                    <sui-item-content class="qr-content">
                      <sui-item-header>{{ urlify(s.code) }}</sui-item-header>
                      <sui-item-meta>
                        <span style="word-wrap:break-word;">{{
                          s.original_url
                        }}</span>
                      </sui-item-meta>
                      <sui-item-extra>
                        <sui-button-group attached="bottom">
                          <div
                            v-if="!s.copying"
                            is="sui-button"
                            content="Copy to clipboard"
                            icon="clipboard"
                            @click="copyQR(index)"
                          />
                          <div
                            v-else
                            color="teal"
                            is="sui-button"
                            content="Copied"
                            icon="check circle"
                          />
                        </sui-button-group>
                      </sui-item-extra>
                    </sui-item-content>
                  </sui-item>
                </sui-item-group>
              </sui-card-description>
            </sui-card-content>
          </sui-card>
        </sui-grid-column>
      </sui-grid-row>
    </sui-grid>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import axios from "axios";

export default {
  name: "home",
  components: {},
  computed: {
    ...mapGetters(["shorteds"])
  },
  data: () => ({
    url: null,
    copying: false,
    error: null,
    shorted: null
  }),
  methods: {
    urlify(code, protocol) {
      return (
        (protocol ? window.location.protocol + "//" : "") +
        window.location.host +
        "/" +
        code
      );
    },
    copyQR(key) {
      this.$copyText(this.urlify(this.shorteds[key].code, true)).then(() => {
        this.shorteds[key]["copying"] = true;
      });
      setTimeout(() => (this.shorteds[key]["copying"] = false), 1500);
    },
    copyToClipboard() {
      this.copying = true;
      setTimeout(() => (this.copying = false), 1000);
    },
    async saveURL() {
      await axios
        .post("/api/v1/shorten", {
          original_url: this.url.toLowerCase()
        })
        .then(
          response => {
            this.shorted = window.location.host + "/" + response.data.url.code;
            this.$store.dispatch("saveShorted", {
              data: response.data.url
            });
          },
          () => {
            this.error = true;
            setTimeout(() => (this.error = false), 5000);
          }
        );
      this.url = "";
    }
  }
};
</script>
