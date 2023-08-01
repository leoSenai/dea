<template>
  <div>
    <button @click="handleClick">
      click
    </button>
  </div>
</template>
<script>
export default {
  data() {
    return {
    };
  },
  methods: {
    handleClick () {
      this.$api.DoctorController.getAll().then(response => {
        console.log(response)
      })
    },

  }
};
</script>
