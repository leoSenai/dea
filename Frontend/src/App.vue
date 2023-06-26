<template>
  <router-view name="header" />
  <router-view />
</template>

<script setup></script>

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
