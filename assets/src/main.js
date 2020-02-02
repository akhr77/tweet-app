import Vue from 'vue'
import App from './App.vue'
import router from './router'
import cognito from './cognito'

import axios from 'axios'
import './plugins/element.js'
Vue.prototype.$axios = axios

Vue.config.productionTip = false

new Vue({
  router,
  cognito,
  render: h => h(App)
}).$mount('#app')
