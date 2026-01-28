import {
	ListCompleted,
	ListPending,
	ListProcessing,
	FormOrder,
} from '@/components'
import { Container, Flex } from '@mantine/core'
import { useOrderSSE } from '@/api'

function App() {
	useOrderSSE()

	return (
		<Container bg="#ffd580" p="xl">
			<Flex justify="space-between" gap="xl">
				<ListPending />
				<FormOrder />
				<ListCompleted />
			</Flex>
			<ListProcessing />
		</Container>
	)
}

export default App
