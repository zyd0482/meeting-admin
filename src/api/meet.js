import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/api/admin/v1/meets',
    method: 'get',
    params: query
  })
}

export function fetchMeet(id) {
  return request({
    url: '/api/admin/v1/meets/' + id,
    method: 'get'
  })
}

export function createMeet(data) {
  return request({
    url: '/api/admin/v1/meets',
    method: 'post',
    data
  })
}

export function updateMeet(id, data) {
  return request({
    url: '/api/admin/v1/meets/' + id,
    method: 'put',
    data
  })
}

export function deleteMeet(id) {
  return request({
    url: '/api/admin/v1/meets/' + id,
    method: 'delete'
  })
}
