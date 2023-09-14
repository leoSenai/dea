<template>
  <div class="quiz-content">
    <div
      v-if="model.hasError"
      class="error quiz"
    >
      {{ model.message }}
    </div>
    <div
      v-for="quiz in model.data"
      v-else
      :key="quiz.IdQuiz"
      class="row quiz"
    >
      <p>
        {{ quiz.Name }}
      </p>
      <div class="quiz-actions">
        <button
          type="button"
          @click="openAddEditModal(quiz)"
        >
          <PhPencil />
        </button>
      </div>
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
  <QuizAddEditModal
    ref="addEdit"
    @close="load"
  />
</template>
<script>
import { PhPlus, PhPencil } from '@phosphor-icons/vue';
import QuizAddEditModal from './QuizAddEditModal.vue';

export default {
  components: {
    PhPlus,
    QuizAddEditModal,
    PhPencil
  },
  data() {
    return {
      model: {
        data: [],
        hasError: false,
        message: ''
      },
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    load() {
      const th = this;
      th.$api.QuizController.getAll().then(({ data }) => {
        th.model = data
      }).catch(({ response }) => {
        th.model = {
          ...response.data,
          hasError: true
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
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: .75rem;
}

.error {
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
  transition: .2s;
}

.add-quiz button:hover {
  filter: brightness(0.8);
  cursor: pointer;
}

.quiz {
  border: 1px solid;
  padding: 1rem;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.quiz p {
  margin: 0;
}

.quiz button {
  border: none;
  background: none;
  cursor: pointer;
  padding: .5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: .1s;
  border-radius: 9999px;
}

.quiz button:hover {
  background: var(--neutral-dark-gray);
}
</style>