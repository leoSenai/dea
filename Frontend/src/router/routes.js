import Home from '../views/HomeView.vue';
import Quiz from '../views/QuizView.vue';
import Header from '../components/HeaderPrimary.vue';

import { PhArticle, PhUsers } from '@phosphor-icons/vue';

const links = [
  // { path: '/pacientes', name: 'Pacientes', icon: PhScooter },
  { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  { path: '/usuarios', name: 'Usuários', icon: PhUsers },
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
    path: '/questionarios',
    components: {
      default: Quiz,
      header: Header,
    },
    props: {
      header: { links },
    },
  },
];
