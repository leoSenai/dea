<template>
  <modal-primary 
    v-model="show" 
    @close="closeModal"
  >
    <template #modal-title>
      Filiados ao questionário '{{ model.Name }}'
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div 
          v-for="(person) in personsList"
          :key="person.IdPerson"
          class="row q-mb-sm quiz"
        >
          <span>{{ person.Name + ' (pessoa próxima)' }}</span>
          <button
            v-if="personsList.length > 0"
            type="button"
            class="remove-person rounded-full"
            @click="removePerson(person)"
          >
            <b>X</b>
          </button>
        </div>
        <div 
          v-for="(patient) in patientsList"
          :key="patient.IdPatient"
          class="row q-mb-sm quiz"
        >
          <span>{{ patient.Name + ' (paciente)' }}</span>
          <button
            v-if="patientsList.length > 0"
            type="button"
            class="remove-patient rounded-full"
            @click="removePatient(patient)"
          >
            <b>X</b>
          </button>
        </div>
        <div 
          v-for="(filiated) in filiatedsToAdd"
          :key="filiated.type=='patient' ? filiated.IdPatient : filiated.IdPerson"
          class="row q-mb-sm quiz"
        >
          <span>{{ filiated.Name + (filiated.type=='patient' ? ' (paciente)' : ' (pessoa próxima)') }}</span>
          <button
            v-if="filiateds.length > 0"
            type="button"
            class="remove-patient rounded-full"
            @click="removeFiliatedToAdd(filiated)"
          >
            <b>X</b>
          </button>
        </div>
        <div class="row addFiliated">
          <button-primary
            class="fill-content"
            outlined
            @click="addFiliated"
          >
            Adicionar filiação
            <template #after-label>
              <PhPlus
                class="primary"
                weight="bold"
              />
            </template>
          </button-primary>
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
          v-if="model.IdQuiz"
          size="sm"
          type="submit"
        >
          Atualizar lista
        </button-primary>
        <button-primary
          v-else
          size="sm"
          type="submit"
          @click="createQuizPerson"
        >
          Cadastrar
        </button-primary>
      </div>
    </template>
  </modal-primary>
</template>
  <script>
import ModalPrimary from '../../components/ModalPrimary.vue';
import ButtonPrimary from '../../components/ButtonPrimary.vue';

//TODO: toda la mierda

//Existe o ProximityHasQuiz
//e o PatientHasQuiz -ainda n tem controller

