<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Cadastro de Pacientes</div>
      </q-card-section>
      <q-card-section>
        <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
          <q-input
            outlined
            v-model="nome"
            label="Nome"
            lazy-rules
            :dense="dense"
            :rules="[(val) => val.length > 0 || 'Nome é obrigatório']"
          />
          <q-input
            outlined
            v-model="cpf"
            label="CPF"
            lazy-rules
            :dense="dense"
            :rules="[(val) => val.length > 0 || 'CPF é obrigatório']"
          />
          <q-input
            outlined
            v-model="rg"
            label="RG"
            lazy-rules
            :dense="dense"
            :rules="[(val) => val.length > 0 || 'RG é obrigatório']"
          />
          <q-input
            outlined
            v-model="email"
            label="Email"
            lazy-rules
            :dense="dense"
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
    return {
      form: {
        nome: '',
        cpf: '',
        rg: '',
        email: '',
        onSubmit () {
          console.log('submit')
        },
        onReset () {}
      },
      dense: ref(true)
    }
  }
})
</script>
