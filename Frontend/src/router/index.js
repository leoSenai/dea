import { createRouter, createWebHistory } from 'vue-router';
import { routes } from './routes';

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to) => {
  const authToken = `; ${document.cookie}`.split('; authToken=').pop().split(';').shift()
  if (!authToken && to.path !== '/login') {
    return { path: '/login' }
  } else if (authToken && to.path === '/login') {
    return { path: '/' }
  }
});

export default router;
