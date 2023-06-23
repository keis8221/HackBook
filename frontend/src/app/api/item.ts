import { baseAPI } from './base'

type Items = {
  id: number
  name: string
}[]

export async function getItems(): Promise<any> {
  return await baseAPI({
    endpoint: `items`,
  })
}
