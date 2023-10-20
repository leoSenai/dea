<template>
    <div class="patient-content">
        <div class="patientView-content">
            <div class="btnVoltar" onclick="window.history.back()">
                <svg class="go-back" version="1.1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 129 129" xmlns:xlink="http://www.w3.org/1999/xlink" enable-background="new 0 0 129 129">
                    <g>
                        <path d="m88.6,121.3c0.8,0.8 1.8,1.2 2.9,1.2s2.1-0.4 2.9-1.2c1.6-1.6 1.6-4.2 0-5.8l-51-51 51-51c1.6-1.6 1.6-4.2 0-5.8s-4.2-1.6-5.8,0l-54,53.9c-1.6,1.6-1.6,4.2 0,5.8l54,53.9z"/>
                    </g>
                </svg>
                <span>Voltar</span>
            </div>
            <div class="patientView-header">
                <div class="patient-info-editable">
                    <q-input v-model="model.Name" class="inputEditable" placeholder="  Nome"></q-input>
                    <q-input v-model="model.Email" class="inputEditable" placeholder="  Email"></q-input>
                    <q-input v-model="model.Address" class="inputEditable" placeholder="  Endereço"></q-input>
                    <q-input v-model="model.Phone" class="inputEditable" placeholder="  Telefone"></q-input>
                    <!--RECEM NASCIDO SELECT-->
                    <q-select
                        class="select-quasar"
                        v-model="opcaoNewBorn"
                        @update:modelValue="changeNewBornValue()"
                        :options="opcoesSimNao"
                        label="&nbsp Recém nascido?"
                        placeholder="Selecione uma opção"
                    />
                    <q-input v-model="model.Cid10" type="number" class="inputEditable" placeholder=" CID10"></q-input>
                    <!--ATIVO SELECT-->
                    <q-select
                        class="select-quasar"
                        v-model="opcaoAtivo"
                        @update:modelValue="changeAtivoValue()"
                        :options="opcoesSimNao"
                        label="&nbsp Ativo?"
                        placeholder="Selecione uma opção"
                    />
                    <p class="alter-password"><a style="font-size:12px;" class="change-pass" @click="changePassword(model.IdPatient)">Alterar senha</a></p>
                </div>
                <div class="patient-info">
                    <h4>{{ model.Name }}</h4>
                    <p>{{ model.Cpf }}</p>
                    <p>{{ model.Email }}</p>
                </div>
                <div class="edit-button-div">
                    <button-primary class="editBtn" @click="editPatient(model.IdPatient)"><PhPencil class="editIcon"/>{{ editOrSave}}</button-primary>
                </div>
            </div>
        </div>
        <section>
            <button-primary class="nextPersonView">Ver pessoas próximas
                <svg class="goNext" version="1.1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 129 129" xmlns:xlink="http://www.w3.org/1999/xlink" enable-background="new 0 0 129 129">
                    <g>
                        <path d="m88.6,121.3c0.8,0.8 1.8,1.2 2.9,1.2s2.1-0.4 2.9-1.2c1.6-1.6 1.6-4.2 0-5.8l-51-51 51-51c1.6-1.6 1.6-4.2 0-5.8s-4.2-1.6-5.8,0l-54,53.9c-1.6,1.6-1.6,4.2 0,5.8l54,53.9z"/>
                    </g>
                </svg>
            </button-primary>
            <h5>Anamnese</h5>
            <q-input
                :disable="campoAnamneseDesabilitado"
                class="textarea"
                label="Escreva a anamnese"
                type="textarea"
                v-model="anamneseText"
                :rows="15"
            />
            <q-checkbox v-model="analiseConclusiva" class="analiseConclusiva">Análise conclusiva</q-checkbox>
            <div class="gerar-laudo-div">
                <button-primary class="gerar-laudo" @click="gerarLaudo()">Gerar Laudo</button-primary>
            </div>
        </section>
    </div>
