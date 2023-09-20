<template>
  <div class="text">
    <span class="heading-1">
      Ooops!
    </span>
    <span class="heading-6">
      Parece que a página que procurou não existe
    </span>
    <div>
      <ButtonPrimary @click="goHome">
        Voltar à página inicial
      </ButtonPrimary>
    </div>
  </div>
  <div
    v-if="!isMobile"
    class="background-404"
  />
</template>
<script>
import ButtonPrimary from '../components/ButtonPrimary.vue'
export default {
  components: {
    ButtonPrimary
  },
  computed: {
    isMobile() {
      return this.$q.screen.xs || this.$q.screen.sm
    }
  },
  methods: {
    goHome() {
      this.$router.push('/')
    }
  },
}
</script>
<style scoped>
.text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 5rem;
  gap: 2rem;
  height: 100%;
}

.background-404 {
  content: '';
  position: absolute;
  width: 35rem;
  height: 100%;
  bottom: 5rem;
  right: 5rem;
  background: url('../assets/imgs/404-background.svg') no-repeat bottom right;
}
</style>