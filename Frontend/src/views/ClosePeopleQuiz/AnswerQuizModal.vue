<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      Responder questionário
    </template>
    <template #modal-content>
      <div
        v-for="question in model.Questions"
        :key="question.id"
        class="questionBox"
      >
        <div class="questionLabelBox">
          {{ question.label }}
        </div>
        <div class="questionAnswerBox">
          <input type="text">
        </div>
      </div>
    </template>
    <template #modal-actions>
      <div class="modal-actions">
        <button-primary
          outlined
          size="sm"
          @click="closeModal"
        >
          Fechar
        </button-primary>
        <button-primary
          size="sm"
          type="submit"
          @click="saveAnswers"
        >
          Cadastrar
        </button-primary>
      </div>
    </template>
  </modal-primary>
</template>
<script>
import ModalPrimary from '../../components/ModalPrimary.vue';
import ButtonPrimary from '../../components/ButtonPrimary.vue';

export default {
  components: {
    ModalPrimary,
    ButtonPrimary,
  },
  emits: ['close', 'error'],
  data() {
    return {
      show: false,
      model: {
        Title: '',
        Questions: []
      }
    }
  },
  methods: {
    async openModal() {
      this.show = true
      this.fetchData()
    },
    closeModal() {
      this.show = false
      this.$emit('close')
    },
    async saveAnswers() {
      try {
        this.$api.QuestionController.insert(this.model.Questions)
        this.closeModal()
      } catch (error) {
        this.$emit('error', error)
      }
    },
    async fetchData () {
        try {
            const response = await this.$api.QuestionController.getByQuizId()
            this.model = response.data.data
        } catch (error) {
            this.$emit('error', error)
        }
    }
  },
};
</script>
<style scoped>
.select {
  width: 100%;
  position: relative;
}
.fill-content {
  width: 100%;
}

.question {
  position: relative;
}

.remove-question {
  position: absolute;
  top: 30%;
  right: 3%;
  width: 1.5rem;
  height: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;

  transform: translateY(-50%);
  color: var(--neutral-dark-gray);
  flex: 1;
  text-align: end;
  cursor: pointer;
  transition: 0.3s;
  background: #fff;
  border: none;
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 99999px;
  cursor: pointer;
  z-index: 1;
}

.remove-question:hover {
  color: var(--others-red-600);
  border-color: var(--others-red-600);
}

.interval-field {
  border: 1px solid black;
  padding: 0.5rem 1rem;
  color: var(--neutral-dark-gray);
  border-radius: 4px;
}

.modal-actions {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
</style>
