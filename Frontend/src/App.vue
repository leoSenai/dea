<template>
  <Header :links="links" />
  <router-view />
</template>

<script setup>
import Header from './components/Header.vue';
import { PhUsers } from '@phosphor-icons/vue';

const links = [
  // { path: '/pacientes', name: 'Pacientes', icon: PhScooter },
  // { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  { path: '/usuarios', name: 'Usuários', icon: PhUsers }
]
</script>

<style scoped></style>
