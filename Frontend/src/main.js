import { createApp } from 'vue';
import { Quasar } from 'quasar';
import router from './router';
import 'toastify-js/src/toastify.css';

import 'quasar/dist/quasar.css';
import './style.css';

import { colorTheme } from './assets/styles/theme';

import api from './api';
import App from './App.vue';

const myApp = createApp(App);

myApp
  .use(Quasar, {
    plugins: {},
    config: {
      brand: { ...colorTheme },
    },
  })
  .use(router)
  .use(api);

router.isReady().then(() => myApp.mount('#app'));
