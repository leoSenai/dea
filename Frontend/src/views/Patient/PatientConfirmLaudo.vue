<template>
  <modal-primary
    v-model="show"
    @close="closeModal"
  >
    <template #modal-title>
      Laudo final de {{ model.Name }}
    </template>
    <template #modal-content>
      <q-form ref="form">
        <div class="row q-mb-sm">
          <q-select
            v-model="modeloAnamnese.Indicative"
            :options="indicativeOptions"
            label="Indicativo"
            option-label="label"
            option-value="value"
            label-color="primary"
            class="q-field select row select-input q-field__control"
            required
            :style="
              modeloAnamnese.Indicative.value == 0
                ? 'width: 100%;'
                : 'width: 50%;'
            "
          />
          <q-select
            v-show="modeloAnamnese.Indicative.value == 1"
            v-model="grau"
            :options="graus"
            label="Grau"
            option-label="label"
            option-value="value"
            label-color="primary"
            class="q-field select row select-input q-field__control"
            required
            style="width: 50%"
          />
        </div>
      </q-form>
    </template>
    <template #modal-actions>
      <div
        class="modal-actions"
        style="gap: 2em"
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
          style="background-color: var(--primary); color: white"
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
      modeloAnamnese: {
        IdAnamnese: null,
        IdPatient: null,
        IdUser: null,
        Notes: '',
        Indicative: {
          label: 'Não possui Transtorno do Espectro Autista',
          value: 0,
        },
      },
      grau: { label: 'Selecione', value: 0 },
      indicativeOptions: [
        { label: 'Não possui Transtorno do Espectro Autista', value: 0 },
        { label: 'Possui Transtorno do Espectro Autista', value: 1 },
      ],
      graus: [
        { label: 'Leve', value: 1 },
        { label: 'Moderado', value: 2 },
        { label: 'Severo', value: 3 },
      ],
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
      RoleEnum,
    };
  },
  computed: {
    ...RoleEnum,
    typeUser() {
      const token = cookie.get('authToken');
      return cookie.getUserType(token);
    },
    idUser() {
      const token = cookie.get('authToken');
      return cookie.getUserId(token);
    },
  },
  methods: {
    openModal(current, modeloAnamnese) {
      const th = this;
      th.show = true;
      if (current) {
        th.model = current;
        th.modeloAnamnese = { ...modeloAnamnese };
        th.modeloAnamnese.Indicative =
          modeloAnamnese.Indicative == 0
            ? { label: 'Não possui Transtorno do Espectro Autista', value: 0 }
            : { label: 'Possui Transtorno do Espectro Autista', value: 1 };
      }
    },
    async generateLaudo() {
      const th = this;
      if (
        (th.grau.value != 0 && th.modeloAnamnese.Indicative.value == 1) ||
        th.modeloAnamnese.Indicative.value == 0
      ) {
        if (th.modeloAnamnese.Indicative.value == 0) {
          th.grau.value = 0;
        }

        await th.$api.AnamneseController.update({
          IdAnamnese: th.modeloAnamnese.IdAnamnese,
          IdPatient: th.modeloAnamnese.IdPatient,
          IdUser: th.modeloAnamnese.IdUser,
          Notes: th.modeloAnamnese.Notes,
          Indicative: th.modeloAnamnese.Indicative.value == 0 ? 0 : 1,
        });

        var response = await th.$api.AnamneseController.getLaudo(
          th.modeloAnamnese.IdUser,
          th.model.IdPatient,
          th.grau.value
        );
        const bytes = this.base64ToArrayBuffer(response.data.data);
        const pdfBlob = new Blob([bytes], { type: 'application/pdf' });
        let link = document.createElement('a');
        link.href = window.URL.createObjectURL(pdfBlob);
        link.style.visibility = 'hidden';
        link.download = pdfBlob.name;
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      } else {
        alert('O laudo só pode ser gerado quando a anamnese for conclusiva!');
      }
    },
    base64ToArrayBuffer(base64) {
      const binaryString = window.atob(base64);
      const binaryLen = binaryString.length;
      const bytes = new Uint8Array(binaryLen);
      for (let i = 0; i < binaryLen; i++) {
        const ascii = binaryString.charCodeAt(i);
        bytes[i] = ascii;
      }

      return bytes;
    },
    closeModal() {
      this.show = false;
      this.modeloAnamnese;
      this.model = [];
      this.$emit('close');
    },
  },
};
</script>
<style>
.q-slide.answer .q-slider__track {
  background: linear-gradient(to right, red, yellow 50%, #04df04) !important;
}
.q-slide.answer .q-slider__selection {
  background: transparent;
}
.q-slide.answer .text-primary {
  color: blue !important;
}
</style>
<style scoped>
.fill-content {
  width: 100%;
}

p {
  margin: 0;
}

.modal-actions {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
</style>
