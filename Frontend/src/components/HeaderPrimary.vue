<template>
  <header class="bg-primary-700">
    <button
      class="hamburguer"
      @click="activeSidebar"
    >
      <PhList />
    </button>
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
  <div :class="['sidebar', 'bg-success', isSidebarActive ? 'active' : '']">
    <template
      v-for="link in links"
      :key="link.path"
    >
      <router-link
        :to="link.path"
        class="link"
      >
        {{ link.name }}
        <component
          :is="link.icon"
          class="link-icon"
        />
      </router-link>
    </template>
  </div>
</template>

<script>
import { PhUserCircle, PhList } from '@phosphor-icons/vue';

export default {
  components: {
    PhUserCircle,
    PhList,
  },
  props: {
    links: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      LogoSrc: 'src/assets/imgs/Logo.png',
      isSidebarActive: false,
    };
  },
  methods: {
    activeSidebar() {
      this.isSidebarActive = !this.isSidebarActive;
    },
  },
};
</script>

<style scoped>
header {
  position: sticky;
  min-width: 100%;
  z-index: 999;

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

.sidebar {
  position: fixed;
  z-index: 999;
  min-height: 100%;
  padding: 2rem 0rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  transition: 0.3s ease-in-out;
  transform: translateX(0%);
}

.link {
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-width: 100%;
  padding: 1rem 2rem;
  gap: 1rem;
  transition: 200ms ease;
  cursor: pointer;
  color: white;
}

.link:hover {
  background: rgba(0, 0, 0, 0.1);
}

.link-icon {
  font-size: 1.25rem;
}

.hamburguer {
  display: flex;
  font-size: 1.5rem;
  background: none;
  border: none;
  cursor: pointer;
}

.sidebar {
  transform: translateX(-100%);
}

.sidebar.active {
  transform: translateX(0%);
}</style>
