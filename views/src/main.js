import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import SuiVue from "semantic-ui-vue";
import "semantic-ui-css/semantic.min.css";
import VueClipboard from 'vue-clipboard2'
import VueQrcode from '@chenfengyuan/vue-qrcode';

Vue.config.productionTip = false;
Vue.use(VueClipboard);
Vue.use(SuiVue);
Vue.component(VueQrcode.name, VueQrcode);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
