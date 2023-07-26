<template>
  <div class="q-pa-md row items-start q-gutter-md">
    <PatientCard />
    <PatientCard />
    <PatientCard />
    <PatientCard />
  </div>
</template>

<script>
import PatientCard from '../components/Cards/PatientCard.vue'

export default {
  name: 'PatientsView',
  components: {
    PatientCard
  },
  setup () {
    return {
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
    }
  },
  methods: {
    teste() {
      console.log(this.valor);
    },
  },
};
</script>
