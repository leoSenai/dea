import { createApp } from 'vue';
import { Quasar } from 'quasar';
import router from './router';
import 'toastify-js/src/toastify.css';

import 'quasar/dist/quasar.css';
import './style.css';

import { colorTheme } from './assets/styles/theme';

import cookies from 'vue3-cookies';

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
  .use(api)
  .use(cookies, {
    expireTimes: '30d',
    sameSite: 'None',
    secure: true
  });

router.isReady().then(() => myApp.mount('#app'));
