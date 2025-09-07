import { PropsWithChildren } from 'react';
import { Button } from '@/component/base/Button';
import {
	DropdownMenu,
	DropdownMenuCheckboxItem,
	DropdownMenuContent,
	DropdownMenuGroup,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuPortal,
	DropdownMenuSeparator,
	DropdownMenuShortcut,
	DropdownMenuSub,
	DropdownMenuSubContent,
	DropdownMenuSubTrigger,
	DropdownMenuTrigger,
} from '@/component/base/dropdown-menu';
import {
	Select,
	SelectContent,
	SelectGroup,
	SelectItem,
	SelectLabel,
	SelectTrigger,
	SelectValue,
} from '@/component/base/select';

export default function AuthLayout({ children }: Readonly<PropsWithChildren>) {
	return (
		<div className='bg-background-burn box-border flex min-h-screen w-full justify-center overflow-hidden p-6'>
		  <div className='border-effects-highlight bg-background-subtle flex w-full shrink-0 flex-col rounded-2xl border'>
				<div className='flex w-full items-center justify-between p-6'>
					<div></div>
					<div className='flex items-center gap-2'>
						<Select>
							<SelectTrigger className='w-[200px]'>
								<SelectValue placeholder='Select a timezone' />
							</SelectTrigger>
							<SelectContent>
								<SelectGroup>
									<SelectItem value='bot'>简体中文</SelectItem>
									<SelectItem value='art'>English</SelectItem>
								</SelectGroup>
							</SelectContent>
						</Select>
						<Button variant='ghost'>
							<img src='/theme.svg' className='size-4' alt='' />
						</Button>
					</div>
				</div>
				<div className='flex w-full grow flex-col items-center justify-center px-6 md:px-[108px]'>{children}</div>
			</div>
		</div>
	);
}
