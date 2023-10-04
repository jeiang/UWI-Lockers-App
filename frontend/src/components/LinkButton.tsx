import { Button, ButtonProps } from '@chakra-ui/react';

export interface LinkButtonProps extends ButtonProps {
  to: string
  from: string
  navigate: (url: string) => void
}

export const LinkButton = ({ navigate, to, from, children, ...props }: LinkButtonProps) => (
  <Button 
    onClick={() => navigate(to)} 
    isActive={to == from} 
    colorScheme='purple' 
    cursor={to == from ? "pointer" : "default"} 
    {...props}
  >
    {children}
  </Button>
)

