<template>
  <div class="quiz-content">
    <div class="quiz-title">
      <div class="title">
        <h3>Questionários</h3>
      </div>
      <div
        v-if="!isMobile"
        class="title-add-quiz"
      >
        <button
          type="button"
          @click="openAddEditModal()"
        >
          Adicionar
          <PhPlus />
        </button>
      </div>
    </div>
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
        <button
          type="button"
          @click="openAddQuizPersons(quiz)"
        >
          <PhUser />
        </button>
      </div>
    </div>
    <div
      v-if="isMobile"
      class="add-quiz"
    >
      <button
        type="button"
        @click="openAddEditModal()"
      >
        <PhPlus />
      </button>
    </div>
    <QuizAddEditModal
      ref="addEdit"
      @close="load"
    />
    <QuizPersonsAddEditModal
      ref="addPersons"
      @close="load"
    />
  </div>
</template>
<script>
import { PhPlus, PhPencil, PhUser } from '@phosphor-icons/vue';
import QuizAddEditModal from './QuizAddEditModal.vue';
import QuizPersonsAddEditModal from './QuizPersonsAddEditModal.vue';

export default {
  components: {
    PhPlus,
    QuizAddEditModal,
    QuizPersonsAddEditModal,
    PhPencil,
    PhUser
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
  computed: {
    isMobile() {
      return this.$q.screen.xs || this.$q.screen.sm
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
    },
    openAddQuizPersons(current){
      this.$refs.addPersons.openModal(current)
    }
  },
}
</script>
<style>
.quiz-actions{
  display: flex;
  gap: 0.6em;
}
.quiz-content {
  padding: 3rem 1.5rem;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: .75rem;
}

.quiz-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.title-add-quiz button {
  background: var(--primary);
  border-radius: 8px;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: .2s;
  gap: .5rem;
  padding: .5rem 1rem;
  cursor: pointer;
}

.title-add-quiz button:hover {
  filter: brightness(0.8);
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
}

.quiz {
  border: 1px solid;
  padding: 1rem;
  border-radius: 4px;
  display: flex;
  cursor:text;
  justify-content: space-between;
  align-items: center;
}

.quiz:hover{
  background-color: rgba(200, 255, 172, 0.041);
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