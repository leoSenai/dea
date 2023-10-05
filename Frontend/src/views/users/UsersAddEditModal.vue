<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      {{ model && model.IdQuiz ? 'Editar' : 'Cadastrar' }} Usuário
    </template>
    <template #modal-content>
      <q-form ref="form">
        {{ model }}
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
            v-model="model.Password"
            label="Senha"
            label-color="primary"
            required
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
          Active: {label: 'Sim', value: 1 },
          Phone: '',
          IdCbo: '',
          IdService: '',
        },
        optionsTypeUser: [
            { label: 'Administrador', value: 'A' },
            { label: 'Outro', value: 'O' },
            { label: 'Psicologo', value: 'P'},
            { label: 'Psiquiatra', value: 'C'}
        ],
        optionsActive: [
            {label: 'Sim', value: 1 },
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
        try {
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
              TypeUser: th.optionsTypeUser.find(( value ) => value === current?.TypeUser),
              Active: th.optionsActive.find(( value ) => value === current?.Active),
              IdCbo: th.optionsCbo.find(( cbo ) => cbo.value === current?.IdCbo),
              IdService: th.optionsServices.find(( value ) => value === current?.IdService),
            };
            console.log('model', model);
            th.model = { ...model };
          }
        } catch (error) {
          console.error('Ocorreu um erro ao buscar os dados.', error);
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
      createQuiz() {
        // const th = this;
        // th.$refs.form.validate().then((success) => {
        //   if (th.questions.some(({ Desc }) => !Desc || Desc.trim() === '')) {
        //     alert(
        //       'Não é possível criar questionários com questões sem descrição!'
        //     );
        //     return;
        //   }
        //   if (success) {
        //     th.$api.QuizController.insert({
        //       ...th.model,
        //       Interval: th.model.Interval.toString(),
        //     })
        //       .then(({ data }) => {
        //         th.model = data.data;
        //         if (th.model.IdQuiz) {
        //           th.questions.forEach((question) => {
        //             const questionDto = {
        //               IdQuiz: th.model.IdQuiz,
        //               IdQuestion: question.IdQuestion,
        //               Desc: question.Desc.trim(),
        //             };
        //             th.$api.QuestionController.insert(questionDto)
        //           });
        //         }
        //         return;
        //       })
        //       .then(() => {
        //         th.closeModal();
        //       })
        //   }
        // });
      },
      updateQuiz() {
        // const th = this;
        // th.$refs.form.validate().then((success) => {
        //   if (success) {
        //     if (th.questions.some(({ Desc }) => !Desc || Desc.trim() === '')) {
        //       alert(
        //         'Não é possível atualizar questionários com questões sem descrição!'
        //       );
        //       return;
        //     }
        //     th.$api.QuizController.update({
        //       ...th.model,
        //       Interval: th.model.Interval.toString(),
        //     })
        //       .then(() => {
        //         th.questions.forEach((question) => {
        //           const questionDto = {
        //             IdQuiz: th.model.IdQuiz,
        //             IdQuestion: question.IdQuestion,
        //             Desc: question.Desc.trim(),
        //           };
        //           if (questionDto.IdQuestion) {
        //             th.$api.QuestionController.update(questionDto);
        //           } else {
        //             th.$api.QuestionController.insert(questionDto)
        //           }
        //         });
        //         return;
        //       })
        //       .then(() => {
        //         th.questionRemoved.forEach(({ IdQuestion }) => {
        //           th.$api.QuestionController.delete(IdQuestion);
        //         });
        //         return;
        //       })
        //       .then(() => {
        //         th.closeModal();
        //       })
        //   }
        // });
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
  