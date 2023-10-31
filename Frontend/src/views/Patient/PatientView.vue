<template>
  <div class="patient-content">
    <div class="patientView-content">
      <div
        class="back-page"
        @click="goBack"
      >
        <PhCaretLeft color="#656565" />
        Voltar
      </div>
      <div class="patientView-header">
        <div class="patient-info">
          <h4>{{ model.Name }}</h4>
          <p>{{ model.Cpf }}</p>
          <p>{{ model.Email }}</p>
          <p class="reset-password" @click="resetPassword">
            Redefinir Senha
          </p>
        </div>
        <div class="edit-button-div">
          <button-primary
            class="editBtn" 
            @click="editPatient"
          >
            <PhPencil class="editIcon" />
            {{ editOrSave }}
          </button-primary>
        </div>
      </div>
    </div>
    <section>
      <button-primary 
        class="nextPersonView"
        @click="showNextPersons()" 
      >
        Ver pessoas próximas
        <PhCaretRight />
      </button-primary>
      <h5>Anamnese</h5>
      <q-editor
        v-model="anamneseModel.Notes"
        :disable="campoAnamneseDesabilitado"
        class="textarea"
        label="Escreva a anamnese"
        type="textarea"
        :rows="15"
      />
      <q-checkbox 
        v-model="analiseConclusiva"
        class="analiseConclusiva"
      >
        Análise conclusiva
      </q-checkbox>
      <div class="gerar-laudo-div">
        <button-primary 
          class="gerar-laudo"
          @click="gerarLaudo()"
        >
          Gerar Laudo
        </button-primary>
      </div>
    </section>
    <PatientsAddEditModal 
      ref="addEdit"
      @close="load" 
    />
  </div>
</template>
<script>
import ButtonPrimary from '../../components/ButtonPrimary.vue';
import { PhCaretLeft, PhCaretRight, PhPencil } from '@phosphor-icons/vue';
import cookie from '../../cookie';
import PatientsAddEditModal from './PatientsAddEditModal.vue';

export default {
  components: {
    ButtonPrimary,
    PhPencil,
    PhCaretLeft,
    PhCaretRight,
    PatientsAddEditModal
},
  data() {
    return {
      opcoesSimNao: ['Sim', 'Não'],
      analiseConclusiva: false,
      edit: false,
      opcaoNewBorn: '',
      opcaoAtivo: '',
      editOrSave: 'Editar',
      model: {
        IdPatient: null,
        Name: '',
        Email: '',
        Cpf: '',
        Address: '',
        Phone: '',
        BornDate: '',
        Sex: '',
        NewBorn: null,
        Cid10: null,
        Active: null,
      },
      anamneseModel: {
        IdAnamnese: null,
        IdPatient: null,
        IdUser: null,
        Notes: '',
        Indicative: 0,
      },
      campoAnamneseDesabilitado: false,
    };
  },
  mounted () {
    const th = this
    th.load()
  },
  methods: {
    goBack () {
      this.$router.push('/pacientes')
    },
    editPatient() {
      const th = this;
      th.$refs.addEdit.openModal(th.model)
    },
    savePatientData() {
      const th = this;
      th.model.Cid10 = th.model.Cid10 ? parseInt(th.model.Cid10) : 0;
      th.$api.PatientController.update(th.model);
    },
    saveAnamnese() {
      const th = this;
      th.anamneseModel.IdPatient = th.model.IdPatient;
      th.anamneseModel.IdUser = cookie.getUserId(cookie.get('authToken'));
      th.$api.AnamneseController.update(th.anamneseModel);
    },
    changePassword(id) {
      alert(id + ' - FUTURA IMPLEMENTACAO');
    },
    changeNewBornValue() {
      const th = this;
      if (th.opcaoNewBorn == 'Sim') {
        th.model.NewBorn = 1;
      } else {
        th.model.NewBorn = 0;
      }
    },
    gerarLaudo() {
      //const th = this
    },
    changeAtivoValue() {
      const th = this;
      if (th.opcaoAtivo == 'Sim') {
        th.model.Active = 1;
      } else {
        th.model.Active = 0;
      }
    },
    showNextPersons() {
      this.$router.push(
        '/paciente/' + this.model.IdPatient + '/pessoas-proximas'
      );
    },
    load () {
      const th = this;
      const idPatient = th.$router.currentRoute.value.query.id
      th.$api.PatientController.getById(idPatient).then(({data}) => {
        th.model = { ...data.data }
      })
    },
    resetPassword () {
      const th = this;
      th.$api.PatientController.resetPassword(th.model.IdPatient)
    }
  },
};
</script>
<style scoped>
.select-quasar {
  margin-bottom: 10px;
}

