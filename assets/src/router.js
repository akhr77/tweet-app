import Vue from 'vue'
import VueRouter from 'vue-router'
import List from './views/List.vue'
import Auth from './views/Auth.vue'
import Detail from './views/Detail.vue'
import Profile from './views/Profile.vue'
import MyPage from './views/MyPage.vue'
import Post from './views/Post.vue'
import Logout from './views/Logout.vue'
import Sorry from './views/Sorry.vue';
import store from './store';

Vue.use(VueRouter)


const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'auth',
      component: Auth,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next('./list')
        } else {
          next()
        }
      }
    },
    {
      path: '/list',
      name: 'list',
      component: List,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next()
        } else {
          next('./')
        }
      }
    },
    {
      path: '/detail',
      name: 'detail',
      component: Detail,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next()
        } else {
          next('./')
        }
      }
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next()
        } else {
          next('./')
        }
      }
    },
    {
      path: '/mypage',
      name: 'mypage',
      component: MyPage,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next()
        } else {
          next('./')
        }
      }
    },
    {
      path: '/post',
      name: 'post',
      component: Post,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next()
        } else {
          next('./')
        }
      }
    },
    {
      path: '/logout',
      name: 'logout',
      component: Logout,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next()
        } else {
          next('./')
        }
      }
    },
    {
      path: '*',
      name: 'sorry',
      component: Sorry,
      beforeEnter(to, from, next) {
        if (store.getters.signIn) {
          next()
        } else {
          next('./')
        }
      }
    }
  ]
})

export default router
