import { ScrollArea, Paper, Text, Stack, Flex } from '@mantine/core'
import { useOrdersGet } from '@/api'
import { Loader } from '@mantine/core'

export const ListPending = () => {
	const { isSuccess, data, isLoading } = useOrdersGet()

	return (
		<Paper display="flex">
			{(['regular', 'vip'] as const).map(type => {
				return (
					<Stack key={type} align="center">
						<Text ta="center" tt="uppercase" td="underline" fz="lg" fw={700}>
							{type}
						</Text>
						<ScrollArea h={400} w={160}>
							{isLoading ? (
								<Flex justify="center">
									<Loader />
								</Flex>
							) : (
								isSuccess &&
								data.pending[type].map(({ id_order }) => {
									return (
										<Text key={id_order} ta="center">
											{id_order.toString().padStart(4, '0')}
										</Text>
									)
								})
							)}
						</ScrollArea>
					</Stack>
				)
			})}
		</Paper>
	)
}
