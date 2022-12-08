<template>
    <q-page padding>
        <q-card>
            <q-card-section>
                <div class="text-h6">Cadastro de Profissionais</div>
            </q-card-section>
            <q-card-section class="col-12">
                <q-form @submit="onSubmit" class="q-gutter-md">
                    <q-input outlined v-model="form.nome" label="Nome" ondemandString
                        :rules="[(val) => val.length > 0 || 'Nome é obrigatório']" />

                    <q-input outlined v-model="form.cpf" label="CPF" mask="###.###.###-##" ondemandString
                        :rules="[(val) => val.length > 0 || 'CPF é obrigatório']" />

                    <q-input outlined v-model="form.email" label="E-mail" ondemandString
                        :rules="[(val) => val.length > 0 || 'E-mail é obrigatório']" />

                    <q-input outlined v-model="form.senha" type="password" label="Senha" ondemandString
                        :rules="[(val) => val.length > 0 || 'Senha é obrigatório']" />

                    <q-card-actions align="right">
                        <q-btn label="Cancelar" color="negative" :to="{ name: 'ListaProfissionais' }" />
                        <q-btn label="Salvar" type="submit" color="primary" />
                    </q-card-actions>
                </q-form>
            </q-card-section>
        </q-card>
    </q-page>
</template>

<script>
import { defineComponent, ref, onMounted } from 'vue'
import { importaMetodosListagemProfissionais } from 'src/services/posts'
import { useQuasar } from 'quasar'
import { useRouter, useRoute } from 'vue-router'
export default defineComponent({
  name: 'CadastroProfissionais',
  setup () {
    const form = ref({
      nome: '',
      cpf: '',
      email: '',
      senha: ''
    })
    const $q = useQuasar()
    const { post, update, getById } = importaMetodosListagemProfissionais()
    const router = useRouter()
    const route = useRoute()
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

    const onSubmit = async () => {
      if (form.value.id) {
        await update(form.value)
      } else {
        await post(form.value)
      }
      $q.notify({
        message: 'Profissional cadastrado com sucesso!',
        color: 'positive',
        position: 'bottom-right'
      })
      router.push({ name: 'ListaProfissionais' })
    }
    return {
      form,
      onSubmit
    }
  }
})
</script>
