<template>
  <q-page>
    <q-page-sticky position="bottom-right" :offset="[26, 26]">
      <q-fab
        v-model="fab2"
        vertical-actions-align="right"
        color="primary"
        icon="add"
        direction="up"
      >
        <q-fab-action  square color="primary" icon="mail" />
        <q-fab-action  square color="secondary" icon="alarm" />
      </q-fab>
    </q-page-sticky>
  </q-page>
</template>
<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'FloatingActionButton'
})
</script>
