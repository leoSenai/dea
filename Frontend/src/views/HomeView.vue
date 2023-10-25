<template>
  <div class="home">
    <div class="home-options">
      <button
        class="menu-item"
        @click="goLinkMenu('pacientes')"
      >
        Pacientes<PhScooter weight="bold" />
      </button>
      <button
        class="menu-item"
        @click="goLinkMenu('questionarios')"
      >
        Questionários<PhBookOpen weight="bold" />
      </button>
      <button
        class="menu-item"
        @click="goLinkMenu('usuarios')"
      >
        Usuários<PhPerson weight="bold" />
      </button>
    </div>
  </div>
</template>
<script>
import { PhBookOpen, PhPerson, PhScooter } from '@phosphor-icons/vue';

export default {
  components: { PhScooter, PhBookOpen, PhPerson },
  data() {
    return {};
  },
  methods: {
    goLinkMenu(linktag) {
      this.$router.push('/' + linktag);
    },
  },
};
</script>
<style scoped>
.home {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: url(../assets/imgs/home-background.svg) no-repeat;
  background-position: 90% 80%;
  background-size: 30rem;
}

.home-options {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 30rem;
  border-radius: 4px;
  background: var(--primary-700);
  padding: 3rem 1rem;
} 

.home-options button {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  padding: 1rem;
  border-radius: 4px;
  border: 1px solid white;
  background: var(--primary-700);
  cursor: pointer;
  transition: .2s;
}

.home-options button:hover {
  filter: brightness(0.8);
}
</style>
