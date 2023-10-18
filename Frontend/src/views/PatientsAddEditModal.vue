<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      {{ model && model.IdPatients ? "Editar" : 'Cadastrar' }} Pacientes
    </template>
    <template #modal-content>
      <q-form ref="form">
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
          <div class="col-12 col-lg-4 q-px-sm">
            <input-primary
              v-model="model.Phone"
              label="Telefone"
              label-color="primary"
              required
            />
          </div>
          <div class="col-12 col-lg-4 q-px-sm">
            <input-primary
              v-model="model.Adress"
              label="CEP"
              label-color="primary"
              required
            />
          </div>
          <div class="col-12 col-lg-4 q-px-sm">
            <input-primary
              v-model="model.Cpf"
              label="CPF"
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
          <div class="col-12 col-lg-4">
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
          <div class="col-12 col-lg-6 q-px-sm">
            <input-primary
              v-model="model.Cns"
              label="CNS"
              label-color="primary"
              required
            />
          </div>
          <div class="col-12 col-lg-6 q-px-sm">
            <select-primary
              v-model="model.NewBorn"
              :label="'Recem nascido?'"
              :options="['sim', 'não']"
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
        size="sm"
        type="submit"
        @click="createPatient"
      >
        Cadastrar
      </button-primary>
    </template>
  </modal-primary>
</template>

<script>
import ModalPrimary from '../components/ModalPrimary.vue';
import InputPrimary from '../components/InputPrimary.vue';
import ButtonPrimary from '../components/ButtonPrimary.vue';
import SelectPrimary from '../components/SelectPrimary.vue';

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
              NewBorn: 'não'
            }
        };
    },
    methods: {
        openModal(current) {
            const th = this;
            th.show = true;
            if (current){
                th.model = { ...current, Interval: +current.Interval };
                th.$api.PatientController.getAll()
                 .then(({ data }) => {
                    const filteredPatients = data.data
                    .filter(({ IdPatient }) => IdPatient === th.model.IdQuiz)
                    .map((el) => {
                        return {
                            IdPatient: el.Id
                        };
                    });
                    th.patients =
                    filteredPatients.length > 0
                    ? filteredPatients
                    : [{ key: 0}]
                 })
            }
        },
        createPatient(){
          const th = this;
              th.$api.PatientController.insert({
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
                  th.$api.PatientController.insert(patientDto)
                });
              }
              return;
            })
        },
        closeModal() {
            this.show = false;
            this.model = {
                Name: '',
                IdPatient: null,
            };
            this.$emit('close');
        },
    }
}
</script>