<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Cadastro de Questionários</div>
      </q-card-section>
      <q-card-section>
        <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
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
            v-model="form.range"
            label="Range"
            lazy-rules
            :dense="dense"
            :rules="[(val) => val.length > 0 || 'Range é obrigatório']"
          />

          <!-- <div class="q-gutter-y-md column">
            <q-rating
              v-model="form.ratingModel"
              max="7"
              size="3em"
              color="green-5"
              color-selected="green-10"
              icon="mdi-checkbox-blank-circle-outline"
              icon-selected="mdi-record-circle"
              no-dimming
            />
          </div> -->

          <q-card-actions align="right">
            <q-btn label="Cancelar" color="negative" :to="{name: 'ListaQuestionario'}" />
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
import { importaMetodosCadastroQuestionarios } from 'src/services/posts'
export default defineComponent({
  name: 'CadastroQuestionario',
  setup () {
    const $q = useQuasar()
    const router = useRouter()
    const { update, post } = importaMetodosCadastroQuestionarios()
    const form = ref({
      nome: '',
      range: '',
      ratingModel: 0
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
      router.push({ name: 'ListaQuestionario' })
    }
    return {
      form,
      onSubmit
    }
  }
})
</script>
