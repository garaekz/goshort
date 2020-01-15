import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import SuiVue from "semantic-ui-vue";
import "semantic-ui-css/semantic.min.css";
import VueClipboard from 'vue-clipboard2'

Vue.config.productionTip = false;
Vue.use(VueClipboard);
Vue.use(SuiVue);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
