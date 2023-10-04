import { render } from 'preact';
import { LocationProvider, Router, Route } from 'preact-iso';

import { Header } from '@/components/Header.jsx';
import { Home } from '@/pages/Home/index.jsx';
import { NotFound } from '@/pages/_404.jsx';
import { ChakraProvider } from '@chakra-ui/react';
import { theme, FontsProvider } from '@/theme';

export function App() {
	return (
		<ChakraProvider theme={theme}>
			<FontsProvider />
			<LocationProvider>
				<Header />
				<main>
					<Router>
						<Route path="/" component={Home} />
						<Route default component={NotFound} />
					</Router>
				</main>
			</LocationProvider>
		</ChakraProvider>
	);
}

render(<App />, document.getElementById('app'));
