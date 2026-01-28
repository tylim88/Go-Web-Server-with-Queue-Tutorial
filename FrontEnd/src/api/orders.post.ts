import { useMutation } from '@tanstack/react-query'
import { client } from './__utils'

export const useOrdersPost = () =>
	useMutation({
		mutationFn: (type: 'vip' | 'regular') =>
			client.post('orders', {
				json: { type },
			}),
	})
