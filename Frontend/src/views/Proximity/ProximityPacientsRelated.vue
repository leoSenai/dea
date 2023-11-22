<template>
  <div class="proximity-patient-content">
    <div class="proximity-patient-title">
      <div class="title">
        <h3>Pacientes de {{ proximity.Name }}</h3>
      </div>
      <div
        v-if="!isMobile"
        class="title-add-proximity-patient"
      />
    </div>
    <div v-if="patients.length==0">
      <i>Não há pacientes vinculados a você.</i>
    </div>
    <div 
      v-for="patient in patients"
      v-else
      :key="patient.IdPatient"
      class="row proximity-patient"
    >
      <p @click="viewQuizzes(patient.IdPatient)">
        {{ patient.Name }}{{ patient.IdPatient }}
      </p>
    </div>
  </div>
</template>
<script>
import cookie from '../../utils/cookie';

export default {
  components: {
  },
  data() {
    return {
      model: {
        data: [],
        hasError: false,
        message: '',
      },
      patients: [],
      proximity: {}
    }
  },
  computed: {
    isMobile() {
      return this.$q.screen.xs || this.$q.screen.sm
    },
    personId () {
      const token = cookie.get('authToken')
      return cookie.getUserId(token)
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    load() {
      const th = this;
      th.$api.PersonController.getById(th.personId).then(response => {
        th.proximity = response.data.data
        th.$api.ProximityController.getByIdPerson(th.personId).then(({data}) => {
          data.data.forEach(({ IdPatient }) => {
            th.$api.PatientController.getById(IdPatient).then(responsePatient => {
              th.patients.push(responsePatient.data.data)
              console.log(responsePatient.data)
            })
          })
        })
      })
    },
    viewQuizzes(patientId) {
      this.$router.push(`/pessoa-proxima/${this.personId}/paciente/${patientId}/questionarios`)
    }
  },
}
</script>
<style>
.proximity-patient-actions {
  display: flex;
  gap: 0.6em;
  width: 20%;
  justify-content: flex-end;
  padding-right: 1rem;
}

.card .row{
  background-color: transparent;
}

.card div.row:hover{
  background-color: transparent !important;
}

.proximity-patient-content {
  padding: 3rem 1.5rem;
  padding-top: 0;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: .75rem;
  background-size: 50% !important;
  background: url(../../assets/imgs/home-background.svg) no-repeat;
  background-position-x:center;
  background-position-y: center;
  height: 100%;
}

.back-page {
  display: flex;
  align-items: center;
  margin-top: 1.5rem;
  margin-left: 0.5rem;
  cursor: pointer;
  transition: 1.5s;
}

.proximity-patient-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.error {
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
  padding: 1rem;
}

.proximity-patient {
  border: 1px solid var(--neutral-dark-gray);
  color: var(--neutral-dark-gray);
  padding: 0;
  border-radius: 4px;
  display: flex;
  cursor: text;
  justify-content: space-between;
  align-items: center;
}

.proximity-patient:hover {
  background-color: rgba(200, 255, 172, 0.041);
}

.proximity-patient p {
  margin: 0;
  width: 80%;
  height: 100%;
  padding-left: 1rem;
  align-items: center;
  display: flex;
  padding-top: 1rem;
  padding-bottom: 1rem;
}

.proximity-patient button {
  border: none;
  background: none;
  cursor: pointer;
  padding: .5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: .1s;
  border-radius: 9999px;
  z-index: 1;
}

.proximity-patient button:hover {
  background: var(--neutral-gray);
}

  .back-page {
    display: flex;
    align-items: center;
    margin-top: 1.5rem;
    margin-left: 0.5rem;
    cursor: pointer;
    width: fit-content;
    transition: 1.5s;
  }
  
  .proximity-patient-title {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  
  .title-add-proximity-patient button {
    background: var(--primary);
    border-radius: 8px;
    border: none;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: .2s;
    gap: .5rem;
    padding: 1rem 2rem;
    color: white;
    cursor: pointer;
  }
  
  .title-add-proximity-patient button:hover {
    filter: brightness(0.8);
  }
  
  .error {
    border: 1px solid var(--neutral-dark-gray);
    border-radius: 4px;
    padding: 1rem;
  }
  
  .add-proximity-patient {
    position: absolute;
    bottom: 1rem;
    right: 1rem;
  }
  
  .add-proximity-patient button {
    background: var(--primary);
    border-radius: 99999px;
    font-size: 2rem;
    border: none;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 3rem;
    height: 3rem;
    transition: .2s;
  }
  
  .add-proximity-patient button:hover {
    filter: brightness(0.8);
  }
  
  .proximity-patient {
    border: 1px solid var(--neutral-dark-gray);
    color: var(--neutral-dark-gray);
    padding: 0;
    border-radius: 4px;
    display: flex;
    cursor: text;
    justify-content: space-between;
    align-items: center;
  }
  
  .proximity-patient:hover {
    background-color: rgba(200, 255, 172, 0.041);
  }
  
  .proximity-patient p {
    margin: 0;
    width: 80%;
    height: 100%;
    padding-left: 1rem;
    align-items: center;
    display: flex;
    padding-top: 1rem;
    padding-bottom: 1rem;
  }
  
  .proximity-patient button {
    border: none;
    background: none;
    cursor: pointer;
    padding: .5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: .1s;
    border-radius: 9999px;
    z-index: 1;
  }
  
  .proximity-patient button:hover {
    background: var(--neutral-gray);
  }
  
  .row.proximity-patient{
    z-index: 0;
    cursor: pointer;
  }
  </style>
