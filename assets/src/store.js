import Vue from 'vue'
import Vuex from 'vuex'
import router from './router'
import { AmplifyEventBus } from "aws-amplify-vue";
import { Auth } from "aws-amplify";

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    userId: 0,
    avater: '',
    userName: '',
    email: '',
    idToken: null,
    signIn: false
  },
  getters: {
    signIn: state => state.signIn,
    userId: state => state.userId,
    avater: state => state.avater,
    userName: state => state.userName,
  },
  mutations: {
    setUser(state, user) {
      state.userName = user.username;
      state.email = user.attributes.email;
      state.idToken = user.signInUserSession.idToken
    },
    setAuthState(state, auth) {
      state.signIn = auth;
    },
    setUserId(state, userId) {
      state.userId = userId;
    },
    setAvater(state, avater) {
      state.avater = avater;
    }

  },
  actions: {
    async login({ commit }) {
      try {
        const user = await Auth.currentAuthenticatedUser();
        commit("setUser", user);
        commit("setAuthState", true)
        router.push('/list')
      } catch (err) {
        commit("setAuthState", false)
      }
    },
    loginCheck({ commit }) {
      AmplifyEventBus.$on("authState", info => {
        if (info === "signedIn") {
          commit("setAuthState", true)
          router.push('/list')
        } else {
          commit("setAuthState", false)
        }
      });
    },
    userId({ commit }, userId) {
      commit("setUserId", userId)
    },
    avater({ commit }, avater) {
      commit("setAvater", avater)
    },
  }
});

export default store
