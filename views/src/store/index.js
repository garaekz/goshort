import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    shorteds: []
  },
  mutations: {
    SAVE_SHORTED(state, payload) {
      payload.data.copying = false;
      state.shorteds.unshift(payload.data);
    }
  },
  getters: {
    shorteds: state => state.shorteds
  },
  actions: {
    saveShorted(context, payload) {
      context.commit("SAVE_SHORTED", payload);
    }
  },
  modules: {}
});
