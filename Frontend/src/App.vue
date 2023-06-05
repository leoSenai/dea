<template>
  <HeaderPrimary />
  <main>
    <SideBar :links="links" />
    <router-view class="content" />
  </main>
</template>

<script setup>
import { PhArticle, PhScooter, PhUsers } from '@phosphor-icons/vue';
import HeaderPrimary from './components/HeaderPrimary.vue';
import SideBar from './components/SideBar.vue';

const links = [
  { path: '/pacientes', name: 'Pacientes', icon: PhScooter },
  { path: '/questionarios', name: 'Questionários', icon: PhArticle },
  { path: '/usuarios', name: 'Usuários', icon: PhUsers }
]
</script>

<style scoped>

main {
  display: flex;
}
.content {
  height: calc(100vh - 30%);
  overflow: hidden;
  margin-top: 10rem;
  margin-left: 16rem;
  margin-right: 5rem;
}

</style>
