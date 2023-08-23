import { createRouter, createWebHashHistory } from 'vue-router';
import HomeView from '@/views/HomeView.vue';
import AboutView from '@/views/AboutView.vue';
import PageNotFoundView from '@/views/PageNotFound.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    component: AboutView,
  },
  {
    path: '/404',
    name: '404',
    component: PageNotFoundView,
  },
  {
    path: '/:catchAll(.*)',
    redirect: () => {
      return {
        path: '/404',
      }
    },
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
  linkActiveClass: "active",
});

export default router;
