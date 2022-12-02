<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Cadastro de Pacientes</div>
      </q-card-section>
      <q-card-section>
        <q-form @submit="OnSubmit" class="q-gutter-md">
          <!-- campo nome com limite de 10 caracteres -->
          <q-input
            outlined
            v-model="form.nome"
            label="Nome"
            ondemandString
            dense
            :rules="[(val) => val.length > 0 || 'Nome é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.cep"
            label="cep"
            ondemandString
            dense
            :rules="[(val) => val.length > 0 || 'cep é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.bairro"
            label="Bairro"
            ondemandString
            dense
            :rules="[(val) => val.length > 0 || 'cep é obrigatório']"
          />
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
      bairro: ''
    })
    const { getCep } = BuscaEnderecoPorCep()
    const { post, update } = importaMetodosCadastroPacientes
    const router = useRouter()
    const { $q } = useQuasar()

    // usar watch para observar mudanças no campo de cep e buscar na api pelo getCep
    watch(form.value, async (val) => {
      if (val.cep.length === 8) {
        try {
          const response = await getCep(val.cep)
          form.value.bairro = response.bairro
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
          message: 'Conselho cadastrado com sucesso!',
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
      dense: ref(true),
      onSubmit
    }
  }
})
</script>
