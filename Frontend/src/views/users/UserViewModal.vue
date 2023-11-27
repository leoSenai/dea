<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      Visualizar Usuário
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div class="row q-mb-sm">
          <input-primary
            v-model="model.Name"
            :disable="true"
            label="Nome"
            label-color="primary"
            required
          />
          <input-primary
            v-model="model.Email"
            :disable="true"
            label="E-mail"
            label-color="primary"
            required
            class="select row q-pt-sm"
          />
          <input-primary
            v-if="!model.IdUser"
            v-model="model.Password"
            :disable="true"
            label="Senha"
            label-color="primary"
            required
            password
            class="select row q-pt-sm"
          />
          <q-select
            v-model="model.TypeUser"
            :disable="true"
            :options="optionsTypeUser"
            outlined
            label="Tipo de Usuário"
            label-color="primary"
            option-label="label"
            option-value="value"
            options-dark
            class="select row q-pt-sm"
          />
          <q-select
            v-model="model.Active"
            :disable="true"
            :options="optionsActive"
            outlined
            label="Ativo"
            label-color="primary"
            option-label="label"
            option-value="value"
            options-dark
            class="select row q-pt-lg q-mb-lg"
          />
          <input-primary
            v-model="model.Phone"
            :disable="true"
            label="Telefone"
            label-color="primary"
            required
            mask="(##) #####-####"
            hint="Exemplo: (##) ##### - ####"
          />
          <q-select
            v-model="model.IdCbo"
            :disable="true"
            :options="optionsCbo"
            outlined
            label="Cbo"
            label-color="primary"
            option-label="label"
            option-value="value"
            options-dark
            class="select row q-pt-sm"
          />
          <!--<q-select
            v-model="model.IdService"
            :disable="true"
            :options="optionsServices"
            outlined
            label="Serviço"
            label-color="primary"
            option-label="label"
            option-value="value"
            options-dark
            class="select row q-pt-lg"
          />-->
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
      </div>
    </template>
  </modal-primary>
</template>
  <script>
  import ModalPrimary from '../../components/ModalPrimary.vue';
  import InputPrimary from '../../components/InputPrimary.vue';
  import ButtonPrimary from '../../components/ButtonPrimary.vue';
  
  export default {
    components: {
      ModalPrimary,
      InputPrimary,
      ButtonPrimary,
    },
    emits: ['close'],
    data() {
      return {
        show: false,
        model: {
          Name: '',
          Email: '',
          Password: '',
          TypeUser: { label: '', value: '' },
          Active: { label: 'Sim', value: 1 },
          Phone: '',
          IdCbo: '',
          IdService: '',
        },
        optionsTypeUser: [
        { label: 'Administrador', value: 'A' },
        { label: 'Usuário', value: 'U' },
      ],
        optionsActive: [
          { label: 'Sim', value: 1 },
          { label: 'Não', value: 0 },
        ],
        optionsServices: [],
        optionsCbo: [],
      };
    },
    methods: {
      async openModal(current) {
        const th = this;
        th.show = true;
          const responseService = await th.$api.ServicesController.getAll();
          th.optionsServices = responseService.data.data.map((service) => ({
            label: service.Desc,
            value: service.IdServices,
          }));
  
          const responseCbo = await th.$api.CboController.getAll();
          th.optionsCbo = responseCbo.data.data.map((cbo) => ({
            label: cbo.Desc,
            value: cbo.IdCbo,
          }));
  
          if (current) {
            const model = {
              ...current,
              TypeUser: th.optionsTypeUser.find(
                (type) => type.value === current?.TypeUser
              ),
              Active: th.optionsActive.find(
                (active) => active.value === current?.Active
              ),
              IdCbo: th.optionsCbo.find((cbo) => cbo.value === current?.IdCbo),
              IdService: th.optionsServices.find(
                (service) => service.value === current?.IdServices
              ),
            };
  
            th.model = { ...model };
          }
      
      },
      closeModal() {
        this.show = false;
        this.model = {
          Name: '',
          Email: '',
          Password: '',
          TypeUser: 'A',
          Active: 1,
          Phone: '',
          IdCbo: '',
          IdService: '',
        };
        this.$emit('close');
      },
    },
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

  .modal-actions button{
    width: 100%;
  }
  </style>
  