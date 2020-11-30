import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store/index';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => {
      store.state.logginState = true;
      return import ('../views/Home.vue')
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => {
      store.state.logginState = false;
      return import ('../views/Login.vue')
    }
  },
  {
    path: '/booksgenre/:genreId',
    name: 'BooksGenre',
    component: () => { 
      store.state.logginState = true;
      return import('../views/BooksGenre.vue')
    },
    props: true
  },
  {
    path: '/review/:id',
    name: 'Review',
    component: () => { 
      store.state.logginState = true;
      return import('../views/Review.vue')
    },
    props: true
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
