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
              required
            />
          </div>
          <div class="col-12 col-lg-6 q-px-sm">
            <input-primary
              v-model="model.Cpf"
              label="CPF"
              label-color="primary"
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
            <input-primary
              v-model="model.Sex"
              label="Sexo"
              label-color="primary"
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
              required
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
              type="number"
              label="CID"
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
    openModal(current) {
      const th = this;
      th.show = true;
      if (current) {

        this.model = { ...current, NewBorn: current.NewBorn ? 'Sim' : 'Não'};
        
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
        if (
          !th.model.Name ||
          !th.model.Email ||
          !th.model.Phone ||
          !th.model.Address ||
          !th.model.Cpf ||
          !th.model.BornDate ||
          !th.model.Sex ||
          !th.model.DadName ||
          !th.model.MomName ||
          !th.model.Cid10 ||
          !th.model.Cns ||
          !th.model.NewBorn
        ) {
          alert('Certifique-se de preencher todos os campos.');
          return;
        }
        th.model.Cid10 = th.model.Cid10 ? parseInt(th.model.Cid10, 10) : 0;
        if (th.model.NewBorn === 'Sim') {
          th.model.NewBorn = 1;
        } else {
          th.model.NewBorn = 0;
        }
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
        var patientId = th.patientExistToAdd.IdPatient
        var userId = cookie.getUserId(cookie.get('authToken'))

        th.$api.PatientHasUserController.insert({IdPatient: patientId, IdUser: userId}).then(() => {
          th.closeModal()
        })

      }

    },
    updatePatient() {
      const th = this;

      if (th.model.NewBorn === 'Sim') {
        th.model.NewBorn = 1;
      } else {
        th.model.NewBorn = 0;
      }

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
