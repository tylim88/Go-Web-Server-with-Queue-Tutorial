import ky from 'ky'
import { QueryClient } from '@tanstack/react-query'

export const prefixUrl = 'http://127.0.0.1:8080'
export const client_api = ky.create({ prefixUrl })
export const client_query = new QueryClient()
