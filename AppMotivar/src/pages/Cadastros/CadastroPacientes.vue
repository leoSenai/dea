<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Cadastro de Pacientes</div>
      </q-card-section>
      <q-card-section class="row">
        <q-form @submit="onSubmit" class="q-gutter-md col-sm-12 col-md-10 col-xs-12" >
          <!-- campo nome com limite de 10 caracteres -->
          <q-input
            outlined
            v-model="form.nome"
            label="Nome"
            ondemandString
            :rules="[(val) => val.length > 0 || 'Nome é obrigatório']"
            class="col-12"
          />

          <q-input
            outlined
            v-model="form.cep"
            label="Cep"
            ondemandString
            :rules="[(val) => val.length > 0 || 'Cep é obrigatório']" />

          <q-input
            outlined
            v-model="form.cidade"
            label="Cidade"
            ondemandString
            :rules="[(val) => val.length > 0 || 'Cidade é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.bairro"
            label="Bairro"
            ondemandString

            :rules="[(val) => val.length > 0 || 'Bairro é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.rua"
            label="Rua"
            ondemandString

            :rules="[(val) => val.length > 0 || 'Rua é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.numero"
            label="Número"
            ondemandString

            :rules="[(val) => val.length > 0 || 'Número é obrigatório']"
          />
          <q-card-actions align="right">
            <q-btn label="Cancelar" color="negative" :to="{name: 'ListaPacientes'}" />
            <q-btn label="Salvar"  type="submit" color="primary" />
          </q-card-actions>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import { defineComponent, ref, watch } from 'vue'
import { BuscaEnderecoPorCep, importaMetodosCadastroPacientes } from 'src/services/posts'
import { useRouter } from 'vue-router'
import { useQuasar } from 'quasar'

export default defineComponent({
  name: 'CadastroPacientes',
  setup () {
    const form = ref({
      cep: '',
      nome: '',
      bairro: '',
      cidade: '',
      rua: '',
      numero: ''
    })

    const { getCep } = BuscaEnderecoPorCep()
    const { post, update } = importaMetodosCadastroPacientes
    const router = useRouter()
    const $q = useQuasar()

    watch(() => form.value.cep, async (val) => {
      if (val.length === 8) {
        try {
          const response = await getCep(val)
          if (response.erro) {
            $q.notify({
              message: 'Não foi possível encontrar o CEP',
              color: 'negative',
              position: 'top'
            })
            return
          }
          form.value.cidade = response.localidade ? response.localidade : ''
          form.value.bairro = response.bairro ? response.bairro : ''
          form.value.rua = response.logradouro ? response.logradouro : ''
        } catch (error) {
          throw new Error(error)
        }
      }
    })

    const onSubmit = async () => {
      try {
        if (form.value.id) {
          await update(form.value)
        } else {
          await post(form.value)
        }
        $q.notify({
          message: 'Paciente cadastrado com sucesso!',
          color: 'positive',
          position: 'bottom-right'
        })
        router.push({ name: 'ListaConselho' })
      } catch (error) {
        throw new Error(error)
      }
    }

    return {
      form,
      onSubmit
    }
  }
})
</script>
