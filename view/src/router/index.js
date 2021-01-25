import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/404',
    name: 'Page Not Found',
    component: () => import(/* webpackChunkName: "about" */ '../views/404.vue'),
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    redirect: '/404',
  },
];

const router = new VueRouter({
  routes,
});

export default router;
