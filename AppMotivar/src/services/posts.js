import useApi from 'src/composables/UseApi'

export default function importaMetodosListagemConselho () {
  const { list, post, update, remove, getById, getCep } = useApi('conselho')
  return { list, post, update, remove, getById, getCep }
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
export function importaMetodosCadastroPacientes () {
  const { list, post, update, remove, getById } = useApi('paciente')
  return { list, post, update, remove, getById }
}

export function importaMetodosListagemPacientes () {
  const { list, remove } = useApi('paciente')
  return { list, remove }
}

export function BuscaEnderecoPorCep () {
  const { getCep } = useApi('')
  return { getCep }
}

export function Login () {
  const { getCep } = useApi('login')
  return { getCep }
}
