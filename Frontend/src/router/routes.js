import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import Header from '../components/HeaderPrimary.vue';
import Patients from '../views/PatientsView.vue';
import InputPrimary from '../components/InputPrimary.vue';

import { PhAppWindow, PhUsers } from '@phosphor-icons/vue';

const links = [
  // { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  {path: '/patients', name: 'Patients', icon: PhAppWindow},
  {path: '/usuarios', name: 'Usuários', icon: PhUsers},
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
  {
    path: '/patients',
    components: {
      default: Patients,
      header: Header,
      input: InputPrimary
    },
    props: {
      header: { links },
    },
  },
];
