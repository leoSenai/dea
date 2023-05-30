<template>
  <Select v-model="valor" label="Label" :options="['google', 'twitter']" />
</template>
<script>
import Header from "../components/Header.vue";
import Select from "../components/Select.vue";

export default {
  data() {
    return {
      valor: "",
    };
  },
  components: { Header, Select },
  methods: {
    teste() {
      console.log(this.valor);
    },
  },
};
</script>
