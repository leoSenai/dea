import Home from '../views/Home.vue'
import Login from '../views/Login.vue'

export const routes = [
  { path: '/', component: Home, name: 'Home' },
  { path: '/home', component: Home, name: 'Home' },
  { path: '/login', component: Login, name: 'Login' }
]