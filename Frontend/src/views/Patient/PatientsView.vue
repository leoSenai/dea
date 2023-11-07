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
              @click="openAddEditModal()"
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
            <div v-if="model.data.length==0">
              <i>Não há pacientes cadastrados ainda.</i>
            </div>
            <div 
              v-for="patient in model.data"
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
                    Visualizar
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
      const th = this;
      th.$api.PatientController.getAll().then(({ data }) => {
        th.model = data
      }).catch(({ response }) => {
        th.model = {
          ...response.data,
          hasError: true
        }
      })
    },
    openAddEditModal(current) {
      this.$refs.addEdit.openModal(current)
    },
    openViewPatient(id){
      this.$router.push('pacienteInfo?id='+id)
    },
    openEditPatient(id){
      this.$router.push('pacienteInfo?id='+ id +'&edit=true')
    },
    resetPassword ({IdPatient}) {
      const th = this;
      th.$api.PatientController.resetPassword(IdPatient)
    }
  }
};
</script>

<style scoped>

.edit-button{
  height: 110%;
  margin-right: 12px;
  cursor: pointer;
}

.patients-content{
  gap: 1em;
}

.row {
  width: 100%;
}

.container {
  height: 100%;
}

.patients {
  height: 89.5vh;
  width: 100%;
  padding: 2rem;
  padding-top: 0;
  margin-top: 5rem;
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
  display: flex;
  justify-content: space-between;
}

.patients-list:hover{
  background-color: rgba(200, 255, 172, 0.041);
}

.patients-list span{
  width: -webkit-fill-available;
  cursor: pointer;
  padding: 0.8em;
  font-size: 1.25rem;
  font-weight: 300;
}

.btn-modal {
  width: 4.5rem;
  height: 4.5rem;
  background: var(--primary-500, #519832);
  border-radius: 2.5rem;
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.btn-modal>.icon-color {
  font-size: 2.5rem;
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

@media screen and (max-width: 992px) {

  .patients {
    margin-top: 0rem;
  }

  .info-patients button{
    display: none;
  }

  .info-patients .patients-title{
    padding-left: 2rem;
    margin-top: 4%;
  }

  .patients-content {
    padding-left: 2rem;
    padding-right: 2rem;
    max-height: 70vh;
    overflow-y: auto;
  }

  .btn-modal {
    display: flex !important;
  }
}
</style>