<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Cadastro de Pessoas Próximas</div>
      </q-card-section>
      <q-card-section>
        <q-form @submit="onSubmit" class="q-gutter-md">
          <q-input
            outlined
            v-model="form.nome"
            label="Nome"
            lazy-rules
            :rules="[(val) => val.length > 0 || 'Nome é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.cpf"
            label="CPF"
            lazy-rules
            mask="###.###.###-##"
            :rules="[(val) => val.length > 13 || 'CPF é obrigatório']"
          />

          <q-input
            outlined
            v-model="form.email"
            label="Email"
            lazy-rules
            :rules="[(val) => val.length > 0 || 'Email é obrigatório']"
          />
          <q-card-actions align="right">
            <q-btn label="Cancelar" color="negative" :to="{name: 'ListaPessoasProximas'}" />
            <q-btn label="Salvar"  type="submit" color="primary" />
          </q-card-actions>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useQuasar } from 'quasar'
import { importaMetodosListagemPessoasProximas } from 'src/services/posts'
export default defineComponent({
  name: 'CadastroPessoasProximas',
  setup () {
    const router = useRouter()
    const $q = useQuasar()
    const { update, post } = importaMetodosListagemPessoasProximas()
    const form = ref({
      nome: '',
      cpf: '',
      email: ''
    })
    const onSubmit = async () => {
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
      router.push({ name: 'ListaPessoasProximas' })
    }

    return {
      form,
      onSubmit
    }
  }
})
</script>
