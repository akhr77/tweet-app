import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import Buefy from 'buefy';
import './plugins/element.js'

import Amplify, * as AmplifyModules from 'aws-amplify'
import { AmplifyPlugin } from 'aws-amplify-vue'
import awsconfig from './aws-exports'
Amplify.configure(awsconfig)
Vue.use(Buefy)
Vue.use(AmplifyPlugin, AmplifyModules)

Vue.prototype.$axios = axios
Vue.config.productionTip = false
// axios.defaults.baseURL = 'http://localhost:8082';

let messageResource = {
  ja: {
    'Sign in to your account': 'サインイン',
    'Username': 'ユーザ名',
    'Enter your Username': 'ユーザ名を入力してください',
    'Password': 'パスワード',
    'Enter your password': 'パスワードを入力してください',
    'Forget your password? ': ' ',
    'Reset password': 'パスワードのリセット',
    'No account? ': ' ',
    'Create account': 'アカウント作成',
    'Sign In': 'サインイン',
    'User does not exist.': '登録がされていないユーザー名です。',
    'Password attempts exceeded': 'パスワードの試行回数を超えました。',
    'Incorrect username or password.': 'ユーザー名またはパスワードが間違っています。'
  }
}
AmplifyModules.I18n.putVocabularies(messageResource)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
