import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    shorteds: [],
  },
  getters: {
    shorteds: (state) => state.shorteds,
  },
  mutations: {
    SAVE_URL_SUCCESS: (state, response) => {
      const { data } = response;
      const item = state.shorteds.find((x) => x.code === data.code);
      if (!item) {
        state.shorteds.unshift(data);
        const parsed = JSON.stringify(state.shorteds);
        localStorage.setItem('shorteds', parsed);
      }
    },
    GET_LOCALSTORAGE_SHORTEDS: (state) => {
      state.shorteds = JSON.parse(localStorage.getItem('shorteds') || '[]');
    },
    DELETE_LOCALSTORAGE_SHORTED: (state, code) => {
      const index = state.shorteds.findIndex((x) => x.code === code);
      if (index >= 0) {
        state.shorteds.splice(index, 1);
        const parsed = JSON.stringify(state.shorteds);
        localStorage.setItem('shorteds', parsed);
      }
    },
  },
  actions: {
    SAVE_URL: ({ commit }, url) => new Promise((resolve, reject) => {
      const main_url = `${window.location.protocol}//${window.location.host}`;
      axios.post(`${main_url}/v1/links`, {
        url,
      })
        .then((response) => {
          resolve(response);
          commit('SAVE_URL_SUCCESS', response);
        })
        .catch((err) => {
          reject(err.response);
        });
    }),
    DELETE_SHORTED: ({ commit }, code) => {
      commit('DELETE_LOCALSTORAGE_SHORTED', code);
    },
    GET_SAVED_SHORTEDS: ({ commit }) => {
      commit('GET_LOCALSTORAGE_SHORTEDS');
    },
  },
  modules: {
  },
});
