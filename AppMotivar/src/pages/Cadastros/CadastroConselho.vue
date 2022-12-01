<template>
  <q-page padding>
    <q-card >
      <q-card-section>
        <div class="text-h6">Cadastro de Conselho</div>
      </q-card-section>
      <q-card-section  class="col-12">
        <q-form @submit="onSubmit" class="q-gutter-md" >
          <q-input
            outlined
            v-model="form.nome"
            label="Nome"
            lazy-rules
            :dense="dense"
            :rules="[(val) => val.length > 0 || 'Nome é obrigatório']"
          />
          <q-input
            outlined
            v-model="form.ultimaAlteracao"
            label="Última Alterção"
            lazy-rules
            :dense="dense"
            :rules="[(val) => val.length > 0 || 'Última Alterção']"
          />
          <q-card-actions align="right">
            <q-btn label="Cancelar" color="negative" :to="{name: 'ListaConselho'}" />
            <q-btn label="Salvar"  type="submit" color="primary" />
          </q-card-actions>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import { defineComponent, ref, onMounted } from 'vue'
import importaMetodosListagemConselho from 'src/services/posts'
import { useQuasar } from 'quasar'
import { useRouter, useRoute } from 'vue-router'
export default defineComponent({
  name: 'CadastroConselho',
  setup () {
    const form = ref({
      nome: '',
      ultimaAlteracao: ''
    })
    const $q = useQuasar()
    const { post, getById, update } = importaMetodosListagemConselho()
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
      onSubmit,
      dense: ref(true)
    }
  }
})
</script>
