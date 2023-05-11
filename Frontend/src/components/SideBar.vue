<template>
  <div class="sidebar bg-success">
    <template v-for="link in links" :key="link.path">
      <router-link :to="link.path" class="link">
        {{ link.name }}
        <i :class="['ph', `ph-${link.icon}`, 'link-icon']"></i>
      </router-link>
    </template>
  </div>
</template>
<script>
export default {
  props: {
    links: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {};
  },
  components: {
    PhScooter: () => import("@phosphor-icons/vue"),
  },
};
</script>
<style scoped>
.sidebar {
  position: fixed;
  min-height: 100%;
  padding: 2rem 0rem;

  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
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
</style>
