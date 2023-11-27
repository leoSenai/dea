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
            mask="(##) #####-####"
            hint="Exemplo: (##) #####-####"
          />
          <input-primary
            v-model="model.IdCbo"
            label="Código Brasileiro de Ocupação (2002)"
            label-color="primary"
            required
            mask="######"
          />
          <!--<q-select
            v-model="model.IdService"
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
        TypeUser: { label: null, value: null },
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
       // const responseService = await th.$api.ServicesController.getAll();
        //th.optionsServices = responseService.data.data.map((service) => ({
        //  label: service.Desc,
        //  value: service.IdServices,
        //}));

        //var responseCboExist = await th.$api.CboController.getAll();
          var responseCbo = await th.$api.CboController.getAll();
          th.optionsCbo = responseCbo.data.data
  
        //    for (let i = 0; i < th.optionsCbo.length; i++) {
        //      const optioncbo = th.optionsCbo[i];
        //      await th.$api.CboController.update(optioncbo);
        //    }
        //  }
        //}


        if (current) {

          var idCboCurrent = current?.IdCbo
          var codeCbo = (th.optionsCbo.map((cbos) => { return {valid: idCboCurrent==cbos.IdCbo, Code: cbos.Code} })).find((result) => result.valid==true).Code

          const model = {
            ...current,
            TypeUser: th.optionsTypeUser.find(
              (type) => type.value === current?.TypeUser
            ),
            Active: th.optionsActive.find(
              (active) => active.value === current?.Active
            ),
            IdCbo: codeCbo,
            IdService: 1,
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
        TypeUser: null,
        Active: undefined,
        Phone: '',
        IdCbo: '',
        IdService: 1,
      };
      this.$emit('close');
    },
    createAndUpdateUser() {
      const th = this;
      th.$refs.form.validate().then((success) => {
        success = th.validations();

        if (success) {
          const IdCbo = parseInt(th.model.IdCbo);
          const IdServices = 1;
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
                if(data.data=='' && !data.message.includes('atualizadas com sucesso')){
                  return 
                }else{
                  th.model = data.data;
                  th.closeModal();
                  return;
                }
              })
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
              if(!data.message.includes('cadastrado com sucesso')){
                return
              }else{
                th.model = data.data;
                th.closeModal();
                return;
              }
            })
            
        }
      });
    },
    validations() {
      const th = this;
      
      var validName;
      try{
        validName = (th.model.Name.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g))[0]==th.model.Name
      }catch{
        validName = false
      }
      
      var validEmail;
      try{
        validEmail = (th.model.Email.match(/^([\w.]{3,})@([a-z]{1,}[.]){1,}([a-z]{1,})/g))[0]==th.model.Email
      }catch{
        validEmail = false
      }
      
      var validCbo;
      try{
        validCbo = (String(th.model.IdCbo).match(/^[0-9]{3,6}/g))[0]==String(th.model.IdCbo)
      }catch{
        validCbo = false
      }

      if (th.model.Name.length < 3 || !validName) {
        alert('O nome deve conter no mínimo 3 caracteres, ter sobrenome, e todas letras iniciais maiúsculas!');
        return false;
      }else if (th.model.Email.length < 3 || !validEmail) {
        alert('O e-mail deve conter no mínimo 3 caracteres, conter "@", e dominio do email!');
        return false;
      }else if (!th.model.IdUser && th.model.Password.length < 6) {
        alert('A senha deve conter no mínimo 6 caracteres!');
        return false;
      }else if(th.model.TypeUser.value==null || th.model.TypeUser==''){
        alert('Selecione um tipo de usuário!');
        return false;
      }else if (!(th.model.Phone.length == 15) || th.model.Phone[5]!='9') {
        alert('O telefone deve conter 11 digitos e um dígito "9" após o DDD!');
        return false;
      }else if (!validCbo) {
        alert('Digite um CBO válido!');
        return false;
      }





      //if (th.model.IdService === '') {
      //  alert('Selecione um Serviço!');
      //  return false;
      //}
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
