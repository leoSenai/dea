<template>
  <div>
    <div>
      <div class="btnVoltar" onclick="window.history.back()">
        <svg class="go-back" version="1.1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 129 129" xmlns:xlink="http://www.w3.org/1999/xlink" enable-background="new 0 0 129 129">
            <g>
                <path d="m88.6,121.3c0.8,0.8 1.8,1.2 2.9,1.2s2.1-0.4 2.9-1.2c1.6-1.6 1.6-4.2 0-5.8l-51-51 51-51c1.6-1.6 1.6-4.2 0-5.8s-4.2-1.6-5.8,0l-54,53.9c-1.6,1.6-1.6,4.2 0,5.8l54,53.9z"/>
            </g>
        </svg>
        <span>Voltar</span>
      </div>
    </div>
    <div class="proximity-content">
      <div class="proximity-title">
        <div class="title">
          <h4>Pessoas próximas de {{patientModel.Name}}</h4>
        </div>
        <div v-if="!isMobile" class="title-add-proximity">
          <button type="button" @click="openAddEditModal()">
            Adicionar
            <PhPlus />
          </button>
        </div>
      </div>
      <div v-if="model.hasError" class="error proximity">
        {{ model.message }}
      </div>
      <div
        v-for="proximity in model.data"
        v-else
        :key="proximity.Idproximity"
        class="row proximity"
      >
        <p>
          {{ proximity.Name }}
        </p>
        <div class="proximity-actions">
          <button type="button" @click="openAddEditModal(proximity)">
            <PhPencil />
          </button>
        </div>
      </div>
      <div v-if="isMobile" class="add-proximity">
        <button type="button" @click="openAddEditModal()">
          <PhPlus />
        </button>
      </div>
    </div>
    <ProximityAddEditModal ref="addEdit" @close="load" />
  </div>
</template>
<script>
import { PhPlus, PhPencil } from '@phosphor-icons/vue';
import ProximityAddEditModal from './ProximityAddEditModal.vue';

/* 

DESC: Mostra as pessoas próximas do paciente, assim como
suas informações, e há a possibilidade de editar as informações.

*/

export default {
  components: {
    PhPlus,
    ProximityAddEditModal,
    PhPencil,
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
    load() {
      const th = this;
      th.$api.ProximityController.getPersonsByIdPatient(th.patientId).then(
        ({ data }) => {
          th.model = data;
        }
      );
      th.$api.PatientController.getById(th.patientId).then((response) => {
        console.log(response)
        th.patientModel.Name = response.data.data.Name
      })
    },
    openAddEditModal(current) {
      this.$refs.addEdit.openModal(current);
    },
    goBack() {
      this.$router.push('/paciente/id');
    },
  },
};
</script>
<style>
.go-back {
  width: auto;
  margin-right: 5px;
  height: 18px;
}
.btnVoltar {
  margin-left: 1em;
  margin-top: 1em;
  margin-bottom: 10px;
  width: fit-content;
  border-radius: 15px;
  padding: 10px;
  display: flex;
  background-color: var(--primary);
  cursor: pointer;
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
  padding: 1rem;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
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
</style>