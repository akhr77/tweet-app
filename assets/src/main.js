import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueCookie from 'vue-cookie'
import axios from 'axios'
import './plugins/element.js'

import Amplify, * as AmplifyModules from 'aws-amplify'
import { AmplifyPlugin } from 'aws-amplify-vue'
import awsconfig from './aws-exports'
Amplify.configure(awsconfig)

Vue.use(AmplifyPlugin, AmplifyModules)
Vue.use(VueCookie)
Vue.prototype.$axios = axios
Vue.config.productionTip = false

let messageResource = {
  ja: {
    'Sign in to your account': 'アカウントにサインイン',
    'Username': 'ユーザ名',
    'Enter your Username': 'ユーザ名を入力してください',
    'Password': 'パスワード',
    'Enter your password': 'パスワードを入力してください',
    'Forget your password? ': 'パスワードを忘れた場合',
    'Reset password': 'パスワードのリセット',
    'No account? ': 'まだアカウントがない場合',
    'Create account': 'アカウント作成',
    'Sign In': 'サインイン',
    'User does not exist.': '登録がされていないユーザー名です。',
    'Password attempts exceeded': 'パスワードの試行回数を超えました。'
  }
}
AmplifyModules.I18n.putVocabularies(messageResource)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
