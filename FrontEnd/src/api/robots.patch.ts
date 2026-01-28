import { useMutation } from '@tanstack/react-query'
import { client_api } from './__utils'

export const useRobotPatch = () =>
	useMutation({
		mutationFn: (count_robots: number) =>
			client_api.patch('robots', {
				json: { count_robots },
			}),
	})
