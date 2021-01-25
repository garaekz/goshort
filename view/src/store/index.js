import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
    FETCH_LINK_BY_CODE: (_, code) => new Promise((resolve, reject) => {
      console.log(code);
      axios({ url: `http://localhost:8080/v1/links/${code}`, method: 'GET' })
        .then((response) => resolve(response))
        .catch((err) => {
          reject(err.response);
        });
    })
    ,
  },
  modules: {
  },
});
