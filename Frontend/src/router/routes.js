import Home from '../views/HomeView.vue';
import Header from '../components/HeaderPrimary.vue';
import Header2 from '../components/HeaderPrimary2.vue';
import Patients from '../views/PatientsView.vue';

import { PhUsers } from '@phosphor-icons/vue';

const links = [
  // { path: '/pacientes', name: 'Pacientes', icon: PhScooter },
  // { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  {path: '/patients', name: 'Pacientes', icon: PhUsers},
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
    path: '/patients',
    components: {
      default: Patients,
      header: Header2,
    },
    props: {
      header: { links },
    },
  }
];
