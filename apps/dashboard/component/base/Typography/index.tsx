import * as React from 'react';
import { cn } from '@/lib/util';
import { Slot } from '@radix-ui/react-slot';
import { cva, type VariantProps } from 'class-variance-authority';

const titleVariants = cva('', {
	variants: {
		size: {
			h1: '',
			h2: 'text-4xl',
			sm: 'h-8 gap-1.5 rounded-md px-3 has-[>svg]:px-2.5',
			lg: 'h-10 rounded-md px-6 has-[>svg]:px-4',
			icon: 'size-9',
		},
	},
	defaultVariants: {
		size: 'h1',
	},
});

const textVariants = cva('', {
	variants: {
		size: {
			h1: '',
			h2: 'text-4xl',
			sm: 'h-8 gap-1.5 rounded-md px-3 has-[>svg]:px-2.5',
			lg: 'h-10 rounded-md px-6 has-[>svg]:px-4',
			icon: 'size-9',
		},
	},
	defaultVariants: {
		size: 'h1',
	},
});

export function Title({
	className,
	size,
	asChild = false,
	children,
}: React.ComponentProps<'div'> & VariantProps<typeof titleVariants> & { asChild?: boolean }) {
	const Comp = asChild ? Slot : 'div';
	return <Comp className={cn(titleVariants({ className, size }))}>{children}</Comp>;
}

export function Text({
	className,
	size,
	asChild = false,
	children,
}: React.ComponentProps<'span'> & VariantProps<typeof textVariants> & { asChild?: boolean }) {
	const Comp = asChild ? Slot : 'span';
	return <Comp className={cn(titleVariants({ className, size }))}>{children}</Comp>;
}
