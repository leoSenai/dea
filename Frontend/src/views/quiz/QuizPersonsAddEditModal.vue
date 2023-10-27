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
          :key="filiated.Name"
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
        <div
          v-for="(filiated, i) in filiatedsToAdd"
          v-show="addFiliatedQSelectVisible"
          :key="i"
          class="select-filiated"
        >
          <q-select
            v-model="filiatedsToAdd[i].Name"
            :options="filiatedToAddList"
            option-label="Name"
            label="&nbsp Adicionar filiado"
            placeholder="Selecione uma opção"
            class="select-quasar newFiliated"
            max-options="3"
            @update:model-value="updateFiliatedToAddList(filiatedsToAdd[i].Name)"
          />
          <button
            v-if="filiateds.length > 0"
            type="button"
            class="remove-person rounded-full"
            @click="closeAddQSelect()"
          >
            <b>X</b>
          </button>
        </div>
        <div
          class="row addFiliated"
        >
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
          @click="updateQuizFiliated"
        >
          Atualizar lista
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
      filiatedToAddList: [],
      filiatedToRemoveList: [],
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
      filiateds: [],
      filiatedsToAdd: [],
    };
  },
  methods: {
    closeAddQSelect(){
      this.filiatedsToAdd = []
    },
    updateFiliatedToAddList(filiated_obj){
      const th = this

      //filiated_obj.Name
      //filiated_obj.Type

      var indexToRemove = th.filiatedToAddList.findIndex((fili) => {
        return (fili.Name === filiated_obj.Name && fili.Type === filiated_obj.Type);
      });
      if(indexToRemove!=-1){console.log(indexToRemove) // se achou o filiado na filiatedToAddList, remove-o
        th.filiatedToAddList.splice(indexToRemove, 1)
        th.filiatedToRemoveList.push(filiated_obj)
        th.filiateds.push(filiated_obj)
      }

      th.addFiliatedQSelectVisible = false
      th.filiatedsToAdd = []
      //th.filiatedToAddList = []
      th.filiatedToRemoveList = []
      th.patientRemoved = []
      th.personRemoved = []

    },
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
      }
    },
    addFiliated() {
      const th = this
      if(th.filiatedsToAdd.length==1){
        return
      }
      th.addFiliatedQSelectVisible = true
      th.filiatedsToAdd.push({ Name: '' })

      if(th.filiatedToAddList.length===0){

        //PROCURA TODOS PACIENTES E ADICIONA A LISTA DE FILIADOS PARA ADICIONAR
        th.$api.PatientController.getAll().then(({ data }) => {
  
          var patientList =  data.data //array com os objetos pacientes
          var alreadyIncluded;
          //adapt each object to a filiated model and push to 'filiatedToAddList'
          for(var i in patientList){
  
            alreadyIncluded = false
            for(var l=0;l<th.filiateds.length;l++){
              if(th.filiateds[l].Name==patientList[i].Name && th.filiateds[l].Type == 'Patient'){
                alreadyIncluded = true;
              }
            }
  
            if(!alreadyIncluded){
              th.filiatedToAddList.push({Name: patientList[i].Name, Type: 'Patient', IdQuiz: th.model.IdQuiz, IdFiliated: patientList[i].IdPatient})
            }
          }
        });
  
        //PROCURA TODOS PESSOAS E ADICIONA A LISTA DE FILIADOS PARA ADICIONAR
        th.$api.PersonController.getAll().then(({ data }) => {
  
          var personList =  data.data //array com os objetos pacientes
          var alreadyIncluded;
          //adapt each object to a filiated model and push to 'filiatedToAddList'
          for(let i in personList){
  
            alreadyIncluded = false
            for(let l=0;l<th.filiateds.length;l++){ 
              if(th.filiateds[l].Name==personList[i].Name && th.filiateds[l].Type == 'Person'){
                alreadyIncluded = true; 
              }
            }
  
            if(!alreadyIncluded){
              th.filiatedToAddList.push({Name: personList[i].Name, Type: 'Person', IdQuiz: th.model.IdQuiz, IdFiliated: personList[i].IdPerson})
            }
          }
        });
      }

      //AGORA TEMOS JA A LISTA DE PACIENTES/PESSOAPROX PARA ADD NOVO
      //ANTES DE CONTINUAR: FAZER UPDATE DA LISTA DE DE ADD caso o maluco resolver excluir algum, ou add, enquanto tem outros q-selects em aberto.

    },
    removeFiliated(filiated){
      console.log('filiado removendo...')

      const th = this
      if(filiated.Name == ''){
        return
      }
      console.log(filiated)
      if(filiated.Type==='Patient'){
        th.removePatient(filiated)
      }else{
        th.removePerson(filiated)
      }
    },
    removePatient(filiated) {

      const th = this
      var indexToRemove = th.filiatedToAddList.findIndex((fili) => {
        return fili.Name === filiated.Name;
      });
      if(indexToRemove==-1){
        th.filiatedToAddList.push({Name: filiated.Name, Type: filiated.Type, IdQuiz: th.model.IdQuiz})
        th.patientRemoved.push(filiated)
      }

      var index2Remove = th.filiateds.findIndex((fili) => {
        return (fili.Name === filiated.Name && fili.Type==='Patient');
      });
      if(index2Remove!=-1){
        th.filiateds.splice(index2Remove, 1)
        th.filiatedsRemoved.push(filiated)
      }

    },
    removePerson(filiated) {
      const th = this
      var indexToRemove = th.filiatedToAddList.findIndex((fili) => {
        return fili.Name === filiated.Name;
      });
      if(indexToRemove==-1){
        th.filiatedToAddList.push({Name: filiated.Name, Type: filiated.Type, IdQuiz: th.model.IdQuiz})
        th.personRemoved.push(filiated)
      }

      var index2Remove = th.filiateds.findIndex((fili) => {
        return (fili.Name === filiated.Name && fili.Type==='Person');
      });
      if(index2Remove!=-1){
        th.filiateds.splice(index2Remove, 1)
        th.filiatedsRemoved.push(filiated)
      }

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
      this.$emit('close');
    },
    updateQuizFiliated() {
      
      const th = this;
      const filiateds = th.filiateds

      console.log(filiateds)

      //th.filiateds = [{Name: '', Type: '', IdQuiz: 0}]
      th.$api.PatientHasQuizController.update(filiateds)
      th.$api.ProximityHasQuizController.update(filiateds)

    },
  },
};
</script>
<style scoped>
.select-filiated{
  display: flex;
  align-items: center;
  margin-bottom: 8px;
  z-index: 0;
}

.newFiliated{
  border: #fff solid 1px;
  border-radius: 5px;
  width: -webkit-fill-available;
  z-index: 0;
}
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
<style>
  .select-filiated .q-field__append.q-field__marginal.row.no-wrap.items-center.q-anchor--skip{
    padding-right: 3rem !important;
  }
</style>