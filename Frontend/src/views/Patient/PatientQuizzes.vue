<template>
  <div class="patient-quiz-content">
    <div
      v-if="typeUser !== RoleEnum.Patient"
      class="back-page"
      onclick="window.history.back()"
    >
      <PhCaretLeft color="#656565" />
      Voltar
    </div>
    <div class="quiz-title">
      <div class="title">
        <h3>Questionários de {{ patient.Name }}</h3>
      </div>
    </div>
    <div
      v-if="model.hasError"
      class="error quiz"
    >
      {{ model.message }}
    </div>
    <div v-if="model.data.length==0">
      <i>Não há questionários vinculados a este paciente.</i>
    </div>
    <div 
      v-for="quiz in model.data"
      v-else
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
          v-if="quiz.Name.includes('EM ABERTO') && typeUser === RoleEnum.Patient"
          type="button"
          @click="answerQuiz(quiz)"
        >
          <q-tooltip>
            Responder
          </q-tooltip>
          <PhArticle color="black" />
        </button>
      </div>
    </div>
    <QuizAddEditModal
      ref="addEdit"
      @close="load"
    />
    <QuizViewModal ref="viewQuiz" />
    <QuizAnswerModal 
      ref="answerQuiz" 
      @close="load"
    />
  </div>
</template>
<script>
import { PhEye, PhCaretLeft, PhArticle } from '@phosphor-icons/vue';
import cookie from '../../utils/cookie';
import { RoleEnum } from '../../utils/Enum';
import QuizAnswerModal from '../quiz/QuizAnswerModal.vue'
import QuizViewModal from './PatientQuizViewModal.vue'

export default {
  components: {
    PhCaretLeft,
    QuizViewModal,
    PhEye,
    PhArticle,
    QuizAnswerModal
  },
  data() {
    return {
      model: {
        data: [],
        hasError: false,
        message: '',
      },
      quizzes: [],
      patient: [],
      RoleEnum
    }
  },
  computed: {
    isMobile() {
      return this.$q.screen.xs || this.$q.screen.sm
    },
    typeUser () {
      const token = cookie.get('authToken')
      return cookie.getUserType(token)
    },
    idUser() {
      const token = cookie.get('authToken')
      return cookie.getUserId(token)
    },
    ...RoleEnum
  },
  mounted() {
    this.load();
    this.fixScreenSize()
  },
  methods: {
    load() {
      const th = this;
      var idPatient = this.$router.currentRoute.value.params.id;

      if(th.idUser!=idPatient){
        alert('Você não tem permissão para acessar essa tela!')
        th.$api.AuthController.logout();
        window.history.back()
        return
      }

      th.$api.PatientController.getById(idPatient).then((data) => {
        th.patient = data.data.data
      })

      th.$api.PatientHasQuizController.getByIdPatient(idPatient).then(async ({ data }) => {

        if (data.data) {

          th.quizzes = data.data.map((item) => {
            return { IdQuiz: item.IdQuiz, Finished: item.Finished };
          })
          th.model.data = []
          for (var id in th.quizzes) {

            var result = await th.$api.QuizController.getById(th.quizzes[id].IdQuiz)
            th.loadQuiz(result, id)
          }
        }

      })
    },
    fixScreenSize () {
      try{
        const contentLoginScreen = document.body.getElementsByClassName('login-screen')
        if(contentLoginScreen.length != 0){
          contentLoginScreen[0].classList.remove('login-screen')
        }
      } finally {
        //
      }
    },
    loadQuiz(data, id) {
      const th = this
      var quizFinished = th.quizzes[id].Finished
      if (quizFinished == 0) {
        data.data.data.Name += ' - EM ABERTO'
      } else {
        data.data.data.Name += ' - RESPONDIDO'
      }
      th.model.data.push(data.data.data)
    },
    openViewModal(currentQuiz) {
      this.$refs.viewQuiz.openModal(currentQuiz, this.patient)
    },
    answerQuiz (currentQuiz) {
      this.$refs.answerQuiz.openModal(currentQuiz)
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

.card .row{
  background-color: transparent;
}

.card div.row:hover{
  background-color: transparent !important;
}

.patient-quiz-content {
  padding: 3rem 1.5rem;
  padding-top: 2rem;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: .75rem;
  background-size: 50% !important;
  background: url(../../assets/imgs/home-background.svg) no-repeat;
  background-position-x:center;
  background-position-y: center;
  height: 100%;
}

.back-page {
  display: flex;
  align-items: center;
  margin-top: 1.5rem;
  margin-left: 0.5rem;
  cursor: pointer;
  transition: 1.5s;
}

.quiz-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.error {
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
  padding: 1rem;
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

  .back-page {
    display: flex;
    align-items: center;
    margin-top: 1.5rem;
    margin-left: 0.5rem;
    cursor: pointer;
    width: fit-content;
    transition: 1.5s;
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
  .add-questionary {
    padding: .75rem 1rem !important;
    font-size: .875rem;  }
</style>
