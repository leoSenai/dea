<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      Laudo final de  {{ model.Name }}
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div class="row q-mb-sm">
          <div style="width: 100%;">
            <h6>Somatório dos resultados:</h6>
            <hr>
            <div
              class="main-results"
            >
              <table
                class="sumresult-table"
              >
                <thead>
                  <tr>
                    <th><b>Nome do Questionário</b></th>
                    <th><b>Pessoa Próxima</b></th>
                    <th><b>Descrição da proximidade</b></th>
                    <th><b>Somatório</b></th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="sumresult in sumresults"
                    :key="sumresult.ProximityName"
                  >
                    <td><i>{{ sumresult.QuizDesc }}</i></td>
                    <td><i>{{ sumresult.ProximityName }}</i></td>
                    <td><i>{{ sumresult.ProximityDesc }}</i></td>
                    <td><i><b>{{ sumresult.Sum }}</b></i></td>
                  </tr>
                </tbody>
              </table>
            </div>
            <hr>
          </div>
          <q-select
            v-model="modeloAnamnese.Indicative"
            :options="indicativeOptions"
            label="Indicativo"
            option-label="label"
            option-value="value"
            label-color="primary"
            class="q-field select row select-input q-field__control"
            required
            :style="modeloAnamnese.Indicative.value==0 ? 'width: 100%;': 'width: 50%;'"
          />
          <q-select
            v-show="modeloAnamnese.Indicative.value==1"
            v-model="grau"
            :options="graus"
            label="Grau"
            option-label="label"
            option-value="value"
            label-color="primary"
            class="q-field select row select-input q-field__control"
            required
            style="width: 50%;"
          />
        </div>
      </q-form>
    </template>
    <template #modal-actions>
      <div
        class="modal-actions"
        style="gap: 2em;"
      >
        <button-primary
          outlined
          class="fill-content"
          size="sm"
          @click="closeModal"
        >
          Fechar
        </button-primary>
        <button-primary
          outlined
          class="fill-content"
          size="sm"
          style="background-color: var(--primary);color: white;"
          @click="generateLaudo"
        >
          Gerar
        </button-primary>
      </div>
    </template>
  </modal-primary>
</template>
    <script>
  import ModalPrimary from '../../components/ModalPrimary.vue';
  import ButtonPrimary from '../../components/ButtonPrimary.vue';
  import { RoleEnum } from '../../utils/Enum';
  import cookie from '../../utils/cookie';
  
    export default {
      components: {
        ModalPrimary,
        ButtonPrimary,
      },
      emits: ['close'],
      data() {
        return {
          show: false,
          sumresults: [{
            ProximityName: '',
            ProximityDesc: '',
            QuizDesc: '',
            Sum: 0
          }],
          modeloAnamnese: {
            IdAnamnese: null,
            IdPatient: null,
            IdUser: null,
            Notes: '',
            Indicative: {label: 'Não possui Transtorno do Espectro Autista', value: 0},
        },
          grau: {label: 'Selecione', value: 0},
          indicativeOptions: [{label: 'Não possui Transtorno do Espectro Autista', value: 0}, {label: 'Possui Transtorno do Espectro Autista', value: 1}],
          graus: [{label: 'Leve', value: 1}, {label: 'Moderado', value: 2}, {label: 'Severo', value: 3}],
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
          RoleEnum
        };
      },
      computed: {
        ...RoleEnum,
        typeUser () {
        const token = cookie.get('authToken')
        return cookie.getUserType(token)
        },
        idUser() {
          const token = cookie.get('authToken')
          return cookie.getUserId(token)
        },
      },
      methods: {
        openModal(current, modeloAnamnese) {
          const th = this;
          th.show = true;
          if (current) {

            th.model = current
            th.modeloAnamnese = {...modeloAnamnese}
            th.modeloAnamnese.Indicative = 
                modeloAnamnese.Indicative == 0 ? {label: 'Não possui Transtorno do Espectro Autista', value: 0} : {label: 'Possui Transtorno do Espectro Autista', value: 1};
          
            th.$api.AnamneseController.getSumResults(th.model.IdPatient).then((data) => {
              th.sumresults = data.data.data
            })
          
          }
        },
        async generateLaudo(){
            const th = this
            if((th.grau.value != 0 && th.modeloAnamnese.Indicative.value== 1) || (th.modeloAnamnese.Indicative.value == 0) ){

                if(th.modeloAnamnese.Indicative.value == 0){
                    th.grau.value = 0
                }

                await th.$api.AnamneseController.update({
                    IdAnamnese: th.modeloAnamnese.IdAnamnese,
                    IdPatient: th.modeloAnamnese.IdPatient,
                    IdUser: th.modeloAnamnese.IdUser,
                    Notes: th.modeloAnamnese.Notes,
                    Indicative: th.modeloAnamnese.Indicative.value==0 ? 0 : 1,
                })

                var response = await th.$api.AnamneseController.getLaudo(th.modeloAnamnese.IdUser, th.model.IdPatient, th.grau.value)
                console.log(response)
            }else{
                alert('O laudo só pode ser gerado quando a anamnese for conclusiva!')
            }
        },
        closeModal() {
          this.show = false;
          this.modeloAnamnese
          this.model = []
          this.$emit('close');
        },
      },
    };
    </script>
  <style>
    .q-slide.answer .q-slider__track{
      background: linear-gradient(to right, red, yellow 50%, #04df04) !important;
    }
    .q-slide.answer .q-slider__selection{
      background: transparent;
    }
    .q-slide.answer .text-primary{
      color: blue !important;
    }
  </style>
    <style scoped>
    .fill-content {
      width: 100%;
    }
  
    p{
      margin: 0;
    }

    .modal-actions {
      width: 100%;
      display: flex;
      justify-content: space-between;
    }

    .main-results{
      width: 100%;
      display: flex;
      justify-content: center;
    }

    .sumresult-table{
      text-align: center;
      overflow:scroll;
      margin: 0.5em;
      display: flexbox;
      max-height: 100px;
      border: 1px solid black !important;
      border-style: dashed;
      background-color: rgb(30, 82, 30);
      width: 100%;
    }

    .sumresult-table th, .sumresult-table td{
      padding: 0.5em;
      text-align: center;
      border: 1% groove black !important;
    }

    .sumresult-table th {
      background-color: rgb(88, 161, 88);
    }

    .sumresult-table td {
      background-color: rgb(243, 255, 243);
    }

    </style>
    