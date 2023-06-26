<template>
  <div>
    <div>
      <SelectPrimary
        v-model="valor"
        label="Label"
        :options="['google', 'twitter']"
      />
    </div>
    <div><HeaderPrimary /></div>
  </div>
</template>
<script>
import HeaderPrimary from '../components/HeaderPrimary.vue';
import SelectPrimary from '../components/SelectPrimary.vue';

export default {
  components: { HeaderPrimary, SelectPrimary },
  data() {
    return {
      valor: '',
    };
  },
  methods: {
    teste() {
      console.log(this.valor);
    },
  },
};
</script>
