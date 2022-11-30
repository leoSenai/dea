import useApi from 'src/composables/UseApi'

export default function postsService () {
  const { list, post, update, remove, getById } = useApi('conselho')
  return { list, post, update, remove, getById }
}

export function teste () {
  const { list, post, update, remove, getById } = useApi('conselho')
  return { list, post, update, remove, getById }
}
