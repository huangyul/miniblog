import { apiGet, apiPost } from "@/utils/request"

interface UserItem {
  name: string
  age: number
}

export const getUserList = () => {
  return apiPost<UserItem[]>({
    url: '/userlist'
  })
}

export const get2 = () => {
  return apiGet<{status: string}>({
    url: '/healthz'
  })
}


export const test = async () => {
  const res = await getUserList()
  console.log(res)
}
