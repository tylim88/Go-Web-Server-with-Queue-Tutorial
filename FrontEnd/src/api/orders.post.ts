import { useMutation } from '@tanstack/react-query'
import { client_api } from './__utils'

export const useOrdersPost = () =>
	useMutation({
		mutationFn: (type: 'vip' | 'regular') =>
			client_api.post('orders', {
				json: { type },
			}),
	})
