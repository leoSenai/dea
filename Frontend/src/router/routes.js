import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import Header from '../components/HeaderPrimary.vue';
import Patients from '../views/PatientsView.vue';
import InputPrimary from '../components/InputPrimary.vue';
import NotFound from '../views/NotFoundView.vue';
import { PhUsers } from '@phosphor-icons/vue';
import { PhArticle } from '@phosphor-icons/vue';

const links = [
  { path: '/usuarios', name: 'Usuários', icon: PhUsers },
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
  {
    path: '/:pathMatch(.*)*',
    components: {
      default: NotFound,
      header: Header,
      input: InputPrimary
    },
    props: {
      header: { links },
    },
  },
  {
    path: '/:pathMatch(.*)*',
    components: {
      default: NotFound,
      header: Header,
    },
    props: {
      header: { links },
    },
  },
];
