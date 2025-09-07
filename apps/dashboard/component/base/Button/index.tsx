import * as React from 'react';
import { cn } from '@/lib/util';
import { Slot } from '@radix-ui/react-slot';
import { cva, type VariantProps } from 'class-variance-authority';

const buttonVariants = cva(
	'inline-flex shrink-0 cursor-pointer select-none items-center justify-center gap-2 whitespace-nowrap rounded-[3px] text-sm font-medium outline-none disabled:pointer-events-none disabled:opacity-50',
	{
		variants: {
			variant: {
				primary: 'bg-primary text-primary-foreground hover:bg-primary/90 shadow-xs',
				secondary:
					'bg-destructive hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60 shadow-xs text-white',
				outline:
					'bg-background hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50 shadow-xs border',
				warning: 'bg-warning text-secondary-foreground hover:bg-secondary/80 shadow-xs',
				danger: 'bg-danger text-secondary-foreground hover:bg-secondary/80 shadow-xs',
				ghost: 'hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50',
				link: 'text-primary underline-offset-4 hover:underline',
			},
		},
		defaultVariants: {
			variant: 'primary',
		},
	}
);

type ButtonProps = React.ComponentProps<'button'> &
	VariantProps<typeof buttonVariants> & {
		asChild?: boolean;
	};

export function Button({ className, variant, asChild = false, ...props }: ButtonProps) {
	const Comp = asChild ? Slot : 'button';

	return <Comp data-slot='button' className={cn(buttonVariants({ variant, className }))} {...props} />;
}

// button biz component
export function ButtonLoading({ className, variant, asChild = false, ...props }: ButtonProps) {
	return <Button>isLoading icon 文案</Button>;
}
