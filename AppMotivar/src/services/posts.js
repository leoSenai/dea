import useApi from 'src/composables/UseApi'

export default function importaMetodosListagemConselho () {
  const { list, post, update, remove, getById } = useApi('conselho')
  return { list, post, update, remove, getById }
}

export function importaMetodosCadastroConselho () {
  const { list, post, update, remove, getById } = useApi('conselho')
  return { list, post, update, remove, getById }
}

export function importaMetodosListagemPessoasProximas () {
  const { list, post, update, remove, getById } = useApi('conselho')
  return { list, post, update, remove, getById }
}

export function importaMetodosCadastroPessoasProximas () {
  const { list, post, update, remove, getById } = useApi('conselho')
  return { list, post, update, remove, getById }
}
