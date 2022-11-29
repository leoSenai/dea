<template>
  <q-btn-dropdown flat color="white" icon="person">
    <q-list>
      <q-item clickable v-close-popup>
        <q-item-section>
          <q-item-label>Configurações</q-item-label>
        </q-item-section>
      </q-item>
      <q-item clickable  :to="LoginUsuario ">
        <q-item-section>
          <q-item-label>Logout</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
  </q-btn-dropdown>
</template>
<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'DropDownLogout',
  setup () {
    return {
      LoginUsuario: { name: 'LoginUsuario' }
    }
  }
})
</script>
