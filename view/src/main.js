import Vue from 'vue';
import Clipboard from 'v-clipboard';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faCopy } from '@fortawesome/free-regular-svg-icons';
import { faCheck, faTimes, faClipboard } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import VueQRCodeComponent from 'vue-qrcode-component';
import App from './App.vue';
import store from './store';
import router from './router';

import './plugins';

Vue.config.productionTip = false;
Vue.use(Clipboard);
library.add(faCopy, faCheck, faTimes, faClipboard);
Vue.component('font-awesome-icon', FontAwesomeIcon);
Vue.component('qr-code', VueQRCodeComponent);

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app');
