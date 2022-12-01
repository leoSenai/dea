<template>
  <q-page padding>
     <q-table
      title="Lista de Conselhos"
      :rows="rows"
      row-key="name"
      :columns="columns"
  >
    <template v-slot:top>
      <q-toolbar>
        <q-toolbar-title>
          Lista de Conselhos
        </q-toolbar-title>
        <q-space />
        <q-btn flat label="Adicionar" color="primary" :to="{name :'CadastroConselho'}" />
      </q-toolbar>
    </template>
    <template v-slot:body-cell-actions="props">
      <q-td :props="props" class="q-gutter-sm">
        <q-btn icon="edit" color="warning" @click="handleEditConselho(props.row.id)" dense/>
        <q-btn  icon="delete" color="negative" @click="handleRemoveConselho(props.row.id)" dense />
      </q-td>
    </template>
  </q-table>
  </q-page>
</template>

<script>
import { defineComponent, ref, onMounted } from 'vue'
import importaMetodosListagemConselho from 'src/services/posts'
import { useQuasar } from 'quasar'
import { useRouter } from 'vue-router'
export default defineComponent({
  name: 'ListaConselho',
  setup () {
    const rows = ref([])
    const columns = [{
      name: 'id',
      field: 'id',
      label: 'id',
      align: 'center',
      sortable: true
    },
    {
      name: 'nome',
      field: 'nome',
      label: 'Conselho',
      align: 'center',
      sortable: true
    },
    {
      name: 'ultimaAlteracao',
      field: 'ultimaAlteracao',
      label: 'Ultima Alteração',
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
    const { list, remove } = importaMetodosListagemConselho()
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

    const handleRemoveConselho = async (id) => {
      try {
        $q.dialog({
          title: 'Remover',
          message: 'Deseja realmente remover esse Conselho ?',
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
            message: 'Conselho removido com sucesso!',
            color: 'positive',
            position: 'bottom-right'
          })
          getRows()
        })
      } catch (error) {
        throw new Error(error)
      }
    }
    const handleEditConselho = (id) => {
      router.push({ name: 'CadastroConselho', params: { id } })
    }

    return {
      rows,
      columns,
      handleRemoveConselho,
      handleEditConselho,
      dense: ref(true)
    }
  }
})
</script>
