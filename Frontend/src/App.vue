<template>
  <router-view name="header" />
  <div class="content">
    <router-view />
  </div>
</template>

<script setup></script>

<style scoped>
.content {
  position: relative;
  display: flex;
  height: calc(100vh - 5rem);
  overflow: hidden;
}

@media (min-width: 768px) {
  .content {
    margin-left: 10rem;
  }
}
</style>
