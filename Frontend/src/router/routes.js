import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import PatientView from '../views/PatientView.vue'
import Header from '../components/HeaderPrimary.vue';
import Patients from '../views/PatientsView.vue';
import Proximity from '../views/Proximity/ProximityView.vue';
import Users from '../views/users/UsersView.vue'
import InputPrimary from '../components/InputPrimary.vue';
import NotFound from '../views/NotFoundView.vue';
import Quiz from '../views/quiz/QuizView.vue';

import { PhUserList, PhUsers } from '@phosphor-icons/vue';
import { PhArticle } from '@phosphor-icons/vue';

const links = [
  { path: '/usuarios', name: 'Usuários', icon: PhUsers },
  { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  { path: '/patients', name: 'Pacientes', icon: PhUserList}
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
    path: '/patientView',
    components: {
      default: PatientView,
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
