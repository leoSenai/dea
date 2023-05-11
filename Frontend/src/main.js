import { createApp } from 'vue'
import { Quasar } from 'quasar'
import router from './router'

import 'quasar/dist/quasar.css'
import './style.css'

import App from './App.vue'

const myApp = createApp(App)

myApp.use(Quasar, {
  plugins: { },
  config: {
    brand: {
      primary: '#519832',
      primary_300: '#6C965A',
      primary_500: '#519832',
      primary_700: '#386923',

      neutral_white: '#FFFFFF',
      neutral_light_gray: '#EBEBEB',
      neutral_gray: '#D9D9D9',
      neutral_dark_gray: '#656565',
      neutral_black: '#000000',
      
      others_red_600: '#B93131',
      others_red_300: '#A44646',

      secondary: '#386923',
      accent: '#6c965a',

      dark: '#1d1d1d',
      dark_page: '#121212',

      positive: '#519832',
      negative: '#b93131',
      info: '#d9d9d9',
      warning: '#F2C037'
    }
  }
}).use(router)

router.isReady().then(() => myApp.mount('#app'))
