<template>
  <div>
    <modal-primary v-model="valor" />
  </div>
</template>
<script>
import ModalPrimary from '../components/ModalPrimary.vue'
export default {
  components: { ModalPrimary },
  data() {
    return {
      valor: true,
    };
  },
  methods: {
    teste() {
      console.log(this.valor);
    },
  },
};
</script>