</template>
<script>
import ButtonPrimary from '../components/ButtonPrimary.vue';
import { PhPencil } from '@phosphor-icons/vue';
import cookie from '../cookie';

export default {
  components: {
    ButtonPrimary,
    PhPencil,
  },
  data() {
      return{
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
            campoAnamneseDesabilitado: false,
            anamneseText: '',
        }
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
            queryParams.forEach(param => {
                const [key, value] = param.split('=');
                params[key] = value;
            });
        }

        // Obtendo o valor do parâmetro 'id'
        const id = params['id'];

        // Obtendo o valor do parâmetro 'edit'
        const edit = params['edit'];
        
        //Style shits
        var element = document.getElementsByClassName('editIcon')[0]
        var elementAlterPassword = document.getElementsByClassName('alter-password')[0]
        if(edit!==undefined && edit=='true'){ //Quando para editar
                element.style.display = 'none'
                elementAlterPassword.style.display = 'block'
                th.edit = true;
                th.editOrSave='Salvar'
                document.getElementsByClassName('patient-info')[0].style.display = 'none'
                document.getElementsByClassName('patient-info-editable')[0].style.display = 'block'
                document.getElementsByClassName('patientView-header')[0].style.display = 'block'
        }else{ //Quando para salvar
            th.edit = false;
        }

        //faz requisição para pegar info do paciente
        th.$api.PatientController.getById(id).then(({ data }) => {
            th.model = data.data
            th.opcaoNewBorn = th.model.NewBorn==1 ? 'Sim' : 'Não';
            th.opcaoAtivo = th.model.Active=='1' ? 'Sim' : 'Não';
        })

    },
    methods: {
        editPatient(id){
            console.log(id)
            //Style shits
            const th = this;
            var element = document.getElementsByClassName('editIcon')[0]
            var elementAlterPassword = document.getElementsByClassName('alter-password')[0]
            if(th.edit==false && th.editOrSave=='Editar'){ //Quando for para editar
                element.style.display = 'none'
                elementAlterPassword.style.display = 'block'
                th.edit = true;
                th.editOrSave='Salvar'
                document.getElementsByClassName('patient-info')[0].style.display = 'none'
                document.getElementsByClassName('patient-info-editable')[0].style.display = 'block'
                document.getElementsByClassName('patientView-header')[0].style.display = 'block'
            }else if(th.edit==true && th.editOrSave=='Salvar'){ //Quando for para salvar
                element = document.getElementsByClassName('editIcon')[0]
                element.style.display = 'block'
                elementAlterPassword.style.display = 'none'
                th.edit = false;
                th.editOrSave='Editar'
                document.getElementsByClassName('patient-info')[0].style.display = 'block'
                document.getElementsByClassName('patient-info-editable')[0].style.display = 'none'
                document.getElementsByClassName('patientView-header')[0].style.display = 'flex'

                this.savePatientData()

            }

        },
        savePatientData(){
            const th = this
            th.model.Cid10 = th.model.Cid10 ? parseInt(th.model.Cid10) : 0;
            th.$api.PatientController.update(
                th.model
            )
            var idUser = cookie.getUserId(cookie.get('authToken'))
            var idPatient = th.model.IdPatient
            th.$api.AnamneseController.getByIdUserPatient({'IdUser': idUser, 'IdPatient': idPatient}).then((response)=>{
                if(response.statusText==='No Content'){
                    //Not Found so insert
                }else{
                    //Found so update
                }
            })

        },
        changePassword(id){
            alert(id+' - FUTURA IMPLEMENTACAO')
        },
        changeNewBornValue(){
            const th = this
            if(th.opcaoNewBorn=='Sim'){
                th.model.NewBorn = 1
            }else{
                th.model.NewBorn = 0
            }
        },
        gerarLaudo(){
            const th = this
            alert(th.model.Active)
        },
        changeAtivoValue(){
            const th = this
            if(th.opcaoAtivo=='Sim'){
                th.model.Active = 1
            }else{
                th.model.Active = 0
            }
        }
    },
};
</script>
<style scoped>
    .select-quasar{
        background-color: #00000063;
        margin-bottom: 10px;
    }
    .inputEditable{
        background-color: #00000063;
        
        margin-bottom: 10px;
    }
    .patient-info-editable{
        display: none;
        width: -webkit-fill-available;
    }
    .patient-info{
      display: block;  
    }
    .alter-password{
        display: none;
    }
    .change-pass{
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
    .edit-button-div{
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .edit-button-div .editBtn{
        height: 55px;
    }

    .gerar-laudo-div{
        width: 100%;
        display: flex;
        justify-content: center;
    }

    .gerar-laudo{
        margin: 10px;
        margin-top: 0;
    }

    .analiseConclusiva{
        margin-top: 10px;
        margin-bottom: 10px;
    }

    .textarea{
        border-radius: 10px;
        border: 2px solid green;
        background-color: rgba(26, 26, 26, 0.671);
        min-height: 30vh;
        padding: 10px;
        font-family:Arial, Helvetica, sans-serif;
        font-size: medium;
    }

    h5{
        padding: 15px;
    }

    .nextPersonView{
        border: 1px solid var(--primary);
        background-color: rgba(26, 26, 26, 0.671);
        color: green;
        margin: 10px;
        width: -webkit-fill-available;
        margin-bottom: 20px;
    }
    .nextPersonView:hover{
        background-color: rgba(43, 43, 43, 0.671); 
    }
    .patient-content{
        width: 99%;
        height: auto;
        padding: 20px;
        padding-top: 0;
        margin-left: 0.5%;
        margin-right: 0.5%;
        font-family: "Roboto", "-apple-system", "Helvetica Neue", Helvetica, Arial, sans-serif;
        display: block;
        justify-content: space-between;
    }
    section{
        display: block;
        background-color: #1f1e1e;
        width: 99%;
        height: fit-content;
        margin-top: 2px;
        margin-left: 0.5%;
        margin-right: 0.5%;
        padding: 5px;
        justify-content: center;
    }
    .patientView-content{
        width: 100%;
        height: auto;
    }
    .go-back{
        width: auto;
        margin-right: 5px;
        height: 18px;
    }
    .goNext{
        width: auto;
        margin-right: 5px;
        height: 18px;
        fill: green;
        transform: scaleX(-1);
    }
    .patientView-header{
        width: 99%;
        height: auto;
        padding: 20px;
        background-color: #1f1e1e;
        margin-left: 0.5%;
        margin-right: 0.5%;
        font-family: "Roboto", "-apple-system", "Helvetica Neue", Helvetica, Arial, sans-serif;
        display: flex;
        justify-content: space-between;
    }

    .patientView-header section{
        display: block;
        height: 300px;
    }
    .patientView-header p{
        margin: 0;
        font-weight: 300;
    }
    .editIcon{
        margin-left: -10px;
    }
    .btnVoltar{
        margin-left: 0.5%;
        margin-top: 10px; 
        margin-bottom: 10px;
        width: fit-content;
        border-radius: 15px;
        padding: 10px;
        display: flex;
        background-color: var(--primary);
        cursor: pointer;
    }
</style>
<style>
.q-field__native, .q-field__prefix, .q-field__suffix, .q-field__input{
        color: white !important;
}
.patientView-content .q-field__native.q-placeholder{
    padding-left: 10px;
}
.q-field__label.no-pointer-events.absolute.ellipsis{
    color: rgba(255, 255, 255, 0.678);
}
.q-checkbox__bg{
    color: var(--primary) !important;
    border-color: var(--primary);
}

.patient-info-editable .q-field__control.relative-position.row.no-wrap{
    border: 1px solid #0000001c;
}
.q-menu.q-position-engine.scroll.q-menu--square{
    background-color: #272727;
}
.select-quasar span{
    margin-left: 10px;
    margin-top: 10px;
}
</style>