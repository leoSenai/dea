<template>
  <header class="bg-primary-700">
    <div class="logo white">
      <q-img :src="LogoSrc" />
      <h6>Clínica Motivar</h6>
    </div>
    <div class="user white">
      <ph-user-circle
        class="user-icon"
        regular
      />
      <span class="text-body">Usuário</span>
    </div>
  </header>
</template>

<script>
import { PhUserCircle } from '@phosphor-icons/vue';

export default {
  components: {
    PhUserCircle,
  },
  data() {
    return {
      LogoSrc: '/src/assets/imgs/Logo.png',
    };
  },
};
</script>

<style scoped>
header {
  position: fixed;
  min-width: 100%;
  z-index: 1;

  display: flex;
  align-items: center;
  justify-content: space-between;

  padding: 0.5rem 1rem;
}

.logo {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 2rem;
}

.user {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-icon {
  height: 100%;
  width: 1.5rem;
}

.q-img {
  height: 100%;
  width: 4rem;
}
</style>
