import Vue from "vue";
import SuiVue from "semantic-ui-vue";
import "semantic-ui-css/semantic.min.css";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import VueClipboard from "vue-clipboard2";
import VueQRCodeComponent from "vue-qrcode-component";

Vue.config.productionTip = false;
Vue.use(VueClipboard);
Vue.use(SuiVue);
Vue.component("qr-code", VueQRCodeComponent);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
