import { getMessageFromJson } from "../helper"

const message = import.meta.glob('./*.json', { eager: true })

export default getMessageFromJson(message)
