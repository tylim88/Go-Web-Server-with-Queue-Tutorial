import { useQuery } from '@tanstack/react-query'
import { client } from './__utils'

export const useOrdersGet = () =>
	useQuery({
		queryKey: ['get all list'],
		queryFn: () =>
			client.get('orders').json<{
				pending: Record<
					'vip' | 'regular',
					{ id_order: number; time_create: string }[]
				>
				processing: Record<
					number,
					{
						id_order: number
						time_create: string
						time_process: string
						time_remaining: string
						type: 'vip' | 'regular'
					}
				>
				completed: {
					id_order: number
					id_robot: number
					time_create: string
					time_process: string
					time_complete: string
					type: 'vip' | 'regular'
				}[]
			}>(),
	})
