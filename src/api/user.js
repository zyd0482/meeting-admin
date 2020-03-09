import { Post }  from '@/utils/request'
import qs from 'qs'

export function login(data) {
  return Post({
    url: '/user/login',
    data
  })
}

export function getInfo(token) {
  return Post({
    url: '/user/info'
  })
}

export function logout() {
  return Post({
    url: '/user/logout',
  })
}
