<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      Visualizar pessoa próxima
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div class="row q-mb-sm">
          <input-primary
            v-model="model.Name"
            :disable="true"
            label="Nome da Pessoa"
            label-color="primary"
            required
          />
        </div>
        <div class="row q-mb-sm">
          <div class="col">
            <input-primary
              v-model="model.Email"
              :disable="true"
              label="E-mail"
              label-color="primary"
              required
            />
          </div>
          <div class="col q-ml-sm">
            <input-primary
              v-model="model.Phone"
              :disable="true"
              label="Celular"
              format="phone"
              mask="(##) #####-####"
              label-color="primary"
              required
            />
          </div>
        </div>
        <div class="row">
          <div class="col ">
            <input-date-primary
              v-model="model.BornDate"
              :disabled="true"
              :disable="true"
              label="Data de Nascimento"
              style="pointer-events: none; cursor:no-drop;"
            />
          </div>
          <div class="col q-ml-sm">
            <select-primary
              v-model="model.DocType"
              :disable="true"
              :options="docTypes"
              label="Tipo do Documento"
            />
          </div>
          <div class="col q-ml-sm">
            <input-primary
              v-model="model.DocNumber"
              :disable="true"
              :format="model.DocType ? model.DocType.toLowerCase() : 'cpf'"
              label="Número do Documento"
              label-color="primary"
            />
          </div>
        </div>
        <div class="row">
          <input-primary
            v-model="model.DescPerson"
            :disable="true"
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
          class="close-btn"
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
          Phone: '',
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
          current.BornDate = current.BornDate.split('-').reverse().join('/')
          th.model = current
        }
        th.descPersonDisabled = th.model.IdPerson ? true : false
      },
      closeModal() {
        this.show = false;
        this.model = {
          IdPerson: 0,
          Name: '',
          Phone: '',
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
  .close-btn{
    width: 100%;
  }
  </style>
  