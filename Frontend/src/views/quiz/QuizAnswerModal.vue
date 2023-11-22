<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      {{ isAnswered ? 'Responder' : 'Visualizar' }} Questionário: {{ Quiz.Name }}
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div
          v-for="(question, i) in model.Questions"
          :key="question.key"
          class="question row q-mb-sm"
        >
          <question-primary
            v-model="model.Answers[i]"
            :question-number="i + 1"
            :is-answered="!isAnswered"
            :interval="Quiz.Interval"
          >
            {{ question.Desc }}
          </question-primary>
        </div>
      </q-form>
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
          v-if="isAnswered"
          size="sm"
          type="submit"
          @click="saveAnswers"
        >
          Responder
        </button-primary>
      </div>
    </template>
  </modal-primary>
</template>
<script>
import ModalPrimary from '../../components/ModalPrimary.vue';
import ButtonPrimary from '../../components/ButtonPrimary.vue';
import QuestionPrimary from '../../components/QuestionPrimary.vue';
import Cookie from '../../utils/cookie'
import { RoleEnum } from '../../utils/Enum'

export default {
  components: {
    ModalPrimary,
    ButtonPrimary,
    QuestionPrimary
  },
  emits: ['close', 'error'],
  data() {
    return {
      show: false,
      model: {
        Title: '',
        Questions: [],
        Answers: []
      },
      aaa: 0,
      Quiz: {}
    }
  },
  computed: {
    typeUser () {
      const authToken = Cookie.get('authToken');
      return Cookie.getUserType(authToken);
    },
    isAnswered () {
      return this.Quiz.Name.includes('EM ABERTO')
    }
  },
  methods: {
    openModal(currentQuiz) {
      this.show = true
      this.loadQuiz(currentQuiz)
    },
    closeModal() {
      this.show = false
      this.$emit('close')
      this.model = {
        Title: '',
        Questions: []
      }
    },
    loadQuiz(currentQuiz) {
      const th = this
      th.Quiz = currentQuiz
      th.$api.QuestionController.getAll()
        .then(({ data }) => {
          const filteredQuestions = data.data.filter(el => el.IdQuiz === currentQuiz.IdQuiz)
          th.model.Questions = filteredQuestions
            if (th.typeUser === RoleEnum.Patient) {
              th.$api.PatientHasQuizController.getByIdQuizPatient(th.Quiz.IdQuiz, th.$route.params.id).then(response => {
                response.data.data[0].Answers.split(',').forEach((el, i) => {
                  th.model.Answers[i] = el
                })
              })
          }
        })
    },
    saveAnswers() {
      const th = this

      const dto = {
        IdPatient: +th.$route.params.id,
        IdQuiz: th.Quiz.IdQuiz,
        Finished: th.model.Questions.length === th.model.Answers.length ? 1 : 0,
        Answers: th.model.Answers.toString(),
        AnswerdIn: new Date()
      }

      if (th.typeUser === RoleEnum.Person) {
        th.$api.ProximityHasQuizController.update(dto).then(response => {
          if (response) th.closeModal()
        })
      } else if (th.typeUser === RoleEnum.Patient) {
        th.$api.PatientHasQuizController.update(dto).then(response => {
          if (response) th.closeModal()
        })
      }
    }
  }
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
