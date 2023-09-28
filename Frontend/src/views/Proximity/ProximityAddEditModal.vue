<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      {{ model && model.IdPerson ? 'Editar' : 'Cadastrar' }} Pessoa
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div class="row q-mb-sm">
          <input-primary
            v-model="model.Name"
            label="Nome da Pessoa"
            label-color="primary"
            required
          />
        </div>
        <div class="row">
          <input-date-primary label="Data de Nascimento" />
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
          v-if="model.IdPerson"
          size="sm"
          type="submit"
          @click="updatePerson"
        >
          Atualizar
        </button-primary>
        <button-primary
          v-else
          size="sm"
          type="submit"
          @click="createPerson"
        >
          Cadastrar
        </button-primary>
      </div>
    </template>
  </modal-primary>
</template>
<script>
import ModalPrimary from '../../components/ModalPrimary.vue';
import InputPrimary from '../../components/InputPrimary.vue';
import ButtonPrimary from '../../components/ButtonPrimary.vue';
import InputDatePrimary from '../../components/InputDatePrimary.vue'

export default {
  components: {
    ModalPrimary,
    InputPrimary,
    ButtonPrimary,
    InputDatePrimary
  },
  emits: ['close'],
  data() {
    return {
      show: false,
      model: {
        IdPerson: 0,
        Name: '',
        BornDate: '',
        DocNumber: '',
        DocType: '',
        Password: '',
        Salt: '',
        IdPatient: 0,
        DescPerson: '',
        Email: ''
      },
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
    closeModal() {
      this.show = false;
      this.model = {
        IdPerson: 0,
        Name: '',
        BornDate: '',
        DocNumber: '',
        DocType: '',
        Password: '',
        Salt: '',
        IdPatient: 0,
        DescPerson: '',
        Email: ''
      };
      this.$emit('close');
    },
    createPerson() {
      const th = this;
      console.log(th)
    },
    updatePerson() {
      const th = this;
      console.log(th)
    },
  },
};
</script>
<style scoped>
.fill-content {
  width: 100%;
}

.modal-actions {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
</style>
