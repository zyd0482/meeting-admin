import { Post }  from '@/utils/request'

export function fetchList(data) {
  return Post({
    url: '/meet/list',
    data
  })
}