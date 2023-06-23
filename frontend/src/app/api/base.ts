type TResponse = {
  statusCode: number
  message: string
}

export async function baseAPI<
  T extends Record<string, unknown> | Array<Record<string, unknown>> | void,
  B = Record<string, any>,
>({
  endpoint,
  method = 'GET',
  body,
}: {
  endpoint: string
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE'
  body?: B
}): Promise<T> {
  const controller = new AbortController()
  const timeoutId = setTimeout(() => {
    controller.abort()
  }, 10 * 1000) // デフォルトでは10秒でタイムアウト

  const res = await fetch(`/api/${endpoint}`, {
    method,
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${localStorage.getItem('accessToken')}`,
    },
    signal: controller.signal,
    ...(body && { body: JSON.stringify(body) }),
  }).finally(() => clearTimeout(timeoutId))

  if (!res.ok) {
    let err: TResponse
    try {
      err = await res.json()
    } catch {
      err = {
        statusCode: res.status,
        message: res.statusText,
      }
    }
    throw err
  }

  try {
    return await res.json()
  } catch {
    return {} as T
  }
}
