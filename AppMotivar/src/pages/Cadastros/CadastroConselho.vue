<template>
  <q-page padding>
     <q-table
      title="Cadastro de Conselhos"
      :rows="rows"
      row-key="name"
      :columns="columns"
    />
  </q-page>
</template>

<script>
import { defineComponent, ref, onMounted } from 'vue'
import { api } from 'boot/axios'
export default defineComponent({
  name: 'CadastroConselho',
  setup () {
    const rows = ref([]) // array of objects
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
    }
    ]
    onMounted(async () => {
      getRows()
    })
    const getRows = async () => {
      try {
        const { data } = await api.get('conselho')
        rows.value = data
      } catch (error) {
        console.log(error)
      }
    }
    return {
      rows,
      columns
    }
  }
})
</script>
