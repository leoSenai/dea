<template>
  <q-checkbox color="primary" :label="label"/>
</template>
<script>
export default {
  props: {
    label: {
      type: String,
      required: true,
    }
  }
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
