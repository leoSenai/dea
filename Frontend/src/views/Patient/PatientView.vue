<template>
  <div class="patient-content">
    <div class="patientView-content">
      <div
        class="back-page"
        onclick="goBack"
      >
        <PhCaretLeft color="#656565" />
        Voltar
      </div>
      <div class="patientView-header">
        <div class="patient-info-editable">
          <q-input
            v-model="model.Name"
            class="inputEditable"
            placeholder="Nome"
          />
          <q-input
            v-model="model.Email"
            class="inputEditable"
            placeholder="Email"
          />
          <q-input
            v-model="model.Address"
            class="inputEditable"
            placeholder="Endereço"
          />
          <q-input
            v-model="model.Phone"
            class="inputEditable"
            placeholder="Telefone"
          />
          <!--RECEM NASCIDO SELECT-->
          <q-select
            v-model="opcaoNewBorn"
            class="select-quasar"
            :options="opcoesSimNao"
            label="&nbsp Recém nascido?"
            placeholder="Selecione uma opção"
            @update:model-value="changeNewBornValue()"
          />
          <q-input
            v-model="model.Cid10"
            type="number"
            class="inputEditable"
            placeholder="CID10"
          />
          <!--ATIVO SELECT-->
          <q-select
            v-model="opcaoAtivo"
            class="select-quasar"
            :options="opcoesSimNao"
            label="&nbsp Ativo?"
            placeholder="Selecione uma opção"
            @update:model-value="changeAtivoValue()"
          />
          <p class="alter-password">
            <a
              style="font-size: 12px"
              class="change-pass"
              @click="changePassword(model.IdPatient)"
            >
              Alterar senha
            </a>
          </p>
        </div>
        <div class="patient-info">
          <h4>{{ model.Name }}</h4>
          <p>{{ model.Cpf }}</p>
          <p>{{ model.Email }}</p>
        </div>
        <div class="edit-button-div">
          <button-primary
            class="editBtn" 
            @click="editPatient(model.IdPatient)"
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
  </div>
</template>
<script>
import ButtonPrimary from '../../components/ButtonPrimary.vue';
import { PhCaretLeft, PhCaretRight, PhPencil } from '@phosphor-icons/vue';
import cookie from '../../cookie';

export default {
  components: {
    ButtonPrimary,
    PhPencil,
    PhCaretLeft,
    PhCaretRight
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
  mounted() {
    const th = this;

    var contentElement = document.getElementsByClassName('content')[0];
    contentElement.style.overflow = 'auto';

    const urlString = window.location.href;

    // Analisando a string manualmente
    const params = {};
    const urlParts = urlString.split('?');
    if (urlParts.length > 1) {
      const queryParams = urlParts[1].split('&');
      queryParams.forEach((param) => {
        const [key, value] = param.split('=');
        params[key] = value;
      });
    }

    // Obtendo o valor do parâmetro 'id'
    const id = params['id'];
    th.model.IdPatient = id;

    // Obtendo o valor do parâmetro 'edit'
    const edit = params['edit'];

    //Style shits
    var element = document.getElementsByClassName('editIcon')[0];
    var elementAlterPassword =
      document.getElementsByClassName('alter-password')[0];
    if (edit !== undefined && edit == 'true') {
      //Quando para editar
      element.style.display = 'none';
      elementAlterPassword.style.display = 'block';
      th.edit = true;
      th.editOrSave = 'Salvar';
      document.getElementsByClassName('patient-info')[0].style.display = 'none';
      document.getElementsByClassName(
        'patient-info-editable'
      )[0].style.display = 'block';
      document.getElementsByClassName('patientView-header')[0].style.display =
        'block';
    } else {
      //Quando para salvar
      th.edit = false;
    }

    //faz requisição para pegar info do paciente
    th.$api.PatientController.getById(id).then(({ data }) => {
      th.model = data.data;
      th.opcaoNewBorn = th.model.NewBorn == 1 ? 'Sim' : 'Não';
      th.opcaoAtivo = th.model.Active == '1' ? 'Sim' : 'Não';
    });

    var idUser = cookie.getUserId(cookie.get('authToken'));
    th.$api.AnamneseController.getByIdUserPatient({
      IdUser: idUser,
      IdPatient: id,
    }).then((response) => {
      if (response.statusText !== 'No Content') {
        //Found so update
        var data = response.data.data;

        //Bind the anamnese object
        th.anamneseModel.Notes = data.Notes;
        th.anamneseModel.IdAnamnese = data.IdAnamnese;
        th.anamneseModel.IdPatient = data.IdPatient;
        th.anamneseModel.IdUser = data.IdUser;
        if (data.Indicative !== 0) {
          th.campoAnamneseDesabilitado = true;
        }
      }
    });
  },
  methods: {
    editPatient(id) {
      console.log(id);
      //Style shits
      const th = this;
      var element = document.getElementsByClassName('editIcon')[0];
      var elementAlterPassword =
        document.getElementsByClassName('alter-password')[0];
      if (th.edit == false && th.editOrSave == 'Editar') {
        //Quando for para editar
        element.style.display = 'none';
        elementAlterPassword.style.display = 'block';
        th.edit = true;
        th.editOrSave = 'Salvar';
        document.getElementsByClassName('patient-info')[0].style.display =
          'none';
        document.getElementsByClassName(
          'patient-info-editable'
        )[0].style.display = 'block';
        document.getElementsByClassName('patientView-header')[0].style.display =
          'block';
      } else if (th.edit == true && th.editOrSave == 'Salvar') {
        //Quando for para salvar
        element = document.getElementsByClassName('editIcon')[0];
        element.style.display = 'block';
        elementAlterPassword.style.display = 'none';
        th.edit = false;
        th.editOrSave = 'Editar';
        document.getElementsByClassName('patient-info')[0].style.display =
          'block';
        document.getElementsByClassName(
          'patient-info-editable'
        )[0].style.display = 'none';
        document.getElementsByClassName('patientView-header')[0].style.display =
          'flex';

        this.savePatientData();
        if (this.campoAnamneseDesabilitado == false) {
          this.saveAnamnese();
        }
      }
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
    goBack() {
      var contentElement = document.getElementsByClassName('content')[0];
      contentElement.style.overflow = 'hidden';
      window.history.back();
    },
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
  height: auto;
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
</style>
