<template>
  <modal-primary
    v-model="show" 
    @close="closeModal"
  >
    <template #modal-title>
      {{ model && model.IdPatient ? 'Editar' : (newPatient ? "Cadastrar" : "Adicionar") }} Paciente
    </template>
    <template #modal-content>
      <div v-if="!(model && model.IdPatient)">
        <q-checkbox 
          v-model="newPatient"
        >
          Paciente novo
        </q-checkbox>
        <q-select
          v-show="!newPatient"
          v-model="patientExistToAdd"
          style="border: 1px solid rgba(0, 0, 0, 0.432);border-bottom: 1px solid gray;border-radius: 5px;padding-left: 10px;"
          label="&nbsp Paciente"
          :options="patientsDisponible"
          option-label="Name"
        />
      </div>
      <q-form
        v-show="(newPatient || (model && model.IdPatient))"
        ref="form"
      >
        <div class="row">
          <div class="col-12 col-lg-6 q-px-sm">
            <input-primary
              v-model="model.Name"
              label="Nome"
              label-color="primary"
              required
            />
          </div>
          <div class="col-12 col-md-6 q-px-sm">
            <input-primary
              v-model="model.Email"
              label="E-mail"
              label-color="primary"
              required
            />
          </div>
        </div>
        <div class="row">
          <div class="col-12 col-lg-6 q-px-sm">
            <input-primary
              v-model="model.Phone"
              label="Telefone"
              label-color="primary"
              mask="(##) #####-####"
              required
            />
          </div>
          <div class="col-12 col-lg-6 q-px-sm">
            <input-primary
              v-model="model.Cpf"
              label="CPF"
              label-color="primary"
              mask="###.###.###-##"
              required
            />
          </div>
        </div>
        <div class="row">
          <div class="col-12 q-px-sm">
            <input-primary
              v-model="model.Address"
              label="Endereço"
              label-color="primary"
              required
            />
          </div>
        </div>
        <div class="row">
          <div class="col-12 col-lg-6 q-px-sm">
            <input-primary
              v-model="model.BornDate"
              type="date"
              label="Data de nascimento"
              label-color="primary"
              required
            />
          </div>
          <div class="col-12 col-lg-6 q-px-sm">
            <q-select
              v-model="model.Sex"
              :options="sex_options"
              label="Sexo"
              option-label="label"
              option-value="value"
              label-color="primary"
              class="q-field select row select-input q-field__control"
              required
            />
          </div>
        </div>
        <div class="row">
          <div class="col-12 col-lg-4 q-px-sm">
            <input-primary
              v-model="model.DadName"
              label="Nome do pai"
              label-color="primary"
            />
          </div>
          <div class="col-12 col-lg-4 q-px-sm">
            <input-primary
              v-model="model.MomName"
              label="Nome da mãe"
              label-color="primary"
              required
            />
          </div>
          <div class="col-12 col-lg-4 q-px-sm">
            <input-primary
              v-model="model.Cid10"
              mask="#########"
              label="Código CID"
              label-color="primary"
              required
            />
          </div>
        </div>
        <div class="row">
          <div class="col-12 col-lg-8 q-px-sm">
            <input-primary
              v-model="model.Cns"
              label="CNS"
              maxlength="15"
              mask="###############"
              label-color="primary"
              required
            />
          </div>
          <div class="col-12 col-lg-4 q-px-sm">
            <select-primary
              v-model="model.NewBorn"
              :label="'Recem nascido?'"
              :options="['Sim', 'Não']"
            />
          </div>
        </div>
      </q-form>
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
        v-if="!model.IdPatient"
        size="sm" 
        type="submit"
        @click="createPatient"
      >
        {{ newPatient ? "Cadastrar" : "Adicionar" }}
      </button-primary>
      <button-primary
        v-else
        size="sm" 
        type="submit"
        @click="updatePatient"
      >
        Editar
      </button-primary>
    </template>
  </modal-primary>
</template>

<script>
import ModalPrimary from '../../components/ModalPrimary.vue';
import InputPrimary from '../../components/InputPrimary.vue';
import ButtonPrimary from '../../components/ButtonPrimary.vue';
import SelectPrimary from '../../components/SelectPrimary.vue';
import cookie from '../../utils/cookie';
import CPF from '../../utils/cpf/validator.js'

