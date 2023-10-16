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
          <div class="col ">
            <input-date-primary
              v-model="model.BornDate"
              label="Data de Nascimento"
            />
          </div>
          <div class="col q-ml-sm">
            <select-primary
              v-model="model.DocType"
              :options="docTypes"
              label="Tipo do Documento"
            />
          </div>
          <div class="col q-ml-sm">
            <input-primary
              v-model="model.DocNumber"
              :format="model.DocType.toLowerCase()"
              label="Número do Documento"
              label-color="primary"
            />
          </div>
        </div>
        <div class="row">
          <input-primary
            v-model="model.DescPerson"
            label="Descreva a proximidade da pessoa para paciente"
            label-color="primary"
            type="textarea"
            autogrow
          />
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
import InputDatePrimary from '../../components/InputDatePrimary.vue';
import SelectPrimary from '../../components/SelectPrimary.vue';

export default {
  components: {
    ModalPrimary,
    InputPrimary,
    ButtonPrimary,
    InputDatePrimary,
    SelectPrimary
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
        DescPerson: '',
        Email: '',
      },
      docTypes: ['CPF', 'CNPJ']
    };
  },
  computed: {
    IdPatient: {
      get() {
        return this.$route.params.id
      }
    }
  },
  methods: {
    openModal(current) {
      const th = this;
      th.show = true;
      if (current) {
        console.log(current)
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
        DescPerson: '',
        Email: '',
      };
      this.$emit('close');
    },
    createPerson() {
      const th = this;
      const BornDate = th.model.BornDate.split('/').reverse().join('-')
      th.$api.PersonController.getByDoc(th.model.DocNumber).then(({ data }) => {
        if (!data.data) {
          th.$api.PersonController.insert({ ...th.model, IdPatient: +th.IdPatient, BornDate })
        } else {
          th.updatePerson();
        }
      })
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
