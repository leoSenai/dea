import { createRouter, createWebHistory } from 'vue-router';
import { routes } from './routes';
import Cookie from '../utils/cookie';

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to) => {
  const authToken = Cookie.get('authToken');
  if (!authToken && to.path !== '/login') {
    return { path: '/login' };
  } else if (authToken && to.path === '/login') {
    return { path: '/' };
  }
});

export default router;
