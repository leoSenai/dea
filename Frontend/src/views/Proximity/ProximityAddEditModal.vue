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
        <div class="row q-mb-sm">
          <div class="col">
            <input-primary
              v-model="model.Email"
              label="E-mail"
              label-color="primary"
              required
            />
          </div>
          <div class="col q-ml-sm">
            <input-primary
              v-model="model.Phone"
              label="Celular"
              format="phone"
              mask="(##) #####-####"
              hint="Exemplo: (##) #####-####"
              label-color="primary"
              required
            />
          </div>
        </div>
        <div class="row">
          <div class="col ">
            <input-primary
              v-model="model.BornDate"
              type="date"
              label-color="primary"
              label="Data de Nascimento"
              required
            />
          </div>
          <div class="col q-ml-sm">
            <select-primary
              v-model="model.DocType"
              :options="docTypes"
              label="Tipo do Documento"
              required
            />
          </div>
          <div class="col q-ml-sm">
            <input-primary
              v-model="model.DocNumber"
              :format="model.DocType=='CNPJ' ? 'cnpj' : 'cpf'"
              label="Número do Documento"
              label-color="primary"
              :mask=" model.DocType=='CNPJ' ? '##.###.###/####-##': '###.###.###-##' "
              required
            />
          </div>
        </div>
        <div class="row">
          <input-primary
            v-model="model.DescPerson"
            :disable="descPersonDisabled"
            label="Descreva a proximidade da pessoa para paciente"
            label-color="primary"
            type="textarea"
            autogrow
            required
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
import SelectPrimary from '../../components/SelectPrimary.vue';
import CPF from '../../utils/cpf/validator.js'
import { alerta } from '../../utils/alert';

export default {
  components: {
    ModalPrimary,
    InputPrimary,
    ButtonPrimary,
    SelectPrimary
  },
  emits: ['close'],
  data() {
    return {
      show: false,
      descPersonDisabled: false,
      model: {
        IdPerson: 0,
        Name: '',
        Phone: '',
        BornDate: '',
        DocNumber: '',
        DocType: 'CPF',
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
        th.model = current
        th.model.DocType = current.DocType
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
    createPerson() {
      const th = this;

      th.$refs.form.validate().then((success) => {
        
        var validName;
        try{
          validName = th.model.Name.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g)[0]==th.model.Name
        }catch{
          validName = false
        }

        
        var validEmail;
        try{
          validEmail = (th.model.Email.match(/^([\w.]{3,})@([a-z]{1,}[.]){1,}([a-z]{1,})/g))[0]==th.model.Email
        }catch{
          validEmail = false
        }

        var date_field = new Date()
        date_field.setTime(Date.parse(th.model.BornDate))
        date_field.setDate(date_field.getDate()+1)
        var date_now = new Date()
        var dateIsValid = (date_field.getDate() <= date_now.getDate() || date_field.getTime() <= date_now.getTime())

        var validDocNumber;
        try{
          validDocNumber = th.model.DocNumber.match(/(^\d{3}.\d{3}.\d{3}-\d{2})|(^\d{2}.\d{3}.\d{3}[/]\d{4}-\d{2})/g)[0]==th.model.DocNumber
          if(th.model.DocType=='CPF'){
            if(th.model.DocNumber && (new CPF().validate(th.model.DocNumber))){
              validDocNumber = true
            }else{
              validDocNumber = false
            }
          }
        }catch{
          validDocNumber = false
        }
        
        if(!validName){
          alerta('Informe um nome válido!', false)
          return
        }else if(!validEmail){
          alerta('Informe um e-mail válido!', false)
          return
        }else if(!(th.model.Phone.length == 15) || th.model.Phone[5]!='9'){
          alerta('Informe um celular válido!', false)
          return
        }else if(!th.model.BornDate || !dateIsValid){
          alerta('Preencha a data de nascimento corretamente!', false)
          return
        }else if(!validDocNumber){
          alerta('Informe um número de documento válido!', false)
          return
        }else if(!th.model.DescPerson || th.model.DescPerson.length < 3){
          alerta('Preencha a descrição de proximidade com o paciente corretamente!', false)
          return
        }

        if(success){

          th.$api.PersonController.getByDoc(th.model.DocNumber).then(({ data }) => {
            if (!data.data) {
              th.$api.PersonController.insert({ ...th.model, IdPatient: +th.IdPatient}).then(() => {
                th.closeModal()
              })
            } else if(data.data.Name != th.model.Name){
              alerta('O CPF inserido já está cadastrado!', false)
            }else{
              th.updatePerson();
            }
          })
        }else{
          alerta('Preencha todos os campos!', false)
        }
      })
    },
    updatePerson() {

      const th = this;

      th.$refs.form.validate().then((success) => {
        
        var validName;
        try{
          validName = th.model.Name.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g)[0]==th.model.Name
        }catch{
          validName = false
        }

        
        var validEmail;
        try{
          validEmail = (th.model.Email.match(/^([\w.]{3,})@([a-z]{1,}[.]){1,}([a-z]{1,})/g))[0]==th.model.Email
        }catch{
          validEmail = false
        }

        var date_field = new Date()
        date_field.setTime(Date.parse(th.model.BornDate))
        date_field.setDate(date_field.getDate()+1)
        var date_now = new Date()
        var dateIsValid = (date_field.getDate() <= date_now.getDate() || date_field.getTime() <= date_now.getTime())

        var validDocNumber;
        try{
          validDocNumber = th.model.DocNumber.match(/(^\d{3}.\d{3}.\d{3}-\d{2})|(^\d{2}.\d{3}.\d{3}[/]\d{4}-\d{2})/g)[0]==th.model.DocNumber

          if(th.model.DocType=='CPF'){
            if(th.model.DocNumber && (new CPF().validate(th.model.DocNumber))){
              validDocNumber = true
            }else{
              validDocNumber = false
            }
          }

        }catch{
          validDocNumber = false
        }
        
        if(!validName){
          alerta('Informe um nome válido!', false)
          return
        }else if(!validEmail){
          alerta('Informe um e-mail válido!', false)
          return
        }else if(!(th.model.Phone.length == 15) || th.model.Phone[5]!='9'){
          alerta('Informe um celular válido!', false)
          return
        }else if(!th.model.BornDate || !dateIsValid){
          alerta('Preencha a data de nascimento corretamente!', false)
          return
        }else if(!validDocNumber){
          alerta('Informe um número de documento válido!', false)
          return
        }else if(!th.model.DescPerson || th.model.DescPerson.length < 3){
          alerta('Preencha a descrição de proximidade com o paciente corretamente!', false)
          return
        }

        if(success){
          const modelPerson = th.model

          th.$api.PersonController.getByDoc(th.model.DocNumber).then(({ data }) => {
            if (data.data && data.data.IdPerson==th.model.IdPerson) {
              th.$api.PersonController.update({ ...modelPerson, IdPatient: +th.IdPatient}).then(() => {
                th.closeModal()
              })
            } else if(data.data.Name != th.model.Name){
              alerta('O CPF inserido já está cadastrado!', false)
            }
          })
        }else{
          alerta('Preencha todos os campos!', false)
        }
      })


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
