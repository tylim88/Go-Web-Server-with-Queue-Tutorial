import { SimpleGrid, Paper, Text, Flex, Loader } from '@mantine/core'
import { useState, useEffect } from 'react'
import { useOrdersGet } from '@/api'

export const ListProcessing = () => {
	const { isSuccess, data, isLoading } = useOrdersGet()
	const [, setBreathing] = useState([])

	useEffect(() => {
		const id = setInterval(() => {
			setBreathing([])
		}, 500)

		return () => {
			clearInterval(id)
		}
	}, [])

	return isLoading ? (
		<Flex justify="center" mt="xl">
			<Loader />
		</Flex>
	) : (
		isSuccess && (
			<SimpleGrid cols={{ base: 2, md: 4 }} mt="xl">
				{Object.entries(data.processing).map(
					([id_robot, { id_order, time_process, time_remaining, type }]) => {
						return (
							<Paper key={id_order} w={180} p="md">
								{[
									{ label: 'robot', value: id_robot },
									{
										label: 'order',
										value: id_order.toString().padStart(4, '0'),
									},
									{ label: 'type', value: type },
									{
										label: 'time left',
										value: `${Math.ceil(
											(new Date(time_remaining).getTime() -
												(Date.now() - new Date(time_process).getTime())) /
												1000,
										)} s`,
									},
								].map(({ label, value }) => (
									<Flex key={label}>
										<Flex justify="flex-end" w="50%">
											<Text mr="xs">{label}:</Text>
										</Flex>
										<Flex justify="flex-start" w="50%">
											<Text>{value}</Text>
										</Flex>
									</Flex>
								))}
								<Flex justify="center" mt="md">
									<Loader color="orange" />
								</Flex>
							</Paper>
						)
					},
				)}
			</SimpleGrid>
		)
	)
}
