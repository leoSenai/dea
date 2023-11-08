import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import PatientView from '../views/Patient/PatientView.vue'
import Header from '../components/HeaderPrimary.vue';
import Patients from '../views/Patient/PatientsView.vue';
import Proximity from '../views/Proximity/ProximityView.vue';
import PatientQuizzes from '../views/Patient/PatientQuizzes.vue'
import Users from '../views/Users/UsersView.vue'
import InputPrimary from '../components/InputPrimary.vue';
import NotFound from '../views/NotFoundView.vue';
import Quiz from '../views/Quiz/QuizView.vue';
import Cookie from '../cookie'

import { PhUserList, PhUsers } from '@phosphor-icons/vue';
import { PhArticle } from '@phosphor-icons/vue';

const authToken = Cookie.get('authToken')
const userType = authToken ? Cookie.getUserType(authToken) : ''

const closedPeopleLinks = [
  { path: '/questionarios', name: 'Questionários', icon: PhArticle }
]

const allUserLinks = [
  { path: '/usuarios', name: 'Usuários', icon: PhUsers },
  { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  { path: '/pacientes', name: 'Pacientes', icon: PhUserList}
]

const links = userType === 'P' ? closedPeopleLinks : allUserLinks

export const routes = [
  {
    path: '/home',
    components: {
      default: Home,
      header: Header,
    },
    props: {
      header: { links },
      default: { links }
    },
  },
  {
    path: '/pacienteInfo',
    components: {
      default: PatientView,
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
    path: '/paciente/:id/questionarios',
    components: {
      default: PatientQuizzes,
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
    path: '/pacientes',
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
