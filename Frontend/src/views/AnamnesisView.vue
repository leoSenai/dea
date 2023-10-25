<template>
  <div>
    <q-editor
      v-model="anamneseContent"
      min-height="5rem"
      toolbar-text-color="secondary"
      toolbar-toggle-color="light-green"
      :style="{ color: textColor }"
    />
  </div>
</template>

<script>
import anamneseController from '../api/AnamneseController'
export default {
  data() {
    return {
      anamneseContent: '',
      textColor: 'black'
    }
  },
  watch: {
    anamneseContent() {
      this.saveAnamnese()
    }
  },
  async mounted () {
    await this.fetchAnamnese()
  },
  methods: {
    async fetchAnamnese () {
      const nelson = await anamneseController.getByUserIdAndPatientId()
      console.log('anamnesefetch', nelson)
      this.anamneseContent = nelson
    },
    async saveAnamnese () {
      await anamneseController.update({
        content: this.anamneseContent,
        userId: this.$store.state.user.id,
        patientId: this.$store.state.patient.id
      })
    }
  }
}
</script>
