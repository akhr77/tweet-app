import Vue from 'vue'
import VueRouter from 'vue-router'
import List from './views/List.vue'
import Login from './views/Login.vue'
import Detail from './views/Detail.vue'
import Profile from './views/Profile.vue'
import MyPage from './views/MyPage.vue'
import Post from './views/Post.vue'
import Sorry from './views/Sorry.vue';

Vue.use(VueRouter)


const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/list',
      name: 'list',
      component: List,
    },
    {
      path: '/detail',
      name: 'detail',
      component: Detail,
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile,
    },
    {
      path: '/mypage',
      name: 'mypage',
      component: MyPage,
    },
    {
      path: '/post',
      name: 'post',
      component: Post,
    },
    {
      path: '*',
      name: 'sorry',
      component: Sorry,
    }
  ]
})

export default router
