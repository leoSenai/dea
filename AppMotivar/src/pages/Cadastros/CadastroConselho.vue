<template>
  <q-page padding>
     <q-table
      title="Cadastro de Conselhos"
      :rows="rows"
      row-key="name"
      :columns="columns"
  >
    <template v-slot:body-cell-actions="props">
      <q-td :props="props">
        <q-btn  icon="delete" color="negative" @click="handleRemovePost(props.row.id), confirm=true"  dense/>
        <q-btn icon="edit" color="negative" dense/>
      </q-td>
    </template>
  </q-table>
  </q-page>
</template>

<script>
import { defineComponent, ref, onMounted } from 'vue'
import postsService from 'src/services/posts'
import { useQuasar } from 'quasar'
export default defineComponent({
  name: 'CadastroConselho',
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
    const { list, remove } = postsService()
    onMounted(async () => {
      getRows()
    })
    const $q = useQuasar()

    const getRows = async () => {
      try {
        const data = await list()
        rows.value = data
      } catch (error) {
        throw new Error(error)
      }
    }

    const handleRemovePost = async (id) => {
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

    return {
      rows,
      columns,
      handleRemovePost
    }
  }
})
</script>
