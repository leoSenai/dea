<template>
  <q-page padding>
     <q-table
      title="Listagem de Pessoas Próximas"
      :rows="rows"
      row-key="name"
      :columns="columns"
  >
    <template v-slot:top>
      <q-toolbar>
        <q-toolbar-title>
          Listagem de Pessoas Próximas
        </q-toolbar-title>
        <q-space />
        <q-btn flat label="Adicionar" color="primary" :to="{name :'CadastroQuestionario'}" />
      </q-toolbar>
    </template>
    <template v-slot:body-cell-actions="props">
      <q-td :props="props" class="q-gutter-sm">
        <q-btn icon="edit" color="warning" @click="handleEditPessoasProximas(props.row.id)" dense />
        <q-btn  icon="delete" color="negative" @click="handleRemovePessoasProximas(props.row.id)" dense/>
      </q-td>
    </template>
  </q-table>
  </q-page>
</template>

<script>
import { defineComponent, ref, onMounted } from 'vue'
import { importaMetodosListagemQuestionarios } from 'src/services/posts'
import { useQuasar } from 'quasar'
import { useRouter } from 'vue-router'
export default defineComponent({
  name: 'ListaQuestionario',
  setup () {
    const rows = ref([])
    const columns = [{
      name: 'nome',
      field: 'nome',
      label: 'Nome',
      align: 'center',
      sortable: true
    },
    {
      name: 'range',
      field: 'range',
      label: 'Range',
      align: 'center',
      sortable: true
    },
    {
      name: 'actions',
      field: 'actions',
      label: 'Ações',
      align: 'right',
      sortable: true
    }
    ]
    const { list, remove } = importaMetodosListagemQuestionarios()
    onMounted(async () => {
      getRows()
    })
    const $q = useQuasar()
    const router = useRouter()

    const getRows = async () => {
      try {
        const data = await list()
        rows.value = data
      } catch (error) {
        throw new Error(error)
      }
    }

    const handleRemovePessoasProximas = async (id) => {
      try {
        $q.dialog({
          title: 'Remover',
          message: 'Deseja realmente remover essa Pessoa ?',
          cancel: {
            label: 'Cancelar',
            color: 'primary',
            flat: true
          },
          ok: {
            label: 'Remover',
            color: 'negative',
            flat: true
          },
          persistent: true
        }).onOk(async () => {
          await remove(id)
          $q.notify({
            message: 'Pessoa removido com sucesso!',
            color: 'positive',
            position: 'bottom-right'
          })
          getRows()
        })
      } catch (error) {
        throw new Error(error)
      }
    }
    const handleEditPessoasProximas = (id) => {
      router.push({ name: 'CadastroQuestionario', params: { id } })
    }

    return {
      rows,
      columns,
      handleRemovePessoasProximas,
      handleEditPessoasProximas,
      dense: ref(true)
    }
  }
})
</script>
