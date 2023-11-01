<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      {{ model && model.IdUser ? 'Editar' : 'Cadastrar' }} Usuário
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div class="row q-mb-sm">
          <input-primary
            v-model="model.Name"
            label="Nome"
            label-color="primary"
            required
          />
          <input-primary
            v-model="model.Email"
            label="E-mail"
            label-color="primary"
            required
            class="select row q-pt-sm"
          />
          <input-primary
            v-if="!model.IdUser"
            v-model="model.Password"
            label="Senha"
            label-color="primary"
            required
            password
            class="select row q-pt-sm"
          />
          <q-select
            v-model="model.TypeUser"
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
            label="Telefone"
            label-color="primary"
            required
            mask="(##) ##### - ####"
            hint="Exemplo: (##) ##### - ####"
          />
          <q-select
            v-model="model.IdCbo"
            :options="optionsCbo"
            outlined
            label="Cbo"
            label-color="primary"
            option-label="label"
            option-value="value"
            options-dark
            class="select row q-pt-sm"
          />
          <q-select
            v-model="model.IdService"
            :options="optionsServices"
            outlined
            label="Serviço"
            label-color="primary"
            option-label="label"
            option-value="value"
            options-dark
            class="select row q-pt-lg"
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
          v-if="model.IdUser"
          size="sm"
          type="submit"
          @click="createAndUpdateUser"
        >
          Atualizar
        </button-primary>
        <button-primary
          v-else
          size="sm"
          type="submit"
          @click="createAndUpdateUser"
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
        TypeUser: { label: 'Administrador', value: 'A' },
        Active: { label: 'Sim', value: 1 },
        Phone: '',
        IdCbo: '',
        IdService: '',
      },
      optionsTypeUser: [
        { label: 'Administrador', value: 'A' },
        { label: 'Outro', value: 'O' },
        { label: 'Psicologo', value: 'P' },
        { label: 'Psiquiatra', value: 'C' },
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
    createAndUpdateUser() {
      const th = this;
      th.$refs.form.validate().then((success) => {
        success = th.validations();

        if (success) {
          const IdCbo = th.model.IdCbo.value;
          const IdServices = th.model.IdService.value;
          const TypeUser = th.model.TypeUser.value;
          const Active = th.model.Active.value;
          const Phone = th.model.Phone.replace(/\D/g, '');

          delete th.model.IdService;

          if (th.model.IdUser) {
            th.$api.UsersController.update({
              ...th.model,
              IdCbo,
              IdServices,
              TypeUser,
              Active,
              Phone,
            })
              .then(({ data }) => {
                th.model = data.data;
                return;
              })
              .then(() => {
                th.closeModal();
              });
            return;
          }

          th.$api.UsersController.insert({
            ...th.model,
            IdCbo,
            IdServices,
            TypeUser,
            Active,
            Phone,
          })
            .then(({ data }) => {
              th.model = data.data;
              return;
            })
            .then(() => {
              th.closeModal();
            });
        }
      });
    },
    validations() {
      const th = this;
      if (!th.model.IdUser && th.model.Password.length < 6) {
        alert('A senha deve conter no mínimo 6 caracteres!');
        return false;
      }

      if (th.model.Name.length < 3) {
        alert('O nome deve conter no mínimo 3 caracteres!');
        return false;
      }

      if (th.model.Phone.length < 14) {
        alert('O telefone deve conter no mínimo 10 caracteres!');
        return false;
      }

      if (th.model.Email.length < 3 && !th.model.Email.includes('@')) {
        alert('O e-mail deve conter no mínimo 3 caracteres e conter @!');
        return false;
      }

      if (th.model.IdCbo === '') {
        alert('Selecione um CBO!');
        return false;
      }

      if (th.model.IdService === '') {
        alert('Selecione um Serviço!');
        return false;
      }
      return true;
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
</style>
