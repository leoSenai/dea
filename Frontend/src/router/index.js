import { createRouter, createWebHistory } from 'vue-router';
import { routes } from './routes';
import Cookie from '../utils/cookie';
import { RoleEnum } from '../utils/Enum';
import { links } from './routes';

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to) => {
  const authToken = Cookie.get('authToken');
  if (authToken) {
    const typeUser = Cookie.getUserType(authToken);
    if (typeUser === RoleEnum.Patient || typeUser === RoleEnum.Person) {
      const userLinkIndex = links.indexOf(
        links.find((el) => el.path === '/usuarios')
      );
      if (userLinkIndex >= 0) {
        links.splice(userLinkIndex, 1);
      }
      const patientLinkIndex = links.indexOf(
        links.find((el) => el.path === '/pacientes')
      );
      if (patientLinkIndex >= 0) {
        links.splice(patientLinkIndex, 1);
      }
    } else if (typeUser === RoleEnum.User) {
      const userLinkIndex = links.indexOf(
        links.find((el) => el.path === '/usuarios')
      );
      if (userLinkIndex >= 0) {
        links.splice(userLinkIndex, 1);
      }
    }
  }

  if (!authToken && to.path !== '/login') {
    return { path: '/login' };
  } else if (authToken && to.path === '/login') {
    return { path: '/home' }
  }
});

export default router;
