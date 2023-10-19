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
    :style="{ background: `url(${Background}) no-repeat bottom right` }"
  />
</template>
<script>
import Background from '/404-background.svg'
import ButtonPrimary from '../components/ButtonPrimary.vue'
export default {
  components: {
    ButtonPrimary
  },
  data() {
    return {
      Background
    }
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
}
</style>