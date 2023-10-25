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
          v-for="(filiated) in filiateds"
          :key="filiated.IdFiliated"
          class="row q-mb-sm quiz"
        >
          <span>{{ filiated.Name + (filiated.Type==='Patient' ? ' (paciente)' : ' (pessoa próxima)') }}</span>
          <button
            v-if="filiateds.length > 0"
            type="button"
            class="remove-person rounded-full"
            @click="removeFiliated(filiated)"
          >
            <b>X</b>
          </button>
        </div>
        <div class="select-filiated">
          <q-select
            v-model="filiatedToAdd"
            :options="filiatedToAddList"
            label="&nbsp Adicionar filiado"
            placeholder="Selecione uma opção"
            class="select-quasar"
          />
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
      addFiliatedQSelectVisible: false,
      filiatedToAdd: '',
      filiatedToAddList: [],
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
      filiatedsRemoved: [],
      filiateds: [/*{
        IdFiliated: ,
        IdQuiz: 0,
        Finished: 0,
        Answers: '',
        AnsweredIn: '',
        Type: '',
      }*/],
      filiatedsToAdd: [],
    };
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

                var personData = data.data;
                var newFiliated = {
                  Name: personData.Name,
                  IdFiliated: personData.IdPerson,
                  IdQuiz: th.model.IdQuiz,
                  Finished: th.proximitysHasQuiz[i].Finished,
                  Answers: th.proximitysHasQuiz[i].Answers,
                  AnsweredIn: th.proximitysHasQuiz[i].AnsweredIn,
                  Type: 'Person'
                }
                th.filiateds.push(newFiliated)

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

                  var patientData = data.data;
                  var newFiliated = {
                    Name: patientData.Name,
                    IdFiliated: patientData.IdPatient,
                    IdQuiz: th.model.IdQuiz,
                    Finished: th.patientsHasQuiz[i].Finished,
                    Answers: th.patientsHasQuiz[i].Answers,
                    AnsweredIn: th.patientsHasQuiz[i].AnsweredIn,
                    Type: 'Patient'
                  }
                  th.filiateds.push(newFiliated)

                });

              }
            }
        });

        //PROCURA TODOS PACIENTES E ADICIONA A LISTA DE FILIADOS PARA ADICIONAR
        th.$api.PatientController.getAll().then(({ data }) => {

          var patientList =  data.data //array com os objetos pacientes
          var alreadyIncluded;
          //adapt each object to a filiated model and push to 'filiatedToAddList'
          for(let i in patientList){

            alreadyIncluded = false
            for(let l=0;l<th.filiateds.length;l++){
              if(th.filiateds[l].IdFiliated===patientList[i].IdPatient && th.filiateds[l].Type === 'Patient'){
                alreadyIncluded = true;
              }
            }

            if(!alreadyIncluded){
              th.filiatedToAddList.push({
                Name: patientList[i].Name,
                IdFiliated: patientList[i].IdPatient,
                IdQuiz: th.model.IdQuiz,
                Finished: 0,
                Answers: '',
                AnsweredIn: '',
                Type: 'Patient',
  
              })
            }
          }
        });

        //PROCURA TODOS PACIENTES E ADICIONA A LISTA DE FILIADOS PARA ADICIONAR
        th.$api.PersonController.getAll().then(({ data }) => {

          var personList =  data.data //array com os objetos pacientes
          var alreadyIncluded;
          //adapt each object to a filiated model and push to 'filiatedToAddList'
          for(let i in personList){

            alreadyIncluded = false
            for(let l=0;l<th.filiateds.length;l++){ 
              //TODO: NAO TA FUNFANDO
              if(th.filiateds[l].IdFiliated===personList[i].IdPerson && th.filiateds[l].Type === 'Person'){
                alreadyIncluded = true; 
              }
            }

            if(!alreadyIncluded){
              th.filiatedToAddList.push({
                Name: personList[i].Name,
                IdFiliated: personList[i].IdPerson,
                IdQuiz: th.model.IdQuiz,
                Finished: 0,
                Answers: '',
                AnsweredIn: '',
                Type: 'Person',

              })
            }
          }
        });

      }
    },
    addFiliated() {
      const th = this
      th.addFiliatedQSelectVisible = true
    },
    removeFiliated(filiated){
      const th = this
      if(filiated.Type==='Patient'){
        th.removePatient(filiated)
      }else{
        th.removePerson(filiated)
      }
    },
    removePatient(filiated) {

      const th = this
      //var indexToRemove = th.patientsList.findIndex((pat) => {
      //  return pat.IdPatient === filiated.IdFiliated;
      //});
      //th.patientsList.splice(indexToRemove, 1)
      th.patientRemoved.push(filiated)

      var index2Remove = th.filiateds.findIndex((fili) => {
        return (fili.IdFiliated === filiated.IdFiliated && fili.Type==='Patient');
      });
      th.filiateds.splice(index2Remove, 1)
      th.filiatedsRemoved.push(filiated)

    },
    removePerson(filiated) {
      const th = this
      //var indexToRemove = th.personsList.findIndex((per) => {
      //  return per.IdPerson === filiated.IdFiliated;
      //});
      //th.personsList.splice(indexToRemove, 1)
      th.personRemoved.push(filiated)

      var index2Remove = th.filiateds.findIndex((fili) => {
        return (fili.IdFiliated === filiated.IdFiliated && fili.Type==='Person');
      });
      th.filiateds.splice(index2Remove, 1)
      th.filiatedsRemoved.push(filiated)

    },
    closeModal() {
      this.show = false;
      this.addFiliatedQSelectVisible = false;
      this.model = {
        Name: '',
        IdQuiz: null,
        Interval: 5,
      };
      this.proximitysHasQuiz = [];
      this.patientsHasQuiz = [];
      this.personsList = [];
      this.patientsList = [];
      this.filiateds = [];
      this.patientRemoved = [];
      this.personRemoved = []
      this.filiatedsRemoved = [];
      this.filiatedsToAdd = [];
      this.filiatedToAddList = [];
      this.filiatedToAdd = '';
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
  