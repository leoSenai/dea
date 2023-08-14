<template>
  <div class="quiz-content">
    <div
      v-if="model.hasError"
      class="error"
    >
      {{ model.errorMessage }}
    </div>
    <div
      v-else
      class="row"
    >
      Questionario 1
    </div>
    <div class="add-quiz">
      <button
        type="button"
        @click="openAddEditModal()"
      >
        <PhPlus />
      </button>
    </div>
  </div>
  <QuizAddEditModal ref="addEdit" />
</template>
<script>
import { PhPlus } from '@phosphor-icons/vue';
import QuizAddEditModal from './QuizAddEditModal.vue';

export default {
  components: {
    PhPlus,
    QuizAddEditModal
  },
  data() {
    return {
      model: {
        questionaries: [],
        hasError: false,
        errorMessage: ''
      }
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    load() {
      const th = this;
      th.$api.QuizController.getAll().then(data => {
        console.log(data)
      }).catch(({ response }) => {
        th.model = {
          questionaries: [],
          hasError: true,
          errorMessage: response.data.message
        }
      })
    },
    openAddEditModal(current) {
      this.$refs.addEdit.openModal(current)
    }
  },
}
</script>
<style>
.quiz-content {
  padding: 3rem .5rem;
  min-height: 100%;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.error {
  text-align: center;
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
  padding: 1rem;
}

.add-quiz {
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.add-quiz button {
  background: var(--primary);
  border-radius: 99999px;
  font-size: 2rem;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 3rem;
  height: 3rem;
}

.add-quiz button:hover {
  filter: brightness(0.8);
  cursor: pointer;
}
</style>