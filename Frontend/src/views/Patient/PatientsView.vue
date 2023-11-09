<template>
  <div class="row">
    <div class="col-12">
      <div class="flex justify-center items-center">
        <div class="patients">
          <div class="info-patients flex justify-between">
            <h3 class="patients-title">
              Pacientes
            </h3>
            <buttonPrimary
              type="button"
              @click="openAddEditModal(model)"
            >
              Adicionar
              <PhPlus
                class="icon-color"
                color="white"
              />
            </buttonPrimary>
          </div>
          <div 
            class="patients-content q-mt-lg flex" 
          >
            <div v-if="patients.length==0">
              <i>Não há pacientes cadastrados ainda.</i>
            </div>
            <div 
              v-for="patient in patients"
              :key="patient.IdPatient"
              class="patients-list"
            >
              <span @click="openViewPatient(patient.IdPatient)">{{ patient.Name }}</span>
              <div class="patients-actions">
                <button
                  type="button"
                  @click="resetPassword(patient)"
                >
                  <q-tooltip>
                    Redefinir Senha
                  </q-tooltip>
                  <PhFingerprintSimple color="black" />
                </button>
                <button
                  class="edit-button"
                  type="button"
                  @click="openAddEditModal(patient)"
                >
                  <q-tooltip>
                    Editar
                  </q-tooltip>
                  <PhPencil color="black" />
                </button>
              </div>
            </div>
          </div>
          <div
            class="btn-modal hidden flex justify-center items-center"
            type="button"
            @click="openAddEditModal()"
          >
            <PhPlus
              class="icon-color"
              color="white"
            />
          </div>
          <PatientsAddEditModal
            ref="addEdit"
            @close="load"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import buttonPrimary from '../../components/ButtonPrimary.vue';
import { PhPlus, PhPencil, PhFingerprintSimple } from '@phosphor-icons/vue';
import PatientsAddEditModal from './PatientsAddEditModal.vue';
import cookie from '../../cookie';

export default {
  components: {
    buttonPrimary,
    PhPlus,
    PhPencil,
    PatientsAddEditModal,
    PhFingerprintSimple
  },
  data() {
    return {
      patientsHasUser: [],
      patients: [],
      model: {
        data: [],
        hasError: false,
        message: ''
      }
    };
  },
  mounted() {
    this.load();
  },
  methods: {
    load() {
      this.patients = []
      this.getPatientsOfUser()
    },
    async getPatientsOfUser(){
      const th = this

      th.patientsHasUser = (await th.$api.PatientHasUserController.getByIdUser(cookie.getUserId(cookie.get('authToken')))).data.data
      th.patientsHasUser = th.patientsHasUser.map((el) => { return el.IdPatient })

      for (let index = 0; index < th.patientsHasUser.length; index++) {
        var dataPatient = (await th.$api.PatientController.getById(th.patientsHasUser[index])).data.data
        th.patients.push(dataPatient)
      }

      th.model.data = th.patients
   
    },
    openAddEditModal(current) {
      this.$refs.addEdit.openModal(current)
    },
    openViewPatient(id) {
      this.$router.push('paciente?id=' + id)
    },
    openEditPatient(id) {
      this.$router.push('paciente?id=' + id + '&edit=true')
    },
    resetPassword({ IdPatient }) {
      const th = this;
      th.$api.PatientController.resetPassword(IdPatient)
    }
  }
};
</script>

<style scoped>
.edit-button {
  height: 110%;
  margin-right: 12px;
  cursor: pointer;
}

.patients-content {
  gap: 1em;
}

.row {
  width: 100%;
  background-color: rgba(255, 255, 255, 0.548);
  background-size: 50% !important;
  background: url(../../assets/imgs/home-background.svg) no-repeat;
  background-position-x: center;
  background-position-y: center;
  height: 100%;
}

.container {
  height: 100%;
}

.patients {
  height: 89.5vh;
  width: 100%;
  padding: 2rem;
  padding-top: 0;
  margin-top: 3.9rem;
}

.patients-title {
  font-size: 3rem;
}

.patients-list {
  border: 1px solid var(--neutral-dark-gray);
  color: var(--neutral-dark-gray);
  width: 100%;
  border-radius: 0.25rem;
  padding: 0;
  background-color: rgba(255, 255, 255, 0.548);
  display: flex;
  justify-content: space-between;
}

.patients-list:hover {
  background-color: rgba(255, 255, 255, 0.89);
}

.patients-list span {
  width: -webkit-fill-available;
  cursor: pointer;
  padding: 0.8em;
  font-size: 1.25rem;
  font-weight: 400;
}

.btn-modal {
  width: 3rem;
  height: 3rem;
  background: var(--primary-500, #519832);
  border-radius: 2.5rem;
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.btn-modal>.icon-color {
  font-size: 2rem;
}

.patients-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.patients-actions button {
  border: none;
  background: none;
  padding: .25rem;
  height: 1.5rem;
  border-radius: 9999px;
  cursor: pointer;
}

.patients-actions button:hover {
  background: var(--neutral-gray);
}

.info-patients button {
  color: white;
}


.patients {
  margin-top: 0rem;
}

.info-patients button {
  display: none;
}

.info-patients .patients-title {
  padding-left: 2rem;
  margin-top: 4%;
}

.patients-content {
  padding-left: 2rem;
  padding-right: 2rem;
  gap: 1rem;
  max-height: 60vh;
  height: 100%;
  overflow-y: auto;
}

.btn-modal {
  display: flex !important;
}
</style>