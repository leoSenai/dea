<template>
  <div>
    <router-view name="header" />
    <div :class="['content', isLoginScreen ? 'login-screen' : '']">
      <router-view />
    </div>
  </div>
</template>

<script setup>
const isLoginScreen = location.href.includes('login')
</script>

<style scoped>
.content {
  position: relative;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  height: calc(100vh - 5rem);
}

.content.login-screen {
  height: 100vh;
}
</style>
