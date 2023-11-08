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
          <PhPlus color="white" />
        </button>
      </div>
    </div>
    <div
      v-if="model.hasError"
      class="error quiz"
    >
      {{ model.message }}
    </div>
    <div v-if="model.data.Quizzes==undefined">
      <i>Não há questionários criados até o momento.</i>
    </div>
    <div 
      class="quiz-list"
    >
      <div 
        v-for="quiz in model.data.Quizzes"
        :key="quiz.IdQuiz"
        class="row quiz"
      >
        <p @click="openViewModal(quiz)">
          {{ quiz.Name }}
        </p>
        <div class="quiz-actions">
          <button
            type="button"
            @click="openViewModal(quiz)"
          >
            <q-tooltip>
              Visualizar
            </q-tooltip>
            <PhEye color="black" />
          </button>
          <button
            v-show="!finished(quiz)"
            type="button"
            @click="openAddEditModal(quiz)"
          >
            <q-tooltip>
              Editar
            </q-tooltip>
            <PhPencil color="black" />
          </button>
          <button
            type="button"
            @click="openAddQuizPersons(quiz)"
          >
            <q-tooltip>
              Filiados
            </q-tooltip>
            <PhUser color="black" />
          </button>
        </div>
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
        <PhPlus color="white" />
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
    <QuizViewModal
      ref="viewQuiz"
      @close="load"
    />
  </div>
</template>
<script>
import { PhPlus, PhPencil, PhUser, PhEye } from '@phosphor-icons/vue';
import QuizAddEditModal from './QuizAddEditModal.vue';
import QuizPersonsAddEditModal from './QuizPersonsAddEditModal.vue';
import QuizViewModal from './QuizViewModal.vue'

export default {
  components: {
    PhPlus,
    QuizAddEditModal,
    QuizPersonsAddEditModal,
    QuizViewModal,
    PhPencil,
    PhUser,
    PhEye
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
    finished(quiz){
      return this.model.data.FinishedQuizzes.map((el)=>{return el.IdQuiz==quiz.IdQuiz}).includes(true)
    },
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
    openAddQuizPersons(current) {
      this.$refs.addPersons.openModal(current)
    },
    openViewModal(currentQuiz){
      this.$refs.viewQuiz.openModal(currentQuiz)
    }
  },
}
</script>
<style>
.quiz-actions {
  display: flex;
  gap: 0.6em;
  width: 20%;
  justify-content: flex-end;
  padding-right: 1rem;
}

.quiz-content {
  padding: 3rem 1.5rem;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: .75rem;
  background: url(../../assets/imgs/home-background.svg) no-repeat;
  background-position-x:center;
  background-position-y: center;
  height: 100%;
  background-size: 50%;
}

.row{
  background-color: rgba(255, 255, 255, 0.548);
}

.login-form .row:hover{
  background-color: transparent !important;
}

.login-form .row{
  background-color: transparent !important;
}

.row:hover{
  background-color: rgba(255, 255, 255, 0.932) !important;
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
  padding: 1rem 2rem;
  color: white;
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

.quiz-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-height: 70vh;
  height: 100%;
  overflow-y: auto;
}

.quiz {
  border: 1px solid var(--neutral-dark-gray);
  color: var(--neutral-dark-gray);
  padding: 0;
  border-radius: 4px;
  display: flex;
  cursor: text;
  justify-content: space-between;
  align-items: center;
}

.quiz:hover {
  background-color: rgba(200, 255, 172, 0.041);
}

.quiz p {
  margin: 0;
  width: 80%;
  height: 100%;
  padding-left: 1rem;
  align-items: center;
  display: flex;
  padding-top: 1rem;
  padding-bottom: 1rem;
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
  z-index: 1;
}

.quiz button:hover {
  background: var(--neutral-gray);
}

.row.quiz{
  z-index: 0;
  cursor: pointer;
}
</style>