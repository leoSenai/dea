import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import Header from '../components/HeaderPrimary.vue';

import { PhAppWindow, PhUsers } from '@phosphor-icons/vue';

const links = [
  // { path: '/pacientes', name: 'Pacientes', icon: PhScooter },
  // { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  {path: '/usuarios', name: 'Usuários', icon: PhUsers},
  {path: '/login', name: 'Login', icon: PhAppWindow},
];

export const routes = [
  {
    path: '/',
    components: {
      default: Home,
      header: Header,
    },
    props: {
      header: { links },
    },
  },
  {
    path: '/home',
    components: {
    default: Home,
    header: Header,
    },
    props: {
    header: { links },
    },  
  },
  {
    path: '/login',
    components: {
      default: Login,
    },
    props: {
      header: { links },
    },
  },
];
