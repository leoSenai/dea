<template>
  <q-checkbox color="primary"/>
</template>
<script>
export default {

}
</script>
<style >

.q-checkbox__inner--truthy .q-checkbox__bg{
  border-color: var(--primary) !important;
}

@media (prefers-color-scheme: dark) {
  .q-checkbox__bg {
    border-color: var(--neutral-dark-gray) !important;
  }
}

</style>
