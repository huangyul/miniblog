import { apiPost } from "@/utils/request"

interface UserItem {
  name: string
  age: number
}

export const getUserList = () => {
  return apiPost<UserItem[]>({
    url: '/userlist'
  })
}
