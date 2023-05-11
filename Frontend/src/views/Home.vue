<template></template>
<script>
import Header from "../components/Header.vue";
import Input from "../components/Input.vue";

export default {
  data() {
    return {
      valor: "",
    };
  },
  components: { Header, Input },
  methods: {
    teste() {
      console.log(this.valor);
    },
  },
};
</script>
