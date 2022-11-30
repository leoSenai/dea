import useApi from 'src/composables/UseApi'

export default function postsService () {
  const { list, post, put, remove } = useApi('conselho')
  return {
    list,
    post,
    put,
    remove
  }
  // const { list, post, put, remove } = useApi('/posts')
  // return {
  //   list,
  //   post,
  //   put,
  //   remove
  // }
}
