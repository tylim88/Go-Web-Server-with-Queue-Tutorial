import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import { QueryClientProvider } from '@tanstack/react-query'
import { MantineProvider } from '@mantine/core'
import { client_query } from '@/api'
import '@mantine/core/styles.css'

createRoot(document.getElementById('root')!).render(
	<>
		<QueryClientProvider client={client_query}>
			<MantineProvider>
				<App />
			</MantineProvider>
		</QueryClientProvider>
	</>,
)
