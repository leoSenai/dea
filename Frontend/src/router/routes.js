import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import Header from '../components/HeaderPrimary.vue';
import Quiz from '../views/quiz/QuizView.vue';
import Proximity from '../views/Proximity/ProximityView.vue';
import Users from '../views/users/UsersView.vue'
import NotFound from '../views/NotFoundView.vue';
import { PhUsers, PhArticle } from '@phosphor-icons/vue';
//import Cookie from '../cookie';

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
    path: '/questionarios',
    components: {
      default: Quiz,
      header: Header,
    },
    props: {
      header: { links },
    },
  },
  {
    path: '/paciente/:id/pessoas-proximas',
    components: {
      default: Proximity,
      header: Header,
    },
    props: {
      header: { links },
    },
  },
  {
    path: '/usuarios',
    components: {
      default: Users,
      header: Header,
    },
    props: {
      header: { links },
    }
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
