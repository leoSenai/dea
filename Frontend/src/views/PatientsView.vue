<template>
  <div class="row">
    <div class="col-12">
      <div class="flex justify-center items-center">
        <div class="patients">
          <div class="info-patients flex justify-between">
            <h3 class="patients-title">Pacientes</h3>
            <buttonPrimary type="button" @click="openAddEditModal()">
              Adicionar
              <PhPlus class="icon-color"></PhPlus>
            </buttonPrimary>
          </div>
          <div class="patients-content q-mt-lg flex" style="gap: 1rem;">
            <div v-for="patient in model.data" :key="patient.IdPatient" class="patients-list q-pa-md">
              <span>{{ patient.Name }}</span>
            </div>
          </div>
          <div type="button" @click="openAddEditModal()" class="btn-modal hidden flex justify-center items-center">
            <PhPlus class="icon-color"></PhPlus>
          </div>
          <PatientsAddEditModal ref="addEdit" @close="load" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import buttonPrimary from '../components/ButtonPrimary.vue';
import { PhPlus } from '@phosphor-icons/vue';
import PatientsAddEditModal from './PatientsAddEditModal.vue';

export default {
  components: {
    buttonPrimary,
    PhPlus,
    PatientsAddEditModal,
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
    }
  }
};
</script>

<style scoped>
.row {
  width: 100%;
}

.container {
  height: 100%;
}

.patients {
  height: 89.5vh;
  width: 965px;
  margin-top: 5rem;
}

.patients-title {
  font-size: 3rem;
}

.patients-list {
  border: 0.094rem solid var(--neutral-dark-gray);
  width: 100%;
  border-radius: 0.25rem;
}

.patients-list span {
  font-weight: bold;
  font-size: 1.25rem;
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

@media screen and (max-width: 992px) {

  .patients {
    margin-top: 0rem;
  }

  .info-patients {
    display: none;
  }

  .patients-content {
    padding-left: 2rem;
    padding-right: 2rem;
  }

  .btn-modal {
    display: flex !important;
  }
}
</style>