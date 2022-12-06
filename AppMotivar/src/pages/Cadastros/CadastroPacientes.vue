<template>
  <q-page padding>
    <q-card>
      <q-card-section class="row items-center justify-center">
        <div class="text-h6">Cadastro de Pacientes</div>
      </q-card-section>
      <q-card-section class="row items-center justify-center">

        <q-form @submit="onSubmit" class="q-gutter-md col-sm-12 col-md-10 col-xs-12" >
          <!-- campo nome com limite de 10 caracteres -->
          <q-input
            outlined
            v-model="form.nome"
            label="Nome"
            ondemandString
            :rules="[(val) => val.length > 0 || 'Nome é obrigatório']"
          />

          <q-input outlined v-model="form.dataNascimento" label="Data de Nascimento" mask="##/##/####" :rules="['##/##/####']">
            <template v-slot:append>
              <q-icon name="event" class="cursor-pointer">
                <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                  <q-date v-model="form.dataNascimento" mask="DD/MM/YYYY">
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
            v-model="form.cpf"
            label="CPF"
            mask="###.###.###-##"
            ondemandString
            :rules="[(val) => val.length > 0 || 'CPF é obrigatório']" />

          <q-input
            outlined
            v-model="form.cep"
            label="CEP"
            ondemandString
            @blur="buscaCep"
            mask="#####-###"
            :rules="[(val) => val.length > 0 && val.length <= 9 || 'CEP é obrigatório']" />

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
import { defineComponent, ref, onMounted } from 'vue'
import { BuscaEnderecoPorCep, importaMetodosCadastroPacientes } from 'src/services/posts'
import { useRouter, useRoute } from 'vue-router'
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
      numero: '',
      dataNascimento: '',
      cpf: ''
    })

    const { getCep } = BuscaEnderecoPorCep()
    const { post, update, getById } = importaMetodosCadastroPacientes()
    const router = useRouter()
    const route = useRoute()
    const $q = useQuasar()
    onMounted(async () => {
      if (route.params.id) getPost(route.params.id)
    })

    const getPost = async (id) => {
      try {
        form.value = await getById(id)
      } catch (error) {
        throw new Error(error)
      }
    }
    const buscaCep = async () => {
      const response = await getCep(form.value.cep)
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
    }

    const onSubmit = async () => {
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
      router.push({ name: 'ListaPacientes' })
    }

    return {
      form,
      onSubmit,
      buscaCep
    }
  }
})
</script>
