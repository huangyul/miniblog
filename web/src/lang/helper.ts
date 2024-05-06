export const getMessageFromJson = (message: Record<string, unknown>): Record<string, string> => {
  let res = {}

  Object.keys(message).forEach(key => {
    if (message[key]) {
      res = { ...res, ...(message[key] as any).default }
    }
  })

  return res
}
