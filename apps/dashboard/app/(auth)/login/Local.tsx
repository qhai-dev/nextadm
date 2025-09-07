'use client';

import { Fragment } from 'react';
import { Button } from '@/component/base/Button';
// import { Menu, MenuButton, MenuItem, MenuItems, Transition } from '@headlessui/react';
import { DropdownMenu, DropdownMenuTrigger } from '@/component/base/dropdown-menu';

// import { GlobeAltIcon } from '@heroicons/react/24/outline';

type ISelectProps = {
	items: Array<{ value: string; name: string }>;
	value?: string;
	className?: string;
	onChange?: (value: string) => void;
};

/* <DropdownMenu>
							<DropdownMenuTrigger asChild>
								<div className='flex w-56 justify-end'>
									<Button variant='outline'>国际化功能</Button>
								</div>
							</DropdownMenuTrigger>
							<DropdownMenuContent className='w-56' align='center'>
								<DropdownMenuItem>English</DropdownMenuItem>
								<DropdownMenuItem>简体中文</DropdownMenuItem>
							</DropdownMenuContent>
						</DropdownMenu> */

export default function LocaleSigninSelect({ items, value, onChange }: ISelectProps) {
	const item = items.filter((item) => item.value === value)[0];

	return (
		<div className='w-56 text-right'>
			<DropdownMenu>
				<DropdownMenuTrigger asChild>
					<Button className='h-[44px]justify-center border-components-button-secondary-border text-text-primary hover:bg-state-base-hover inline-flex w-full items-center rounded-lg border px-[10px] py-[6px] text-[13px] font-medium'>
						{/* <GlobeAltIcon className='mr-1 h-5 w-5' aria-hidden='true' /> */}
						{item?.name}
					</Button>
				</DropdownMenuTrigger>
				{/* <Transition
					as={Fragment}
					enter='transition ease-out duration-100'
					enterFrom='transform opacity-0 scale-95'
					enterTo='transform opacity-100 scale-100'
					leave='transition ease-in duration-75'
					leaveFrom='transform opacity-100 scale-100'
					leaveTo='transform opacity-0 scale-95'
				> */}
				{/* <MenuItems className='divide-divider-regular border-components-panel-border bg-components-panel-bg-blur absolute right-0 z-10 mt-2 w-[200px] origin-top-right divide-y rounded-xl border-[0.5px] shadow-lg focus:outline-none'> */}
				{/* <div className='max-h-96 overflow-y-auto px-1 py-1 [mask-image:linear-gradient(to_bottom,transparent_0px,black_8px,black_calc(100%-8px),transparent_100%)]'>
						{items.map((item) => {
							return (
								<MenuItem key={item.value}>
									<button
										className={
											'text-text-secondary data-[active]:bg-state-base-hover group flex w-full items-center rounded-lg px-3 py-2 text-sm'
										}
										onClick={(evt) => {
											evt.preventDefault();
											onChange && onChange(item.value);
										}}
									>
										{item.name}
									</button>
								</MenuItem>
							);
						})}
					</div>
				</MenuItems> */}
				{/* </Transition> */}
			</DropdownMenu>
		</div>
	);
}