export default {
  components: {
    ModalPrimary,
    ButtonPrimary,
  },
  emits: ['close'],
  data() {
    return {
      show: false,
      model: {
        Name: '',
        IdQuiz: null,
        Interval: 5,
        Created: null,
        Updated: null,
      },
      proximitysHasQuiz: [{
        ProximityIdPatient: 0,
        ProximityIdPerson: 0,
        IdQuiz: 0,
        Finished: 0,
        Answers: '',
        AnsweredIn: '',
      }],
      patientsHasQuiz: [{
        IdPatient: 0,
        IdQuiz: 0,
        Finished: 0,
        Answers: '',
        AnsweredIn: '',
      }],
      patientsList: [],
      personsList: [],
      patientRemoved: [],
      personRemoved: [],
      filiateds: [/*{
        IdPatient: 0,
        IdPerson: 0,
        IdQuiz: 0,
        Finished: 0,
        AnsweredIn: '',
      }*/],
      filiatedsToAdd: [],
    };
  },
  mounted(){
    
  },
  methods: {
    openModal(current) {
      
      const th = this;
      th.show = true;

      if (current) {
        th.model = { ...current, Interval: +current.Interval };

        //PROCURA PESSOAS PROXIMAS QUE TENHAM FILIACAO COM O QUIZ
        th.$api.ProximityHasQuizController.getByIdQuiz(th.model.IdQuiz).then(({ data }) => {
          if(data.data!==''){
           
            const filteredProximityHasQuiz = data.data.map((el) => {

              

              return {
                ProximityIdPatient: el.ProximityIdPatient,
                ProximityIdPerson: el.ProximityIdPerson,
                IdQuiz: el.IdQuiz,
                Finished: el.Finished,
                Answers: el.Answers,
                AnsweredIn: el.AnsweredIn
              };
            });
            
            th.proximitysHasQuiz =
            filteredProximityHasQuiz.length > 0
              ? filteredProximityHasQuiz
              : [];

            for(let i in th.proximitysHasQuiz){

              th.$api.PersonController.getById(th.proximitysHasQuiz[i].ProximityIdPerson).then(({data})=> {
                th.personsList.push(data.data)
              });

            }

          }
        });

        //PROCURA PACIENTES QUE TENHAM FILIACAO COM O QUIZ
        th.$api.PatientHasQuizController.getByIdQuiz(th.model.IdQuiz).then(({ data }) => {
          if(data.data!==''){
            
            const filteredPatientHasQuiz = data.data.map((el) => {
              return {
                IdPatient: el.IdPatient,
                IdQuiz: el.IdQuiz,
                Finished: el.Finished,
                Answers: el.Answers,
                AnsweredIn: el.AnsweredIn
              };
            });
            
            th.patientsHasQuiz =
            filteredPatientHasQuiz.length > 0
              ? filteredPatientHasQuiz
              : [];

              for(let i in th.patientsHasQuiz){

                th.$api.PatientController.getById(th.patientsHasQuiz[i].IdPatient).then(({data})=> {
                  th.patientsList.push(data.data)
                });

              }
            }
        });
      }

      

    },
    addFiliated() {
      if (this.filiateds.length>0) {
        this.questions.push({
          
        });
      }
    },
    removeFiliatedToAdd(){

    },
    removePatient(patient) {
      const th = this
      var indexToRemove = th.patientsList.findIndex((pat) => {
        return pat.IdPatient === patient.IdPatient;
      });
      th.patientsList.splice(indexToRemove, 1)
      th.patientRemoved.push(patient)
    },
    removePerson(person) {
      const th = this
      var indexToRemove = th.personsList.findIndex((per) => {
        return per.IdPerson === person.IdPerson;
      });
      th.patientsList.splice(indexToRemove, 1)
      th.personRemoved.push(person)
    },
    closeModal() {
      this.show = false;
      this.model = {
        Name: '',
        IdQuiz: null,
        Interval: 5,
      };
      this.proximitysHasQuiz = [];
      this.patientsHasQuiz = [];
      this.personsList = [];
      this.patientsList = [];
      this.$emit('close');
    },
    createQuizPerson() {
      const th = this;
      th.$refs.form.validate().then((success) => {
        if (th.questions.some(({ Desc }) => !Desc || Desc.trim() === '')) {
          alert(
            'Não é possível criar questionários com questões sem descrição!'
          );
          return;
        }
        if (success) {
          th.$api.QuizController.insert({
            ...th.model,
            Interval: th.model.Interval.toString(),
          })
            .then(({ data }) => {
              th.model = data.data;
              if (th.model.IdQuiz) {
                th.questions.forEach((question) => {
                  const questionDto = {
                    IdQuiz: th.model.IdQuiz,
                    IdQuestion: question.IdQuestion,
                    Desc: question.Desc.trim(),
                  };
                  th.$api.QuestionController.insert(questionDto);
                });
              }
              return;
            })
            .then(() => {
              th.closeModal();
            });
        }
      });
    },
    updateQuiz() {
      const th = this;
      th.$refs.form.validate().then((success) => {
        if (success) {
          if (th.questions.some(({ Desc }) => !Desc || Desc.trim() === '')) {
            alert(
              'Não é possível atualizar questionários com questões sem descrição!'
            );
            return;
          }
          th.$api.QuizController.update({
            ...th.model,
            Interval: th.model.Interval.toString(),
          })
            .then(() => {
              th.questions.forEach((question) => {
                const questionDto = {
                  IdQuiz: th.model.IdQuiz,
                  IdQuestion: question.IdQuestion,
                  Desc: question.Desc.trim(),
                };
                if (questionDto.IdQuestion) {
                  th.$api.QuestionController.update(questionDto);
                } else {
                  th.$api.QuestionController.insert(questionDto);
                }
              });
              return;
            })
            .then(() => {
              th.questionRemoved.forEach(({ IdQuestion }) => {
                th.$api.QuestionController.delete(IdQuestion);
              });
              return;
            })
            .then(() => {
              th.closeModal();
            });
        }
      });
    },
  },
};
</script>
<style scoped>
.addFiliated {
  margin-top: 20px;
}
.remove-person, .remove-patient{
  position: absolute;
  right: 1em;
  width: 1.5rem;
  height: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--neutral-dark-gray);
  flex: 1;
  text-align: end;
  cursor: pointer;
  transition: 0.3s;
  background: #fff;
  border: none;
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 99999px;
  cursor: pointer;
  z-index: 1;
  padding: 4px
}

.fill-content {
  width: 100%;
}

.question {
  position: relative;
}

.remove-question {
  position: absolute;
  top: 1.8em;
  right: 1em;
  width: 1.5rem;
  height: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;

  transform: translateY(-50%);
  color: var(--neutral-dark-gray);
  flex: 1;
  text-align: end;
  cursor: pointer;
  transition: 0.3s;
  background: #fff;
  border: none;
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 99999px;
  cursor: pointer;
  z-index: 1;
}

.remove-question:hover {
  color: var(--others-red-600);
  border-color: var(--others-red-600);
}

.interval-field {
  border: 1px solid black;
  padding: 0.5rem 1rem;
  color: var(--neutral-dark-gray);
  border-radius: 4px;
}

.modal-actions {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
</style>
  