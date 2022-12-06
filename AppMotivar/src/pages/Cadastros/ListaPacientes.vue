<template>
  <q-page padding>
     <q-table
      title="Lista de Pacientes"
      :rows="rows"
      row-key="name"
      :columns="columns"
  >
    <template v-slot:top>
      <q-toolbar>
        <q-toolbar-title>
          Lista de Pacientes
        </q-toolbar-title>
        <q-space />
        <q-btn flat label="Adicionar" color="primary" :to="{name :'CadastroPacientes'}" />
      </q-toolbar>
    </template>
    <template v-slot:body-cell-actions="props">
      <q-td :props="props" class="q-gutter-sm">
        <q-btn icon="edit" color="warning" @click="handleEditPaciente(props.row.id)" dense/>
        <q-btn  icon="delete" color="negative" @click="handleRemovePaciente(props.row.id)" dense />
      </q-td>
    </template>
  </q-table>
  </q-page>
</template>

<script>
import { defineComponent, ref, onMounted } from 'vue'
import { importaMetodosListagemPacientes } from 'src/services/posts'
import { useQuasar } from 'quasar'
import { useRouter } from 'vue-router'
export default defineComponent({
  name: 'ListaPacientes',
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
      name: 'cpf',
      field: 'cpf',
      label: 'CPF',
      align: 'center',
      sortable: true
    },
    {
      name: 'dataNascimento',
      field: 'dataNascimento',
      label: 'Data de Nascimento',
      align: 'center',
      sortable: true
    },
    {
      name: 'cep',
      field: 'cep',
      label: 'CEP',
      align: 'center',
      sortable: true
    },
    {
      name: 'cidade',
      field: 'cidade',
      label: 'Cidade',
      align: 'center',
      sortable: true
    },
    {
      name: 'bairro',
      field: 'bairro',
      label: 'Bairro',
      align: 'center',
      sortable: true
    },
    {
      name: 'rua',
      field: 'rua',
      label: 'Rua',
      align: 'center',
      sortable: true
    },
    {
      name: 'numero',
      field: 'numero',
      label: 'Número',
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
    const { list, remove } = importaMetodosListagemPacientes()
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

    const handleRemovePaciente = async (id) => {
      try {
        $q.dialog({
          title: 'Remover',
          message: 'Deseja realmente remover esse Paciente ?',
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
            message: 'Paciente removido com sucesso!',
            color: 'positive',
            position: 'bottom-right'
          })
          getRows()
        })
      } catch (error) {
        throw new Error(error)
      }
    }
    const handleEditPaciente = (id) => {
      router.push({ name: 'CadastroPacientes', params: { id } })
    }

    return {
      rows,
      columns,
      handleRemovePaciente,
      handleEditPaciente,
      dense: ref(true)
    }
  }
})
</script>
