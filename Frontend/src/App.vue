<template>
  <div>
    <router-view name="header" />
    <div class="content">
      <router-view />
    </div>
  </div>
</template>

<script setup></script>

<style scoped>

main {
  display: flex;
}
.content {
  position: relative;
  height: calc(100vh - 30%);
  overflow: hidden;
}

@media (min-width: 768px) {
  .content {
    margin-left: 10rem;
  }
}

</style>
