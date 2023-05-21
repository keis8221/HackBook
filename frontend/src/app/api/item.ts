type Items = {
  id: number
  name: string
}[]

export async function getItems(): Promise<any> {
  try {
    const response = await fetch('http://localhost:8080/items', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })
    return response.json()
  } catch (error) {
    return []
  }
}
