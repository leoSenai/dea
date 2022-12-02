<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Cadastro de Pacientes</div>
      </q-card-section>
      <q-card-section>
        <q-form @submit="onSubmit" class="q-gutter-md">
          <q-input
            outlined
            v-model="form.nome"
            label="Nome"
            lazy-rules
            dense
            :rules="[(val) => val.length > 0 || 'Nome é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.cpf"
            label="CPF"
            lazy-rules
            dense
            :rules="[(val) => val.length > 0 || 'CPF é obrigatório']"
          />

          <q-input outlined v-model="form.dataNascimento" mask="##/##/####" :rules="['##/##/####']">
            <template v-slot:append>
              <q-icon name="event" class="cursor-pointer">
                <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                  <q-date v-model="form.dataNascimento" mask="DD-MM-YYYY" dense>
                    <div class="row items-center justify-end">
                      <q-btn v-close-popup label="Fechar" color="primary" flat />
                    </div>
                  </q-date>
                </q-popup-proxy>
              </q-icon>
            </template>
          </q-input>

          <q-input
            outlined
            v-model="form.email"
            label="Email"
            lazy-rules
            dense
            :rules="[(val) => val.length > 0 || 'Email é obrigatório']"
          />
          <div class="">
            <q-btn label="Salvar" type="submit" color="primary" />
            <q-btn
              label="Limpar"
              type="reset"
              color="primary"
              flat
              class="q-ml-sm"
            />
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'IndexPage',
  setup () {
    const form = ref({
      nome: '',
      cpf: '',
      dataNascimento: '',
      email: ''
    })

    return {
      form
    }
  }
})
</script>
