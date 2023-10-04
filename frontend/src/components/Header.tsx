import { useLocation } from 'preact-iso';
import { ButtonGroup, Flex, Heading, Spacer, Stack } from '@chakra-ui/react';
import { LinkButton } from '@/components/LinkButton';

// TODO: make a prop for 
export function Header() {
	const { url, route } = useLocation();
	const MenuButton = ({ to, children }) => (
		<LinkButton to={to} from={url} navigate={route}>
			{children}
		</LinkButton>
	);

	return (
		<Flex direction="column">
			<Flex as="header" direction="row" padding={3} background="purple.600">
				<Heading>
					SAC Locker Room App
				</Heading>
				<Spacer />
				<Stack
					direction={{ base: "column", md: "row" }}
					display="flex"
					alignItems="center"
				>
					<ButtonGroup flexDirection={{ base: "column", md: "row" }}>
						<MenuButton to="/">
							Home
						</MenuButton>
						<MenuButton to="/404">
							404
						</MenuButton>
					</ButtonGroup>
				</Stack>
			</Flex>
			<Spacer padding={2} />
		</Flex>
	);
}
