<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      {{ model && model.IdQuiz ? 'Editar' : 'Cadastrar' }} Questionário
    </template>
    <template #modal-content>
      <q-form ref="form">
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
          class="question row q-mb-sm"
        >
          <button
            v-if="questions.length > 1"
            class="remove-question rounded-full"
            @click="removeQuestion(i)"
          >
            <PhX weight="bold" />
          </button>
          <input-primary
            v-model="question.Desc"
            :name="`${i}`"
            :label="`Pergunta ${i + 1}`"
            label-color="primary"
          />
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
          v-if="model.IdQuiz"
          size="sm"
          type="submit"
          @click="updateQuiz"
        >
          Atualizar
        </button-primary>
        <button-primary
          v-else
          size="sm"
          type="submit"
          @click="createQuiz"
        >
          Cadastrar
        </button-primary>
      </div>
    </template>
  </modal-primary>
</template>
<script>
import ModalPrimary from '../components/ModalPrimary.vue';
import InputPrimary from '../components/InputPrimary.vue';
import ButtonPrimary from '../components/ButtonPrimary.vue';
import { PhPlus, PhX } from '@phosphor-icons/vue';

export default {
  components: {
    ModalPrimary,
    InputPrimary,
    ButtonPrimary,
    PhPlus,
    PhX,
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
      questionRemoved: [],
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
        const [question] = this.questions.splice(index, 1);
        this.questionRemoved.push(question);
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
      th.$refs.form.validate().then((success) => {
        if (th.questions.some(({ Desc }) => !Desc || Desc.trim() === '')) {
          alert(
            'Não é possível criar questionários com questões sem descrição!'
          );
          return;
        }
        if (success) {
          th.$api.QuizController.insert({
            ...th.model,
            Interval: th.model.Interval.toString(),
          })
            .then(({ data }) => {
              th.model = data.data;
              if (th.model.IdQuiz) {
                th.questions.forEach((question) => {
                  const questionDto = {
                    IdQuiz: th.model.IdQuiz,
                    IdQuestion: question.IdQuestion,
                    Desc: question.Desc.trim(),
                  };
                  th.$api.QuestionController.insert(questionDto)
                });
              }
              return;
            })
            .then(() => {
              th.closeModal();
            })
        }
      });
    },
    updateQuiz() {
      const th = this;
      th.$refs.form.validate().then((success) => {
        if (success) {
          if (th.questions.some(({ Desc }) => !Desc || Desc.trim() === '')) {
            alert(
              'Não é possível atualizar questionários com questões sem descrição!'
            );
            return;
          }
          th.$api.QuizController.update({
            ...th.model,
            Interval: th.model.Interval.toString(),
          })
            .then(() => {
              th.questions.forEach((question) => {
                const questionDto = {
                  IdQuiz: th.model.IdQuiz,
                  IdQuestion: question.IdQuestion,
                  Desc: question.Desc.trim(),
                };
                if (questionDto.IdQuestion) {
                  th.$api.QuestionController.update(questionDto);
                } else {
                  th.$api.QuestionController.insert(questionDto)
                }
              });
              return;
            })
            .then(() => {
              th.questionRemoved.forEach(({ IdQuestion }) => {
                th.$api.QuestionController.delete(IdQuestion);
              });
              return;
            })
            .then(() => {
              th.closeModal();
            })
        }
      });
    },
  },
};
</script>
<style scoped>
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
