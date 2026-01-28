import { ScrollArea, Paper, Text } from '@mantine/core'
import { useOrdersGet } from '@/api'
import { Loader } from '@mantine/core'

export const ListCompleted = () => {
	const { isSuccess, data, isLoading } = useOrdersGet()

	return (
		<Paper display="flex" style={{ flexDirection: 'column' }}>
			<Text ta="center" tt="uppercase" td="underline" fz="lg" fw={700} w="100%">
				completed
			</Text>
			<ScrollArea h={400} w={160}>
				{isLoading ? (
					<Loader />
				) : (
					isSuccess &&
					data.completed.map(({ id_order, type }) => {
						return (
							<Text key={id_order} ta="center">
								{id_order.toString().padStart(4, '0')}
								{type === 'vip' ? ' (vip)' : ''}
							</Text>
						)
					})
				)}
			</ScrollArea>
		</Paper>
	)
}
