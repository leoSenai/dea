<template>
<q-toggle
  v-model="darkMode"
  checked-icon="mdi-weather-night"
  color="green"
  unchecked-icon="mdi-white-balance-sunny" />
</template>
<script>
import { defineComponent, ref, watch } from 'vue'
import { useQuasar } from 'quasar'
export default defineComponent({
  name: 'DarkModeToglle',
  setup () {
    const darkMode = ref(false)
    const $q = useQuasar()
    watch(darkMode, (darkMode) => {
      $q.dark.set(darkMode)
    })
    return {
      darkMode
    }
  }
})

</script>
