import { useQuery } from '@tanstack/react-query'
import { client_api, client_query, prefixUrl } from './__utils'
import { useEffect } from 'react'
import { fetchEventSource } from '@microsoft/fetch-event-source'

const key = 'get all list'

export const useOrdersGet = () =>
	useQuery({
		queryKey: [key],
		queryFn: () => client_api.get('orders').json<InitialData>(),
		refetchOnWindowFocus: false,
		retry: false,
	})

export const useOrderSSE = () =>
	useEffect(() => {
		fetchEventSource(`${prefixUrl}/ordersSSE`, {
			onmessage(msg) {
				console.log({ msg })
				if (msg.event !== 'initial_data') {
					const data = JSON.parse(msg.data) as
						| (Pending & {
								type: 'vip' | 'regular'
								queue: 'pending'
								action: 'add' | 'remove'
						  })
						| (Processing & {
								id_robot: number
								queue: 'processing'
								action: 'add' | 'remove'
						  })
						| (Completed & {
								queue: 'completed'
						  })
					client_query.setQueryData([key], (old: InitialData) => {
						switch (data.queue) {
							case 'pending':
								switch (data.action) {
									case 'add':
										old.pending[data.type].push(data)
										break
									case 'remove': {
										const index = old.pending[data.type].findIndex(
											({ id_order }) => data.id_order === id_order,
										)
										old.pending[data.type].splice(index, 1)
										break
									}
									default:
										data.action satisfies never
								}
								break
							case 'processing':
								switch (data.action) {
									case 'add':
										old.processing[data.id_robot] = data
										break
									case 'remove':
										delete old.processing[data.id_robot]
										break
									default:
										data.action satisfies never
								}
								break
							case 'completed':
								old.completed.unshift(data)
								break
							default:
								data satisfies never
						}

						return { ...old }
					})
				}
			},
		})
	}, [])

type Pending = { id_order: number; time_create: string }

type Processing = {
	id_order: number
	time_create: string
	time_process: string
	time_remaining: string
	type: 'vip' | 'regular'
}

type Completed = {
	id_order: number
	id_robot: number
	time_create: string
	time_process: string
	time_complete: string
	type: 'vip' | 'regular'
}

type InitialData = {
	pending: Record<'vip' | 'regular', Pending[]>
	processing: Record<number, Processing>
	completed: Completed[]
}
