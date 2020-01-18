<template>
  <div class="home">
    <sui-grid>
      <sui-grid-row>
        <sui-grid-column :computer="2" :tablet="1" class=""></sui-grid-column>
        <sui-grid-column centered :computer="12" :tablet="14" :mobile="16">
          <div class="ui one column stackable center aligned page grid logo">
            <img :src="require('@/assets/goshort.svg')" alt="">
          </div>
          <div
            v-if="shorted"
            class="ui one column stackable center aligned page grid shortened-res"
          >
            <div>
              <h1 v-clipboard:copy="shorted" @click="copyToClipboard()">
                <transition name="fade">
                  <sui-icon class="copy-icon" name="copy outline" v-if="!copying" />
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
    <sui-grid :columns="2">
      <sui-grid-row>
        <sui-grid-column>
          <sui-card class="fluid">
            <sui-card-content>
              <sui-card-header>Cute Dog</sui-card-header>
              <sui-card-meta>2 days ago</sui-card-meta>
              <sui-card-description>
                <sui-item-group>
                  <sui-item>
                    <qrcode value="http://dagacoding.com" :options="{ width: 200 }"></qrcode>
                    <sui-item-content>
                      <sui-item-header>Header</sui-item-header>
                      <sui-item-meta>
                        <span>Description</span>
                      </sui-item-meta>
                      <sui-item-description>
                        <p>
                          Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor
                          invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.
                        </p>
                      </sui-item-description>
                      <sui-item-extra>
                        Additional Details
                      </sui-item-extra>
                    </sui-item-content>
                  </sui-item>
                </sui-item-group>
              </sui-card-description>
            </sui-card-content>
          </sui-card>
        </sui-grid-column>
        <sui-grid-column>
          <sui-card class="fluid">
              <sui-card-content>
                <sui-card-header>Cute Dog</sui-card-header>
                <sui-card-meta>2 days ago</sui-card-meta>
                <sui-card-description>
                  <sui-item-group>
                    <sui-item>
                      <qrcode value="http://dagacoding.com" :options="{ width: 200 }"></qrcode>
                      <sui-item-content>
                        <sui-item-header>Header</sui-item-header>
                        <sui-item-meta>
                          <span>Description</span>
                        </sui-item-meta>
                        <sui-item-description>
                          <p>
                            Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor
                            invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.
                          </p>
                        </sui-item-description>
                        <sui-item-extra>
                          Additional Details
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
import axios from "axios";

export default {
  name: "home",
  components: {},
  data: () => ({
    url: null,
    copying: false,
    error: null,
    shorted: null
  }),
  methods: {
    copyToClipboard() {
      this.copying = true;
      setTimeout(() => (this.copying = false), 1000);
    },
    async saveURL() {
      await axios
        .post("/api/v1/shorten", {
          original_url: this.url
        })
        .then(
          response => {
            // TODO: Implement env with domain name
            this.shorted = window.location.host + "/" + response.data.url.code;
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
