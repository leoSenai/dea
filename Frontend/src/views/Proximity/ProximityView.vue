<template>
  <div>
    <div>
      <div
        class="back-page"
        onclick="window.history.back()"
      >
        <PhCaretLeft color="#656565" />
        Voltar
      </div>
    </div>
    <div class="proximity-content">
      <div class="proximity-title">
        <div class="title">
          <h4>Pessoas próximas de {{ patientModel.Name }}</h4>
        </div>
        <div
          v-if="!isMobile"
          class="title-add-proximity"
        >
          <button
            type="button"
            @click="openAddEditModal()"
          >
            Adicionar
            <PhPlus />
          </button>
        </div>
      </div>
      <div
        v-if="model.hasError"
        class="error proximity"
      >
        {{ model.message }}
      </div>
      <div
        v-for="proximity in model.data"
        v-else
        :key="proximity.IdPerson"
        class="row proximity"
      >
        <p>
          {{ proximity.Name }}
        </p>
        <div class="proximity-actions">
          <button
            type="button"
            @click="viewQuizzes(proximity.IdPerson)"
          >
            <PhTable color="black" />
          </button>
          <button
            type="button"
            @click="resetPassword(proximity)"
          >
            <q-tooltip>
              Redefinir Senha
            </q-tooltip>
            <PhFingerprintSimple color="black" />
          </button>
          <button
            type="button"
            @click="openAddEditModal(proximity)"
          >
            <PhPencil color="black" />
          </button>
          <button
            type="button"
            @click="openViewModal(proximity)"
          >
            <PhEye color="black" />
          </button>
        </div>
      </div>
      <div
        v-if="isMobile"
        class="add-proximity"
      >
        <button
          type="button"
          @click="openAddEditModal()"
        >
          <PhPlus color="white" />
        </button>
      </div>
    </div>
    <ProximityAddEditModal
      ref="addEdit"
      @close="load"
    />
    <ProximityViewModal
      ref="viewModal"
      @close="load"
    />
  </div>
</template>
<script>
import { PhPlus, PhPencil, PhFingerprintSimple, PhCaretLeft, PhEye, PhTable } from '@phosphor-icons/vue';
import ProximityAddEditModal from './ProximityAddEditModal.vue';
import ProximityViewModal from './ProximityViewModal.vue'

/* 

DESC: Mostra as pessoas próximas do paciente, assim como
suas informações, e há a possibilidade de editar as informações.

*/

export default {
  components: {
    PhPlus,
    ProximityAddEditModal,
    ProximityViewModal,
    PhPencil,
    PhFingerprintSimple,
    PhCaretLeft,
    PhEye,
    PhTable
},
  data() {
    return {
      model: {
        data: [],
        hasError: false,
        message: '',
      },
      patientModel: {
        Name: '',
      }
    };
  },
  computed: {
    isMobile() {
      return this.$q.screen.xs || this.$q.screen.sm;
    },
    patientId() {
      return this.$router.currentRoute.value.params.id;
    },
  },
  mounted() {
    this.load();
  },
  methods: {
    viewQuizzes(idProximity) {
      this.$router.push(
        '/pessoas-proximas/' + idProximity + '/questionarios'
      );
    },
    load() {
      const th = this;
      th.$api.ProximityController.getPersonsByIdPatient(th.patientId).then(
        ({ data }) => {
          th.model = data;
        }
      );
      th.$api.PatientController.getById(th.patientId).then((response) => {
        th.patientModel.Name = response.data.data.Name
      })
    },
    openAddEditModal(current) {
      this.$refs.addEdit.openModal(current);
    },
    goBack() {
      this.$router.push('/paciente/id')
    },
    resetPassword({ IdPerson, Email }) {
      this.$api.PersonController.resetPassword({ IdPerson, Email })
    },
    openViewModal(current) {
      this.$refs.viewModal.openModal(current);
    }
  },
};
</script>
<style>
.go-back {
  width: auto;
  margin-right: 5px;
  height: 18px;
}

.back-page {
  display: flex;
  align-items: center;
  margin-top: 1.5rem;
  margin-left: 1.5rem;
  cursor: pointer;
  width: fit-content;
  transition: 1.5s;
}

.back-page:hover {
  filter: brightness(0.2);
}

.proximity-content {
  padding: 0rem 1.5rem;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.proximity-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.title-add-proximity button {
  background: var(--primary);
  border-radius: 8px;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: 0.2s;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  cursor: pointer;
}

.title-add-proximity button:hover {
  filter: brightness(0.8);
}

.title-go-back button {
  background: transparent;
  border: none;
  height: 1rem;
  width: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.proximity-title:has(.title-go-back) {
  justify-content: flex-start;
  align-items: center;
  gap: 1rem;
}

.error {
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
  padding: 1rem;
}

.add-proximity {
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.add-proximity button {
  background: var(--primary);
  border-radius: 99999px;
  font-size: 2rem;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 3rem;
  height: 3rem;
  transition: 0.2s;
}

.add-proximity button:hover {
  filter: brightness(0.8);
  cursor: pointer;
}

.proximity {
  border: 1px solid;
  padding: 0.8rem;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.proximity:hover {
  background-color: rgba(200, 255, 172, 0.041);
}

.proximity p {
  margin: 0;
}

.proximity button {
  border: none;
  background: none;
  cursor: pointer;
  padding: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: 0.1s;
  border-radius: 9999px;
}

.proximity button:hover {
  background: var(--neutral-dark-gray);
}

.proximity-actions {
  display: flex;
  align-items: center;
}
</style>