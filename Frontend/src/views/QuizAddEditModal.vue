<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      {{ model && model.IdQuiz ? 'Editar' : 'Cadastrar' }} Questionário
    </template>
    <template #modal-content>
      <div class="row q-mb-sm">
        <input-primary
          v-model="model.Name"
          label="Nome do Questionário"
          label-color="primary"
          required
        />
      </div>
      <div class="row q-mb-sm interval-field">
        <span>Intervalo de resposta: {{ model.Interval }}</span>
        <q-slider
          v-model="model.Interval"
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
        class="row q-mb-sm"
      >
        <input-primary
          v-model="question.Desc"
          :name="`${i}`"
          :label="`Pergunta ${i + 1}`"
          label-color="primary"
        />
        <button
          v-if="questions.length > 1"
          class="remove-question"
          @click="removeQuestion(i)"
        >
          Remover
        </button>
      </div>
      <div class="row">
        <button-primary
          class="fill-content"
          outlined
          @click="addQuestion"
        >
          Adicionar pergunta
          <template #after-label>
            <PhPlus
              class="primary"
              weight="bold"
            />
          </template>
        </button-primary>
      </div>
    </template>
    <template #modal-actions>
      <button-primary
        outlined
        size="sm"
        @click="closeModal"
      >
        Fechar
      </button-primary>
      <button-primary
        v-if="model.IdQuiz"
        size="sm"
        @click="updateQuiz"
      >
        Atualizar
      </button-primary>
      <button-primary
        v-else
        size="sm"
        @click="createQuiz"
      >
        Cadastrar
      </button-primary>
    </template>
  </modal-primary>
</template>
<script>
import ModalPrimary from '../components/ModalPrimary.vue';
import InputPrimary from '../components/InputPrimary.vue';
import ButtonPrimary from '../components/ButtonPrimary.vue';
import { PhPlus } from '@phosphor-icons/vue';

export default {
  components: {
    ModalPrimary,
    InputPrimary,
    ButtonPrimary,
    PhPlus,
  },
  emits: ['close'],
  data() {
    return {
      show: false,
      model: {
        Name: '',
        IdQuiz: null,
        Interval: 5,
        Created: null,
        Updated: null,
      },
      questions: [{ key: 0, Desc: '', IdQuestion: null }],
    };
  },
  methods: {
    openModal(current) {
      const th = this;
      th.show = true;
      if (current) {
        th.model = { ...current, Interval: +current.Interval };
        th.$api.QuestionController.getAll()
          .then(({ data }) => {
            console.log(
              data.data.filter(({ IdQuiz }) => IdQuiz === th.model.IdQuiz)
            );
            const filteredQuestions = data.data
              .filter(({ IdQuiz }) => IdQuiz === th.model.IdQuiz)
              .map((el) => {
                return {
                  IdQuestion: el.IdQuestion,
                  Desc: el.Desc,
                  key: el.IdQuestion,
                };
              });
            th.questions =
              filteredQuestions.length > 0
                ? filteredQuestions
                : [{ key: 0, Desc: '' }];
          })
          .catch((err) => {
            console.log(err);
          });
      }
    },
    addQuestion() {
      if (!this.questions.includes('')) {
        this.questions.push({
          Desc: '',
          IdQuestion: null,
          key: this.questions[this.questions.length - 1].key + 1,
        });
      }
    },
    removeQuestion(index) {
      if (index !== -1) {
        this.questions.splice(index, 1);
      }
    },
    closeModal() {
      this.show = false;
      this.model = {
        Name: '',
        IdQuiz: null,
        Interval: 5,
      };
      this.questions = [{ key: 0, Desc: '', IdQuestion: null }];
      this.$emit('close');
    },
    createQuiz() {
      const th = this;
      th.$api.QuizController.insert({
        ...th.model,
        Interval: th.model.Interval.toString(),
      })
        .then((responseQuiz) => {
          console.log(responseQuiz);
          th.questions.forEach((question) => {
            const questionDto = {
              IdQuiz: th.model.IdQuiz,
              IdQuestion: question.IdQuestion,
              Desc: question.Desc,
            };
            th.$api.QuestionController.insert(questionDto)
              .then((responseQuestion) => {
                console.log(responseQuestion);
              })
              .catch((err) => {
                console.log(err);
              });
          });
          return;
        })
        .then(() => {
          th.closeModal();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    updateQuiz() {
      const th = this;
      th.$api.QuizController.update({
        ...th.model,
        Interval: th.model.Interval.toString(),
      })
        .then((response) => {
          console.log(response);
          th.questions.forEach((question) => {
            const questionDto = {
              IdQuiz: th.model.IdQuiz,
              IdQuestion: question.IdQuestion,
              Desc: question.Desc,
            };
            if (questionDto.IdQuestion) {
              th.$api.QuestionController.update(questionDto).then(
                (responseQuestion) => [console.log(responseQuestion)]
              );
            } else {
              th.$api.QuestionController.insert(questionDto)
                .then((responseQuestion) => {
                  console.log(responseQuestion);
                })
                .catch((err) => {
                  console.log(err);
                });
            }
          });
          return;
        })
        .then(() => {
          th.closeModal();
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>
<style scoped>
.fill-content {
  width: 100%;
}

.remove-question {
  color: var(--others-red-300);
  flex: 1;
  text-align: end;
  cursor: pointer;
  transition: 0.3s;
  border: none;
  background: none;
}

.remove-question:hover {
  filter: brightness(0.8);
  text-decoration: underline;
}

.interval-field {
  border: 1px solid black;
  padding: 0.5rem 1rem;
  color: var(--neutral-dark-gray);
  border-radius: 4px;
}
</style>
