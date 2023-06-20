<template>
  <div class="container">
    <ButtonPrimary
      outlined
      class="btn-link"
    >
      <router-link
        to="/pacientes"
        class="link"
      >
        Pacientes
      </router-link>
      <template #after-label>
        <PhScooter size="20" />
      </template>
    </ButtonPrimary>
    <ButtonPrimary
      outlined
      class="btn-link"
    >
      <router-link
        to="/questionarios"
        class="link"
      >
        Questionários
      </router-link>
      <template #after-label>
        <PhClipboard size="20" />
      </template>
    </ButtonPrimary>
    <ButtonPrimary
      outlined
      class="btn-link"
    >
      <router-link
        to="/usuarios"
        class="link"
      >
        Usuários
      </router-link>
      <template #after-label>
        <PhUsers size="20" />
      </template>
    </ButtonPrimary>
  </div>
</template>
<script>
import ButtonPrimary from '../components/ButtonPrimary.vue';
import { PhUsers, PhScooter, PhClipboard } from '@phosphor-icons/vue';
export default {
    components: {
    ButtonPrimary,
    PhUsers,
    PhScooter,
    PhClipboard
}
};
</script>
<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 50vh;
}

.btn-link {
  width: 50%;
  margin-bottom: 1rem;
}

.btn-link a.link {
  font-size: 1.25rem;
  color: var(--primary-500);
}
</style>
