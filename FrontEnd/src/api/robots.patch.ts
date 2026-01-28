import { useMutation } from '@tanstack/react-query'
import { client } from './__utils'

export const useRobotPatch = () =>
	useMutation({
		mutationFn: (count_robots: number) =>
			client.patch('robots', {
				json: { count_robots },
			}),
	})
