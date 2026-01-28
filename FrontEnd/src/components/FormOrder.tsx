import { NumberInput, Paper, Button, Flex, Text } from '@mantine/core'
import { useState } from 'react'
import { useOrdersPost, useRobotPatch } from '@/api'

export const FormOrder = () => {
	const { mutateAsync: mutateAsync_orders, isPending: isPending_orders } =
		useOrdersPost()
	const { mutateAsync: mutateAsync_robots, isPending: isPending_robots } =
		useRobotPatch()
	const [number, setNumber] = useState(1)

	return (
		<Paper
			px="md"
			style={{ flexGrow: 1, flexDirection: 'column', alignItems: 'center' }}
			display="flex"
		>
			<Text ta="center" tt="uppercase" td="underline" fz="lg" fw={700} w="100%">
				settings
			</Text>
			<NumberInput
				label="Number of Robots (0-100)"
				value={number}
				onChange={v => setNumber(Number(v))}
				min={0}
				max={100}
				w="12rem"
				allowDecimal={false}
			/>
			<Button
				mt="xs"
				bg="orange"
				w="12rem"
				loading={isPending_robots}
				onClick={() => {
					mutateAsync_robots(number)
				}}
			>
				Update
			</Button>
			<Flex
				w="100%"
				align="center"
				style={{ flexGrow: 1 }}
				justify="space-evenly"
			>
				<Button
					w="8rem"
					loading={isPending_orders}
					onClick={() => {
						mutateAsync_orders('regular')
					}}
				>
					Regular Order
				</Button>
				<Button
					w="8rem"
					bg="green"
					loading={isPending_orders}
					onClick={() => {
						mutateAsync_orders('vip')
					}}
				>
					Vip Order
				</Button>
			</Flex>
		</Paper>
	)
}
