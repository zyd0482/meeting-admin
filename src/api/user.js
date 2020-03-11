import { Post }  from '@/utils/request'

export function login(data) {
  return Post({
    url: '/user/login',
    data
  })
}

export function logout() {
  return Post({
    url: '/user/logout',
  })
}
