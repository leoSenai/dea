<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      Vizualizar questionário de {{ patient.Name }}
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div class="row q-mb-sm">
          <input-primary
            v-model="model.Name"
            :disable="true"
            label="Nome do Questionário"
            label-color="primary"
            required
          />
        </div>
        <div class="row q-mb-sm interval-field">
          <span>Intervalo de resposta: {{ model.Interval }}</span>
          <q-slider
            v-model="model.Interval"
            :disable="true"
            :min="3"
            :max="7"
            :step="1"
            label
            color="primary"
          />
        </div>
        <div
          v-for="(question, i) in questions"
          :key="question.key"
          class="question row q-mb-sm"
        >
          <div
            v-if="question.Answer"
            style="display: flex; width: 100%;gap: 5px;"
          >  
            <QuestionPrimary
              :model-value="parseInt(question.Answer)"
              :answer-range="model.Interval"
              :question-number="i + 1"
              :is-answered="true"
            >
              <i>{{ question.Desc }}</i>
            </QuestionPrimary>
          </div>
          <input-primary
            v-else
            v-model="question.Desc"
            :disable="true"
            :name="`${i}`"
            :label="`Pergunta ${i + 1}`"
            label-color="primary"
          />
        </div>
      </q-form>
    </template>
    <template #modal-actions>
      <div class="modal-actions">
        <button-primary
          outlined
          class="fill-content"
          size="sm"
          @click="closeModal"
        >
          Fechar
        </button-primary>
      </div>
    </template>
  </modal-primary>
</template>
  <script>
import ModalPrimary from '../../components/ModalPrimary.vue';
import InputPrimary from '../../components/InputPrimary.vue';
import ButtonPrimary from '../../components/ButtonPrimary.vue';
import QuestionPrimary from '../../components/QuestionPrimary.vue';

  export default {
    components: {
      ModalPrimary,
      InputPrimary,
      ButtonPrimary,
      QuestionPrimary,
    },
    emits: ['close'],
    data() {
      return {
        show: false,
        patient: [],
        model: {
          Name: '',
          IdQuiz: null,
          Interval: 5,
          Created: null,
          Updated: null,
        },
        questions: [{ key: 0, Desc: '', IdQuestion: null, Answer: ''}],
      };
    },
    methods: {
      openModal(current, patient) {
        const th = this;
        th.patient = patient
        th.show = true;
        if (current) {
          th.model = { ...current, Interval: +current.Interval };
          th.$api.QuestionController.getAll()
            .then(async ({ data }) => {

              const awnsers = await th.$api.PatientHasQuizController.getByIdQuizPatient(th.model.IdQuiz, patient.IdPatient)
              const answers = awnsers.data.data[0].Answers.split(',')

              const filteredQuestions = data.data
                .filter(({ IdQuiz }) => IdQuiz === th.model.IdQuiz)
                .map((el, index) => {

                  return {
                    IdQuestion: el.IdQuestion,
                    Desc: el.Desc,
                    Answer: answers[index]  ? answers[index] : null,
                    key: el.IdQuestion,
                  };
                });
              th.questions =
                filteredQuestions.length > 0
                  ? filteredQuestions
                  : [{ key: 0, Desc: '' }];
            })
        }
      },
      goBack() {
        this.$router.push('/pacienteInfo?id='+this.patient.IdPatient)
      },
      closeModal() {
        this.show = false;
        this.model = {
          Name: '',
          IdQuiz: null,
          Answer: '',
          Interval: 5,  
        };
        this.$emit('close');
      },
    },
  };
  </script>
<style>
  .q-slide.answer .q-slider__track{
    background: linear-gradient(to right, red, yellow 50%, #04df04) !important;
  }
  .q-slide.answer .q-slider__selection{
    background: transparent;
  }
  .q-slide.answer .text-primary{
    color: blue !important;
  }
</style>
  <style scoped>
  .fill-content {
    width: 100%;
  }

  p{
    margin: 0;
  }
  
  .question {
    position: relative;
  }
  
  .remove-question {
    position: absolute;
    top: 1.8em;
    right: 1em;
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
  