.inputEditable {
  margin-bottom: 10px;
}

.patient-info-editable {
  display: none;
  width: -webkit-fill-available;
}

.patient-info {
  display: block;
}

.alter-password {
  display: none;
}

.change-pass {
  text-decoration: underline;
  text-decoration-color: var(--primary);
  cursor: pointer;
  width: 120px;
  display: block;
  margin-top: 7px;
  padding-top: 7px;
  padding-bottom: 5px;
  font-weight: bold;
  color: var(--primary);
}

.edit-button-div {
  display: flex;
  justify-content: center;
  align-items: center;
}

.edit-button-div .editBtn {
  height: 55px;
}

.gerar-laudo-div {
  width: 100%;
  display: flex;
  justify-content: center;
}

.gerar-laudo {
  margin: 10px;
  margin-top: 0;
}

.analiseConclusiva {
  margin-top: 10px;
  margin-bottom: 10px;
}

.textarea {
  border-radius: 10px;
  border: 2px solid green;
  min-height: 30vh;
  padding: 10px;
  font-family: Arial, Helvetica, sans-serif;
  font-size: medium;
}

h5 {
  padding: 15px;
}

.nextPersonView {
  border: 1px solid var(--primary);
  background-color: var(--neutral-light-gray);
  color: green;
  margin: 10px;
  width: -webkit-fill-available;
  margin-bottom: 20px;
  transition: .5s;
}

.nextPersonView:hover {
  filter: brightness(0.8);
}

.patient-content {
  width: 99%;
  overflow-y: auto;
  padding: 20px;
  padding-top: 0;
  margin-left: 0.5%;
  margin-right: 0.5%;
  font-family: 'Roboto', '-apple-system', 'Helvetica Neue', Helvetica, Arial,
    sans-serif;
  display: block;
  justify-content: space-between;
}

section {
  display: block;
  width: 99%;
  height: fit-content;
  margin-top: 2px;
  margin-left: 0.5%;
  margin-right: 0.5%;
  padding: 5px;
  justify-content: center;
}

.patientView-content {
  width: 100%;
  height: auto;
}

.go-back {
  width: auto;
  margin-right: 5px;
  height: 18px;
}

.goNext {
  width: auto;
  margin-right: 5px;
  height: 18px;
  fill: green;
  transform: scaleX(-1);
}

.patientView-header {
  width: 99%;
  height: auto;
  padding: 20px;
  margin-left: 0.5%;
  margin-right: 0.5%;
  font-family: 'Roboto', '-apple-system', 'Helvetica Neue', Helvetica, Arial,
    sans-serif;
  display: flex;
  justify-content: space-between;
}

.patientView-header section {
  display: block;
  height: 300px;
}

.patientView-header p {
  margin: 0;
  font-weight: 300;
}

.editIcon {
  margin-left: -10px;
}

.back-page {
  display: flex;
  align-items: center;
  margin-top: 1.5rem;
  cursor: pointer;
  transition: 1.5s;
}

.back-page:hover {
  filter: brightness(0.2);
}

p.reset-password {
  color: var(--primary);
  font-weight: 600;
  font-size: .875rem;
  cursor: pointer;
}

p.reset-password:hover {
  filter: brightness(0.8);
}
</style>