export default {
  components: {
    ModalPrimary,
    InputPrimary,
    ButtonPrimary,
    SelectPrimary,
  },
  emits: ['close'],
  data() {
    return {
      newPatient: false,
      patientExistToAdd: [],
      patientsVisible: [],
      patientsDisponible: [],
      sex_options: [
        {label: 'Masculino', value: 'M'},
        {label: 'Feminino', value: 'F'},
        {label: 'Outros', value: 'O'},
        {label: 'Prefiro não dizer', value: 'P'}
      ],
      show: false,
      aaa: 0,
      model: {
        IdPatient: null,
        Name: '',
        Email: '',
        Phone: '',
        Address: '',
        Cpf: '',
        BornDate: null,
        Sex: '',
        DadName: '',
        MomName: '',
        Cid10: 0,
        Cns: '',
        NewBorn: '',
      },
    };
  },
  mounted() {
    const th = this
    th.$api.PatientController.getAll().then((data)=> {
      th.patientsDisponible = data.data.data
    })
  },
  methods: {
    getSex(current){
      
      if(current.Sex=='M'){
        return {label: 'Masculino', value: 'M'}
      }else if(current.Sex=='F'){
        return {label: 'Feminino', value: 'F'}
      }else if(current.Sex=='O'){
        return {label: 'Outros', value: 'O'}
      }else{
        return {label: 'Prefiro não dizer', value: 'P'}
      }
        
    },
    openModal(current) {
      const th = this;
      th.show = true;
      if (current) {

        var sexo = th.getSex(current)

        this.model = { ...current, NewBorn: current.NewBorn ? 'Sim' : 'Não', Sex: sexo};
        
        if(!(th.model && th.model.IdPatient)){
          th.patientsVisible = current.data
          th.filterPatients()
        }

      }
    },
    filterPatients(){
      const th = this

      for (let index = 0; index < th.model.data.length; index++) {
        for (let index2 = 0; index2 < th.patientsDisponible.length; index2++) {
          if(th.model.data[index].IdPatient == th.patientsDisponible[index2].IdPatient){
            th.patientsDisponible.splice(index2, 1)
          }
        }
      }
   
    },
    createPatient() { //create or add
      const th = this;

      if(th.newPatient){ //criar novo paciente

        var validPatientName;
        try{
          validPatientName = (th.model.Name.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g))[0]==th.model.Name
        }catch{
          validPatientName = false
        }

        var validEmail;
        try{
          validEmail = (th.model.Email.match(/^([\w.]{3,})@([a-z]{1,}[.]){1,}([a-z]{1,})/g))[0]==th.model.Email
        }catch{
          validEmail = false
        }

        var validMomName;
        try{
          validMomName = (th.model.MomName.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g))[0]==th.model.MomName
        }catch{
          validMomName = false
        }

        var validDadName;
        try{
          if(!th.model.DadName){   
            validDadName = true 
          }else{
            validDadName = (th.model.DadName.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g))[0]==th.model.DadName
          }
        }catch{
          validDadName = false
        }

        var date_field = new Date()
        date_field.setTime(Date.parse(th.model.BornDate))
        date_field.setDate(date_field.getDate()+1)
        var date_now = new Date()
        var dateIsValid = (date_field.getDate() <= date_now.getDate() || date_field.getTime() <= date_now.getTime())

        var validCpf = false

        if(th.model.Cpf && (new CPF().validate(th.model.Cpf))){
          validCpf = true
        }

        if(!validPatientName || th.model.Name.length > 90){
          alert('Preencha o nome do paciente corretamente!')
          return
        }else if(!validEmail){
          alert('Preencha o e-mail do paciente corretamente!')
          return
        }else if(!th.model.Phone || !(th.model.Phone.length == 15) || th.model.Phone[5]!='9'){
          alert('Preencha o telefone do paciente corretamente!')
          return
        }else if(!th.model.Address || th.model.Address.length > 255){
          alert('Preencha o endereço do paciente corretamente!')
          return
        }else if(!th.model.Cpf || !(th.model.Cpf.length == 14) || !validCpf){
          alert('Preencha o CPF do paciente corretamente!')
          return
        }else if(!th.model.BornDate || !dateIsValid){
          alert('Preencha a data de nascimento do paciente corretamente, ou verifique se a data corresponde a um recém nascido!')
          return
        }else if(!th.model.Sex){
          alert('Preencha o sexo do paciente corretamente!')
          return
        }else if(!validDadName){
          alert('Preencha o nome do pai do paciente corretamente!')
          return
        }else if (!validMomName || th.model.MomName.length > 80){
          alert('Preencha o nome da mãe do paciente corretamente!')
          return
        }else if(!th.model.Cid10 || !(th.model.Cid10 <= 999999999)){
          alert('Preencha o CID10 do paciente corretamente!')
          return
        }else if(!th.model.Cns || !(th.model.Cns.length == 15) || isNaN(th.model.Cns)){
          alert('Preencha o CNS do paciente corretamente!')
          return
        }else if(!th.model.NewBorn || (th.model.NewBorn=='Não' && (date_field.toDateString()==date_now.toDateString())) || (th.model.NewBorn=='Sim' && (date_field.toDateString()!=date_now.toDateString()))){
          alert('Preencha o campo "Recem nascido" corretamente!')
          return
        }

        th.model.Cid10 = th.model.Cid10 ? parseInt(th.model.Cid10) : 0;
        if (th.model.NewBorn === 'Sim') {
          th.model.NewBorn = 1;
        } else {
          th.model.NewBorn = 0;
        }

        th.model.Sex = th.model.Sex.value

        th.$api.PatientController.insert({
          Patient: th.model, IdUser: cookie.getUserId(cookie.get('authToken'))
        })
        .then(({ data }) => {
          th.model = data.data;
          if (th.model.IdPatient) {
            th.patients.forEach((patient) => {
              const patientDto = {
                IdPatient: th.model.IdPatient,
                Desc: patient.Desc.trim(),
              };
              th.$api.PatientController.insert(patientDto);
            });
          }
          return;
        })
        .then(() => {
          th.closeModal();
        });
      }else{ //adicionar existente

        if(th.patientExistToAdd.length==0){
          alert('Selecione um paciente para adicionar/vincular!')
          return
        }
        var patientId = th.patientExistToAdd.IdPatient
        var userId = cookie.getUserId(cookie.get('authToken'))

        th.$api.PatientHasUserController.insert({IdPatient: patientId, IdUser: userId}).then(() => {
          th.closeModal()
        })

      }

    },
    updatePatient() {
      const th = this;

      var validPatientName;
      try{
        validPatientName = (th.model.Name.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g))[0]==th.model.Name
      }catch{
        validPatientName = false
      }

      var validEmail;
      try{
        validEmail = (th.model.Email.match(/^([\w.]{3,})@([a-z]{1,}[.]){1,}([a-z]{1,})/g))[0]==th.model.Email
      }catch{
        validEmail = false
      }

      var validMomName;
      try{
        validMomName = (th.model.MomName.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g))[0]==th.model.MomName
      }catch{
        validMomName = false
      }

      var validDadName;
      try{
        if(!th.model.DadName){   
          validDadName = true 
        }else{
          validDadName = (th.model.DadName.match(/\b([A-Z][a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,}){1,} ([A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})(( [A-Z]{0,}[a-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]{1,})?){0,}\b/g))[0]==th.model.DadName
        }
      }catch{
        validDadName = false
      } 

      var date_field = new Date()
      date_field.setTime(Date.parse(th.model.BornDate))
      date_field.setDate(date_field.getDate()+1)
      var date_now = new Date()
      var dateIsValid = (date_field.getDate() <= date_now.getDate() || date_field.getTime() <= date_now.getTime())

      var validCpf = false

        if(th.model.Cpf && (new CPF().validate(th.model.Cpf))){
          validCpf = true
        }

      if(!validPatientName || th.model.Name.length > 90){
        alert('Preencha o nome do paciente corretamente!')
        return
      }else if(!validEmail){
        alert('Preencha o e-mail do paciente corretamente!')
        return
      }else if(!th.model.Phone || !(th.model.Phone.length == 15) || th.model.Phone[5]!='9'){
        alert('Preencha o telefone do paciente corretamente!')
        return
      }else if(!th.model.Address || th.model.Address.length > 255){
        alert('Preencha o endereço do paciente corretamente!')
        return
      }else if(!th.model.Cpf || !(th.model.Cpf.length == 14) || !validCpf){
        alert('Preencha o CPF do paciente corretamente!')
        return
      }else if(!th.model.BornDate || !dateIsValid){
        alert('Preencha a data de nascimento do paciente corretamente, ou verifique se a data corresponde a um recém nascido!')
        return
      }else if(!th.model.Sex){
        alert('Preencha o sexo do paciente corretamente!')
        return
      }else if(!validDadName){
          alert('Preencha o nome do pai do paciente corretamente!')
          return
      }else if (!validMomName || th.model.MomName.length > 80){
        alert('Preencha o nome da mãe do paciente corretamente!')
        return
      }else if(!th.model.Cid10 || !(th.model.Cid10 <= 999999999)){
        alert('Preencha o CID10 do paciente corretamente!')
        return
      }else if(!th.model.Cns || !(th.model.Cns.length == 15) || isNaN(th.model.Cns)){
        alert('Preencha o CNS do paciente corretamente!')
        return
      }else if(!th.model.NewBorn || (th.model.NewBorn=='Não' && (date_field.toDateString()==date_now.toDateString())) || (th.model.NewBorn=='Sim' && (date_field.toDateString()!=date_now.toDateString()))){
        alert('Preencha o campo "Recem nascido" corretamente!')
        return
      }

      if (th.model.NewBorn === 'Sim') {
        th.model.NewBorn = 1;
      } else {
        th.model.NewBorn = 0;
      }

      th.model.Sex = th.model.Sex.value

      th.$api.PatientController.update({
        ...th.model,
      })
        .then(({ data }) => {
          th.model = data.data;
          if (th.model.IdPatient) {
            th.patients.forEach((patient) => {
              const patientDto = {
                IdPatient: th.model.IdPatient,
                Desc: patient.Desc.trim(),
              };
              th.$api.PatientController.insert(patientDto);
            });
          }
          return;
        })
        .then(() => {
          th.closeModal();
        });
    },
    closeModal() {
      this.show = false;
      this.patients = []
      this.patientsVisible = []
      this.patientExistToAdd = []
      this.$emit('close');
    },
  },
};
</script>
