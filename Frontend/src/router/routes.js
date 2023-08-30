import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import Header from '../components/HeaderPrimary.vue';
import Quiz from '../views/QuizView.vue';
import { PhAppWindow, PhUsers, PhArticle } from '@phosphor-icons/vue';

const links = [
  { path: '/usuarios', name: 'Usuários', icon: PhUsers },
  { path: '/login', name: 'Login', icon: PhAppWindow },
  { path: '/questionarios', name: 'Questionários', icon: PhArticle },
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
