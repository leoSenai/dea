import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';
import PatientView from '../views/Patient/PatientView.vue'
import Header from '../components/HeaderPrimary.vue';
import Patients from '../views/Patient/PatientsView.vue';
import Proximity from '../views/Proximity/ProximityView.vue';
import PatientQuizzes from '../views/Patient/PatientQuizzes.vue'
import ProximityQuizzes from '../views/Proximity/ProximityQuizzes.vue' 
import Users from '../views/Users/UsersView.vue'
import InputPrimary from '../components/InputPrimary.vue';
import NotFound from '../views/NotFoundView.vue';
import Quiz from '../views/Quiz/QuizView.vue';
import Cookie from '../cookie'

import { PhUserList, PhUsers } from '@phosphor-icons/vue';
import { PhArticle } from '@phosphor-icons/vue';

const typeUser = Cookie.getUserType(Cookie.get('authToken'))

function getLinks(){
  if(typeUser=='U'){
    return [
      { path: '/usuarios', name: 'Usuários', icon: PhUsers },
      { path: '/questionarios', name: 'Questionários', icon: PhArticle },
      { path: '/pacientes', name: 'Pacientes', icon: PhUserList}
    ];
  }else if(typeUser=='PA'){
    return [
      { path: '/questionarios-paciente', name: 'Ver questionários', icon: PhArticle },
    ];
  }else{ //PR
    return [
      { path: '/questionarios-pessoa-proxima', name: 'Ver questionários', icon: PhArticle },
    ];
  }
}

const links = getLinks()

const routesU = [
  {
    path: '/',
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
    path: '/paciente',
    components: {
      default: PatientView,
      header: Header,
    },
    props: {
      header: { links },
    }
  },
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
    }
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
    path: '/pessoas-proximas/:id/questionarios',
    components: {
      default: ProximityQuizzes,
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
    },
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

const routesPA = [
  {
    path: '/',
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
    path: '/login',
    components: {
      default: Login,
    },
  },
  {
    path: '/questionarios-paciente',
    components: {
      default: PatientQuizzes,
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

const routesPR = [
  {
    path: '/',
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
    path: '/login',
    components: {
      default: Login,
    },
  },
  {
    path: '/questionarios-pessoa-proxima',
    components: {
      default: Quiz,
      header: Header,
    },
    props: {
      header: { links },
    },
  },
  {
    path: '/pessoa-proxima/:id/questionarios',
    components: {
      default: PatientQuizzes,
      header: Header,
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

function getRouter(){
  console.log(typeUser)
  if(typeUser=='U'){
    return routesU
  }else if(typeUser=='PA'){
    return routesPA
  }else{
    return routesPR
  }
}

export const routes = getRouter()