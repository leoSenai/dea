<template>
  <div class="quiz-content">
    <div class="quiz-title">
      <div class="title">
        <h3>Questionários</h3>
      </div>
    </div>
    <div
      v-if="model.hasError"
      class="error quiz"
    >
      {{ model.message }}
    </div>
    <div v-if="model.data.length==0">
      <i>Não há questionários cadastrados até o momento.</i>
    </div>
    <div
      v-else
      class="quiz-list"
    >
      <div
        v-for="quiz in model.data"
        :key="quiz.Id"
        class="row quiz"
      >
        <p @click="openViewModal(quiz)">
          {{ quiz.Title }}
        </p>
        <div class="quiz-actions">
          <button
            type="button"
            @click="openAnswerQuizModal(user)"
          >
            <q-tooltip>
              Editar
            </q-tooltip>
            <PhPencil color="black" />
          </button>
        </div>
      </div>
    </div>

    <AnswerQuizModal
      ref="answerModal"
      @close="load"
      @error="error"
    />
  </div>
</template>
<script>
import { PhPencil } from '@phosphor-icons/vue'
import AnswerQuizModal from './AnswerQuizModal.vue'

export default {
  components: {
    PhPencil,
    AnswerQuizModal
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
    openViewModal(currentUser){
      this.$refs.viewUser.openModal(currentUser)
    },
    fetchData () {
      const response = this.$api.QuizController.getAll()
      this.model.data = response.data
    },
    openAnswerQuizModal(current) {
      this.$refs.answerModal.openModal(current)
    },
    error (errorMessage) {
        this.model.hasError = true
        this.model.message = errorMessage
    }
  },
}
</script>
<style scoped>
.quiz-content {
  padding: 3rem 1.5rem;
  width: 100%;
  display: flex;
  background: url(../../assets/imgs/home-background.svg) no-repeat;
  flex-direction: column;
  background-position-x:center;
  background-position-y: center;
  background-size: 50%;
  height: 100%;
  gap: .75rem;
}

.row{
  background-color: rgba(255, 255, 255, 0.548);
}

.quiz-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.row.user{
  padding: 0;
  width: 100%;
  display: flex;
  cursor: pointer;
}

.row.user p{
  padding: 1em;
  width: -webkit-fill-available;
}

.row.user .user-actions{
  position: auto;
  display: flex;
  padding-right: 1em;
  width: auto;
  height: auto;
  margin-left: -60%;
}

.title-add-user button {
  background: var(--primary);
  border-radius: 8px;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: .2s;
  gap: .5rem;
  padding: 1rem 2rem;
  color: white;
  cursor: pointer;
}

.title-add-user button:hover {
  filter: brightness(0.8);
}

.error {
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
  padding: 1rem;
}

.add-user {
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.add-user button {
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

.add-user button:hover {
  filter: brightness(0.8);
  cursor: pointer;
}

.user-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-height: 70vh;
  height: 100%;
  overflow-y: auto;
}

.user {
  border: 1px solid var(--neutral-dark-gray);
  color: var(--neutral-dark-gray);
  padding: 0.8rem;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.user p {
  margin: 0;
}

.user button {
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

.user button:hover {
  background: var(--neutral-gray);
}

.user-actions{
  display: flex;
  gap: 1em;
}
</style>